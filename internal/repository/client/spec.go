package client

import (
	"context"

	"github.com/phil/auth-service/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domain.Client, error)
}
