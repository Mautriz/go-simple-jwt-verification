package helpers

import (
	"fmt"
	"bp-auth-interceptor/env"
	"github.com/dgrijalva/jwt-go"
)

// VerifyToken verifica la validità di un jwt
func  VerifyToken(token string) (bool, error) {
	_, err := jwt.Parse(token, func (parsed *jwt.Token) (interface{}, error) {
		_, ok := parsed.Method.(*jwt.SigningMethodHMAC);

		if !ok {
			return nil, fmt.Errorf("Signing method non accettato")
		}

		mapClaims := parsed.Claims.(jwt.MapClaims)
		exp := int64(mapClaims["exp"].(float64))
		isExpired := CurrTimestampSecond() > exp

		if isExpired {
			return nil, fmt.Errorf("Il token è expired")
		}

		return env.JwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}