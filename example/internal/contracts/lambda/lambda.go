package lambda

//go:generate genny   -pkg $GOPACKAGE     -in=../../../../genny/sarulabsdi/func-types.go -out=gen-func-$GOFILE gen "FuncType=GenerateUUID"

type (
	// GenerateUUID ...
	GenerateUUID func() string
)
