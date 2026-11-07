package domain

import "time"

// User là tài khoản backoffice. Mật khẩu lưu dạng bcrypt.
// Role/scope KHÔNG lưu ở đây: authz-service quản lý, auth-service hỏi khi phát token.
type User struct {
	Username     string `gorm:"primaryKey"`
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (User) TableName() string { return "users" }
