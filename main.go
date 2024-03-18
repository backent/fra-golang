package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/injector"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfError(err)

	TIMEZONE := os.Getenv("TIMEZONE")

	os.Setenv("TZ", TIMEZONE)
	// Load the timezone
	time.LoadLocation(TIMEZONE)

	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "8080"
	}

	router := injector.InitializeRouter()

	server := http.Server{
		Addr:    ":" + APP_PORT,
		Handler: router,
	}

	fmt.Println("serving on :" + APP_PORT)

	server.ListenAndServe()
}
