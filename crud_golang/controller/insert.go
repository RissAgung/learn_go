package controller

import (
	"context"
	db "crud/DB"
	"log"
)

func Insert(name string) {
	db := db.ConnectDB()

	defer db.Close()

	context := context.Background()

	query := "INSERT INTO categories(`name`, `created_at`, `updated_at`) VALUES (?, NOW(), NOW())"

	_, err := db.ExecContext(context, query, name)

	if err != nil {
		panic(err)
	}

	log.Println("Success Insert Data")
}
