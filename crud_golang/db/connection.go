package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/go_products")

	if err != nil {
		panic(err) // return error
	}

	db.SetMaxIdleConns(10)  // minimal koneksi dibuat
	db.SetMaxOpenConns(100) // maximal koneksi dibuat
	// db.SetConnMaxIdleTime(time.Minute * 5) // saat user tidak melakukan koneksi ke db selama 5 menit maka akan di close/perbarui
	db.SetConnMaxLifetime(time.Minute * 60) // max user melakukan koneksi ke db

	log.Println("Database connected")

	return db
}
