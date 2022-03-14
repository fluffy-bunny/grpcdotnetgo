package lambda

import (
	contracts_lambda "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/lambda"
	di "github.com/fluffy-bunny/sarulabsdi"
	uuid "github.com/satori/go.uuid"
)

// AddGenerateSatoriUUIDFunc adds a singleton of Now to the container
func AddGenerateSatoriUUIDFunc(builder *di.Builder) {
	contracts_lambda.AddGenerateUUIDFunc(builder, generateSatoriUUID)
}

func generateSatoriUUID() string {
	// Creating UUID Version 4
	// panic on error
	u4 := uuid.NewV4()
	return u4.String()
}
