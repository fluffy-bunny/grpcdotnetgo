package claimsprincipal

import (
	"fmt"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

type claimsPrincipal struct {
	claims  map[string][]string
	fastMap map[string]map[string]bool
}

// NewIClaimsPrincipal for outside of the DI
func NewIClaimsPrincipal() claimsprincipalContracts.IClaimsPrincipal {
	obj := &claimsPrincipal{}
	obj.Ctor()
	return obj
}

func (c *claimsPrincipal) Ctor() {
	c.claims = make(map[string][]string)
	c.fastMap = make(map[string]map[string]bool)
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
	claimParent, ok := c.fastMap[claim.Type]
	if !ok {
		return false
	}
	_, ok = claimParent[claim.Value]
	return ok
}

func (c *claimsPrincipal) addFastMapClaim(claim claimsprincipalContracts.Claim) {
	claimParent, ok := c.fastMap[claim.Type]
	if !ok {
		claimParent = make(map[string]bool)
		c.fastMap[claim.Type] = claimParent
	}
	claimParent[claim.Value] = true
}

// AddClaim ...
func (c *claimsPrincipal) AddClaim(claim claimsprincipalContracts.Claim) {
	if len(claim.Type) == 0 {
		return
	}
	if c.HasClaim(claim) {
		return
	}

	claims, ok := c.claims[claim.Type]
	if !ok {
		c.claims[claim.Type] = []string{}
		claims, _ = c.claims[claim.Type]
	}
	claims = append(claims, claim.Value)
	c.claims[claim.Type] = claims
	c.addFastMapClaim(claim)
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
