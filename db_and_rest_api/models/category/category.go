package category

import (
	"context"
	"db_api/db"
)

type category struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAll() ([]category, error) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM categories ORDER BY id DESC"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []category

	for rows.Next() {
		var tmp category
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.CreatedAt, &tmp.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, tmp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func GetOne(id string) ([]category, error) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM categories WHERE id = ?"

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []category

	for rows.Next() {
		var tmp category
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.CreatedAt, &tmp.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, tmp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func Store(args ...interface{}) (int64, error) {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO categories(`name`, `created_at`, `updated_at`) VALUES (?, NOW(), NOW())"

	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func Update(args ...interface{}) error {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "UPDATE categories SET name = ?, updated_at = NOW() WHERE id = ?"

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func Destroy(id string) error {
	db := db.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "DELETE FROM `categories` WHERE id = ?"

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
