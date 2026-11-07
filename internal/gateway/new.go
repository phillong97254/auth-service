package gateway

import (
	"fmt"

	"github.com/phillong97254/auth-service/config"
	"github.com/phillong97254/auth-service/internal/gateway/authz"
	"github.com/phillong97254/auth-service/internal/signer"
)

type impl struct {
	authz authz.Gateway
}

func New(cfg *config.Config, s signer.Signer) (Gateway, error) {
	authzGwy, err := authz.New(cfg, s)
	if err != nil {
		return nil, fmt.Errorf("init authz gateway: %w", err)
	}
	return &impl{authz: authzGwy}, nil
}

func (i *impl) Authz() authz.Gateway { return i.authz }
func (i *impl) Close() error         { return i.authz.Close() }
