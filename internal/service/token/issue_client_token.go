package token

import (
	"context"
	"errors"

	clientRepo "github.com/phillong97254/auth-service/internal/repository/client"
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
	return i.sign(signInput{
		subject:  "service:" + c.ID,
		scope:    narrowScope(c.Scope, input.Scope),
		clientID: c.ID,
	})
}
