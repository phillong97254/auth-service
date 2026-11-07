package domain

// User là người dùng backoffice. Demo lưu mật khẩu thô; thực tế phải hash bcrypt/argon2.
type User struct {
	Username string
	Password string
	Roles    []string
	Scope    string // các scope cách nhau bằng dấu cách
}
