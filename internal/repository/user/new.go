package user

import "github.com/phil/auth-service/internal/domain"

type impl struct {
	users map[string]domain.User
}

func New() Repository {
	return &impl{users: seed()}
}

// Dữ liệu demo. Thực tế: lưu DB.
func seed() map[string]domain.User {
	return map[string]domain.User{
		"alice": {Username: "alice", Password: "alice123", Roles: []string{"admin"}, Scope: "wallet:read users:read"},
		"bob":   {Username: "bob", Password: "bob123", Roles: []string{"admin", "superadmin"}, Scope: "wallet:read wallet:write users:read"},
		"carol": {Username: "carol", Password: "carol123", Roles: []string{"viewer"}, Scope: "users:read"}, // không có quyền ví
	}
}
