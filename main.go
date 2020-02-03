package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("static/")))

	log.Println("Server started on port " + port)

	http.ListenAndServe(":" + port, nil)

}
