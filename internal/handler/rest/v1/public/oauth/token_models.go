package oauth

import (
	"net/http"

	"github.com/phil/auth-service/internal/service/token"
)

const (
	grantPassword          = "password"
	grantClientCredentials = "client_credentials"
)

// TokenRequest đọc từ form (application/x-www-form-urlencoded) theo chuẩn OAuth2.
type TokenRequest struct {
	GrantType    string
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
	Scope        string
}

func bindTokenRequest(r *http.Request) (TokenRequest, error) {
	if err := r.ParseForm(); err != nil {
		return TokenRequest{}, err
	}
	req := TokenRequest{
		GrantType: r.PostForm.Get("grant_type"),
		Username:  r.PostForm.Get("username"),
		Password:  r.PostForm.Get("password"),
		Scope:     r.PostForm.Get("scope"),
	}
	// client authentication: ưu tiên HTTP Basic, fallback sang form.
	if id, secret, ok := r.BasicAuth(); ok {
		req.ClientID, req.ClientSecret = id, secret
	} else {
		req.ClientID, req.ClientSecret = r.PostForm.Get("client_id"), r.PostForm.Get("client_secret")
	}
	return req, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func toTokenResponse(out token.TokenOutput) TokenResponse {
	return TokenResponse{
		AccessToken: out.AccessToken,
		TokenType:   out.TokenType,
		ExpiresIn:   out.ExpiresIn,
		Scope:       out.Scope,
	}
}
