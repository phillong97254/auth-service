package public

import "github.com/phillong97254/auth-service/internal/handler/rest/v1/public/oauth"

// Handler cho các endpoint không cần access token.
type Handler interface {
	GetOAuthHandler() oauth.Handler
}
