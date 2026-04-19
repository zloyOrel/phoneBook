package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

// --- Config ---

type Config struct {
	Host        string   `json:"host"`
	Port        int      `json:"port"`
	User        string   `json:"user"`
	Password    string   `json:"password"`
	DBName      string   `json:"dbname"`
	Listen      string   `json:"listen"`
	TLSCert     string   `json:"tls_cert"`
	TLSKey      string   `json:"tls_key"`
	Copyright   string   `json:"copyright"`
	SiteTitle   string   `json:"site_title"`
	HeaderLines []string `json:"header_lines"`
}

func exeDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

func configPath() string { return filepath.Join(exeDir(), "config.json") }

func loadConfig() Config {
	cfg := Config{Host: "localhost", Port: 5432, User: "postgres", DBName: "phonebook", Listen: ":8080"}
	data, err := os.ReadFile(configPath())
	if err != nil {
		saveConfig(cfg)
		return cfg
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Printf("Ошибка чтения config.json: %v", err)
	}
	return cfg
}

func saveConfig(cfg Config) {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(configPath(), data, 0600)
}

// --- Models ---

type Department struct {
	ID             int
	ParentID       int // 0 = корень
	Name           string
	SortOrder      int
	OrganizationID int
	Contacts       []Contact
	Children       []Department
}

type Organization struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

type DeptOption struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	OrgID   int    `json:"org_id"`
	OrgName string `json:"org_name"`
}

type Contact struct {
	ID            int
	DepartmentID  int
	Room          string
	Position      string
	FullName      string
	PhoneCity     string
	PhoneMobile   string
	PhoneInternal string
	Email         string
}

// --- DB init ---

var db *sql.DB
var appConfig Config

func initDB() {
	appConfig = loadConfig()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.Host, appConfig.Port, appConfig.User, appConfig.Password, appConfig.DBName)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка открытия БД:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Не могу подключиться к PostgreSQL: %v", err)
	}
	migrate()
	log.Println("База данных подключена")
}

func migrate() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS organizations (
			id         SERIAL PRIMARY KEY,
			name       TEXT NOT NULL,
			is_default BOOLEAN NOT NULL DEFAULT false
		);
		CREATE TABLE IF NOT EXISTS departments (
			id         SERIAL PRIMARY KEY,
			parent_id  INT  REFERENCES departments(id) ON DELETE CASCADE,
			name       TEXT NOT NULL,
			sort_order INT  NOT NULL DEFAULT 0
		);
		CREATE TABLE IF NOT EXISTS contacts (
			id             SERIAL PRIMARY KEY,
			department_id  INT  NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
			room           TEXT NOT NULL DEFAULT '',
			position       TEXT NOT NULL DEFAULT '',
			full_name      TEXT NOT NULL DEFAULT '',
			phone_city     TEXT NOT NULL DEFAULT '',
			phone_mobile   TEXT NOT NULL DEFAULT '',
			phone_internal TEXT NOT NULL DEFAULT '',
			email          TEXT NOT NULL DEFAULT ''
		);
		CREATE TABLE IF NOT EXISTS admins (
			id            SERIAL PRIMARY KEY,
			username      TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Ошибка создания таблиц:", err)
	}
	// Миграции: добавить новые поля если таблицы уже существовали
	db.Exec(`ALTER TABLE departments ADD COLUMN IF NOT EXISTS parent_id INT REFERENCES departments(id) ON DELETE CASCADE`)
	db.Exec(`ALTER TABLE departments ADD COLUMN IF NOT EXISTS organization_id INT REFERENCES organizations(id) ON DELETE SET NULL`)
}

// --- Admins ---

func getAdminHash(username string) (string, error) {
	var hash string
	err := db.QueryRow(`SELECT password_hash FROM admins WHERE username=$1`, username).Scan(&hash)
	return hash, err
}

func createAdmin(username, passwordHash string) error {
	_, err := db.Exec(`INSERT INTO admins (username, password_hash) VALUES ($1, $2)
	                   ON CONFLICT (username) DO UPDATE SET password_hash=EXCLUDED.password_hash`,
		username, passwordHash)
	return err
}

