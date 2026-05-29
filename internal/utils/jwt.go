package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secretkey = []byte("mysecretkey")

func GenerateJWT(userID int, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"roel":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(Secretkey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
