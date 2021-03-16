package domain

import "errors"

var (
	ErrNoRecord = errors.New("domain: no matching record found")
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
