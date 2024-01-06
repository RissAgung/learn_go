package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@(localhost:3306)/rest_api_gin"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}
