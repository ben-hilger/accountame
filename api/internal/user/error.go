package user

type existsError struct{}

func (u *existsError) Error() string {
	return "Account exists"
}

type invalidAccountError struct{}

func (u *invalidAccountError) Error() string {
	return "invalid Account"
}

type invalidPasswordError struct{}

func (u *invalidPasswordError) Error() string {
	return "invalid password"
}
