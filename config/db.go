package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/onlineshop?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("DB connected")
}
