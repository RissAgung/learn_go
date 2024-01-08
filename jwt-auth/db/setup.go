package db

import (
	"errors"
	"fmt"
	"jwt-auth/migrations"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JWT_SECRET_KEY string

func GetConnection() {
	config, err := godotenv.Read()
	if err != nil {
		panic("Error reading .env file")
	}
	// Create the data source name (DSN) using the environment variables
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DATABASE_HOST"],
		config["DB_DATABASE"],
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&migrations.User{})
	db.AutoMigrate(&migrations.BlackListToken{})
	DB = db
}

func GetJwtKey() (string, error) {

	config, err := godotenv.Read()
	if err != nil {
		return "", errors.New("Error reading .env file")
	}

	secretKey, ok := config["JWT_SECRET_KEY"]
	if !ok || secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY not found")
	}

	return secretKey, nil
}
