package db

import "log"

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=yourdatabase sslmode=disable password=yourpassword host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
}
