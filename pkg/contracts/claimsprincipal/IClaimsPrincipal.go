package claimsprincipal

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClaimsPrincipal"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IClaimsPrincipal

const (
	// ClaimTypeAndValue ...
	ClaimTypeAndValue ClaimFactDirective = 0
	// ClaimType ...
	ClaimType = 1
)

type (
	// ClaimFactDirective tells if we want only the type validated vs type and value
	ClaimFactDirective int64
	// Claim ...
	Claim struct {
		Type  string `json:"type" mapstructure:"TYPE"`
		Value string `json:"value" mapstructure:"VALUE"`
	}
	// ClaimFact used for authorization
	ClaimFact struct {
		Claim     Claim
		Directive ClaimFactDirective
	}
	// IClaimsPrincipal interface
	IClaimsPrincipal interface {
		GetClaims() []Claim
		HasClaim(claim Claim) bool
		AddClaim(claim ...Claim)
		RemoveClaim(claim ...Claim)
		GetClaimsByType(claimType string) []Claim
		HasClaimType(claimType string) bool
	}
)
