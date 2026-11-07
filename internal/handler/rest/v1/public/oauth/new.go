package oauth

import (
	"log/slog"

	"github.com/phil/auth-service/internal/service/token"
)

type impl struct {
	logger   *slog.Logger
	tokenSvc token.Service
}

func New(logger *slog.Logger, tokenSvc token.Service) Handler {
	return &impl{logger: logger.With("handler", "oauth"), tokenSvc: tokenSvc}
}
