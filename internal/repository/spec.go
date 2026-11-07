package repository

import (
	"github.com/phil/auth-service/internal/repository/client"
	"github.com/phil/auth-service/internal/repository/user"
)

// Registry gom các repository của service. Hiện lưu in-memory, đổi sang DB thì chỉ sửa tầng này.
type Registry interface {
	User() user.Repository
	Client() client.Repository
}
