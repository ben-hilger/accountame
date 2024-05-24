package user

type UserExistsError struct{}

func (u *UserExistsError) Error() string {
	return "user exists"
}

type InvalidUserError struct{}

func (u *InvalidUserError) Error() string {
	return "invalid user"
}
