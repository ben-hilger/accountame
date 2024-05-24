package user

import "github.com/google/uuid"

type User struct {
	Id             uuid.UUID
	Name           string
	Username       string
	HashedPassword string
	Email          string
}

func (u User) isMissingInformation() bool {
	return u.Id.String() == "" ||
		u.Name == "" ||
		u.Username == "" ||
		u.Email == "" ||
		u.HashedPassword == ""
}
