package singleton

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ISingleton"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/$GOPACKAGE ISingleton

type (
	// ISingleton ...
	ISingleton interface {
		SetName(in string)
		GetName() string
	}
)
