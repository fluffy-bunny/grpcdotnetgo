package claimfact

import (
	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

type (
	// IClaimFact interface
	IClaimFact interface {
		HasClaim(claimsprincipal contracts_claimsprincipal.IClaimsPrincipal) bool
		Expression() string
	}
)
