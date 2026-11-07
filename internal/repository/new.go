package repository

import (
	"github.com/phillong97254/auth-service/config"
	"gorm.io/gorm"
	"sync"

	"github.com/phillong97254/auth-service/internal/repository/client"
	"github.com/phillong97254/auth-service/internal/repository/user"
)

type impl struct {
	cfg *config.Config
	db  *gorm.DB

	user     user.Repository
	userOnce sync.Once

	client     client.Repository
	clientOnce sync.Once
}

func New(cfg *config.Config, db *gorm.DB) Registry {
	return &impl{cfg: cfg, db: db}
}

func (i *impl) User() user.Repository {
	i.userOnce.Do(func() { i.user = user.New(i.db) })
	return i.user
}

func (i *impl) Client() client.Repository {
	i.clientOnce.Do(func() { i.client = client.New(i.cfg.OAuth.BackofficeClientSecret) })
	return i.client
}
