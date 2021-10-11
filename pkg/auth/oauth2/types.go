package oauth2

import (
	"context"
	"net/url"

	jwxk "github.com/lestrrat-go/jwx/jwk"
	jwxt "github.com/lestrrat-go/jwx/jwt"
)

// CtxClaimsPrincipalKeyStruct struct
type CtxClaimsPrincipalKeyStruct struct{}

// CtxClaimsPrincipalKey key
var CtxClaimsPrincipalKey = &CtxClaimsPrincipalKeyStruct{}

// OAuth2DiscoveryOptions ...
type OAuth2DiscoveryOptions struct {
	JWKSURL string
}

// DiscoveryDocumentOptions ...
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
func ClaimsPrincipalFromClaimsMap(claimsMap map[string]interface{}) *ClaimsPrincipal {
	result := ClaimsPrincipal{
		Claims:  []Claim{},
		FastMap: make(map[string]map[string]bool),
	}
	var addFastMapClaim = func(key string, value string) {
		claimParent, ok := result.FastMap[key]
		if !ok {
			claimParent = make(map[string]bool)
			result.FastMap[key] = claimParent
		}
		claimParent[value] = true
	}

	for key, element := range claimsMap {
		switch value := element.(type) {
		case string:
			addFastMapClaim(key, value)
			result.Claims = append(result.Claims, Claim{Type: key, Value: value})

		case []interface{}:
			for _, value := range value {
				switch claimValue := value.(type) {
				case string:
					addFastMapClaim(key, claimValue)
					result.Claims = append(result.Claims, Claim{Type: key, Value: claimValue})
				}
			}
		case []string:
			for _, value := range value {
				addFastMapClaim(key, value)
				result.Claims = append(result.Claims, Claim{Type: key, Value: value})
			}
		}

	}
	return &result
}
