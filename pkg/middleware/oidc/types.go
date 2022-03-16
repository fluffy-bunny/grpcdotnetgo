package oidc

import (
	"io"
	"net/url"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

// ClaimsConfig ...
type ClaimsConfig struct {
	OR      []claimsprincipalContracts.Claim `mapstructure:"OR"`
	AND     []claimsprincipalContracts.Claim `mapstructure:"AND"`
	ORTYPE  []string                         `mapstructure:"OR_TYPE"`
	ANDTYPE []string                         `mapstructure:"AND_TYPE"`
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
	CronRefreshSchedule string                      `mapstructure:"CRON_REFRESH_SCHEDULE"`
	EntryPoints         map[string]EntryPointConfig `mapstructure:"ENTRY_POINTS"`
}
type IOIDCConfig interface {
	GetAuthority() string
	GetCronRefreshSchedule() string
	GetEntryPoints() map[string]EntryPointConfig
}

func (c *OIDCConfig) GetAuthority() string {
	return c.Authority
}
func (c *OIDCConfig) GetCronRefreshSchedule() string {
	return c.CronRefreshSchedule
}
func (c *OIDCConfig) GetEntryPoints() map[string]EntryPointConfig {
	return c.EntryPoints
}

type IOIDCConfigAccessor interface {
	GetOIDCConfig() IOIDCConfig
}

var (
	TypeIOIDCConfig         = di.GetInterfaceReflectType((*IOIDCConfig)(nil))
	TypeIOIDCConfigAccessor = di.GetInterfaceReflectType((*IOIDCConfigAccessor)(nil))
)

type JSONWebKeyResponse struct {
	Keys []JSONWebKey `json:"keys"`
}

type JSONWebKey struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type DiscoveryDocument struct {
	DiscoveryURL          url.URL
	Algorithms            []string `json:"id_token_signing_alg_values_supported"`
	IntrospectionEndpoint string
	Issuer                string `json:"issuer"`
	JWKSURL               string `json:"jwks_uri"`
	KeyResponse           *JSONWebKeyResponse
}

type User struct {
	Claims jwt.MapClaims
}

type NewOIDCAuthenticationOptions struct {
	Authority *url.URL
}
type NewJWTValidationMiddlewareOptions struct {
	Out      io.Writer
	LogLevel logrus.Level

	DiscoveryURL *url.URL
}

type NewGinIntrospectionValidationMiddlewareOptions struct {
	Out      io.Writer
	LogLevel logrus.Level

	DiscoveryURL *url.URL
	ClientID     string
	ClientSecret string
}
