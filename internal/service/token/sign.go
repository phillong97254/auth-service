package token

import (
	"crypto/subtle"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	auth "github.com/phillong97254/go-lib/authkit"
)

type signInput struct {
	subject  string
	roles    []string
	scope    string
	clientID string
}

func (i *impl) sign(in signInput) (TokenOutput, error) {
	now := time.Now().UTC()
	claims := auth.Claims{
		Scope: in.scope, Roles: in.roles, ClientID: in.clientID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    i.cfg.OAuth.Issuer,
			Subject:   in.subject,
			Audience:  jwt.ClaimStrings{i.cfg.OAuth.Audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(i.ttl)),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tok.Header["kid"] = i.cfg.OAuth.KeyID
	signed, err := tok.SignedString(i.key)
	if err != nil {
		return TokenOutput{}, fmt.Errorf("sign token: %w", err)
	}
	return TokenOutput{
		AccessToken: signed,
		TokenType:   "Bearer",
		ExpiresIn:   int(i.ttl.Seconds()),
		Scope:       in.scope,
	}, nil
}

func secretEqual(a, b string) bool { return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1 }

// narrowScope: client được xin bớt scope, không được xin thêm.
func narrowScope(have, want string) string {
	if want == "" {
		return have
	}
	var out []string
	for _, s := range strings.Fields(want) {
		if strings.Contains(" "+have+" ", " "+s+" ") {
			out = append(out, s)
		}
	}
	return strings.Join(out, " ")
}
