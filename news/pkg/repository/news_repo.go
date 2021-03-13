package repository

import (
	"context"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strconv"
	"time"
)

type NewsRepository struct {
	Pool *pgxpool.Pool
}

func NewNewsRepository(Pool *pgxpool.Pool) NewsRepositoryInterface {
	return &NewsRepository{Pool: Pool}
}

func (r *NewsRepository) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, $3, $4) RETURNING id`
	var id int
	integerExpires, err := strconv.Atoi(expires)
	if err != nil {
		log.Fatal(err)
	}
	//Using queryRow in order to get ID with the SCAN
	row := r.Pool.QueryRow(
		context.Background(), stmt, title, content,
		time.Now(), time.Now().AddDate(0, 0, integerExpires)).
		Scan(&id)

	if row != nil {
		log.Fatal(row)
	}

	return id, nil
}

func (r *NewsRepository) Get(id int) (*domain.News, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > now() AND id = $1`

	snippet := &domain.News{}

	row := r.Pool.QueryRow(context.Background(), stmt, id).
		Scan(&snippet.ID, &snippet.Title,
			&snippet.Content, &snippet.Created, &snippet.Expires)

	if row != nil {
		if row.Error() == "no rows in result set" {
			return nil, domain.ErrNoRecord
		} else {
			return nil, row
		}
	}

	return snippet, nil
}

func (r *NewsRepository) Latest() ([]*domain.News, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > now() ORDER BY created DESC LIMIT 10`

	rows, err := r.Pool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*domain.News

	for rows.Next() {
		s := &domain.News{}
		err = rows.Scan(
			&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
