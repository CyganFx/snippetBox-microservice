package models

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      int
	Title   sql.NullString
	Content sql.NullString
	Created time.Time
	Expires time.Time
}
