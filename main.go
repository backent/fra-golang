package main

import (
	"net/http"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/libs"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	helpers.PanifIfError(err)

	router := libs.NewRouter()
	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
