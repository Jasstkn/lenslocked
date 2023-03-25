package models

import "database/sql"

type User struct {
	ID           int
	Email        string // TODO: validate email
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}
