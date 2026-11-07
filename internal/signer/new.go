package signer

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	auth "github.com/phillong97254/go-lib/authkit"

	"github.com/phillong97254/auth-service/config"
)

type impl struct {
	cfg *config.Config
	key *rsa.PrivateKey
	ttl time.Duration
}

func New(cfg *config.Config, key *rsa.PrivateKey) Signer {
	return &impl{cfg: cfg, key: key, ttl: time.Duration(cfg.OAuth.TokenTTL) * time.Second}
}

func (i *impl) Sign(in SignInput) (SignOutput, error) {
	now := time.Now().UTC()
	claims := auth.Claims{
		Scope: in.Scope, Roles: in.Roles, ClientID: in.ClientID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    i.cfg.OAuth.Issuer,
			Subject:   in.Subject,
			Audience:  jwt.ClaimStrings{i.cfg.OAuth.Audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(i.ttl)),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tok.Header["kid"] = i.cfg.OAuth.KeyID
	signed, err := tok.SignedString(i.key)
	if err != nil {
		return SignOutput{}, fmt.Errorf("sign token: %w", err)
	}
	return SignOutput{Token: signed, ExpiresIn: int(i.ttl.Seconds())}, nil
}

func (i *impl) JWKS() auth.JWKS {
	return auth.JWKS{Keys: []auth.JWK{auth.NewJWK(i.cfg.OAuth.KeyID, &i.key.PublicKey)}}
}
