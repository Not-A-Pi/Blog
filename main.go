package main

import (
	"log"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
	"./handlers"
)

func main() {
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	log.Println("Server started on port " + port)
	http.ListenAndServe(":" + port, nil)

}
