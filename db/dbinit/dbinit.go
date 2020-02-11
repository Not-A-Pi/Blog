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
		posts_exist string
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DBUSER"), os.Getenv("PASSWD"), os.Getenv("DBNAME"))
	db, _ := sql.Open("postgres", psqlInfo)
	defer db.Close()

	err := db.QueryRow("select exists (select from information_schema.tables where table_name = 'users')").Scan(&user_exist)
	checkerr(err)
	err = db.QueryRow("select exists (select from information_schema.tables where table_name = 'posts')").Scan(&posts_exist)
	log.Println("user table exist: " + user_exist)
	log.Println("posts table exist: " + posts_exist)
	if user_exist == "false" {
		_, err = db.Exec("create extension if not exists \"uuid-ossp\"")
		checkerr(err)
		_, err = db.Exec(`create table users (
			user_id uuid DEFAULT uuid_generate_v4(), 
			username text NOT NULL UNIQUE, 
			password text NOT NULL, 
			creation_date timestamp NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'), 
			PRIMARY KEY (user_id))`)
		checkerr(err)
	}
	if posts_exist == "false" {
		_, err = db.Exec(`create table posts (
		tags text[],
		author_uuid uuid NOT NULL,
		title text NOT NULL,
		body text NOT NULL,
		creation_date timestamp NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
		post_id serial,
		PRIMARY KEY (post_id))`)
		checkerr(err)
	}
	return
}
