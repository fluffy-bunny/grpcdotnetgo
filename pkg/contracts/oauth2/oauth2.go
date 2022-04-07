package oauth2

import (
	"context"

	"golang.org/x/oauth2"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IOAuth2Authenticator"
//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/func-types.go -out=gen-func-$GOFILE gen "FuncType=GetOAuth2AuthenticatorConfig"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IOAuth2Authenticator

type (

	// IOAuth2Authenticator ...
	IOAuth2Authenticator interface {
		AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
		Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
		GetTokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource
	}
	// GetOAuth2AuthenticatorConfig ...
	GetOAuth2AuthenticatorConfig func() *oauth2.Config
)
