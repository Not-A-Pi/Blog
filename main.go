package main

import (
	"log"
	"net/http"
	"html/template"
)

type LoginPage struct {
	Title string
	User string
	Pass string
}


func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := LoginPage{Title: "Login"}
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, p)
	case "POST":
		log.Println("Post request!")
	}
}

func main() {
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.HandleFunc("/login", loginHandler)

	log.Println("Server started on port " + port)

	http.ListenAndServe(":" + port, nil)

}
