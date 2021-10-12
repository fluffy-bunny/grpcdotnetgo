package claimsprincipal

// Claim type
type Claim struct {
	Type  string
	Value string
}

// IClaimsPrincipal interface
type IClaimsPrincipal interface {
	GetClaims() []Claim
	HasClaim(claim Claim) bool
	AddClaim(claim Claim)
	RemoveClaim(claim Claim)
}
