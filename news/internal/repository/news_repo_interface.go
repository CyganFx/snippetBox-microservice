package repository

import (
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"time"
)

type NewsRepositoryInterface interface {
	Insert(title, content string, expires time.Time) (int, error)
	GetById(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}
