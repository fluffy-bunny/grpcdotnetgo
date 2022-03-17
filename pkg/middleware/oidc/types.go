package oidc

import (
	"io"
	"net/url"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type (
	// ClaimFactDirective ...
	ClaimFactDirective int64
	// ClaimFact ...
	ClaimFact struct {
		Claim     claimsprincipalContracts.Claim
		Directive ClaimFactDirective
	}
)

// NewClaimFactTypeAndValueClaim ...
func NewClaimFactTypeAndValueClaim(claim claimsprincipalContracts.Claim) ClaimFact {
	return ClaimFact{
		Claim:     claim,
		Directive: ClaimTypeAndValue,
	}
}

// NewClaimFactTypeAndValue ...
func NewClaimFactTypeAndValue(claimType string, value string) ClaimFact {
	return NewClaimFactTypeAndValueClaim(claimsprincipalContracts.Claim{
		Type:  claimType,
		Value: value,
	})
}

// NewClaimFactType ...
func NewClaimFactType(claimType string) ClaimFact {
	return ClaimFact{
		Claim: claimsprincipalContracts.Claim{
			Type: claimType,
		},
		Directive: ClaimType,
	}
}

const (
	// ClaimTypeAndValue ...
	ClaimTypeAndValue ClaimFactDirective = 0
	// ClaimType ...
	ClaimType = 1
)

// ClaimsConfig ...
type ClaimsConfig struct {
	OR  []ClaimFact `mapstructure:"OR"`
	AND []ClaimFact `mapstructure:"AND"`
}

// EntryPointConfig ...
type EntryPointConfig struct {
	FullMethodName string       `mapstructure:"FULL_METHOD_NAME"`
	ClaimsConfig   ClaimsConfig `mapstructure:"CLAIMS_CONFIG"`
}

// OIDCConfig  env:OIDC_CONFIG
type OIDCConfig struct {
	Authority string `mapstructure:"AUTHORITY"`
	// CronRefreshSchedule i.e. @every 0h1m0s
	CronRefreshSchedule string                       `mapstructure:"CRON_REFRESH_SCHEDULE"`
	EntryPoints         map[string]*EntryPointConfig `mapstructure:"ENTRY_POINTS"`
}

// IOIDCConfig ...
type IOIDCConfig interface {
	GetAuthority() string
	GetCronRefreshSchedule() string
	GetEntryPoints() map[string]*EntryPointConfig
}

func assertImplementation() {
	var _ IOIDCConfig = (*OIDCConfig)(nil)
}

// GetAuthority ...
func (c *OIDCConfig) GetAuthority() string {
	return c.Authority
}

// GetCronRefreshSchedule ...
func (c *OIDCConfig) GetCronRefreshSchedule() string {
	return c.CronRefreshSchedule
}

// GetEntryPoints ...
func (c *OIDCConfig) GetEntryPoints() map[string]*EntryPointConfig {
	return c.EntryPoints
}

// IOIDCConfigAccessor ...
type IOIDCConfigAccessor interface {
	GetOIDCConfig() IOIDCConfig
}

var (
	// TypeIOIDCConfig ...
	TypeIOIDCConfig = di.GetInterfaceReflectType((*IOIDCConfig)(nil))
	// TypeIOIDCConfigAccessor ...
	TypeIOIDCConfigAccessor = di.GetInterfaceReflectType((*IOIDCConfigAccessor)(nil))
)

// JSONWebKeyResponse ...
type JSONWebKeyResponse struct {
	Keys []JSONWebKey `json:"keys"`
}

// JSONWebKey ...
type JSONWebKey struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// DiscoveryDocument ...
type DiscoveryDocument struct {
	DiscoveryURL          url.URL
	Algorithms            []string `json:"id_token_signing_alg_values_supported"`
	IntrospectionEndpoint string
	Issuer                string `json:"issuer"`
	JWKSURL               string `json:"jwks_uri"`
	KeyResponse           *JSONWebKeyResponse
}

// User ...
type User struct {
	Claims jwt.MapClaims
}

// NewOIDCAuthenticationOptions ...
type NewOIDCAuthenticationOptions struct {
	Authority *url.URL
}

// NewJWTValidationMiddlewareOptions ...
type NewJWTValidationMiddlewareOptions struct {
	Out      io.Writer
	LogLevel logrus.Level

	DiscoveryURL *url.URL
}

// NewGinIntrospectionValidationMiddlewareOptions ...
type NewGinIntrospectionValidationMiddlewareOptions struct {
	Out      io.Writer
	LogLevel logrus.Level

	DiscoveryURL *url.URL
	ClientID     string
	ClientSecret string
}
