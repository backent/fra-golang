package main

import (
	"log"
	"net"
	"os"

	"github.com/backent/fra-golang/helpers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	helpers.PanicIfError(err)

	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")

	log.Println("Try to dial ", host+":"+port)

	_, err = net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Println("dial error :", err)
	} else {
		log.Println("dial success")
	}

}
