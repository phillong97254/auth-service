// Package signer giữ private key và ký JWT (RS256). Dùng chung cho tầng service
// (phát token cho client) và tầng gateway (token của chính auth-service khi gọi authz-service).
package signer

import auth "github.com/phillong97254/go-lib/authkit"

type Signer interface {
	Sign(input SignInput) (SignOutput, error)
	JWKS() auth.JWKS
}

type SignInput struct {
	Subject  string
	Roles    []string
	Scope    string
	ClientID string
}

type SignOutput struct {
	Token     string
	ExpiresIn int // giây
}
