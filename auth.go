package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// --- Sessions ---

const sessionTTL = 8 * time.Hour
const cookieName = "pbsession"

var (
	sessionsMu sync.Mutex
	sessions   = map[string]time.Time{}
)

func newToken() string {
	b := make([]byte, 24)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func createSession() string {
	t := newToken()
	sessionsMu.Lock()
	sessions[t] = time.Now().Add(sessionTTL)
	sessionsMu.Unlock()
	return t
}

func validSession(token string) bool {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()
	exp, ok := sessions[token]
	if !ok {
		return false
	}
	if time.Now().After(exp) {
		delete(sessions, token)
		return false
	}
	return true
}

func dropSession(token string) {
	sessionsMu.Lock()
	delete(sessions, token)
	sessionsMu.Unlock()
}

// --- Middleware ---

func adminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Пропускаем страницу логина
		if r.URL.Path == "/login" || r.URL.Path == "/logout" {
			next.ServeHTTP(w, r)
			return
		}
		// Только /admin и вложенные
		cookie, err := r.Cookie(cookieName)
		if err != nil || !validSession(cookie.Value) {
			http.Redirect(w, r, "/login?next="+r.URL.Path, http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// --- Handlers ---

func handleLogin(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	if next == "" {
		next = "/admin"
	}

	if r.Method == http.MethodGet {
		// Уже залогинен?
		if cookie, err := r.Cookie(cookieName); err == nil && validSession(cookie.Value) {
			http.Redirect(w, r, next, http.StatusFound)
			return
		}
		errMsg := ""
		if r.URL.Query().Get("err") == "1" {
			errMsg = "Неверный логин или пароль"
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		writeLoginPage(w, errMsg, next)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		nextURL := r.FormValue("next")
		if nextURL == "" {
			nextURL = "/admin"
		}

		hash, err := getAdminHash(username)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
			http.Redirect(w, r, "/login?err=1&next="+nextURL, http.StatusFound)
			return
		}

		token := createSession()
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(sessionTTL.Seconds()),
		})
		http.Redirect(w, r, nextURL, http.StatusFound)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie(cookieName); err == nil {
		dropSession(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{Name: cookieName, Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusFound)
}

// --- Password hashing ---

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

// --- Login page ---

func writeLoginPage(w http.ResponseWriter, errMsg, next string) {
	errBlock := ""
	if errMsg != "" {
		errBlock = `<div class="error">` + errMsg + `</div>`
	}
	w.Write([]byte(`<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>Вход — Администрирование</title>
<style>
*{box-sizing:border-box;margin:0;padding:0}
body{font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;background:#eef0f4;display:flex;align-items:center;justify-content:center;min-height:100vh}
.card{background:#fff;border:1px solid #d0d5e8;border-radius:14px;padding:38px 36px;width:340px;box-shadow:0 4px 24px rgba(0,0,0,.10);display:flex;flex-direction:column;gap:18px}
.logo-row{display:flex;align-items:center;gap:12px;justify-content:center;margin-bottom:4px}
.logo-row img{width:40px;height:40px;object-fit:contain}
h1{font-size:17px;font-weight:700;color:#1a1f36;text-align:center}
.sub{font-size:12px;color:#8a93b2;text-align:center;margin-top:-10px}
label{font-size:12px;color:#5a6380;font-weight:500;display:flex;flex-direction:column;gap:5px}
input{background:#f4f6fb;border:1px solid #d0d5e8;border-radius:7px;color:#1a1f36;padding:9px 12px;font-size:14px;width:100%;outline:none}
input:focus{border-color:#3b5bdb;background:#fff}
button{background:#3b5bdb;color:#fff;border:none;border-radius:8px;padding:11px;font-size:14px;font-weight:600;cursor:pointer;margin-top:2px}
button:hover{background:#2f4ac9}
.error{background:#fff5f5;border:1px solid #ffc9c9;color:#e03131;font-size:13px;padding:10px 14px;border-radius:8px;text-align:center}
</style>
</head>
<body>
<form class="card" method="POST" action="/login">
  <input type="hidden" name="next" value="` + next + `"/>
  <div class="logo-row">
    <img src="/logo" onerror="this.style.display='none'" alt=""/>
  </div>
  <h1>Администрирование</h1>
  <p class="sub">Телефонный справочник</p>
  ` + errBlock + `
  <label>Логин
    <input type="text" name="username" autocomplete="username" autofocus required/>
  </label>
  <label>Пароль
    <input type="password" name="password" autocomplete="current-password" required/>
  </label>
  <button type="submit">Войти</button>
</form>
</body>
</html>`))
}
