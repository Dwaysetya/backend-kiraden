package helpers

import (
	"backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET","secret_key"))

func GenerateToken(name string) string{
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject: name,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	return token
}