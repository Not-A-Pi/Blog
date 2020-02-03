package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"./env"
	"database/sql"
	_ "github.com/lib/pq"
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
		log.Println(r.FormValue("username"))
		log.Println(r.FormValue("password"))
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", 
	env.Host, env.Port, env.User, env.Passwd, env.Dbname)
	db, _ := sql.Open("postgres", psqlInfo)
	defer db.Close()
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.HandleFunc("/login", loginHandler)

	log.Println("Server started on port " + port)

	http.ListenAndServe(":" + port, nil)

}
