package service

import (
	"context"

	"github.com/mohammed1146/skelton/internal/domain"
	"github.com/mohammed1146/skelton/internal/repository"
)

type SpacecraftService interface {
	ListSpacecrafts(ctx context.Context, name string, class string, status string) ([]domain.Spacecraft, error)
}

type spacecraftService struct {
	repo repository.SpacecraftRepository
}

func NewSpacecraftService(repo repository.SpacecraftRepository) SpacecraftService {
	return &spacecraftService{repo: repo}
}

// ListSpacecrafts get all spacecrafts filtered by name, class and status.
func (s *spacecraftService) ListSpacecrafts(ctx context.Context, name string, class string, status string) ([]domain.Spacecraft, error) {
	return s.repo.ListSpacecrafts(ctx, name, class, status)
}
