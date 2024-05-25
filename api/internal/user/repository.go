package user

import (
	"database/sql"
	"fmt"
)

type repository interface {
	DoesUserExist(email string) (bool, error)
	RegisterUser(user Account) error
	GetUser(email string) (Account, error)
}

type PostgresClient struct {
	database *sql.DB
}

func NewPostgresClient(database *sql.DB) PostgresClient {
	return PostgresClient{database: database}
}

func (p PostgresClient) DoesUserExist(email string) (bool, error) {
	rows, err := p.database.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return true, fmt.Errorf("unable to get Account: %q: %v", email, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	return rows.Next(), nil
}

func (p PostgresClient) RegisterUser(user Account) error {
	if user.isMissingInformation() {
		return &invalidAccountError{}
	}
	_, err := p.database.Exec("INSERT INTO users (id, name, email, password_hash, username) VALUES ($1, $2, $3, $4, $5)",
		user.Id, user.Name, user.Email, user.HashedPassword, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresClient) GetUser(email string) (Account, error) {
	user := Account{}

	result, err := p.database.Query("SELECT id, name, email, username, password_hash FROM users WHERE email = $1", email)
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(result)

	if err != nil {
		return user, err
	}
	if !result.Next() {
		return user, &existsError{}
	}
	err = result.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.HashedPassword)
	return user, nil
}
