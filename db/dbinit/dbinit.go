package dbinit

import (
	"log"
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func checkerr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FirstDB() {
	var (
		user_exist string
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DBUSER"), os.Getenv("PASSWD"), os.Getenv("DBNAME"))
	db, _ := sql.Open("postgres", psqlInfo)
	defer db.Close()

	err := db.QueryRow("select exists (select from information_schema.tables where table_name = 'users')").Scan(&user_exist)
	checkerr(err)
	log.Println(user_exist)
	return
}
