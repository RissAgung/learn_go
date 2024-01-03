package controller

import (
	"context"
	db "crud/DB"
	"log"
)

func Delete(id int) {
	db := db.ConnectDB()

	defer db.Close()

	context := context.Background()

	query := "DELETE FROM categories WHERE id = ?"

	_, err := db.ExecContext(context, query, id)

	if err != nil {
		panic(err)
	}

	log.Println("Success Delete Data")
}
