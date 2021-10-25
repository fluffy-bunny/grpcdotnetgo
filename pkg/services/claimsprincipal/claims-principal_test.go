package claimsprincipal

import (
	"testing"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	"github.com/stretchr/testify/assert"
)

func Test_add_bad_claim(t *testing.T) {
	cp := NewIClaimsPrincipal()
	cp.AddClaim(claimsprincipalContracts.Claim{
		Type:  "",
		Value: "",
	})
	assert.Empty(t, cp.GetClaims())
}
func Test_add_claim(t *testing.T) {
	cp := NewIClaimsPrincipal()
	claim := claimsprincipalContracts.Claim{
		Type:  "a",
		Value: "b",
	}
	cp.AddClaim(claim)
	claims := cp.GetClaims()
	assert.NotEmpty(t, claims)
	assert.Equal(t, 1, len(claims))
	assert.True(t, cp.HasClaim(claim))
	assert.False(t, cp.HasClaim(claimsprincipalContracts.Claim{
		Type:  "junk",
		Value: "junk",
	}))

	cp.RemoveClaim(claim)
	claims = cp.GetClaims()
	assert.Empty(t, claims)

}
