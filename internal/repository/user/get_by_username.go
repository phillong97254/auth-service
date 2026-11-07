package user

import (
	"context"

	"github.com/phil/auth-service/internal/domain"
)

func (i *impl) GetByUsername(_ context.Context, username string) (domain.User, error) {
	u, ok := i.users[username]
	if !ok {
		return domain.User{}, ErrNotFound
	}
	return u, nil
}
