package user

import (
	"database/sql"
	"fmt"
)

type Store interface {
	DoesUserExist(email string) (bool, error)
	InsertUser(user User) error
	GetUser(email string) (User, error)
}

type PostgresClient struct {
	database *sql.DB
}

func (p PostgresClient) DoesUserExist(email string) (bool, error) {
	rows, err := p.database.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return true, fmt.Errorf("unable to get user: %q: %v", email, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	return rows.Next(), nil
}

func (p PostgresClient) InsertUser(user User) error {
	if user.isMissingInformation() {
		return fmt.Errorf("missing information to create the user")
	}
	_, err := p.database.Exec("INSERT INTO users SET id = ?, name = ?, email = ?, username = ?",
		user.Id, user.Name, user.Email, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresClient) GetUser(email string) (User, error) {
	var user User

	result, err := p.database.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return user, err
	}
	if result.Next() {
		err = result.Scan(&user)
	}
	return user, nil
}
