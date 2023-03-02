package claimsprincipal

import (
	"testing"

	contracts_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/auth"
	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"

	"github.com/stretchr/testify/assert"
)

func TestClaimsEmpty(t *testing.T) {
	// ()
	perms := ClaimsAST{}

	assert.Equal(t, "()", perms.String())
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D")))
}

func TestClaimsRootOnly(t *testing.T) {
	// (A && B)
	perms := ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "A",
			},
			{
				Type:  "permissions",
				Value: "B",
			},
		},
	}

	assert.Equal(t, "(permissions|A && permissions|B)", perms.String())
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("B", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "C", "D")))
}

func TestClaimsBranchAnd(t *testing.T) {
	// ((A && B))
	var ands []contracts_auth.IClaimsValidator
	ands = append(ands, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "A",
			},
			{
				Type:  "permissions",
				Value: "B",
			},
		},
	})
	perms := ClaimsAST{
		And: ands,
	}

	assert.Equal(t, "((permissions|A && permissions|B))", perms.String())
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("B", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "C", "D")))

	_ = perms
}

func TestClaimsBranchOr(t *testing.T) {
	// ((A || B))
	var ors []contracts_auth.IClaimsValidator
	ors = append(ors, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "A",
			},
			{
				Type:  "permissions",
				Value: "B",
			},
		},
	})
	perms := ClaimsAST{
		Or: ors,
	}

	assert.Equal(t, "((permissions|A || permissions|B))", perms.String())
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("B", "C", "D")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("X", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("C", "D")))
}

func TestClaimsBranchNot(t *testing.T) {
	// Not inherits the operand from it's parent, and in this case
	// (!(A && B))
	var nots []contracts_auth.IClaimsValidator
	nots = append(nots, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "A",
			},
			{
				Type:  "permissions",
				Value: "B",
			},
		},
	})
	perms := ClaimsAST{
		Not: nots,
	}

	assert.Equal(t, "(!(permissions|A && permissions|B))", perms.String())
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("C", "D", "E", "F")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "C", "D")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("B", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B")))
}

func TestClaimsBranchNotNested(t *testing.T) {
	// (!((A || B)))
	var ors []contracts_auth.IClaimsValidator
	ors = append(ors, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "A",
			},
			{
				Type:  "permissions",
				Value: "B",
			},
		},
	})

	var nots []contracts_auth.IClaimsValidator
	nots = append(nots, &ClaimsAST{
		Or: ors,
	})
	perms := ClaimsAST{
		Not: nots,
	}

	assert.Equal(t, "(!((permissions|A || permissions|B)))", perms.String())
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("C", "D", "E", "F")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("B", "C", "D")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B")))
}

func TestClaimsDocSample(t *testing.T) {
	// (A && B && (C || D) && (E || F || (G && H)) && !((I && J)))
	var ands1 []contracts_auth.IClaimsValidator
	ands1 = append(ands1, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "G",
			},
			{
				Type:  "permissions",
				Value: "H",
			},
		},
	})
	var ors []contracts_auth.IClaimsValidator
	ors = append(ors, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "C",
			},
			{
				Type:  "permissions",
				Value: "D",
			},
		},
	})
	ors = append(ors, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "E",
			},
			{
				Type:  "permissions",
				Value: "F",
			},
		},
		And: ands1,
	})

	var ands2 []contracts_auth.IClaimsValidator
	ands2 = append(ands2, &ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "I",
			},
			{
				Type:  "permissions",
				Value: "J",
			},
		},
	})
	var nots []contracts_auth.IClaimsValidator
	nots = append(nots, &ClaimsAST{
		And: ands2,
	})
	perms := ClaimsAST{
		Claims: []contracts_claimsprincipal.Claim{
			{
				Type:  "permissions",
				Value: "A",
			},
			{
				Type:  "permissions",
				Value: "B",
			},
		},
		Or:  ors,
		Not: nots,
	}
	/*
		perms := Claims{
			Values: []string{"A", "B"},
			Or: []*Claims{
				{Values: []string{"C", "D"}},
				{
					Values: []string{"E", "F"},
					And: []*Claims{
						{Values: []string{"G", "H"}},
					},
				},
			},
			Not: []*Claims{
				{
					And: []*Claims{
						{Values: []string{"I", "J"}},
					},
				},
			},
		}
	*/
	assert.Equal(t, "(permissions|A && permissions|B && (permissions|C || permissions|D) && (permissions|E || permissions|F || (permissions|G && permissions|H)) && !((permissions|I && permissions|J)))", perms.String())

	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "E", "F", "G", "H")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "E", "F", "G", "H", "I")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "E", "F", "G", "H", "J")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "E", "G", "H")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "E")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "D", "F")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "G", "H")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "F", "G", "H", "I")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "E", "G", "H")))
	assert.True(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "D", "F", "G", "H", "J")))

	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("B", "C", "D", "E", "F", "G", "H")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "C", "D", "E", "F", "G", "H")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "E", "F", "G", "H", "I")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "H", "J")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "H")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "G")))
	assert.False(t, perms.Validate(NewmockClaimsPrincipalToken("A", "B", "C", "D", "F", "G", "H", "I", "J")))
}
func NewmockClaimsPrincipalToken(perms ...string) contracts_claimsprincipal.IClaimsPrincipal {

	cp := NewIClaimsPrincipal()

	for _, perm := range perms {
		claim := contracts_claimsprincipal.Claim{
			Type:  "permissions",
			Value: perm,
		}
		cp.AddClaim(claim)
	}
	return cp
}
