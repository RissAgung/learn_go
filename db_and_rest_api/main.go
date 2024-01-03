package main

import (
	"db_api/controller/category"
	"db_api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.GetConnection()

	router := gin.Default()
	router.GET("/categories", category.GetAll)
	router.GET("/categories/:id", category.GetOne)
	router.POST("/categories", category.Store)
	router.PUT("/categories/:id", category.Update)
	router.DELETE("/categories/:id", category.Destroy)
	router.Run("localhost:8080")
}
