package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtKey = []byte(viper.GetString("jwtSecret"))

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}
