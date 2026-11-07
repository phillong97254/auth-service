package token

import (
	"context"

	auth "github.com/phillong97254/go-lib/authkit"
)

func (i *impl) GetJWKS(_ context.Context) GetJWKSOutput {
	return GetJWKSOutput{
		JWKS: auth.JWKS{Keys: []auth.JWK{auth.NewJWK(i.cfg.OAuth.KeyID, &i.key.PublicKey)}},
	}
}
