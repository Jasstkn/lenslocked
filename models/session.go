package models

import (
	"database/sql"
	"fmt"

	"github.com/Jasstkn/lenslocked/rand"
)

const (
	// MinBytesPerToken is the minimum number of bytes to be used for each session token
	// Ref: https://owasp.org/www-community/vulnerabilities/Insufficient_Session-ID_Length
	MinBytesPerToken = 32
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. When lookup a session
	// this will be left empty, as we only store the hash of a session token
	// in our database, and we cannot reverse it into a raw token
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// session token. If the value is not set or is less than the MinBytesPerToken const
	// it will be ignored and MinBytesPerToken will be used instead.
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := MaxInt(ss.BytesPerToken, MinBytesPerToken)

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	// TODO: Hash the session token
	session := Session{
		UserID: userID,
		Token:  token,
		// Set the TokenHash
	}

	// TODO: Store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: implement SessionService.User
	return nil, nil
}

// MaxInt function return the maximum between 2 integers
// it takes 2 int and return the biggest
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
