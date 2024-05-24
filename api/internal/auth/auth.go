package auth

type Authentication interface {
	CreateAuthenticationToken(userId string) (string, error)
	ValidatePassword(hashedPassword, plainTextPassword string) error
	ValidateAuthenticationToken(tokenString string) error
}
