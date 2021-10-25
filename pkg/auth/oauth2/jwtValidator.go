package oauth2

import (
	"context"
	"time"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	jwxt "github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
)

// JWTValidatorOptions is a struct for specifying configuration options.
type JWTValidatorOptions struct {
	OAuth2Document   *OAuth2Document
	ClockSkewMinutes int
}

// JWTValidator struct
type JWTValidator struct {
	Options *JWTValidatorOptions
}

const (
	optionsCannotBeNil = "options cannot be nil"
)

// NewJWTValidator creates a new *JWTValidator
func NewJWTValidator(options *JWTValidatorOptions) *JWTValidator {
	if options == nil {
		log.Fatal().Msg(optionsCannotBeNil)
		panic(optionsCannotBeNil)
	}

	return &JWTValidator{
		Options: options,
	}
}

// ParseToken validates an produces a claims principal
func (jwtValidator *JWTValidator) ParseToken(ctx context.Context, accessToken string) (claimsprincipalContracts.IClaimsPrincipal, error) {
	var validationOpts []jwxt.ValidateOption
	// Parse the JWT
	jwkSet, err := jwtValidator.Options.OAuth2Document.fetchJwks(ctx)
	if err != nil {
		return nil, err
	}

	token, err := jwxt.ParseString(accessToken, jwxt.WithKeySet(jwkSet))
	if err != nil {
		return nil, err
	}

	// This set had a key that worked
	validationOpts = append(validationOpts, jwxt.WithIssuer(jwtValidator.Options.OAuth2Document.Issuer))

	// Allow clock skew
	validationOpts = append(validationOpts, jwxt.WithAcceptableSkew(time.Minute*time.Duration(jwtValidator.Options.ClockSkewMinutes)))

	opts := validationOpts
	err = jwxt.Validate(token, opts...)
	if err != nil {
		return nil, err
	}
	claimsMap, err := token.AsMap(ctx)
	if err != nil {
		return nil, err
	}
	result := ClaimsPrincipalFromClaimsMap(claimsMap)

	return result, nil
}
