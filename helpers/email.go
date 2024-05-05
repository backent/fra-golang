package helpers

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
	"strings"
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

func SendMailWithoutAuth(mailInterface Mail) error {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	c, err := smtp.Dial(host + ":" + port)
	if err != nil {
		return err
	}
	defer c.Close()
	from := os.Getenv("MAIL_FROM")
	if err = c.Mail(r.Replace(from)); err != nil {
		return err
	}
	to := mailInterface.GetTo()
	c.Rcpt(to)

	w, err := c.Data()
	if err != nil {
		return err
	}

	subject := mailInterface.GetSubject()

	msg := getMessage(from, to, subject, mailInterface.GetHTML())

	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
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
