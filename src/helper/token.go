package helper

import (
	"time"

	"github.com/dreamnajababy/go-ecom/src/models"
	"github.com/golang-jwt/jwt"
)

func GetUserToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	setUserToken(token, user)
	cipherToken, err := encryptToken(token)
	return cipherToken, err
}

func encryptToken(token *jwt.Token) (string, error) {
	t, err := token.SignedString([]byte("secret"))
	return t, err
}

func setUserToken(token *jwt.Token, user models.User) {
	// set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
}
