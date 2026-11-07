package client

import "github.com/phil/auth-service/internal/domain"

type impl struct {
	clients map[string]domain.Client
}

func New() Repository {
	return &impl{clients: seed()}
}

// Dữ liệu demo. Thực tế: secret lấy từ DB/Vault, không hard-code.
func seed() map[string]domain.Client {
	return map[string]domain.Client{
		"backoffice": {ID: "backoffice", Secret: "backoffice-secret", Scope: "users:read"},
	}
}
