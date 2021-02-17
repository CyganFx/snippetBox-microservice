package models

import (
	"errors"
	"time"
)

var (
	// Add a new ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

//sql.Nullstring doesn't work

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"hashedPassword"`
	Created        time.Time `json:"created"`
	Active         bool      `json:"active"`
}

type JWT struct {
	Token string `json:"token"`
}

//type Error struct {
//	Message string `json:"message"`
//}
