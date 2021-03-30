package repository

import (
	"snippetBox-microservice/news/pkg/domain"
	"time"
)

type NewsInterface interface {
	Insert(title, content string, expires time.Time) (int, error)
	GetById(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}
