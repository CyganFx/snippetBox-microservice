package service

import "github.com/CyganFx/snippetBox-microservice/user_details/pkg/domain"

type UserServiceInterface interface {
	Save(name, email, password string) error
	Authenticate(email, password string) (*domain.User, error)
}
