package handlers

import (
	"fmt"
	"net/http"
	"html/template"
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


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := Registerpage{Title: "Register"}
		t, _ := template.ParseFiles("templates/register.html")
		t.Execute(w, p)
	case "POST":
		godotenv.Load("../.env")
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DBUSER"), os.Getenv("PASSWD"), os.Getenv("DBNAME"))
		db, err := sql.Open("postgres", psqlInfo)
		checkErr(err)
		defer db.Close()

		if err = db.Ping(); err != nil {
			panic(err)
		}

		rows, err := db.Query("SELECT $1 FROM users", r.FormValue("username"))
		checkErr(err)
		defer rows.Close()
	}
}
