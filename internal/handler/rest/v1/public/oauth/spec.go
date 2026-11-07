package oauth

import "net/http"

type Handler interface {
	// POST /oauth/token  grant_type=password | client_credentials
	Token(w http.ResponseWriter, r *http.Request)
	// GET /.well-known/jwks.json
	GetJWKS(w http.ResponseWriter, r *http.Request)
}
