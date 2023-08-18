package main

import (
	"fmt"

	"github.com/Jasstkn/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = ""
	password = ""
)

func main() {
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err := es.ForgotPassword("test@gmail.com", "https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}

	fmt.Println("email sent")
}
