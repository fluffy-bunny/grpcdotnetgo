package oauth2

import (
	"context"
	"net/url"

	jwxk "github.com/lestrrat-go/jwx/jwk"
	jwxt "github.com/lestrrat-go/jwx/jwt"
)

const (
	CtxClaimsPrincipalKey = "ClaimsPrincipal"
)

type OAuth2DiscoveryOptions struct {
	JWKSURL string
}
type DiscoveryDocumentOptions struct {
	Authority              string
	OAuth2DiscoveryOptions OAuth2DiscoveryOptions
}
type OAuth2Document struct {
	Options      *OAuth2DiscoveryOptions
	Issuer       string `json:"issuer"`
	JWKSURL      string `json:"jwks_uri"`
	jwksAR       *jwxk.AutoRefresh
	jwksCancelAR context.CancelFunc
}
type DiscoveryDocument struct {
	OAuth2Document        *OAuth2Document
	Options               *DiscoveryDocumentOptions
	DiscoveryURL          url.URL
	Algorithms            []string `json:"id_token_signing_alg_values_supported"`
	IntrospectionEndpoint string
	Issuer                string `json:"issuer"`
	JWKSURL               string `json:"jwks_uri"`
}
type Claim struct {
	Type  string
	Value string
}
type ClaimsPrincipal struct {
	Token   jwxt.Token
	Claims  []Claim
	FastMap map[string]map[string]bool
}
type OAuth2Context struct {
	OAuth2Document *OAuth2Document
	JWTValidator   *JWTValidator
	Scheme         string
	Config         *GrpcFuncAuthConfig
}
type MethodClaims struct {
	OR  []Claim
	AND []Claim
}
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

func NewGrpcFuncAuthConfig(authority string, expectedScheme string, clockSkewMinutes int) *GrpcFuncAuthConfig {
	return &GrpcFuncAuthConfig{
		Authority:              authority,
		ExpectedScheme:         expectedScheme,
		ClockSkewMinutes:       clockSkewMinutes,
		FullMethodNameToClaims: make(map[string]MethodClaims),
	}
}
