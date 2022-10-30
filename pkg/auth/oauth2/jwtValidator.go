package oauth2

import (
	"context"
	"time"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	jwxt "github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
)

// JWTValidatorOptions is a struct for specifying configuration options.
type JWTValidatorOptions struct {
	OAuth2Document    *OAuth2Document
	ClockSkewMinutes  int
	ValidateSignature *bool
	ValidateIssuer    *bool
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

func (jwtValidator *JWTValidator) shouldValidateSignature() bool {
	if jwtValidator.Options.ValidateSignature == nil {
		return true
	}
	return *jwtValidator.Options.ValidateSignature
}

func (jwtValidator *JWTValidator) shouldValidateIssuer() bool {
	if jwtValidator.Options.ValidateIssuer == nil {
		return true
	}
	return *jwtValidator.Options.ValidateIssuer
}

// ParseTokenRaw validates an produces an inteface to the raw token artifacts
func (jwtValidator *JWTValidator) ParseTokenRaw(ctx context.Context, accessToken string) (jwxt.Token, error) {
	// Parse the JWT
	parseOptions := []jwxt.ParseOption{}
	if jwtValidator.shouldValidateSignature() {
		jwkSet, err := jwtValidator.Options.OAuth2Document.fetchJwks(ctx)
		if err != nil {
			return nil, err
		}
		parseOptions = append(parseOptions, jwxt.WithKeySet(jwkSet))
	}

	token, err := jwxt.ParseString(accessToken, parseOptions...)
	if err != nil {
		return nil, err
	}

	// This set had a key that worked
	var validationOpts []jwxt.ValidateOption
	if jwtValidator.shouldValidateIssuer() {
		validationOpts = append(validationOpts, jwxt.WithIssuer(jwtValidator.Options.OAuth2Document.Issuer))
	}
	// Allow clock skew
	validationOpts = append(validationOpts, jwxt.WithAcceptableSkew(time.Minute*time.Duration(jwtValidator.Options.ClockSkewMinutes)))

	opts := validationOpts
	err = jwxt.Validate(token, opts...)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ParseToken validates an produces a claims principal
func (jwtValidator *JWTValidator) ParseToken(ctx context.Context, accessToken string) (contracts_claimsprincipal.IClaimsPrincipal, error) {
	token, err := jwtValidator.ParseTokenRaw(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	claimsMap, err := token.AsMap(ctx)
	if err != nil {
		return nil, err
	}
	result := services_claimsprincipal.ClaimsPrincipalFromClaimsMap(claimsMap)

	return result, nil
}
