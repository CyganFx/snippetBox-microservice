package repository

import (
	"context"
	"errors"
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/domain"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(Pool *pgxpool.Pool) UserRepositoryInterface {
	return &UserRepository{Pool: Pool}
}

func (r *UserRepository) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES($1, $2, $3, $4)`

	_, err = r.Pool.Exec(context.Background(), stmt, name, email, string(hashedPassword), time.Now())
	if err != nil {
		postgresError := err.(*pgconn.PgError)
		if errors.As(err, &postgresError) {
			if postgresError.Code == "23505" &&
				strings.Contains(postgresError.Message, "users_uc_email") {
				return domain.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (r *UserRepository) Authenticate(email, password string) (*domain.User, error) {
	var id int
	var username string
	var hashedPassword []byte
	stmt := "SELECT id, hashed_password, name FROM users WHERE email = $1 AND active = TRUE"
	row := r.Pool.QueryRow(context.Background(), stmt, email)
	err := row.Scan(&id, &hashedPassword, &username)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, domain.ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, domain.ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	user := &domain.User{
		ID:   id,
		Name: username,
	}

	return user, nil
}
