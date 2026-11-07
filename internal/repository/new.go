package repository

import (
	"sync"

	"github.com/phil/auth-service/internal/repository/client"
	"github.com/phil/auth-service/internal/repository/user"
)

type impl struct {
	user     user.Repository
	userOnce sync.Once

	client     client.Repository
	clientOnce sync.Once
}

func New() Registry {
	return &impl{}
}

func (i *impl) User() user.Repository {
	i.userOnce.Do(func() { i.user = user.New() })
	return i.user
}

func (i *impl) Client() client.Repository {
	i.clientOnce.Do(func() { i.client = client.New() })
	return i.client
}
