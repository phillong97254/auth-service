package repository

import (
	"github.com/phillong97254/auth-service/config"
	"sync"

	"github.com/phillong97254/auth-service/internal/repository/client"
	"github.com/phillong97254/auth-service/internal/repository/user"
)

type impl struct {
	cfg *config.Config

	user     user.Repository
	userOnce sync.Once

	client     client.Repository
	clientOnce sync.Once
}

func New(cfg *config.Config) Registry {
	return &impl{cfg: cfg}
}

func (i *impl) User() user.Repository {
	i.userOnce.Do(func() { i.user = user.New() })
	return i.user
}

func (i *impl) Client() client.Repository {
	i.clientOnce.Do(func() { i.client = client.New(i.cfg.OAuth.BackofficeClientSecret) })
	return i.client
}
