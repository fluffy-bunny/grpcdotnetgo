package oidc

import (
	"context"

	"errors"
	"reflect"

	auth_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/auth/oidc"
	contracts_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oidc"

	"github.com/coreos/go-oidc/v3/oidc"
	di "github.com/fluffy-bunny/sarulabsdi"
	"golang.org/x/oauth2"
)

type (
	service struct {
		*oidc.Provider
		oauth2.Config
		GetOIDCAuthenticatorConfig contracts_oidc.GetOIDCAuthenticatorConfig `inject:""`
		oidcProviderEx             *auth_oidc.Provider
		issuer                     string
	}
)

func assertImplementation() {
	var _ contracts_oidc.IOIDCAuthenticator = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIOIDCAuthenticator registers the *service as a singleton.
func AddSingletonIOIDCAuthenticator(builder *di.Builder) {
	contracts_oidc.AddSingletonIOIDCAuthenticator(builder, reflectType)
}
func (s *service) Ctor() {
	config := s.GetOIDCAuthenticatorConfig()
	s.issuer = "https://" + config.Domain + "/"
	oidcProviderEx, err := auth_oidc.NewProvider(context.Background(), s.issuer)
	if err != nil {
		panic(err)
	}
	s.oidcProviderEx = oidcProviderEx
	provider, err := oidc.NewProvider(
		context.Background(),
		s.issuer,
	)
	if err != nil {
		panic(err)
	}
	s.Provider = provider

	conf := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess, "profile"},
	}
	s.Config = conf
}

func (s *service) ValidateJWTAccessToken(accessToken string) (*auth_oidc.AccessToken, error) {
	verifier := auth_oidc.NewJWTAccessTokenVerifier(s.issuer, s.oidcProviderEx.GetRemoteKeySet(), &oidc.Config{
		SkipClientIDCheck: true,
	})
	return verifier.Verify(context.Background(), accessToken)
}

func (s *service) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: s.ClientID,
	}
	return s.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (s *service) GetTokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	ts := s.Config.TokenSource(ctx, token)
	return ts
}
