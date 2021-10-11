package claimsprincipal

import (
	"fmt"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

type claimsPrincipal struct {
	claims map[string][]string
}

func removeIndex(s []string, index int) []string {
	if index >= len(s) {
		panic(fmt.Errorf("len:%v, index:%v out of range", len(s), index))
	}
	s[index] = s[len(s)-1]
	s[len(s)-1] = ""
	s = s[:len(s)-1]
	return s
}

// RemoveClaim removes a claims
func (c *claimsPrincipal) RemoveClaim(claim claimsprincipalContracts.Claim) {
	claims, ok := c.claims[claim.Type]
	if !ok {
		return
	}

	var foundidx *int
	for idx, value := range claims {
		if value == claim.Value {
			foundidx = &idx
			break
		}
	}
	if foundidx != nil {
		c.claims[claim.Type] = removeIndex(claims, *foundidx)
	}
}

// HasClaim ...
func (c *claimsPrincipal) HasClaim(claim claimsprincipalContracts.Claim) bool {
	claims, ok := c.claims[claim.Type]
	if !ok {
		return false
	}

	for _, value := range claims {
		if value == claim.Value {
			return true
		}
	}
	return false
}

// AddClaim ...
func (c *claimsPrincipal) AddClaim(claim claimsprincipalContracts.Claim) {
	if c.HasClaim(claim) {
		return
	}

	if len(claim.Type) == 0 || len(claim.Value) == 0 {
		panic("invalid claim input")
	}
	claims, ok := c.claims[claim.Type]
	if !ok {
		c.claims[claim.Type] = []string{}
		claims, _ = c.claims[claim.Type]
	}
	claims = append(claims, claim.Value)
	c.claims[claim.Type] = claims
}

// GetClaims ...
func (c *claimsPrincipal) GetClaims() []claimsprincipalContracts.Claim {
	var result []claimsprincipalContracts.Claim
	for claimType, claimValues := range c.claims {
		for _, claimValue := range claimValues {
			result = append(result, claimsprincipalContracts.Claim{
				Type: claimType, Value: claimValue,
			})
		}
	}
	return result
}
