package auth

import (
	"time"
)

type JWTService struct{}

func NewJwtService() JWTService {
	return JWTService{}
}

func (j JWTService) ValidateAuthenticationToken(tokenString string) (string, error) {
	claims, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}

	expires := int64(claims["expires"].(float64))

	if expires < time.Now().Unix() {
		return "", &tokenExpired{}
	}

	userId := claims["userId"].(string)
	return userId, nil
}
