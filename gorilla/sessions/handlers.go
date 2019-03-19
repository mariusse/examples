package main

import (
	"fmt"
	"net/http"
)

// indexHandler checks authentication
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET") // rfc2616#section-10.4.6
		http.Error(w, http.StatusText(405), 405)
		return
	}

	session, _ := store.Get(r, cookie)
	if _, ok := session.Values["authenticated"]; !ok {
		tpl.ExecuteTemplate(w, "login.html.tmpl", nil)
		return
	}

	tpl.ExecuteTemplate(w, "index.html.tmpl", nil)
	return
}

// authHandler authenticates users
func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST") // rfc2616#section-10.4.6
		http.Error(w, http.StatusText(405), 405)
		return
	}

	user := r.FormValue("user")
	pass := r.FormValue("pass")

	// do authentication stuff here

	session, _ := store.Get(r, cookie)
	session.Values["authenticated"] = true
	session.Values["user"] = user
	if err := session.Save(r, w); err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 302)
}
