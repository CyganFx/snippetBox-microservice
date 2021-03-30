package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"snippetBox-microservice/news/pkg/domain"
	"time"
)

type news struct {
	Pool *pgxpool.Pool
}

func News(Pool *pgxpool.Pool) NewsInterface {
	return &news{Pool: Pool}
}

func (r *news) Insert(title, content string, expires time.Time) (int, error) {
	stmt := `INSERT INTO news (title, content, created, expires)
	VALUES($1, $2, $3, $4) RETURNING id`
	var id int
	//Using queryRow in order to get ID with the SCAN
	err := r.Pool.QueryRow(
		context.Background(), stmt, title, content, time.Now(), expires).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *news) GetById(id int) (*domain.News, error) {
	stmt := `SELECT id, title, content, created, expires FROM news
	WHERE expires > now() AND id = $1`

	news := &domain.News{}

	err := r.Pool.QueryRow(context.Background(), stmt, id).
		Scan(&news.ID, &news.Title,
			&news.Content, &news.Created, &news.Expires)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, domain.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return news, nil
}

func (r *news) Latest() ([]*domain.News, error) {
	stmt := `SELECT id, title, content, created, expires FROM news
	WHERE expires > now() ORDER BY created DESC LIMIT 10`

	rows, err := r.Pool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []*domain.News

	for rows.Next() {
		s := &domain.News{}
		err = rows.Scan(
			&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		news = append(news, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}
