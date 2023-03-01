package auth

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClaimsValidator"
//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE  go.mapped.dev/gtm/pkg/contracts/$GOPACKAGE IClaimsValidator
import (
	core_contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

type (
	Operand int

	IClaimsValidator interface {
		Validate(claimsPrincipal core_contracts_claimsprincipal.IClaimsPrincipal) bool
		ValidateWithOperand(claimsPrincipal core_contracts_claimsprincipal.IClaimsPrincipal, op Operand) bool
		String() string
		StringWithOperand(op Operand) string
	}
)
