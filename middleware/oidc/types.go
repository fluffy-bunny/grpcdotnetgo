package oidc

import (
	"io"
	"net/url"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Type  string `mapstructure:"TYPE"`
	Value string `mapstructure:"VALUE"`
}

type ClaimsConfig struct {
	OR  []Claim `mapstructure:"OR"`
	AND []Claim `mapstructure:"AND"`
}
type EntryPointConfig struct {
	ClaimsConfig ClaimsConfig `mapstructure:"CLAIMS_CONFIG"`
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
