package repository

import "github.com/CyganFx/snippetBox-microservice/news/pkg/domain"

type NewsRepositoryInterface interface {
	Insert(title, content, expires string) (int, error)
	Get(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}
