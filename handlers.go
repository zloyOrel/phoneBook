package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var viewTmpl = template.Must(template.New("view").Parse(viewPage))
var adminTmpl = template.Must(template.New("admin").Parse(adminPage))

func seeOther(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// --- Public view ---

func handleView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	orgParam := r.URL.Query().Get("org")

	orgs, err := allOrgs()
	if err != nil {
		log.Printf("allOrgs error: %v", err)
	}

	// Определяем текущую организацию
	orgID := 0
	orgName := ""
	if orgParam == "" {
		// Нет параметра — использовать организацию по умолчанию
		for _, o := range orgs {
			if o.IsDefault {
				orgID = o.ID
				orgName = o.Name
				break
			}
		}
	} else if orgParam != "0" {
		id, _ := strconv.Atoi(orgParam)
		orgID = id
		for _, o := range orgs {
			if o.ID == id {
				orgName = o.Name
				break
			}
		}
	}

	depts, err := deptTree("", orgID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	orgsJSON, _ := json.Marshal(orgs)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	viewTmpl.Execute(w, map[string]interface{}{
		"Departments": depts,
		"Search":      q,
		"Orgs":        orgs,
		"OrgsJSON":    template.JS(orgsJSON),
		"CurrentOrg":  orgID,
		"OrgName":     orgName,
		"Copyright":   appConfig.Copyright,
		"SiteTitle":   appConfig.SiteTitle,
		"HeaderLines": appConfig.HeaderLines,
	})
}

// --- Admin view ---

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	orgParam := r.URL.Query().Get("org")

	orgs, orgsErr := allOrgs()
	if orgsErr != nil {
		log.Printf("allOrgs error: %v", orgsErr)
	}

	orgID := 0
	orgName := ""
	if orgParam != "" && orgParam != "0" {
		id, _ := strconv.Atoi(orgParam)
		orgID = id
		for _, o := range orgs {
			if o.ID == id {
				orgName = o.Name
				break
			}
		}
	}

	depts, err := deptTree("", orgID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	roots, _ := allRootDepts()
	deptOpts, _ := allDeptsFlat()
	deptOptsJSON, _ := json.Marshal(deptOpts)
	orgsJSON, _ := json.Marshal(orgs)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	adminTmpl.Execute(w, map[string]interface{}{
		"Departments":  depts,
		"RootDepts":    roots,
		"Orgs":         orgs,
		"OrgsJSON":     template.JS(orgsJSON),
		"DeptOptsJSON": template.JS(deptOptsJSON),
		"Search":       q,
		"CurrentOrg":   orgID,
		"OrgName":      orgName,
		"Copyright":    appConfig.Copyright,
		"SiteTitle":    appConfig.SiteTitle,
		"HeaderLines":  appConfig.HeaderLines,
	})
}

// --- Department actions ---

func handleDeptAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	name := strings.TrimSpace(r.FormValue("name"))
	parentID, _ := strconv.Atoi(r.FormValue("parent_id"))
	sort, _ := strconv.Atoi(r.FormValue("sort_order"))
	orgID, _ := strconv.Atoi(r.FormValue("organization_id"))
	if name != "" {
		insertDept(name, parentID, sort, orgID)
	}
	seeOther(w, r, "/admin")
}

func handleDeptEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := strings.TrimSpace(r.FormValue("name"))
	parentID, _ := strconv.Atoi(r.FormValue("parent_id"))
	sort, _ := strconv.Atoi(r.FormValue("sort_order"))
	orgID, _ := strconv.Atoi(r.FormValue("organization_id"))
	if id > 0 && name != "" {
		updateDept(id, name, parentID, sort, orgID)
	}
	seeOther(w, r, "/admin")
}

func handleDeptDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	if id > 0 {
		deleteDept(id)
	}
	seeOther(w, r, "/admin")
}

// --- Contact actions ---

func handleContactAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	deptID, _ := strconv.Atoi(r.FormValue("dept_id"))
	if deptID > 0 {
		insertContact(Contact{
			DepartmentID:  deptID,
			Room:          r.FormValue("room"),
			Position:      r.FormValue("position"),
			FullName:      r.FormValue("full_name"),
			PhoneCity:     r.FormValue("phone_city"),
			PhoneMobile:   r.FormValue("phone_mobile"),
			PhoneInternal: r.FormValue("phone_internal"),
			Email:         r.FormValue("email"),
		})
	}
	seeOther(w, r, "/admin")
}

func handleContactEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	deptID, _ := strconv.Atoi(r.FormValue("dept_id"))
	if id > 0 {
		updateContact(Contact{
			ID:            id,
			DepartmentID:  deptID,
			Room:          r.FormValue("room"),
			Position:      r.FormValue("position"),
			FullName:      r.FormValue("full_name"),
			PhoneCity:     r.FormValue("phone_city"),
			PhoneMobile:   r.FormValue("phone_mobile"),
			PhoneInternal: r.FormValue("phone_internal"),
			Email:         r.FormValue("email"),
		})
	}
	seeOther(w, r, "/admin")
}

func handleContactDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	if id > 0 {
		deleteContact(id)
	}
	seeOther(w, r, "/admin")
}

// --- Organization actions ---

func handleOrgAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	name := strings.TrimSpace(r.FormValue("name"))
	if name != "" {
		insertOrg(name)
	}
	seeOther(w, r, "/admin")
}

func handleOrgEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := strings.TrimSpace(r.FormValue("name"))
	if id > 0 && name != "" {
		updateOrg(id, name)
	}
	seeOther(w, r, "/admin")
}

func handleOrgDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	if id > 0 {
		deleteOrg(id)
	}
	seeOther(w, r, "/admin")
}

func handleOrgSetDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	setDefaultOrg(id)
	seeOther(w, r, "/admin")
}
