package client

import (
	"context"

	"github.com/phil/auth-service/internal/domain"
)

func (i *impl) GetByID(_ context.Context, id string) (domain.Client, error) {
	c, ok := i.clients[id]
	if !ok {
		return domain.Client{}, ErrNotFound
	}
	return c, nil
}
