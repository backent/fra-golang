package main

import (
	"net/http"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/libs"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.PanifIfError(err)

	router := libs.NewRouter()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
