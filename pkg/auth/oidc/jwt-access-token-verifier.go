package oidc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	go_oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	jose_jwt "gopkg.in/square/go-jose.v2/jwt"
)

const (
	issuerGoogleAccounts         = "https://accounts.google.com"
	issuerGoogleAccountsNoScheme = "accounts.google.com"
)

// JWTAccessTokenVerifier provides verification for ID Tokens.
type JWTAccessTokenVerifier struct {
	keySet go_oidc.KeySet
	config *go_oidc.Config
	issuer string
}

// NewJWTAccessTokenVerifier returns a verifier manually constructed from a key set and issuer URL.
//
// It's easier to use provider discovery to construct an JWTAccessTokenVerifier than creating
// one directly. This method is intended to be used with provider that don't support
// metadata discovery, or avoiding round trips when the key set URL is already known.
//
// This constructor can be used to create a verifier directly using the issuer URL and
// JSON Web Key Set URL without using discovery:
//
//		keySet := oidc.NewRemoteKeySet(ctx, "https://www.googleapis.com/oauth2/v3/certs")
//		verifier := oidc.NewVerifier("https://accounts.google.com", keySet, config)
//
// Since KeySet is an interface, this constructor can also be used to supply custom
// public key sources. For example, if a user wanted to supply public keys out-of-band
// and hold them statically in-memory:
//
//		// Custom KeySet implementation.
//		keySet := newStatisKeySet(publicKeys...)
//
//		// Verifier uses the custom KeySet implementation.
//		verifier := oidc.NewVerifier("https://auth.example.com", keySet, config)
//
func NewJWTAccessTokenVerifier(issuerURL string, keySet go_oidc.KeySet, config *go_oidc.Config) *JWTAccessTokenVerifier {
	return &JWTAccessTokenVerifier{keySet: keySet, config: config, issuer: issuerURL}
}

// Verifier returns an JWTAccessTokenVerifier that uses the provider's key set to verify JWTs.
//
// The returned JWTAccessTokenVerifier is tied to the Provider's context and its behavior is
// undefined once the Provider's context is canceled.
func (p *Provider) Verifier(config *go_oidc.Config) *JWTAccessTokenVerifier {
	if len(config.SupportedSigningAlgs) == 0 && len(p.algorithms) > 0 {
		// Make a copy so we don't modify the config values.
		cp := &go_oidc.Config{}
		*cp = *config
		cp.SupportedSigningAlgs = p.algorithms
		config = cp
	}
	return NewJWTAccessTokenVerifier(p.issuer, p.remoteKeySet, config)
}

func parseJWT(p string) ([]byte, error) {
	parts := strings.Split(p, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("oidc: malformed jwt, expected 3 parts got %d", len(parts))
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("oidc: malformed jwt payload: %v", err)
	}
	return payload, nil
}

func contains(sli []string, ele string) bool {
	for _, s := range sli {
		if s == ele {
			return true
		}
	}
	return false
}

type accessToken struct {
	Issuer    string    `json:"iss"`
	Subject   string    `json:"sub"`
	Audience  audience  `json:"aud"`
	Expiry    jsonTime  `json:"exp"`
	IssuedAt  jsonTime  `json:"iat"`
	NotBefore *jsonTime `json:"nbf"`
}

// Verify parses a raw ID Token, verifies it's been signed by the provider, performs
// any additional checks depending on the Config, and returns the payload.
//
// Verify does NOT do nonce validation, which is the callers responsibility.
//
// See: https://openid.net/specs/openid-connect-core-1_0.html#IDTokenValidation
//
//    oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
//    if err != nil {
//        // handle error
//    }
//
//    // Extract the ID Token from oauth2 token.
//    rawIDToken, ok := oauth2Token.Extra("id_token").(string)
//    if !ok {
//        // handle error
//    }
//
//    token, err := verifier.Verify(ctx, rawIDToken)
//
type CustomClaims struct {
	*jose_jwt.Claims
	AnyJSONObjectClaim map[string]interface{} `json:"anyJSONObjectClaim"`
}

