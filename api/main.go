package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	_, err := setupDatabase()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handleHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupDatabase() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		return nil, err
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	return sql.Open("postgres", psqlInfo)
}

func handleHello(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
	}
}
