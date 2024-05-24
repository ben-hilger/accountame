package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("BEARER_SALT")
		return []byte(secret), nil
	})
	if err != nil {
		return jwt.MapClaims{}, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid - unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("claims of unauthorized type")
	}

	return claims, nil
}
