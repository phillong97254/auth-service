package oauth

import (
	"errors"
	"net/http"

	"github.com/phillong97254/auth-service/internal/service/token"
	auth "github.com/phillong97254/go-lib/authkit"
)

func (i *impl) Token(w http.ResponseWriter, r *http.Request) {
	req, err := bindTokenRequest(r)
	if err != nil {
		writeOAuthError(w, "invalid_request")
		return
	}

	var out token.TokenOutput
	switch req.GrantType {
	case grantPassword:
		out, err = i.tokenSvc.IssuePasswordToken(r.Context(), token.IssuePasswordTokenInput{
			Username: req.Username, Password: req.Password, Scope: req.Scope,
		})
	case grantClientCredentials:
		out, err = i.tokenSvc.IssueClientToken(r.Context(), token.IssueClientTokenInput{
			ClientID: req.ClientID, ClientSecret: req.ClientSecret, Scope: req.Scope,
		})
	default:
		writeOAuthError(w, "unsupported_grant_type")
		return
	}

	switch {
	case err == nil:
		auth.WriteJSON(w, http.StatusOK, toTokenResponse(out))
	case errors.Is(err, token.ErrInvalidGrant), errors.Is(err, token.ErrInvalidClient), errors.Is(err, token.ErrInvalidRequest):
		writeOAuthError(w, err.Error())
	case errors.Is(err, token.ErrTemporarilyUnavailable):
		i.logger.Error("issue token", "grant_type", req.GrantType, "error", err)
		auth.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "temporarily_unavailable"})
	default:
		i.logger.Error("issue token", "grant_type", req.GrantType, "error", err)
		auth.WriteError(w, http.StatusInternalServerError, "internal error")
	}
}

// writeOAuthError trả lỗi theo định dạng OAuth2: {"error":"<code>"}.
// invalid_client => 401, còn lại => 400.
func writeOAuthError(w http.ResponseWriter, code string) {
	status := http.StatusBadRequest
	if code == "invalid_client" {
		status = http.StatusUnauthorized
	}
	auth.WriteJSON(w, status, map[string]string{"error": code})
}
