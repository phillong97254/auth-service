package authz

import (
	"context"
	"fmt"

	authzv1 "github.com/phillong97254/auth-service/pkg/gen/authz/v1"
)

func (i *impl) GetUserRoles(ctx context.Context, input GetUserRolesInput) (GetUserRolesOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeout)
	defer cancel()
	resp, err := i.client.GetUserRoles(ctx, &authzv1.GetUserRolesRequest{UserId: input.UserID})
	if err != nil {
		return GetUserRolesOutput{}, fmt.Errorf("authz GetUserRoles: %w", err)
	}
	return GetUserRolesOutput{Roles: resp.GetRoles(), Scope: resp.GetScope()}, nil
}
