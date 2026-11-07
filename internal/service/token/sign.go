package token

import (
	"crypto/subtle"
	"strings"

	"github.com/phillong97254/auth-service/internal/signer"
)

func (i *impl) sign(in signer.SignInput) (TokenOutput, error) {
	out, err := i.signer.Sign(in)
	if err != nil {
		return TokenOutput{}, err
	}
	return TokenOutput{AccessToken: out.Token, TokenType: "Bearer", ExpiresIn: out.ExpiresIn, Scope: in.Scope}, nil
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
