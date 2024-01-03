package controller

import (
	"context"
	db "crud/DB"
	"crud/entities"
	"fmt"
	"log"
)

func Select() {
	db := db.ConnectDB()
	defer db.Close()

	context := context.Background()

	query := "SELECT * FROM categories"

	rows, err := db.QueryContext(context, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// INI KODE UNTUK REST API NANTI
	// var categories []entities.Category

	// for rows.Next() {
	// 	var tmp entities.Category
	// 	err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.CreatedAt, &tmp.UpdatedAt)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	categories = append(categories, tmp)
	// }

	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// return categories

	fmt.Println("Categories:")
	fmt.Println("| ID | Name | Created At | Updated At |")

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("| %d | %s | %s | %s |\n", category.Id, category.Name, category.CreatedAt, category.UpdatedAt)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
