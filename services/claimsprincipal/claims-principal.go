package claimsprincipal

type Claim struct {
	Type  string
	Value string
}
type IClaimsPrincipal interface {
	GetClaims() []Claim
	HasClaim(claim Claim) bool
	AddClaim(claim Claim)
	RemoveClaim(claim Claim)
}
type claimsPrincipal struct {
	claims map[string][]string
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func (c *claimsPrincipal) RemoveClaim(claim Claim) {
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

func (c *claimsPrincipal) HasClaim(claim Claim) bool {
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
func (c *claimsPrincipal) AddClaim(claim Claim) {
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
}

func (c *claimsPrincipal) GetClaims() []Claim {
	var result []Claim
	for claimType, claimValues := range c.claims {
		for _, claimValue := range claimValues {
			result = append(result, Claim{
				Type: claimType, Value: claimValue,
			})
		}

	}
	return result
}
