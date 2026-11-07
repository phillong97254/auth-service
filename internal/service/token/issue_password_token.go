package token

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	authzGwy "github.com/phillong97254/auth-service/internal/gateway/authz"
	userRepo "github.com/phillong97254/auth-service/internal/repository/user"
	"github.com/phillong97254/auth-service/internal/signer"
)

// IssuePasswordToken: grant_type=password, admin đăng nhập từ backoffice.
// Kiểm tra mật khẩu ở đây, còn role/scope lấy từ authz-service tại thời điểm phát token.
func (i *impl) IssuePasswordToken(ctx context.Context, input IssuePasswordTokenInput) (TokenOutput, error) {
	u, err := i.repo.User().GetByUsername(ctx, input.Username)
	if errors.Is(err, userRepo.ErrNotFound) {
		return TokenOutput{}, ErrInvalidGrant
	}
	if err != nil {
		return TokenOutput{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password)) != nil {
		return TokenOutput{}, ErrInvalidGrant
	}

	perm, err := i.gwy.Authz().GetUserRoles(ctx, authzGwy.GetUserRolesInput{UserID: u.Username})
	if err != nil {
		return TokenOutput{}, fmt.Errorf("%w: %w", ErrTemporarilyUnavailable, err)
	}
	return i.sign(signer.SignInput{
		Subject:  u.Username,
		Roles:    perm.Roles,
		Scope:    narrowScope(perm.Scope, input.Scope),
		ClientID: "backoffice",
	})
}
