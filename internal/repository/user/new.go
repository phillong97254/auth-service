package user

import "github.com/phillong97254/auth-service/internal/domain"

type impl struct {
	users map[string]domain.User
}

func New() Repository {
	return &impl{users: seed()}
}

// Dữ liệu demo (chỉ tài khoản + mật khẩu; role ở authz-service). Thực tế: lưu DB.
func seed() map[string]domain.User {
	return map[string]domain.User{
		"alice": {Username: "alice", Password: "alice123"},
		"bob":   {Username: "bob", Password: "bob123"},
		"carol": {Username: "carol", Password: "carol123"},
	}
}
