package service

import (
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/repository"
)

type NewsService struct {
	NewsRepository repository.NewsRepositoryInterface
}

func NewNewsService(NewsRepository repository.NewsRepositoryInterface) NewsServiceInterface {
	return &NewsService{NewsRepository: NewsRepository}
}

func (s *NewsService) Save(title, content, expires string) (int, error) {
	return s.NewsRepository.Insert(title, content, expires)
}

func (s *NewsService) FindById(id int) (*domain.News, error) {
	return s.NewsRepository.GetById(id)
}

func (s *NewsService) Latest() ([]*domain.News, error) {
	return s.NewsRepository.Latest()
}
