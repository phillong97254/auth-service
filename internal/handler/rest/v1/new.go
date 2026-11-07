package v1

import (
	"log/slog"

	"github.com/phillong97254/auth-service/internal/handler/rest/v1/public"
	"github.com/phillong97254/auth-service/internal/service/token"
)

type impl struct {
	public public.Handler
}

func New(logger *slog.Logger, tokenSvc token.Service) Handler {
	return &impl{public: public.New(logger, tokenSvc)}
}

func (i *impl) GetPublicHandler() public.Handler { return i.public }
