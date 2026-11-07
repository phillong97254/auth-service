package client

import (
	"context"

	"github.com/phillong97254/auth-service/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (domain.Client, error)
}
