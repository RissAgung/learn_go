package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@(localhost:3306)/go_products")

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)                  // min user
	db.SetMaxOpenConns(100)                 // max user
	db.SetConnMaxLifetime(60 * time.Minute) // max life time user

	log.Println("Database Connected")

	return db
}
