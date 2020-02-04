package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"os"
	//"../structs/Registerpage"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type Registerpage struct {
	Title string
	User string
	Pass string
}



func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := Registerpage{Title: "Register"}
		t, _ := template.ParseFiles("templates/register.html")
		t.Execute(w, p)
	case "POST":
		godotenv.Load("../.env")
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWD"), os.Getenv("DBNAME"))
		db, _ := sql.Open("postgres", psqlInfo)
		defer db.Close()
		log.Println(r.FormValue("username"))
		log.Println(r.FormValue("password"))
	}
}
