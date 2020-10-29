package helpers

import (
	"bp-auth-interceptor/env"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken creates a jwt with the given user
func GenerateToken(userid string) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = userid

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := at.SignedString(env.JwtSecret)

	if err != nil {
		return "", err
	}

	return token, nil

}