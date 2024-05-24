package main

import (
	"fmt"
	"github.com/ben-hilger/accountame-api/internal/database"
	"github.com/ben-hilger/accountame-api/internal/user"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}

	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/", handleHello)

	userHandler := user.NewHandler(user.NewPostgresClient(db))
	userHandler.RegisterUserRouteHandlers(muxServer)

	log.Fatal(http.ListenAndServe(":8080", muxServer))
}

func handleHello(response http.ResponseWriter, _ *http.Request) {
	_, err := response.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
	}
}
