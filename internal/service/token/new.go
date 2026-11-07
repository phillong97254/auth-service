package token

import (
	"crypto/rsa"
	"log/slog"
	"time"

	"github.com/phillong97254/auth-service/config"
	"github.com/phillong97254/auth-service/internal/repository"
)

type impl struct {
	cfg    *config.Config
	logger *slog.Logger
	repo   repository.Registry
	key    *rsa.PrivateKey
	ttl    time.Duration
}

func New(cfg *config.Config, logger *slog.Logger, repo repository.Registry, key *rsa.PrivateKey) Service {
	return &impl{
		cfg:    cfg,
		logger: logger.With("service", "token"),
		repo:   repo,
		key:    key,
		ttl:    time.Duration(cfg.OAuth.TokenTTL) * time.Second,
	}
}
