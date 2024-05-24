package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type JWTAuthentication struct{}

func NewJwtAuthentication() JWTAuthentication {
	return JWTAuthentication{}
}

func (j JWTAuthentication) ValidateAuthenticationToken(tokenString string) error {
	claims, err := parseToken(tokenString)
	if err != nil {
		return err
	}

	expires := int64(claims["expires"].(float64))

	if expires < time.Now().Unix() {
		return fmt.Errorf("authentication token expired")
	}

	return nil
}

func (j JWTAuthentication) ValidatePassword(hashedPassword, plainTextPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err
}

func (j JWTAuthentication) CreateAuthenticationToken(userId string) (string, error) {
	now := time.Now()
	validUtil := now.Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user":    userId,
		"expires": validUtil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims, nil)
	secret := os.Getenv("BEARER_SALT")

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
