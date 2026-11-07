package client

import "github.com/phillong97254/auth-service/internal/domain"

type impl struct {
	clients map[string]domain.Client
}

func New(backofficeSecret string) Repository {
	return &impl{clients: seed(backofficeSecret)}
}

// Dữ liệu demo. Secret lấy từ config/env (thực tế: DB/Vault), không hard-code.
func seed(backofficeSecret string) map[string]domain.Client {
	return map[string]domain.Client{
		"backoffice": {ID: "backoffice", Secret: backofficeSecret, Scope: "users:read"},
	}
}
