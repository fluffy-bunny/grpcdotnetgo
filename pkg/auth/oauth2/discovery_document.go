package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	jwxk "github.com/lestrrat-go/jwx/jwk"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	OptionsCannotBeNil = "options cannot be nil"
)

func (p OAuth2Document) MarshalZerologObject(e *zerolog.Event) {
	e.Str("JWKSURL", p.JWKSURL)
}
func (p DiscoveryDocument) MarshalZerologObject(e *zerolog.Event) {
	e.Str("Issuer", p.Issuer).
		Str("JWKSURL", p.JWKSURL)
}
func newOAuth2Document(options *OAuth2DiscoveryOptions) (*OAuth2Document, error) {
	if options == nil {
		log.Fatal().Msg(OptionsCannotBeNil)
		panic(OptionsCannotBeNil)
	}

	return &OAuth2Document{
		Options: options,
	}, nil
}
func newDiscoveryDocument(options *DiscoveryDocumentOptions) (*DiscoveryDocument, error) {
	if options == nil {
		log.Fatal().Msg(OptionsCannotBeNil)
		panic(OptionsCannotBeNil)
	}
	u, err := url.Parse(options.Authority)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/.well-known/openid-configuration")

	return &DiscoveryDocument{
		Options:      options,
		DiscoveryURL: *u,
	}, nil
}
func (document *DiscoveryDocument) fetchJwks(ctx context.Context) (jwxk.Set, error) {
	return document.OAuth2Document.fetchJwks(ctx)

}
func (document *OAuth2Document) fetchJwks(ctx context.Context) (jwxk.Set, error) {
	return document.jwksAR.Fetch(ctx, document.JWKSURL)

}
func (document *OAuth2Document) initialize() error {

	var ctx context.Context
	ctx, document.jwksCancelAR = context.WithCancel(context.Background())
	document.jwksAR = jwxk.NewAutoRefresh(ctx)
	document.JWKSURL = document.Options.JWKSURL
	document.jwksAR.Configure(document.JWKSURL, jwxk.WithMinRefreshInterval(time.Minute*5))

	_, err := document.jwksAR.Refresh(ctx, document.JWKSURL)
	if err != nil {
		log.Error().Err(err).
			Str("uri", document.JWKSURL).
			Msg("Initial fetch of JWKS - will try again in the background and when a request is received")
		return err
	}
	jwkSet, err := document.jwksAR.Fetch(ctx, document.JWKSURL)
	if err != nil {
		log.Error().Err(err).Str("jwks", document.JWKSURL).Msg("Fetching JWKS at auth time")
		return err
	}
	log.Debug().Int("keys", jwkSet.Len())
	return nil
}

func (document *DiscoveryDocument) initialize() error {
	err := document.loadDiscoveryDocument()
	if err != nil {
		return fmt.Errorf("error loading discovery document: %w", err)
	}
	document.Options.OAuth2DiscoveryOptions.JWKSURL = document.JWKSURL
	document.OAuth2Document, err = newOAuth2Document(&(document.Options.OAuth2DiscoveryOptions))
	if err != nil {
		return fmt.Errorf("error newOAuth2Document: %w", err)
	}
	err = document.OAuth2Document.initialize()
	if err != nil {
		return fmt.Errorf("error initializing OAuth2Document: %w", err)
	}
	return nil
}

func (document *DiscoveryDocument) loadDiscoveryDocument() error {
	resp, err := http.Get(document.DiscoveryURL.String())
	if err != nil {
		return fmt.Errorf("could not fetch discovery url: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(document)
	if err != nil {
		return fmt.Errorf("error decoding discovery document: %w", err)
	}

	return nil
}
