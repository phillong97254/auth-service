package oauth

import (
	"net/http"

	auth "github.com/phil/auth-service/pkg/authkit"
)

func (i *impl) GetJWKS(w http.ResponseWriter, r *http.Request) {
	out := i.tokenSvc.GetJWKS(r.Context())
	auth.WriteJSON(w, http.StatusOK, out.JWKS)
}
