package claimsprincipal

import (
	"fmt"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

type claimsPrincipal struct {
	claims  map[string][]string
	fastMap map[string]map[string]bool
}

// NewIClaimsPrincipal for outside of the DI
func NewIClaimsPrincipal() contracts_claimsprincipal.IClaimsPrincipal {
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
func (c *claimsPrincipal) RemoveClaim(claims ...contracts_claimsprincipal.Claim) {
	for _, claim := range claims {
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
}
func (c *claimsPrincipal) HasClaimType(claimType string) bool {
	_, ok := c.claims[claimType]
	return ok
}

func (c *claimsPrincipal) GetClaimsByType(claimType string) []contracts_claimsprincipal.Claim {
	claimParent, ok := c.claims[claimType]
	if !ok {
		return []contracts_claimsprincipal.Claim{}
	}
	var result []contracts_claimsprincipal.Claim
	for _, claimValue := range claimParent {
		result = append(result, contracts_claimsprincipal.Claim{
			Type:  claimType,
			Value: claimValue,
		})
	}
	return result
}

// HasClaim ...
func (c *claimsPrincipal) HasClaim(claim contracts_claimsprincipal.Claim) bool {
	claimParent, ok := c.fastMap[claim.Type]
	if !ok {
		return false
	}
	_, ok = claimParent[claim.Value]
	return ok
}

func (c *claimsPrincipal) addFastMapClaim(claim contracts_claimsprincipal.Claim) {
	claimParent, ok := c.fastMap[claim.Type]
	if !ok {
		claimParent = make(map[string]bool)
		c.fastMap[claim.Type] = claimParent
	}
	claimParent[claim.Value] = true
}

// AddClaim ...
func (c *claimsPrincipal) AddClaim(claims ...contracts_claimsprincipal.Claim) {
	for _, claim := range claims {
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
}

// GetClaims ...
func (c *claimsPrincipal) GetClaims() []contracts_claimsprincipal.Claim {
	var result []contracts_claimsprincipal.Claim
	for claimType, claimValues := range c.claims {
		for _, claimValue := range claimValues {
			result = append(result, contracts_claimsprincipal.Claim{
				Type: claimType, Value: claimValue,
			})
		}
	}
	return result
}

// NewBoolClaim ...
func NewBoolClaim(claimType string, value bool) contracts_claimsprincipal.Claim {
	return contracts_claimsprincipal.Claim{
		Type:  claimType,
		Value: fmt.Sprintf("%v", value),
	}
}

// NewStringClaim ...
func NewStringClaim(claimType string, value string) contracts_claimsprincipal.Claim {
	return contracts_claimsprincipal.Claim{
		Type:  claimType,
		Value: value,
	}
}

// NewFloat64Claim ...
func NewFloat64Claim(claimType string, value float64) contracts_claimsprincipal.Claim {
	return contracts_claimsprincipal.Claim{
		Type:  claimType,
		Value: fmt.Sprintf("%v", value),
	}
}

// ClaimsPrincipalFromClaimsMap ...
func ClaimsPrincipalFromClaimsMap(claimsMap map[string]interface{}) contracts_claimsprincipal.IClaimsPrincipal {
	principal := NewIClaimsPrincipal()
	for key, element := range claimsMap {
		switch value := element.(type) {
		case float64:
			principal.AddClaim(NewFloat64Claim(key, value))
		case bool:
			principal.AddClaim(NewBoolClaim(key, value))
		case string:
			principal.AddClaim(NewStringClaim(key, value))
		case []interface{}:
			for _, value := range value {
				switch claimValue := value.(type) {
				case float64:
					principal.AddClaim(NewFloat64Claim(key, claimValue))
				case bool:
					principal.AddClaim(NewBoolClaim(key, claimValue))
				case string:
					principal.AddClaim(NewStringClaim(key, claimValue))
				}
			}
		case []string:
			for _, value := range value {
				principal.AddClaim(contracts_claimsprincipal.Claim{
					Type:  key,
					Value: value,
				})
			}
		}
	}
	return principal
}
