package scoped

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IScoped"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/$GOPACKAGE IScoped

type (
	// IScoped ...
	IScoped interface {
		SetName(in string)
		GetName() string
	}
)
