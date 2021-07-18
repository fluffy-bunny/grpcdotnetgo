package oauth2

import (
	"context"
	"time"

	jwxt "github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
)

// Options is a struct for specifying configuration options.
type JWTValidatorOptions struct {
	OAuth2Document   *OAuth2Document
	ClockSkewMinutes int
}
type JWTValidator struct {
	Options *JWTValidatorOptions
}

func NewJWTValidator(options *JWTValidatorOptions) *JWTValidator {
	if options == nil {
		log.Fatal().Msg("options cannot be nil")
		panic("options cannot be nil")
	}

	return &JWTValidator{
		Options: options,
	}
}
func (jwtValidator *JWTValidator) ParseToken(ctx context.Context, accessToken string) (*ClaimsPrincipal, error) {
	var validationOpts []jwxt.ValidateOption
	// Parse the JWT
	jwkSet, err := jwtValidator.Options.OAuth2Document.fetchJwks(ctx)
	token, err := jwxt.ParseString(accessToken, jwxt.WithKeySet(jwkSet))
	if err != nil {
		return nil, err
	}
	if err == nil {
		// This set had a key that worked
		validationOpts = append(validationOpts, jwxt.WithIssuer(jwtValidator.Options.OAuth2Document.Issuer))

	}
	// Allow clock skew
	validationOpts = append(validationOpts, jwxt.WithAcceptableSkew(time.Minute*time.Duration(jwtValidator.Options.ClockSkewMinutes)))

	opts := validationOpts
	err = jwxt.Validate(token, opts...)

	if err != nil {
		return nil, err
	}
	result := ClaimsPrincipal{
		Claims: []Claim{},
		Token:  token,
	}
	claimsMap, err := token.AsMap(ctx)
	for key, element := range claimsMap {
		switch c := element.(type) {
		case string:
			result.Claims = append(result.Claims, Claim{Type: key, Value: element.(string)})
			break
		case []interface{}:
			for _, value := range c {
				switch value.(type) {
				case string:
					result.Claims = append(result.Claims, Claim{Type: key, Value: value.(string)})
					break
				}
			}
			break
		case []string:
			for _, value := range c {
				result.Claims = append(result.Claims, Claim{Type: key, Value: value})
			}
			break
		}

	}

	return &result, nil
}
