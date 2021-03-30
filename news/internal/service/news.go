package service

import (
	"fmt"
	"snippetBox-microservice/news/internal/repository"
	"snippetBox-microservice/news/pkg/domain"
	"snippetBox-microservice/news/pkg/validator"
)

type news struct {
	NewsRepository repository.NewsInterface
}

func News(NewsRepository repository.NewsInterface) NewsInterface {
	return &news{NewsRepository: NewsRepository}
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

	return s.NewsRepository.Insert(title, content, expires)
}

func (s *news) FindById(id int) (*domain.News, error) {
	return s.NewsRepository.GetById(id)
}

func (s *news) Latest() ([]*domain.News, error) {
	return s.NewsRepository.Latest()
}
