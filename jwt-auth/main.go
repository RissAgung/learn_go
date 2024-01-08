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
	guest_access := router.Group("/api")
	auth_access := router.Group("/api").Use(jwt.AuthMiddleware())

	db.GetConnection()

	guest_access.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "connected api"})
	})
	guest_access.POST("/login", auth.Login)

	auth_access.POST("/register", auth.Register)
	auth_access.POST("/extractToken", auth.ExtractData)
	auth_access.POST("/logout", auth.Logout)
	router.Run()
}
