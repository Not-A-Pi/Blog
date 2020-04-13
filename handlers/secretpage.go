package handlers

import (
	"net/http"
	"html/template"
	"log"
)

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	log.Println(c.Value)
	log.Println(Sessions["spencer"])
	if (c.Value == Sessions["spencer"]) {
		log.Println("passed")
	}
	t, _ := template.ParseFiles("templates/secret.html")
	t.Execute(w, "empty")
}
