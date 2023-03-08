package claimfact

import (
	"fmt"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

type (
	// Directive tells if we want only the type validated vs type and value
	Directive int64
)

const (
	// ClaimTypeAndValue ...
	ClaimTypeAndValue Directive = 0
	// ClaimType ...
	ClaimType = 1
)

// ClaimFact used for authorization
type ClaimFact struct {
	Claim     contracts_claimsprincipal.Claim
	Directive Directive
}

func init() {
	var _ contracts_claimsprincipal.IClaimFact = &ClaimFact{}
}

// NewClaimFact ...
func NewClaimFact(claim contracts_claimsprincipal.Claim) contracts_claimsprincipal.IClaimFact {
	return &ClaimFact{
		Claim:     claim,
		Directive: ClaimTypeAndValue,
	}
}
func NewClaimFactType(claimType string) contracts_claimsprincipal.IClaimFact {
	return &ClaimFact{
		Claim: contracts_claimsprincipal.Claim{
			Type: claimType,
		},
		Directive: contracts_claimsprincipal.ClaimType,
	}
}

// HasClaim ...
func (s *ClaimFact) HasClaim(claimsPrincipal contracts_claimsprincipal.IClaimsPrincipal) bool {
	if s.Directive == contracts_claimsprincipal.ClaimType {
		return claimsPrincipal.HasClaimType(s.Claim.Type)
	}
	return claimsPrincipal.HasClaim(s.Claim)
}

const (
	// ClaimTypeAndValueExpression ...
	ClaimTypeAndValueExpression = "has_claim(%s|%s)"
	// ClaimTypeExpression ...
	ClaimTypeExpression = "has_claim_type(%s)"
)

// Expression ...
func (s *ClaimFact) Expression() string {
	if s.Directive == contracts_claimsprincipal.ClaimType {
		return fmt.Sprintf(ClaimTypeExpression, s.Claim.Type)
	}
	return fmt.Sprintf(ClaimTypeAndValueExpression, s.Claim.Type, s.Claim.Value)
}
