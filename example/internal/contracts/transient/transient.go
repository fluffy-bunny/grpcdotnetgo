package transient

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITransient"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/$GOPACKAGE ITransient

type (
	// ITransient ...
	ITransient interface {
		SetName(in string)
		GetName() string
	}
)
