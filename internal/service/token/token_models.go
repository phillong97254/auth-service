package token

// TokenOutput dùng chung cho mọi grant type.
type TokenOutput struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int
	Scope       string
}
