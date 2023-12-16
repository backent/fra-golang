package main

import (
	"net/http"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/injector"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfError(err)

	router := injector.InitializeRouter()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
