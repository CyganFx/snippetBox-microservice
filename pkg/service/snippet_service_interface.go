package service

import "alexedwards.net/snippetbox/pkg/domain"

type SnippetServiceInterface interface {
	Save(title, content, expires string) (int, error)
	Find(id int) (*domain.Snippet, error)
	Latest() ([]*domain.Snippet, error)
}
