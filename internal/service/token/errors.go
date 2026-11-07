package token

import "errors"

// Mã lỗi theo chuẩn OAuth2 (RFC 6749 §5.2).
var (
	ErrInvalidRequest = errors.New("invalid_request")
	ErrInvalidGrant   = errors.New("invalid_grant")
	ErrInvalidClient  = errors.New("invalid_client")
)
