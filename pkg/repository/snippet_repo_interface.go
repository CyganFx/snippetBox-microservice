package repository

import "alexedwards.net/snippetbox/pkg/domain"

type SnippetRepositoryInterface interface {
	Insert(title, content, expires string) (int, error)
	Get(id int) (*domain.Snippet, error)
	Latest() ([]*domain.Snippet, error)
}
