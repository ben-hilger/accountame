package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(user RegisterUser) error
}

type ServiceImpl struct {
	userRepository Repository
}

func NewService(userRepository Repository) ServiceImpl {
	return ServiceImpl{
		userRepository: userRepository,
	}
}

func (s ServiceImpl) RegisterUser(user RegisterUser) error {
	if user.isMissingInformation() {
		return &InvalidUserError{}
	}

	userExists, err := s.userRepository.DoesUserExist(user.Email)

	if err != nil {
		return err
	}

	if userExists {
		return &UserExistsError{}
	}

	newUser, err := newUser(user)
	if err != nil {
		return err
	}

	err = s.userRepository.RegisterUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func newUser(user RegisterUser) (User, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	return User{
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
