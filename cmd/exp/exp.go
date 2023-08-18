package main

import (
	"github.com/Jasstkn/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = ""
	password = ""
)

func main() {
	email := models.Email{
		From:    "test@lenslocked.com",
		To:      "test@gmail.com",
		Subject: "This is a test email",
		// PlainText: "This is the body of the email",
		// HTML: `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`,
	}

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err := es.Send(email)
	if err != nil {
		panic(err)
	}

}
