package common

import (
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int, config configuration.Config) string {
	secret := config.Get("JWT_SECRET")

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * time.Duration(2)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	exception.PanicLogging(err)

	return signed
}
