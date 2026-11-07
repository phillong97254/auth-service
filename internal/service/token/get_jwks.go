package token

import "context"

func (i *impl) GetJWKS(_ context.Context) GetJWKSOutput {
	return GetJWKSOutput{JWKS: i.signer.JWKS()}
}
