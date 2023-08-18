package models

import (
	"errors"
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	DefaultSender = "support@lenslocked.com"
)

var (
	errBodyConfigEmpty = errors.New("both PlainText and HTML are empty")
)

type Email struct {
	From      string
	To        string
	Subject   string
	PlainText string
	HTML      string
}

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
		DefaultSender: DefaultSender,
		dialer:        mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}

	return &es
}

// SendEmail sends the provided email to the provided recipient.
func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()

	es.setFromHeader(msg, email)

	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)

	switch {
	case email.PlainText != "" && email.HTML != "":
		msg.SetBody("text/plain", email.PlainText)
		msg.AddAlternative("text/html", email.HTML)
	case email.PlainText != "":
		msg.SetBody("text/plain", email.PlainText)
	case email.HTML != "":
		msg.SetBody("text/html", email.HTML)
	default:
		return fmt.Errorf("failed to set body, %w", errBodyConfigEmpty)
	}

	err := es.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// setFrom sets the from header based on the email's from field.
// If the email's from field is empty, it will use the default sender from the email service.
// If the email service's default sender is empty, it will use the DefaultSender constant.
func (es *EmailService) setFromHeader(msg *mail.Message, email Email) {
	var from string

	switch {
	case email.From != "":
		from = email.From
	case es.DefaultSender != "":
		from = es.DefaultSender
	default:
		from = DefaultSender
	}

	msg.SetHeader("From", from)
}
