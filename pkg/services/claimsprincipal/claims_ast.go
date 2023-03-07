package claimsprincipal

import (
	"fmt"
	"strings"

	contracts_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/auth"
	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	"github.com/rs/zerolog/log"
)

const (
	and contracts_auth.Operand = 1
	or  contracts_auth.Operand = 2
)

// ClaimsAST is a light-weight AST that allows for logical collections of claims to
// be defined and tested by GTM based services. Grouping is implicit in the tree's structure
// such that the root arrays form grouped AND operations, and branches are processed by
// their placement in the parent. For example:
// ```
//
//	ClaimsAST{
//		Values: []string{"A", "B"},
//		Or: []ClaimsAST{
//			{Values: []string{"C", "D"}},
//			{
//				Values: []string{"E", "F"},
//				And: []ClaimsAST{
//					{Values: []string{"G", "H"}},
//				},
//			},
//		},
//		Not: []ClaimsAST{
//			{
//				Or: []ClaimsAST{
//					{Values: []string{"I", "J"}},
//				},
//			},
//		},
//	}
//
// ```
//
// Is the equivalent to:
// if A && B && ((C || D) || (E || F || (G && H))) && !(I || J)
type ClaimsAST struct {
	ClaimFacts []contracts_claimsprincipal.ClaimFact

	And []contracts_auth.IClaimsValidator
	Or  []contracts_auth.IClaimsValidator
	Not []contracts_auth.IClaimsValidator
}

// Validate the assumptions made in a Claims object
func (p *ClaimsAST) Validate(claimsPrincipal contracts_claimsprincipal.IClaimsPrincipal) bool {
	// Root is processed as an AND operation
	return p.validate(claimsPrincipal, and)
}

// ValidateWithOperand ...
func (p *ClaimsAST) ValidateWithOperand(claimsPrincipal contracts_claimsprincipal.IClaimsPrincipal, op contracts_auth.Operand) bool {
	return p.validate(claimsPrincipal, op)
}

func (p *ClaimsAST) validate(claimsPrincipal contracts_claimsprincipal.IClaimsPrincipal, op contracts_auth.Operand) bool {
	switch op {
	case and:
		// Return false on the first false, true if everything is true

		// Values
		for _, val := range p.ClaimFacts {
			switch val.Directive {
			case contracts_claimsprincipal.ClaimTypeAndValue:
				if !claimsPrincipal.HasClaim(val.Claim) {
					return false
				}
			case contracts_claimsprincipal.ClaimType:
				if !claimsPrincipal.HasClaimType(val.Claim.Type) {
					return false
				}
			}
		}

		// Ands
		for _, andVal := range p.And {
			if !andVal.ValidateWithOperand(claimsPrincipal, and) {
				return false
			}
		}

		// Ors
		for _, orVal := range p.Or {
			if !orVal.ValidateWithOperand(claimsPrincipal, or) {
				return false
			}
		}

		// Nots - processed with our op, but negated (we are and an, so fail on true)
		for _, notVal := range p.Not {
			if notVal.ValidateWithOperand(claimsPrincipal, op) {
				return false
			}
		}

		// All good
		return true
	case or:
		// Return true on the first true, false if everything is false

		// Values
		for _, val := range p.ClaimFacts {
			switch val.Directive {
			case contracts_claimsprincipal.ClaimTypeAndValue:
				if claimsPrincipal.HasClaim(val.Claim) {
					return true
				}
			case contracts_claimsprincipal.ClaimType:
				if claimsPrincipal.HasClaimType(val.Claim.Type) {
					return true
				}
			}
		}

		// Ands
		for _, andVal := range p.And {
			if andVal.ValidateWithOperand(claimsPrincipal, and) {
				return true
			}
		}

		// Ors
		for _, orVal := range p.Or {
			if orVal.ValidateWithOperand(claimsPrincipal, or) {
				return true
			}
		}

		// Nots - processed with our op, but negated (we are an or, so true on false)
		for _, notVal := range p.Not {
			if !notVal.ValidateWithOperand(claimsPrincipal, op) {
				return true
			}
		}

		// Nothing was true
		return false
	}

	log.Fatal().Int("op", int(op)).Msg("invalid operand")
	return false
}

// String ...
func (p *ClaimsAST) String() string {
	return p._string(and)
}

// StringWithOperand ...
func (p *ClaimsAST) StringWithOperand(op contracts_auth.Operand) string {
	return p._string(op)
}
func (p *ClaimsAST) _string(op contracts_auth.Operand) string {
	var groups []string

	// Values
	for _, claimFacts := range p.ClaimFacts {
		val := fmt.Sprintf("%s|%s", claimFacts.Claim.Type, claimFacts.Claim.Value)
		groups = append(groups, val)
	}

	// Ands
	for _, andVal := range p.And {
		groups = append(groups, andVal.StringWithOperand(and))
	}

	// Ors
	for _, orVal := range p.Or {
		groups = append(groups, orVal.StringWithOperand(or))
	}

	// Nots - processed with our op, but negated (we are and an, so fail on true)
	for _, notVal := range p.Not {
		groups = append(groups, "!"+notVal.StringWithOperand(op))
	}

	switch op {
	case and:
		return "(" + strings.Join(groups, " && ") + ")"
	case or:
		return "(" + strings.Join(groups, " || ") + ")"
	}

	log.Fatal().Int("op", int(op)).Msg("invalid operand")
	return ""
}
