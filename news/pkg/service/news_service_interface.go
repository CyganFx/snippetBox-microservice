package service

import (
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"time"
)

type NewsServiceInterface interface {
	Save(title, content string, expires time.Time) (int, error)
	FindById(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}
