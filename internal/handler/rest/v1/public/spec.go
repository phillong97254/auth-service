package public

import "github.com/phil/auth-service/internal/handler/rest/v1/public/oauth"

// Handler cho các endpoint không cần access token.
type Handler interface {
	GetOAuthHandler() oauth.Handler
}
