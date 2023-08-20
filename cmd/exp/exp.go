package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Jasstkn/lenslocked/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err = es.ForgotPassword("test@gmail.com", "https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}

	fmt.Println("email sent")
}
