package main

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"./handlers"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", 
	os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWD"), os.Getenv("DBNAME"))
	db, _ := sql.Open("postgres", psqlInfo)
	defer db.Close()
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	log.Println("Server started on port " + port)

	http.ListenAndServe(":" + port, nil)

}
