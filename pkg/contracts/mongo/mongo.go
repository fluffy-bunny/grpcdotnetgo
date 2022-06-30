package mongo

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IMongoConnectionStore"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IMongoConnectionStore

type (
	// ConnectionConfig ...
	ConnectionConfig struct {
		URI        string `json:"uri" mapstructure:"URI"`
		DB         string `json:"db" mapstructure:"DB"`
		Collection string `json:"collection" mapstructure:"COLLECTION"`
		Timeout    int64  `json:"timeout" mapstructure:"TIMEOUT"`
	}
	// IMongoConnectionConfigStore ...
	IMongoConnectionConfigStore interface {
		GetConnectionConfig() *ConnectionConfig
	}
)
