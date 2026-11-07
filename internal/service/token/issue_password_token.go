package token

import (
	"context"
	"errors"

	userRepo "github.com/phil/auth-service/internal/repository/user"
)

// IssuePasswordToken: grant_type=password, admin đăng nhập từ backoffice.
func (i *impl) IssuePasswordToken(ctx context.Context, input IssuePasswordTokenInput) (TokenOutput, error) {
	u, err := i.repo.User().GetByUsername(ctx, input.Username)
	if errors.Is(err, userRepo.ErrNotFound) {
		return TokenOutput{}, ErrInvalidGrant
	}
	if err != nil {
		return TokenOutput{}, err
	}
	if !secretEqual(u.Password, input.Password) {
		return TokenOutput{}, ErrInvalidGrant
	}
	return i.sign(signInput{
		subject:  u.Username,
		roles:    u.Roles,
		scope:    narrowScope(u.Scope, input.Scope),
		clientID: "backoffice",
	})
}
