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

const (
	optionsCannotBeNil = "options cannot be nil"
)

func NewJWTValidator(options *JWTValidatorOptions) *JWTValidator {
	if options == nil {
		log.Fatal().Msg(optionsCannotBeNil)
		panic(optionsCannotBeNil)
	}

	return &JWTValidator{
		Options: options,
	}
}

func (jwtValidator *JWTValidator) NewEmptyClaimsPrincipal() *ClaimsPrincipal {
	return &ClaimsPrincipal{}
}
func (jwtValidator *JWTValidator) ParseToken(ctx context.Context, accessToken string) (*ClaimsPrincipal, error) {
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

	result := ClaimsPrincipal{
		Claims:  []Claim{},
		Token:   token,
		FastMap: make(map[string]map[string]bool),
	}

	var addFastMapClaim = func(key string, value string) {
		claimParent, ok := result.FastMap[key]
		if !ok {
			claimParent = make(map[string]bool)
			result.FastMap[key] = claimParent
		}
		claimParent[value] = true
	}
	claimsMap, err := token.AsMap(ctx)
	if err != nil {
		return nil, err
	}

	for key, element := range claimsMap {
		switch c := element.(type) {
		case string:
			addFastMapClaim(key, element.(string))
			result.Claims = append(result.Claims, Claim{Type: key, Value: element.(string)})

		case []interface{}:
			for _, value := range c {
				switch claimValue := value.(type) {
				case string:
					addFastMapClaim(key, claimValue)
					result.Claims = append(result.Claims, Claim{Type: key, Value: claimValue})
				}
			}
		case []string:
			for _, value := range c {
				addFastMapClaim(key, value)
				result.Claims = append(result.Claims, Claim{Type: key, Value: value})
			}
		}

	}

	return &result, nil
}
