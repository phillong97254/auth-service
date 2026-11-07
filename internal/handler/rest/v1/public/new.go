package public

import (
	"log/slog"

	"github.com/phil/auth-service/internal/handler/rest/v1/public/oauth"
	"github.com/phil/auth-service/internal/service/token"
)

type impl struct {
	oauth oauth.Handler
}

func New(logger *slog.Logger, tokenSvc token.Service) Handler {
	return &impl{oauth: oauth.New(logger, tokenSvc)}
}

func (i *impl) GetOAuthHandler() oauth.Handler { return i.oauth }
