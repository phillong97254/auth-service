package authz

import "context"

type Gateway interface {
	GetUserRoles(ctx context.Context, input GetUserRolesInput) (GetUserRolesOutput, error)
	Close() error
}