// Verify ...
func (v *JWTAccessTokenVerifier) Verify(ctx context.Context, rawToken string) (*AccessToken, error) {
	log.Trace().Msg("ENTER - JWTAccessTokenVerifier.Verify")
	defer log.Trace().Msg("EXIT - JWTAccessTokenVerifier.Verify")
	parsedJWT, err := jose_jwt.ParseSigned(rawToken)
	if err != nil {
		log.Trace().Err(err).Msg("Failed to parse JWT")
		return nil, fmt.Errorf("oidc: malformed jwt: %v", err)
	}
	_, err = v.keySet.VerifySignature(ctx, rawToken)
	if err != nil {
		log.Trace().Err(err).Msg("failed to verify signature")
		return nil, fmt.Errorf("failed to verify signature: %v", err)
	}
	resultCl := CustomClaims{}
	rawCL := map[string]interface{}{}
	err = parsedJWT.UnsafeClaimsWithoutVerification(&resultCl)
	if err != nil {
		log.Trace().Err(err).Msg("UnsafeClaimsWithoutVerification(&resultCl) failed to parse claims")
		return nil, fmt.Errorf("oidc: malformed jwt: %v", err)
	}
	err = parsedJWT.UnsafeClaimsWithoutVerification(&rawCL)
	if err != nil {
		log.Trace().Err(err).Msg("UnsafeClaimsWithoutVerification(&rawCL) failed to parse claims")
		return nil, fmt.Errorf("oidc: malformed jwt: %v", err)
	}
	// Throw out tokens with invalid claims before trying to verify the token. This lets
	// us do cheap checks before possibly re-syncing keys.
	payload, err := parseJWT(rawToken)
	if err != nil {
		log.Trace().Err(err).Msg("Failed to parse JWT")
		return nil, fmt.Errorf("oidc: malformed jwt: %v", err)
	}
	var token accessToken
	if err := json.Unmarshal(payload, &token); err != nil {
		log.Trace().Err(err).Msg("Failed to unmarshal access token")
		return nil, fmt.Errorf("oidc: failed to unmarshal claims: %v", err)
	}

	t := &AccessToken{
		Issuer:   resultCl.Issuer,
		Subject:  resultCl.Subject,
		Audience: []string(resultCl.Audience),
		Expiry:   resultCl.Expiry.Time(),
		IssuedAt: resultCl.IssuedAt.Time(),

		Claims: rawCL,
	}

	// Check issuer.
	if !v.config.SkipIssuerCheck && t.Issuer != v.issuer {
		// Google sometimes returns "accounts.google.com" as the issuer claim instead of
		// the required "https://accounts.google.com". Detect this case and allow it only
		// for Google.
		//
		// We will not add hooks to let other providers go off spec like this.
		if !(v.issuer == issuerGoogleAccounts && t.Issuer == issuerGoogleAccountsNoScheme) {
			log.Trace().Err(err).Msg("Issuer mismatch")
			return nil, fmt.Errorf("oidc: id token issued by a different provider, expected %q got %q", v.issuer, t.Issuer)
		}
	}

	// If a client ID has been provided, make sure it's part of the audience. SkipClientIDCheck must be true if ClientID is empty.
	//
	// This check DOES NOT ensure that the ClientID is the party to which the ID Token was issued (i.e. Authorized party).
	if !v.config.SkipClientIDCheck {
		if v.config.ClientID != "" {
			if !contains(t.Audience, v.config.ClientID) {
				log.Trace().Err(err).Msg("ClientID mismatch")

				return nil, fmt.Errorf("oidc: expected audience %q got %q", v.config.ClientID, t.Audience)
			}
		} else {
			return nil, fmt.Errorf("oidc: invalid configuration, clientID must be provided or SkipClientIDCheck must be set")
		}
	}

	// If a SkipExpiryCheck is false, make sure token is not expired.
	if !v.config.SkipExpiryCheck {
		now := time.Now
		if v.config.Now != nil {
			now = v.config.Now
		}
		nowTime := now()

		if t.Expiry.Before(nowTime) {
			log.Trace().Err(err).Msg("Token is expired")

			return nil, fmt.Errorf("oidc: token is expired (Token Expiry: %v)", t.Expiry)
		}

		// If nbf claim is provided in token, ensure that it is indeed in the past.
		if token.NotBefore != nil {
			nbfTime := time.Time(*token.NotBefore)
			leeway := 1 * time.Minute

			if nowTime.Add(leeway).Before(nbfTime) {
				log.Trace().Err(err).Msg("Token is not yet valid")
				return nil, fmt.Errorf("oidc: current time %v before the nbf (not before) time: %v", nowTime, nbfTime)
			}
		}
	}

	return t, nil
}

// Nonce returns an auth code option which requires the ID Token created by the
// OpenID Connect provider to contain the specified nonce.
func Nonce(nonce string) oauth2.AuthCodeOption {
	return oauth2.SetAuthURLParam("nonce", nonce)
}
