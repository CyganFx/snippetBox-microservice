package repository

import "alexedwards.net/snippetbox/pkg/domain"

type UserRepositoryInterface interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (*domain.User, error)
}
