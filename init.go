package main

import (
	"./db/dbinit"
	"log"
	_"github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Initializing Not-A-Pi")
	dbinit.FirstDB()
	log.Println("Initialized!")
}
