package oauth

import (
	"log/slog"

	"github.com/phillong97254/auth-service/internal/service/token"
)

type impl struct {
	logger   *slog.Logger
	tokenSvc token.Service
}

func New(logger *slog.Logger, tokenSvc token.Service) Handler {
	return &impl{logger: logger.With("handler", "oauth"), tokenSvc: tokenSvc}
}
