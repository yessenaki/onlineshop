package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(source string) {
	var err error
	DB, err = sql.Open("postgres", source)
	if err != nil {
		log.Panic(err)
	}
	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
