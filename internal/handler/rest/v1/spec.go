package v1

import "github.com/phillong97254/auth-service/internal/handler/rest/v1/public"

type Handler interface {
	GetPublicHandler() public.Handler
}
