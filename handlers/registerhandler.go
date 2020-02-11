package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
	"log"
	"time"
	//"../structs/Registerpage"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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

		if r.FormValue("password") != r.FormValue("confirmpassword") {
			return
		}

		var (
			username string
		)

		err = db.QueryRow("SELECT username FROM users WHERE username=$1",
			r.FormValue("username")).Scan(&username)

		if err == sql.ErrNoRows {
			passhash, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)
			regtime := time.Now().UTC()
			_, err = db.Exec("insert into users (username, password, creation_date) values ($1, $2, $3)",
				r.FormValue("username"), string(passhash), regtime.Format("01-02-2006 15:04:05"))
		} else {
			log.Println("user " + r.FormValue("username") + " is registered")
		}
	}
}
