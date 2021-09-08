package mockoidc

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func standardRootClaims(config *Config, ttl time.Duration, now time.Time) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Audience:  config.ClientID,
		ExpiresAt: now.Add(ttl).Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    config.Issuer,
		NotBefore: now.Unix(),
	}
}
func MakeAccessToken(claims *jwt.StandardClaims, kp *Keypair, now time.Time) (string, error) {
	return kp.SignJWT(claims)
}
