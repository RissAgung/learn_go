package main

import (
	"jwt-auth/controllers/auth"
	"jwt-auth/db"
	"jwt-auth/middleware/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	guest := api.Use(jwt.AuthMiddleware())

	db.GetConnection()

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "connected api"})
	})

	guest.POST("/register", auth.Register)
	api.POST("/login", auth.Login)

	// api.POST("/verifyToken", jwt.Verify)

	router.Run()
}
