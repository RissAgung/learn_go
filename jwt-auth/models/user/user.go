package user

import (
	"errors"
	"fmt"
	"jwt-auth/db"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		Id_User   string    `gorm:"primary_key;type:varchar(11);not null" json:"id_user"`
		Username  string    `gorm:"type:varchar(20);not null" json:"username"`
		Password  string    `gorm:"type:varchar(70);not null" json:"password"`
		CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"created_at"`
		UpdatedAt time.Time `gorm:"autoUpdateTime;not null" json:"updated_at"`
	}

	JwtTokenPayload struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}

	Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"` // omitempty akan membuat key data tidak di tampilkan apabila nilainya kosong
	}
)

func (user *User) HashPassword(pass string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) ValidateAccount(username string, password string) error {
	// check username
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (user *User) GenerateId() error {

	var tmpUser User
	formatDateNow := time.Now().Format("060102") //format "YYMMDD"

	if err := db.DB.Table("users").Select("id_user").Order("created_at DESC").First(&tmpUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.Id_User = "US" + formatDateNow + "1"
			return nil
		}
		return err
	}

	regNumber, err := strconv.Atoi(tmpUser.Id_User[8:])
	if err != nil {
		return err
	}

	user.Id_User = fmt.Sprintf("US%s%d", formatDateNow, regNumber+1)
	return nil
}

func (user *User) GetToken() (Response, error) {
	var response Response
	var token JwtTokenPayload

	config, err := godotenv.Read()
	if err != nil {
		return response, err
	}

	lifeTimeToken, err := strconv.Atoi(config["LIFE_TIME_TOKEN"])
	if err != nil {
		return response, err
	}

	secretKey, errGetKey := db.GetJwtKey()
	if errGetKey != nil {
		return response, errGetKey
	}

	token.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "my-app",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(lifeTimeToken))),
	}

	token.Username = user.Username
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenResult, _ := _token.SignedString([]byte(secretKey))

	response.Message = "Success Login"
	response.Data = gin.H{
		"username":     user.Username,
		"access_token": tokenResult,
	}

	return response, nil
}