func adminCount() int {
	var n int
	db.QueryRow(`SELECT COUNT(*) FROM admins`).Scan(&n)
	return n
}

// --- Department queries ---

// allRootDepts возвращает только корневые подразделения (без parent_id) — для селектора в форме
func allRootDepts() ([]Department, error) {
	rows, err := db.Query(`SELECT id, name FROM departments WHERE parent_id IS NULL ORDER BY sort_order, id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var depts []Department
	for rows.Next() {
		var d Department
		rows.Scan(&d.ID, &d.Name)
		depts = append(depts, d)
	}
	return depts, nil
}

// deptTree строит дерево подразделений с контактами. При поиске — фильтрует по контактам.
// orgID > 0 — показывать только подразделения выбранной организации.
func deptTree(search string, orgID int) ([]Department, error) {
	// Загружаем отделы (с фильтром по организации если задан)
	var rows *sql.Rows
	var err error
	if orgID > 0 {
		rows, err = db.Query(`
			SELECT id, COALESCE(parent_id,0), name, sort_order, COALESCE(organization_id,0)
			FROM departments
			WHERE organization_id=$1
			   OR parent_id IN (SELECT id FROM departments WHERE organization_id=$1)
			ORDER BY sort_order, id`, orgID)
	} else {
		rows, err = db.Query(`
			SELECT id, COALESCE(parent_id,0), name, sort_order, COALESCE(organization_id,0)
			FROM departments ORDER BY sort_order, id`)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	byID := map[int]*Department{}
	var order []int
	for rows.Next() {
		var d Department
		rows.Scan(&d.ID, &d.ParentID, &d.Name, &d.SortOrder, &d.OrganizationID)
		cp := d
		byID[d.ID] = &cp
		order = append(order, d.ID)
	}
	rows.Close()

	// Загружаем контакты
	for id := range byID {
		contacts, err := contactsByDept(id, search)
		if err != nil {
			return nil, err
		}
		byID[id].Contacts = contacts
	}

	// Строим дерево
	var roots []Department
	for _, id := range order {
		d := byID[id]
		if d.ParentID == 0 {
			// корень — будет добавлен ниже
			continue
		}
		if parent, ok := byID[d.ParentID]; ok {
			parent.Children = append(parent.Children, *d)
		}
	}
	for _, id := range order {
		d := byID[id]
		if d.ParentID != 0 {
			continue
		}
		// обновляем Children из map (они уже добавлены)
		roots = append(roots, *d)
	}

	// Если поиск — оставляем только те ветки, где есть контакты
	if search != "" {
		var filtered []Department
		for _, root := range roots {
			// фильтруем детей
			var filteredChildren []Department
			for _, child := range root.Children {
				if len(child.Contacts) > 0 {
					filteredChildren = append(filteredChildren, child)
				}
			}
			root.Children = filteredChildren
			if len(root.Contacts) > 0 || len(root.Children) > 0 {
				filtered = append(filtered, root)
			}
		}
		return filtered, nil
	}
	return roots, nil
}

func getDeptByID(id int) (Department, error) {
	var d Department
	var pid sql.NullInt64
	var oid sql.NullInt64
	err := db.QueryRow(`SELECT id, COALESCE(parent_id,0), name, sort_order, organization_id FROM departments WHERE id=$1`, id).
		Scan(&d.ID, &pid, &d.Name, &d.SortOrder, &oid)
	if pid.Valid {
		d.ParentID = int(pid.Int64)
	}
	if oid.Valid {
		d.OrganizationID = int(oid.Int64)
	}
	return d, err
}

func insertDept(name string, parentID int, sortOrder int, orgID int) error {
	var orgVal interface{}
	if orgID > 0 {
		orgVal = orgID
	}
	if parentID == 0 {
		_, err := db.Exec(`INSERT INTO departments (name, sort_order, organization_id) VALUES ($1, $2, $3)`, name, sortOrder, orgVal)
		return err
	}
	_, err := db.Exec(`INSERT INTO departments (name, parent_id, sort_order, organization_id) VALUES ($1, $2, $3, $4)`, name, parentID, sortOrder, orgVal)
	return err
}

func updateDept(id int, name string, parentID int, sortOrder int, orgID int) error {
	var orgVal interface{}
	if orgID > 0 {
		orgVal = orgID
	}
	if parentID == 0 {
		_, err := db.Exec(`UPDATE departments SET name=$1, parent_id=NULL, sort_order=$2, organization_id=$3 WHERE id=$4`, name, sortOrder, orgVal, id)
		return err
	}
	_, err := db.Exec(`UPDATE departments SET name=$1, parent_id=$2, sort_order=$3, organization_id=$4 WHERE id=$5`, name, parentID, sortOrder, orgVal, id)
	return err
}

func deleteDept(id int) error {
	_, err := db.Exec(`DELETE FROM departments WHERE id=$1`, id)
	return err
}

// --- Contact queries ---

func contactsByDept(deptID int, search string) ([]Contact, error) {
	q := `SELECT id, department_id, room, position, full_name, phone_city, phone_mobile, phone_internal, email
	      FROM contacts WHERE department_id=$1`
	args := []interface{}{deptID}

	if search != "" {
		s := "%" + strings.ToLower(search) + "%"
		q += ` AND (lower(full_name) LIKE $2 OR lower(position) LIKE $2
		             OR phone_city LIKE $2 OR phone_mobile LIKE $2
		             OR phone_internal LIKE $2 OR lower(email) LIKE $2)`
		args = append(args, s)
	}
	q += ` ORDER BY id`

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var c Contact
		rows.Scan(&c.ID, &c.DepartmentID, &c.Room, &c.Position,
			&c.FullName, &c.PhoneCity, &c.PhoneMobile, &c.PhoneInternal, &c.Email)
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func insertContact(c Contact) error {
	_, err := db.Exec(
		`INSERT INTO contacts (department_id, room, position, full_name, phone_city, phone_mobile, phone_internal, email)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		c.DepartmentID, c.Room, c.Position, c.FullName,
		c.PhoneCity, c.PhoneMobile, c.PhoneInternal, c.Email,
	)
	return err
}

