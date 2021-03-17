package service

import (
	"fmt"
	"github.com/CyganFx/snippetBox-microservice/news/internal/repository"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/news/utils/validator"
)

type NewsService struct {
	NewsRepository repository.NewsRepositoryInterface
}

func NewNewsService(NewsRepository repository.NewsRepositoryInterface) NewsServiceInterface {
	return &NewsService{NewsRepository: NewsRepository}
}

func (s *NewsService) Save(news *domain.News) (int, error) {
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

func (s *NewsService) FindById(id int) (*domain.News, error) {
	return s.NewsRepository.GetById(id)
}

func (s *NewsService) Latest() ([]*domain.News, error) {
	return s.NewsRepository.Latest()
}
