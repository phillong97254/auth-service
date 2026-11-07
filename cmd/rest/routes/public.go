package routes

import (
	"net/http"

	v1 "github.com/phillong97254/auth-service/internal/handler/rest/v1"
)

// Endpoint chuẩn OAuth2/OIDC, không có prefix /v1 để client (oauth2 lib, Keycloak-compatible) dùng thẳng.
func v1PublicRoutes(mux *http.ServeMux, h v1.Handler) {
	oauthHandler := h.GetPublicHandler().GetOAuthHandler()
	mux.HandleFunc("POST /oauth/token", oauthHandler.Token)
	mux.HandleFunc("GET /.well-known/jwks.json", oauthHandler.GetJWKS)
}
