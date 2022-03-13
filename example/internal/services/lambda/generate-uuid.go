package lambda

import (
	contracts_lambda "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/lambda"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/google/uuid"
)

// AddGenerateUUIDFunc adds a singleton of Now to the container
func AddGenerateUUIDFunc(builder *di.Builder) {
	contracts_lambda.AddGenerateUUIDFunc(builder, generateUUID)
}

func generateUUID() string {
	uuidWithHyphen := uuid.New()
	return uuidWithHyphen.String()
}
