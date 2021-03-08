package service

import (
	"alexedwards.net/snippetbox/pkg/domain"
	"alexedwards.net/snippetbox/pkg/repository"
)

type SnippetService struct {
	SnippetRepository repository.SnippetRepositoryInterface
}

func NewSnippetService(SnippetRepository repository.SnippetRepositoryInterface) SnippetServiceInterface {
	return &SnippetService{SnippetRepository: SnippetRepository}
}

func (s *SnippetService) Save(title, content, expires string) (int, error) {
	return s.SnippetRepository.Insert(title, content, expires)
}

func (s *SnippetService) Find(id int) (*domain.Snippet, error) {
	return s.SnippetRepository.Get(id)
}

func (s *SnippetService) Latest() ([]*domain.Snippet, error) {
	return s.SnippetRepository.Latest()
}
