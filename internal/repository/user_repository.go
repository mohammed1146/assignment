package repository

import (
	"context"
	"database/sql"

	"github.com/mohammed1146/skelton/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser adds a new user to the database
func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	return err
}

// GetUserByEmail retrieves a user by their email address
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = ?`
	row := r.db.QueryRowContext(ctx, query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
