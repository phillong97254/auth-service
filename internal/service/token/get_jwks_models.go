package token

import auth "github.com/phil/auth-service/pkg/authkit"

type GetJWKSOutput struct {
	JWKS auth.JWKS
}
