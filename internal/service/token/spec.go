package token

import "context"

// Service phát hành access token (JWT RS256) và public key (JWKS) để service khác verify.
type Service interface {
	IssuePasswordToken(ctx context.Context, input IssuePasswordTokenInput) (TokenOutput, error)
	IssueClientToken(ctx context.Context, input IssueClientTokenInput) (TokenOutput, error)
	GetJWKS(ctx context.Context) GetJWKSOutput
}
