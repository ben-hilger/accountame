package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

func SetupDatabase() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		return nil, err
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	return sql.Open("postgres", psqlInfo)
}
