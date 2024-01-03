package controller

import (
	"context"
	db "crud/DB"
	"log"
)

func Update(id int, newName string) {
	db := db.ConnectDB()

	defer db.Close()

	context := context.Background()

	query := "UPDATE categories SET name = ?, updated_at = NOW() WHERE id = ?"

	_, err := db.ExecContext(context, query, newName, id)

	if err != nil {
		panic(err)
	}

	log.Println("Success Update Data")
}
