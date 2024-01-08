package auth

import (
	"fmt"
	"jwt-auth/db"
	"jwt-auth/middleware/jwt"
	"jwt-auth/migrations"
	"jwt-auth/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ResponseJson struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"` // omitempty akan membuat key data tidak di tampilkan apabila nilainya kosong
	}
)

func Register(c *gin.Context) {
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.GenerateId(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "success registered", "data": user})
}

func Login(c *gin.Context) {
	var user user.User
	var response ResponseJson

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.ValidateAccount(user.Username, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Incorrect username or password"})
		return
	}

	token, err := user.GetToken()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error env"})
		return
	}

	response.Message = "Login success"
	response.Data = gin.H{
		"username":     user.Username,
		"token_access": token,
	}

	c.JSON(http.StatusOK, response)
}

func ExtractData(c *gin.Context) {
	payloadData, err := jwt.GetPayloadIdUser(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, payloadData)
}

func Logout(c *gin.Context) {
	var blacklistToken migrations.BlackListToken
	token, exists := c.Get("token")
	if !exists {
		fmt.Printf("Unauthorized")
		return
	}

	tokenString, ok := token.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Invalid claims type"})
		return
	}
	blacklistToken.Token = tokenString

	db.DB.Create(&blacklistToken)

	c.JSON(http.StatusOK, gin.H{"message": "success logout", "data": blacklistToken})
}
