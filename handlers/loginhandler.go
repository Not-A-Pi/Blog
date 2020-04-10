package handlers

import (
	"net/http"
	"html/template"
	"log"
	"../structs/Loginpage"
	"database/sql"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := Loginpage.LoginPage{Title: "Login"}
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, p)
	case "POST":
		godotenv.Load("../.env")
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DBUSER"), os.Getenv("PASSWD"), os.Getenv("DBNAME"))
		db, _ := sql.Open("postgres", psqlInfo)
		defer db.Close()
		var passhash string
		_ = db.QueryRow("select password from users where username=$1", r.FormValue("username")).Scan(&passhash)
		err := bcrypt.CompareHashAndPassword([]byte(passhash), []byte(r.FormValue("password")))
		if err == nil {
			log.Println("Logged in")
		} else {
			log.Println("Not logged in")
		}
	}
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")

	if err != nil {
		panic(err)
	}

	cache = conn
}
