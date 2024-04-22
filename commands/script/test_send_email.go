package main

import (
	"flag"
	"log"

	"github.com/backent/fra-golang/helpers"
	"github.com/joho/godotenv"
)

type testRecipient struct {
	Email   string
	Subject string
}

func (implementation testRecipient) GetHTML() string {

	return `
	<div>
  <p>Dear Recipient, </p>
  <p>This is a test email intended solely for the purpose of evaluation. Kindly disregard its contents, as it holds no significance or actionable information.</p>
  <p>Thank you for your understanding and cooperation.</p>
  <p>Best regards</p>
</div>
	`
}

func (implementation testRecipient) GetTo() string {
	return implementation.Email
}

func (implementation testRecipient) GetSubject() string {
	return implementation.Subject
}

func main() {
	err := godotenv.Load(".env")
	helpers.PanicIfError(err)

	var emailTo string

	flag.StringVar(&emailTo, "emailTo", "", "email address to send test email")
	flag.Parse()
	if emailTo == "" {
		log.Fatal("please provide arg emailTo")
	}

	recipient := testRecipient{
		Email:   emailTo,
		Subject: "Testing Email",
	}

	log.Println("Try to sending email, please wait ...")
	err = helpers.SendMail(recipient)
	if err != nil {
		log.Println("Error while sending email to :", recipient, err)
	} else {
		log.Println("success send email to :", recipient)
	}

}
