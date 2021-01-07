package psql

import (
	"alexedwards.net/snippetbox/pkg/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, now(), now() + INTERVAL '7' DAY) RETURNING id`

	var id int

	//Using queryRow in order to get ID with the SCAN
	row := m.DB.QueryRow(
		context.Background(), stmt, title, content).
		Scan(&id)

	if row != nil {
		log.Fatal(row)

	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > now() AND id = $1`

	snippet := &models.Snippet{}

	row := m.DB.QueryRow(context.Background(), stmt, id).
		Scan(&snippet.ID, &snippet.Title,
			&snippet.Content, &snippet.Created, &snippet.Expires)

	if row != nil {
		if errors.Is(row, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, row
		}
	}

	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > now() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*models.Snippet

	for rows.Next() {
		s := &models.Snippet{}
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
