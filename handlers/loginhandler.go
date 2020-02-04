package handlers

import (
	"net/http"
	"html/template"
	"log"
)

type LoginPage struct {
	Title string
	User string
	Pass string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := LoginPage{Title: "Login"}
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, p)
	case "POST":
		log.Println(r.FormValue("username"))
		log.Println(r.FormValue("password"))
	}
}
