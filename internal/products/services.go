package products

import (
	"context"

	repo "github.com/SaranHiruthikM/ecomm-go/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) error {
	products, err := s.repo.ListProducts(ctx)
}
