package service

import (
	"snippetBox-microservice/news/pkg/domain"
)

type NewsInterface interface {
	Save(news *domain.News) (int, error)
	FindById(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}
