package user

import (
	"errors"
	"fmt"
	"jwt-auth/db"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		Id_User   string    `json:"id_user"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	JwtTokenPayload struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		jwt.RegisteredClaims
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

func (user *User) GetToken() (string, error) {
	var token JwtTokenPayload

	config, err := godotenv.Read()
	if err != nil {
		return "", err
	}

	lifeTimeToken, err := strconv.Atoi(config["LIFE_TIME_TOKEN"])
	if err != nil {
		return "", err
	}

	secretKey, errGetKey := db.GetJwtKey()
	if errGetKey != nil {
		return "", errGetKey
	}

	token.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "my-app",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(lifeTimeToken))),
	}

	token.Username = user.Username
	token.Id = user.Id_User
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenResult, _ := _token.SignedString([]byte(secretKey))

	return tokenResult, nil
}
