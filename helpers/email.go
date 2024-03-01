package helpers

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
)

func SendMail(mailInterface Mail) error {

	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")

	auth := smtp.PlainAuth("", username, password, host)

	to := mailInterface.GetTo()
	from := os.Getenv("MAIL_FROM")
	subject := mailInterface.GetSubject()
	msg := getMessage(from, to, subject, mailInterface.GetHTML())

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, msg)

	return err
}

func getMessage(from string, to string, subject string, bodyEmail string) []byte {
	return []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +

			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" +

			"\r\n" +

			bodyEmail + "\r\n")
}

type Mail interface {
	GetHTML() string
	GetTo() string
	GetSubject() string
}

type RecipientRegistration struct {
	Name    string
	Email   string
	Status  string
	Subject string
}

func (implementation RecipientRegistration) GetHTML() string {

	content, err := os.ReadFile("assets/template/template_email_registration_approval.html")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Parse(string(content))
	if err != nil {
		panic(err)
	}

	var resultTemplate bytes.Buffer

	tmpl.Execute(&resultTemplate, implementation)

	return resultTemplate.String()
}

func (implementation RecipientRegistration) GetTo() string {
	return implementation.Email
}

func (implementation RecipientRegistration) GetSubject() string {
	return implementation.Subject
}
