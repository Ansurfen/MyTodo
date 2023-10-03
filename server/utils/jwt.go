package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_todo_key")

func ReleaseToken(id uint) (string, error) {
	now := time.Now()
	expirationTime := now.Add(7 * 24 * time.Hour)
	claims := jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", id),
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    "org.my_todo",
		Subject:   "user token",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenString string) (*jwt.Token, jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
