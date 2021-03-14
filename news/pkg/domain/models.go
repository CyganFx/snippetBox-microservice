package domain

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("domain: no matching record found")
)

type News struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}
