package v1

import "github.com/phil/auth-service/internal/handler/rest/v1/public"

type Handler interface {
	GetPublicHandler() public.Handler
}
