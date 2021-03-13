package domain

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("domain: no matching record found")
	ErrInvalidCredentials = errors.New("domain: invalid credentials")
	ErrDuplicateEmail     = errors.New("domain: duplicate email")
)

type User struct {
	ID             int       `json:"user_id"`
	Name           string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"hashedPassword"`
	Created        time.Time `json:"created"`
	Active         bool      `json:"active"`
	Role           string    `json:"role"`
}
