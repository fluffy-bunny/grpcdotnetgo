package jwt

import (
	"time"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
)

type (
	MintUnsignedOptons struct {
		NotBefore  *time.Time
		IssuedAt   *time.Time
		Expiration *time.Duration
	}
)

// MintUnsignedToken creates a new unsigned token
func MintUnsignedToken(subject string, extraClaims jwt.MapClaims, options *MintUnsignedOptons) (string, error) {
	if !utils.IsEmptyOrNil(subject) {
		extraClaims["sub"] = subject
	}
	extraClaims["iat"] = time.Now().Unix()
	extraClaims["nbf"] = time.Now().Unix()
	extraClaims["exp"] = time.Now().Add(time.Hour).Unix()
	if options != nil {
		if options.Expiration != nil {
			extraClaims["exp"] = time.Now().Add(*options.Expiration).Unix()
		}
		if options.IssuedAt != nil {
			extraClaims["iat"] = options.IssuedAt.Unix()
		}
		if options.NotBefore != nil {
			extraClaims["nbf"] = options.NotBefore.Unix()
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, extraClaims)
	return token.SignedString(jwt.UnsafeAllowNoneSignatureType)
}

// DecodeUnsignedToken decodes an unsigned token
func DecodeUnsignedToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt.UnsafeAllowNoneSignatureType, nil
	})
	return token, err
}

// ClaimsPrincipalFromUnsignedToken decodes an unsigned token
func ClaimsPrincipalFromUnsignedToken(tokenString string) (contracts_claimsprincipal.IClaimsPrincipal, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt.UnsafeAllowNoneSignatureType, nil
	})
	if err != nil {
		return nil, err
	}
	return services_claimsprincipal.ClaimsPrincipalFromClaimsMap(token.Claims.(jwt.MapClaims)), err
}
