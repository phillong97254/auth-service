package token

import "testing"

func TestNarrowScope(t *testing.T) {
	tests := []struct{ have, want, out string }{
		{"wallet:read users:read", "", "wallet:read users:read"},
		{"wallet:read users:read", "wallet:read", "wallet:read"},
		{"wallet:read users:read", "wallet:read wallet:write", "wallet:read"}, // không được xin thêm
		{"wallet:read", "wallet", ""},                                         // không khớp một phần
	}
	for _, tt := range tests {
		if got := narrowScope(tt.have, tt.want); got != tt.out {
			t.Errorf("narrowScope(%q, %q) = %q, want %q", tt.have, tt.want, got, tt.out)
		}
	}
}
