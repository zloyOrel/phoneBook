package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	addAdmin := flag.String("add-admin", "", "Создать/обновить администратора: -add-admin логин:пароль")
	flag.Parse()

	initDB()
	defer db.Close()

	// Режим создания администратора
	if *addAdmin != "" {
		parts := strings.SplitN(*addAdmin, ":", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			fmt.Fprintln(os.Stderr, "Формат: -add-admin логин:пароль")
			os.Exit(1)
		}
		hash, err := hashPassword(parts[1])
		if err != nil {
			log.Fatal("Ошибка хеширования пароля:", err)
		}
		if err := createAdmin(parts[0], hash); err != nil {
			log.Fatal("Ошибка создания администратора:", err)
		}
		fmt.Printf("Администратор «%s» сохранён.\n", parts[0])
		return
	}

	// Предупреждение если нет ни одного администратора
	if adminCount() == 0 {
		log.Println("ВНИМАНИЕ: нет администраторов. Создайте первого:")
		log.Printf("  ./phonebook -add-admin логин:пароль")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/logo", handleLogo)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/logout", handleLogout)

	// Публичный просмотр
	mux.HandleFunc("/", handleView)

	// /admin и вложенные — за авторизацией
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin", handleAdmin)
	adminMux.HandleFunc("/admin/dept/add", handleDeptAdd)
	adminMux.HandleFunc("/admin/dept/edit", handleDeptEdit)
	adminMux.HandleFunc("/admin/dept/delete", handleDeptDelete)
	adminMux.HandleFunc("/admin/contact/add", handleContactAdd)
	adminMux.HandleFunc("/admin/contact/edit", handleContactEdit)
	adminMux.HandleFunc("/admin/contact/delete", handleContactDelete)
	adminMux.HandleFunc("/admin/org/add", handleOrgAdd)
	adminMux.HandleFunc("/admin/org/edit", handleOrgEdit)
	adminMux.HandleFunc("/admin/org/delete", handleOrgDelete)
	adminMux.HandleFunc("/admin/org/setdefault", handleOrgSetDefault)

	mux.Handle("/admin", adminAuth(adminMux))
	mux.Handle("/admin/", adminAuth(adminMux))

	if appConfig.TLSCert != "" && appConfig.TLSKey != "" {
		log.Printf("Сервер запущен: https://localhost%s", appConfig.Listen)
		log.Fatal(http.ListenAndServeTLS(appConfig.Listen, appConfig.TLSCert, appConfig.TLSKey, mux))
	} else {
		log.Printf("Сервер запущен: http://localhost%s", appConfig.Listen)
		log.Fatal(http.ListenAndServe(appConfig.Listen, mux))
	}
}

func handleLogo(w http.ResponseWriter, r *http.Request) {
	dir := exeDir()
	for _, name := range []string{"logo.png", "logo.jpg", "logo.jpeg", "logo.svg", "logo.gif"} {
		path := filepath.Join(dir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		ext := strings.ToLower(filepath.Ext(name))
		ct := mime.TypeByExtension(ext)
		if ct == "" {
			ct = "image/png"
		}
		w.Header().Set("Content-Type", ct)
		w.Header().Set("Cache-Control", "max-age=3600")
		w.Write(data)
		return
	}
	http.NotFound(w, r)
}
