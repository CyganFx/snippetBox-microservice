package repository

import "snippetBox-microservice/catalog/pkg/domain"

type ICatalogRepository interface {
	Insert(title, category, description string, price float32) (int, error)
	GetById(id int) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	GetByCategory(category string) ([]*domain.Product, error)
}
