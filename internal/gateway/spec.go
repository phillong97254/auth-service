package gateway

import "github.com/phillong97254/auth-service/internal/gateway/authz"

// Gateway gom các client gọi sang service khác.
type Gateway interface {
	Authz() authz.Gateway // authz-service (gRPC): role/scope của user
	Close() error
}
