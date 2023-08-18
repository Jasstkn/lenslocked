package models

import "github.com/go-mail/mail/v2"

const (
	DefaultSender = "support@lenslocked.com"
)

type EmailService struct {
	// DefaultSender is used as the sender for all emails
	// when one is not provided for an email.
	// It is also used in functions where the email is a predefined
	// like the forgot password email.
	DefaultSender string

	dialer *mail.Dialer
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// NewEmailService creates a new email service using the provided
// smtp information.
func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}

	return &es
}
