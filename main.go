package main

import (
	"log"
	"net/http"
	"html/template"
)

type LoginPage struct {
	Title string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	p := LoginPage{Title: "Login"}
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, p)
}

func main() {
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.HandleFunc("/login", loginHandler)

	log.Println("Server started on port " + port)

	http.ListenAndServe(":" + port, nil)

}
