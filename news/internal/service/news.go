package service

import (
	"fmt"
	"snippetBox-microservice/news/api/controller"
	"snippetBox-microservice/news/pkg/domain"
	"snippetBox-microservice/news/pkg/validator"
	"time"
)

type news struct {
	repo NewsRepositoryInterface
}

type NewsRepositoryInterface interface {
	Insert(title, content string, expires time.Time) (int, error)
	GetById(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}

func News(NewsRepository NewsRepositoryInterface) controller.NewsServiceInterface {
	return &news{repo: NewsRepository}
}

func (s *news) Save(news *domain.News) (int, error) {
	title := news.Title
	content := news.Content
	expires := news.Expires

	v := validator.New()
	v.MaxLength(title, 100)

	if !v.Valid() {
		return -1, fmt.Errorf("news validation error")
	}

	return s.repo.Insert(title, content, expires)
}

func (s *news) FindById(id int) (*domain.News, error) {
	return s.repo.GetById(id)
}

func (s *news) Latest() ([]*domain.News, error) {
	return s.repo.Latest()
}
