package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Token is only set when PasswordReset is being created.
	// When lookup a PasswordReset this will be left empty.
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// password reset token. If the value is not set or is less than the MinBytesPerToken const
	// it will be ignored and MinBytesPerToken will be used instead.
	BytesPerToken int
	// Duration is the amount of time that a PasswordReset is valid for.
	// Defaults to DefaultResetDuration
	Duration time.Duration
}

func (pwr *PasswordResetService) Create(email string) (*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordService.Create")
}

func (pwr *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Consume")
}