package oidc

import (
	"context"

	services_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/auth/oidc"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IOIDCAuthenticator"
//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/func-types.go -out=gen-func-$GOFILE gen "FuncType=GetOIDCAuthenticatorConfig"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IOIDCAuthenticator

type (
	// AuthenticatorConfig ...
	AuthenticatorConfig struct {
		Domain       string `json:"domain" mapstructure:"DOMAIN"`
		ClientID     string `json:"client_id" mapstructure:"CLIENT_ID"`
		ClientSecret string `json:"client_secret" mapstructure:"CLIENT_SECRET"`
		CallbackURL  string `json:"callback_url" mapstructure:"CALLBACK_URL"`
	}
	// IOIDCAuthenticator ...
	IOIDCAuthenticator interface {
		GetTokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource
		VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error)
		ValidateJWTAccessToken(accessToken string) (*services_oidc.AccessToken, error)
		AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
		Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	}
	// GetOIDCAuthenticatorConfig ...
	GetOIDCAuthenticatorConfig func() *AuthenticatorConfig
)
