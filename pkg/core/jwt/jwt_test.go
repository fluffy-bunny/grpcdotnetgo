package jwt

import (
	"testing"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

func TestUnsignedTok(t *testing.T) {
	jwtClaims := jwt.MapClaims{
		"permissions": []string{"read", "write"},
		"client_id":   "my-client",
		"some_number": 123,
		"some_bool":   true,
	}
	token, err := MintUnsignedToken("bob", jwtClaims, nil)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	jwtToken, err := DecodeUnsignedToken(token)
	require.NoError(t, err)
	require.NotNil(t, jwtToken)
	require.Equal(t, "bob", jwtToken.Claims.(jwt.MapClaims)["sub"])
	require.Contains(t, jwtToken.Claims.(jwt.MapClaims)["permissions"], "read")
	require.Contains(t, jwtToken.Claims.(jwt.MapClaims)["permissions"], "write")

	claimsPrincipal, err := ClaimsPrincipalFromUnsignedToken(token)
	require.NoError(t, err)
	require.NotNil(t, claimsPrincipal)
	claims := claimsPrincipal.GetClaims()
	require.NotEmpty(t, claims)

	for name, jwtClaim := range jwtClaims {
		switch jwtClaim.(type) {
		case bool:
			require.True(t, claimsPrincipal.HasClaim(services_claimsprincipal.NewBoolClaim(name, jwtClaim.(bool))))
		case int:
			val := jwtClaim.(int)
			f64 := float64(val)
			require.True(t, claimsPrincipal.HasClaim(services_claimsprincipal.NewFloat64Claim(name, f64)))
		case float64:
			require.True(t, claimsPrincipal.HasClaim(services_claimsprincipal.NewFloat64Claim(name, jwtClaim.(float64))))
		case string:
			require.True(t, claimsPrincipal.HasClaim(services_claimsprincipal.NewStringClaim(name, jwtClaim.(string))))

		case []string:
			for _, value := range jwtClaim.([]string) {
				require.True(t, claimsPrincipal.HasClaim(contracts_claimsprincipal.Claim{
					Type:  name,
					Value: value,
				}))
			}
		}
	}

}
