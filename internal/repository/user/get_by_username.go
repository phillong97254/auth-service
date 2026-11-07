package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/phillong97254/auth-service/internal/domain"
)

func (i *impl) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var u domain.User
	err := i.db.WithContext(ctx).Where("username = ?", username).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, ErrNotFound
	}
	return u, err
}
