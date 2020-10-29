package helpers

import (
	"fmt"
	"errors"
	"bp-auth-interceptor/env"
	"github.com/dgrijalva/jwt-go"
)

// VerifyToken verifica la validità di un jwt
func  VerifyToken(token string) (bool, error) {
	_, err := jwt.Parse(token, func (parsed *jwt.Token) (interface{}, error) {
		_, ok := parsed.Method.(*jwt.SigningMethodHMAC);

		if !ok {
			return nil, errors.Unwrap(fmt.Errorf("Verifica jwt fallita"))
		}

		return env.JwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}