package user

import (
	"context"

	"github.com/phil/auth-service/internal/domain"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (domain.User, error)
}
