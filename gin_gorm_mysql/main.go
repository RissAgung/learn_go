package main

import (
	"gin_gorm_mysql/controllers/productController"
	"gin_gorm_mysql/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")

	models.ConnectDB()

	api.GET("/products", productController.Index)
	api.GET("/product/:id", productController.Show)
	api.POST("/product", productController.Create)
	api.PUT("/product/:id", productController.Update)
	api.DELETE("/product", productController.Delete)

	router.Run()
}
