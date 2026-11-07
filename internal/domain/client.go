package domain

// Client là OAuth client dùng cho giao tiếp service-to-service (client_credentials).
type Client struct {
	ID     string
	Secret string
	Scope  string
}
