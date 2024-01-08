package jwt

import (
	"errors"
	"fmt"
	"jwt-auth/db"
	"jwt-auth/migrations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type (
	PayloadData struct {
		Id       string `json:"message"`
		Username string `json:"username"`
	}
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var blacklistToken migrations.BlackListToken
		// Mendapatkan nilai token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Mengekstrak token dari header Authorization
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		token := parts[1]

		if err := db.DB.Where("token = ?", token).First(&blacklistToken).Error; err == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token / udah di blacklist"})
			return
		}

		// Verifikasi token menggunakan fungsi VerifyToken
		claims, err := VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Simpan payload di dalam konteks
		c.Set("claims", claims)
		c.Set("token", token)

		// Jika token valid, lanjutkan ke handler berikutnya
		c.Next()
	}
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {

	secretKey, errGetKey := db.GetJwtKey()
	if errGetKey != nil {
		return nil, errGetKey
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Kembalikan claims (payload) dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func GetPayloadIdUser(c *gin.Context) (interface{}, error) {

	claims, exists := c.Get("claims")
	if !exists {
		return "", errors.New("Unauthorized")
	}

	// Cetak tipe data dari claims
	fmt.Printf("Type of claims: %T\n", claims)

	// Tipe asertion untuk memastikan bahwa claims adalah map[string]interface{}
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid claims type")
		return "", errors.New("Invalid claims type")
	}

	// Mengambil nilai ID dari map
	id, exists := claimsMap["id"].(string)
	if !exists {
		fmt.Println("ID not found in the map")
		return "", errors.New("ID not found in the map")
	}

	return id, nil
}
