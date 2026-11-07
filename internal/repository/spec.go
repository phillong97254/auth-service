package repository

import (
	"github.com/phillong97254/auth-service/internal/repository/client"
	"github.com/phillong97254/auth-service/internal/repository/user"
)

// Registry gom các repository của service. users: Postgres (gorm); clients: từ config (secret qua env).
type Registry interface {
	User() user.Repository
	Client() client.Repository
}