func updateContact(c Contact) error {
	_, err := db.Exec(
		`UPDATE contacts SET room=$1, position=$2, full_name=$3,
		 phone_city=$4, phone_mobile=$5, phone_internal=$6, email=$7 WHERE id=$8`,
		c.Room, c.Position, c.FullName,
		c.PhoneCity, c.PhoneMobile, c.PhoneInternal, c.Email, c.ID,
	)
	return err
}

func deleteContact(id int) error {
	_, err := db.Exec(`DELETE FROM contacts WHERE id=$1`, id)
	return err
}

// --- Organization queries ---

func allDeptsFlat() ([]DeptOption, error) {
	rows, err := db.Query(`
		SELECT d.id, d.name, COALESCE(d.organization_id,0), COALESCE(o.name,'— Без организации —')
		FROM departments d
		LEFT JOIN organizations o ON o.id = d.organization_id
		ORDER BY COALESCE(o.name,''), d.sort_order, d.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var opts []DeptOption
	for rows.Next() {
		var opt DeptOption
		rows.Scan(&opt.ID, &opt.Name, &opt.OrgID, &opt.OrgName)
		opts = append(opts, opt)
	}
	return opts, nil
}

func allOrgs() ([]Organization, error) {
	rows, err := db.Query(`SELECT id, name, is_default FROM organizations ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orgs []Organization
	for rows.Next() {
		var o Organization
		rows.Scan(&o.ID, &o.Name, &o.IsDefault)
		orgs = append(orgs, o)
	}
	return orgs, nil
}

func insertOrg(name string) error {
	_, err := db.Exec(`INSERT INTO organizations (name) VALUES ($1)`, name)
	return err
}

func updateOrg(id int, name string) error {
	_, err := db.Exec(`UPDATE organizations SET name=$1 WHERE id=$2`, name, id)
	return err
}

func deleteOrg(id int) error {
	_, err := db.Exec(`DELETE FROM organizations WHERE id=$1`, id)
	return err
}

func setDefaultOrg(id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`UPDATE organizations SET is_default=false`); err != nil {
		tx.Rollback()
		return err
	}
	if id > 0 {
		if _, err := tx.Exec(`UPDATE organizations SET is_default=true WHERE id=$1`, id); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
