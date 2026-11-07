package public

import (
	"log/slog"

	"github.com/phillong97254/auth-service/internal/handler/rest/v1/public/oauth"
	"github.com/phillong97254/auth-service/internal/service/token"
)

type impl struct {
	oauth oauth.Handler
}

func New(logger *slog.Logger, tokenSvc token.Service) Handler {
	return &impl{oauth: oauth.New(logger, tokenSvc)}
}

func (i *impl) GetOAuthHandler() oauth.Handler { return i.oauth }
