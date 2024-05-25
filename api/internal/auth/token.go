package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("BEARER_SALT")
		return []byte(secret), nil
	})
	if err != nil {
		return jwt.MapClaims{}, &unauthorizedToken{}
	}

	if !token.Valid {
		return nil, &invalidToken{}
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, &unauthorizedToken{}
	}

	return claims, nil
}
