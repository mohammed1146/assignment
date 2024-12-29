package service

import (
	"context"
	"errors"

	"github.com/mohammed1146/skelton/internal/domain"
	"github.com/mohammed1146/skelton/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, email, password string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Register creates a new user
func (s *userService) Register(ctx context.Context, user *domain.User) error {
	existingUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}
	return s.repo.CreateUser(ctx, user)
}

// Login authenticates a user and returns a token
func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil || user.Password != password {
		return "", errors.New("invalid email or password")
	}
	return "valid-token", nil
}
