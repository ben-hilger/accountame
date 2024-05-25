package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type accountService interface {
	RegisterUser(user registerAccount) error
	GetUser(id string) (Account, error)
	LoginUser(loginRequest loginAccountRequest) (string, error)
	CreateAuthenticationToken(accountId string) (string, error)
}

type serviceImpl struct {
	accountRepository repository
}

func newService(userRepository repository) serviceImpl {
	return serviceImpl{
		accountRepository: userRepository,
	}
}

func (s serviceImpl) CreateAuthenticationToken(userId string) (string, error) {
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

func (s serviceImpl) LoginUser(loginRequest loginAccountRequest) (string, error) {
	if loginRequest.isMissingInformation() {
		return "", &invalidAccountError{}
	}

	userExists, err := s.accountRepository.DoesUserExist(loginRequest.Email)
	if err != nil {
		return "", err
	}
	if !userExists {
		return "", &existsError{}
	}
	account, err := s.accountRepository.GetUser(loginRequest.Email)
	if err != nil {
		return "", err
	}

	err = validatePassword(account.HashedPassword, loginRequest.Password)
	if err != nil {
		return "", &invalidPasswordError{}
	}

	return account.Id, nil
}

func (s serviceImpl) GetUser(id string) (Account, error) {
	return Account{}, nil
}

func (s serviceImpl) RegisterUser(user registerAccount) error {
	if user.isMissingInformation() {
		return &invalidAccountError{}
	}

	userExists, err := s.accountRepository.DoesUserExist(user.Email)

	if err != nil {
		return err
	}

	if userExists {
		return &existsError{}
	}

	newUser, err := newUser(user)
	if err != nil {
		return err
	}

	err = s.accountRepository.RegisterUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func newUser(user registerAccount) (Account, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return Account{}, err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return Account{}, err
	}
	return Account{
		Id:             id.String(),
		Username:       user.Username,
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: hashedPassword,
	}, nil

}

func hashPassword(plainTextPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func validatePassword(hashedPassword, plainTextPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
}
