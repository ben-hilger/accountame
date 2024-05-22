package main

import (
	"fmt"
	"github.com/ben-hilger/accountame-api/env"
	"log"
	"net/http"
)

func loadEnvironment(env env.Environment) {
	err := env.LoadEnvironment()
	if err != nil {
		panic(fmt.Sprintf("unable to load environment"))
	}
}

func main() {
	fmt.Println("Starting... agains")
	loadEnvironment(env.EnvironmentFile{
		File: ".env",
	})

	http.HandleFunc("/", handleHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHello(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
	}
}
