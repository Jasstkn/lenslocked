package models

type User struct {
	ID           int
	Email        string // TODO: validate email
	PasswordHash string
}
