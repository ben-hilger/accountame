package auth

type tokenExpired struct{}

func (u *tokenExpired) Error() string {
	return "token expired"
}

type unauthorizedToken struct{}

func (u *unauthorizedToken) Error() string {
	return "token unauthorized"
}

type invalidToken struct{}

func (u *invalidToken) Error() string {
	return "token unauthorized"
}
