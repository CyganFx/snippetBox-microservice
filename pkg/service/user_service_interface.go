package service

import "alexedwards.net/snippetbox/pkg/domain"

type UserServiceInterface interface {
	Save(name, email, password string) error
	Authenticate(email, password string) (*domain.User, error)
}
