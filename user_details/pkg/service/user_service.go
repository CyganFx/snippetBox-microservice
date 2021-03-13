package service

import (
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/repository"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(UserRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{UserRepository: UserRepository}
}

func (s *UserService) Save(name, email, password string) error {
	return s.UserRepository.Insert(name, email, password)
}
func (s *UserService) Authenticate(email, password string) (*domain.User, error) {
	return s.UserRepository.Authenticate(email, password)
}
