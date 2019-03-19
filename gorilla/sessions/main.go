package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	tpl      *template.Template
	store    *sessions.CookieStore
	cookie   = "sessioncookie-name"
	bindUser = os.Getenv("PROD_BINDER")
	bindPass = os.Getenv("PROD_PASS")
)

func main() {
	tpl = template.Must(template.ParseGlob("template/*.tmpl"))

	store = sessions.NewCookieStore(
		// []byte(os.Getenv("AUTH_KEY_ONE")),
		// []byte(os.Getenv("ENCRYPTION_KEY_ONE")),
		[]byte("AUTH_KEY_ONE"),
		[]byte("ENCRYPTION_KEY_ONE______________"),
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 60 * 1,
		HttpOnly: true,
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth", authHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
