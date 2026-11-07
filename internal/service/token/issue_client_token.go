package token

import (
	"context"
	"errors"

	clientRepo "github.com/phillong97254/auth-service/internal/repository/client"
	"github.com/phillong97254/auth-service/internal/signer"
)

// IssueClientToken: grant_type=client_credentials, service tự xin token cho chính nó.
func (i *impl) IssueClientToken(ctx context.Context, input IssueClientTokenInput) (TokenOutput, error) {
	c, err := i.repo.Client().GetByID(ctx, input.ClientID)
	if errors.Is(err, clientRepo.ErrNotFound) {
		return TokenOutput{}, ErrInvalidClient
	}
	if err != nil {
		return TokenOutput{}, err
	}
	if !secretEqual(c.Secret, input.ClientSecret) {
		return TokenOutput{}, ErrInvalidClient
	}
	return i.sign(signer.SignInput{
		Subject:  "service:" + c.ID,
		Scope:    narrowScope(c.Scope, input.Scope),
		ClientID: c.ID,
	})
}
