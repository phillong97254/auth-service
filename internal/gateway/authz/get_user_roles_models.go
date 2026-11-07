package authz

type GetUserRolesInput struct {
	UserID string
}

type GetUserRolesOutput struct {
	Roles []string
	Scope string
}
