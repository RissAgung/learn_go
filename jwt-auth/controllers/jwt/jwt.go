package jwt

import (
	"jwt-auth/middleware/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

func Verify(c *gin.Context) {

	var tokenRequest TokenRequest

	if err := c.BindJSON(&tokenRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := jwt.VerifyToken(tokenRequest.Token); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bisa Cuy"})
}
