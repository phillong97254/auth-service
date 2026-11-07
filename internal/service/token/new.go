package token

import (
	"log/slog"

	"github.com/phillong97254/auth-service/config"
	"github.com/phillong97254/auth-service/internal/gateway"
	"github.com/phillong97254/auth-service/internal/repository"
	"github.com/phillong97254/auth-service/internal/signer"
)

type impl struct {
	cfg    *config.Config
	logger *slog.Logger
	repo   repository.Registry
	gwy    gateway.Gateway
	signer signer.Signer
}

func New(cfg *config.Config, logger *slog.Logger, repo repository.Registry, gwy gateway.Gateway, s signer.Signer) Service {
	return &impl{
		cfg:    cfg,
		logger: logger.With("service", "token"),
		repo:   repo,
		gwy:    gwy,
		signer: s,
	}
}
