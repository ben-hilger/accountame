package main

import (
	"fmt"
	"github.com/ben-hilger/accountame-api/internal/auth"
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

	muxHandler := http.NewServeMux()

	//muxHandler.HandleFunc("/", handleHello)

	authMiddleware := auth.NewMiddleware(auth.NewJwtService())

	userHandler := user.NewHandler(user.NewPostgresClient(db))
	muxHandler.HandleFunc("POST /user/create", userHandler.RegisterUserHandler)
	muxHandler.HandleFunc("POST /user/login", userHandler.LoginUserHandler)
	muxHandler.Handle("GET /user", authMiddleware.Protect(userHandler.RegisterUserHandler))

	muxHandler.Handle("/protected-hello", authMiddleware.Protect(handleHello))

	log.Fatal(http.ListenAndServe(":8080", muxHandler))
}

func handleHello(response http.ResponseWriter, _ *http.Request) {
	_, err := response.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
	}
}
