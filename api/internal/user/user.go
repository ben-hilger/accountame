package user

type Account struct {
	Id             string
	Name           string
	Username       string
	HashedPassword string
	Email          string
}

func (u Account) isMissingInformation() bool {
	return u.Id == "" ||
		u.Name == "" ||
		u.Username == "" ||
		u.Email == "" ||
		u.HashedPassword == ""
}

type registerAccount struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u registerAccount) isMissingInformation() bool {
	return u.Name == "" ||
		u.Username == "" ||
		u.Email == "" ||
		u.Password == ""
}

type loginAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l loginAccountRequest) isMissingInformation() bool {
	return l.Email == "" || l.Password == ""
}

type loginAccountResponse struct {
	AuthToken string `json:"authToken"`
}
