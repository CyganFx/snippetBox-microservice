package repository

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
	"snippetBox-microservice/catalog/pkg/domain"
)

type CatalogRepository struct {
	Pool *pgxpool.Pool
}

func NewCatalogRepository(Pool *pgxpool.Pool) ICatalogRepository {
	return &CatalogRepository{Pool: Pool}
}

func (c CatalogRepository) Insert(title, category, description string, price float32) (int, error) {
	stmt := `INSERT INTO products(title, category, description, price)
			 VALUES($1, $2, $3, $4) RETURNING id`

	var id int

	err := c.Pool.QueryRow(
		context.Background(), stmt, title, category, description, price).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (c CatalogRepository) GetById(id int) (*domain.Product, error) {
	stmt := `SELECT id, title, category, description, price
			 FROM products
			 WHERE id = $1`

	product := &domain.Product{}

	err := c.Pool.QueryRow(context.Background(), stmt, id).
		Scan(&product.ID, &product.Title,
			&product.Category, &product.Description,
			&product.Price)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, domain.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return product, nil
}

func (c CatalogRepository) GetAll() ([]*domain.Product, error) {
	stmt := `SELECT id, title, category, description, price
			 FROM products ORDER BY category`

	rows, err := c.Pool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		s := &domain.Product{}
		err = rows.Scan(
			&s.ID, &s.Title, &s.Category,
			&s.Description, &s.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (c CatalogRepository) GetByCategory(category string) ([]*domain.Product, error) {
	stmt := `SELECT id, title, category, description, price
			 FROM products
			 WHERE category = $1`

	rows, err := c.Pool.Query(context.Background(), stmt, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		s := &domain.Product{}
		err = rows.Scan(
			&s.ID, &s.Title, &s.Category,
			&s.Description, &s.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
