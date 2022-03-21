package oauth2

import (
	"context"
	"net/url"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"

	jwxk "github.com/lestrrat-go/jwx/jwk"
)

// OAuth2DiscoveryOptions ...
type OAuth2DiscoveryOptions struct {
	JWKSURL string
}

// DiscoveryDocumentOptions ...
type DiscoveryDocumentOptions struct {
	Authority              string
	OAuth2DiscoveryOptions OAuth2DiscoveryOptions
}

// OAuth2Document ...
type OAuth2Document struct {
	Options      *OAuth2DiscoveryOptions
	Issuer       string `json:"issuer"`
	JWKSURL      string `json:"jwks_uri"`
	jwksAR       *jwxk.AutoRefresh
	jwksCancelAR context.CancelFunc
}

// DiscoveryDocument ...
type DiscoveryDocument struct {
	OAuth2Document        *OAuth2Document
	Options               *DiscoveryDocumentOptions
	DiscoveryURL          url.URL
	Algorithms            []string `json:"id_token_signing_alg_values_supported"`
	IntrospectionEndpoint string
	Issuer                string `json:"issuer"`
	JWKSURL               string `json:"jwks_uri"`
}

// OAuth2Context ...
type OAuth2Context struct {
	OAuth2Document *OAuth2Document
	JWTValidator   *JWTValidator
	Scheme         string
	Config         *GrpcFuncAuthConfig
}

// MethodClaims ...
type MethodClaims struct {
	OR  []claimsprincipalContracts.Claim
	AND []claimsprincipalContracts.Claim
}

// GrpcFuncAuthConfig ...
type GrpcFuncAuthConfig struct {
	Authority        string
	ExpectedScheme   string
	ClockSkewMinutes int
	/*
		FuncMapping["/helloworld.Greeter/SayHello"] = []oauth2.Claim{
			{"a", "b"},
			{"a", "c"},
			{"a", "d"},
			{"a", "e"},
			{"a", "f"},
		}
	*/

	FullMethodNameToClaims map[string]MethodClaims
}

// NewGrpcFuncAuthConfig ...
func NewGrpcFuncAuthConfig(authority string, expectedScheme string, clockSkewMinutes int) *GrpcFuncAuthConfig {
	return &GrpcFuncAuthConfig{
		Authority:              authority,
		ExpectedScheme:         expectedScheme,
		ClockSkewMinutes:       clockSkewMinutes,
		FullMethodNameToClaims: make(map[string]MethodClaims),
	}
}
