package domain

import (
	"errors"
	"github.com/gin-gonic/gin"
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

type NewsController interface {
	Home(c *gin.Context)
	ShowNews(c *gin.Context)
	CreateNews(news *News) (int, error)
}
