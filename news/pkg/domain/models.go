package domain

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("domain: no matching record found")
	ErrInvalidCredentials = errors.New("domain: invalid credentials")
	ErrDuplicateEmail     = errors.New("domain: duplicate email")
)

type News struct {
	ID      int
	Title   sql.NullString
	Content sql.NullString
	Created time.Time
	Expires time.Time
}
