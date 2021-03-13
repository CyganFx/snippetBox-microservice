package repository

import "github.com/CyganFx/snippetBox-microservice/user_details/pkg/domain"

type UserRepositoryInterface interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (*domain.User, error)
}
