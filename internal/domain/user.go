package domain

// User là người dùng backoffice. Demo lưu mật khẩu thô; thực tế phải hash bcrypt/argon2.
// Role/scope KHÔNG lưu ở đây: authz-service quản lý, auth-service hỏi khi phát token.
type User struct {
	Username string
	Password string
}
