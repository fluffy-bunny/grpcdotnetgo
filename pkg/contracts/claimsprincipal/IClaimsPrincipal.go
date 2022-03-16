package claimsprincipal

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IClaimsPrincipal"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IClaimsPrincipal

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
	GetClaimsByType(claimType string) []Claim
	HasClaimType(claimType string) bool
}
