package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/Jasstkn/lenslocked/rand"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID int
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

func (s *PasswordResetService) Create(email string) (*PasswordReset, error) {
	// verify that we have a valid email address for a user.
	email = strings.ToLower(email)

	// build the PasswordReset
	bytesPerToken := s.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	duration := s.Duration
	if duration == 0 {
		duration = DefaultResetDuration
	}

	pwReset := PasswordReset{
		Token:     token,
		TokenHash: s.hash(token),
		ExpiresAt: time.Now().Add(duration),
	}

	// Insert the PasswordReset into the DB
	row := s.DB.QueryRow(`
		INSERT INTO password_resets (user_id, token_hash, expires_at)
		SELECT id, $2, $3 FROM users WHERE email = $1
		ON CONFLICT (user_id) 
		DO UPDATE SET token_hash = $2, expires_at = $3
		RETURNING id;`, email, pwReset.TokenHash, pwReset.ExpiresAt)
	err = row.Scan(&pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &pwReset, nil
}

func (ss *PasswordResetService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func (s *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Consume")
}
