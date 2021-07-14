package internal

import (
	"encoding/json"

	_ "github.com/fluffy-bunny/grpcdotnetgo" // ensure that go mod vendor brings everything
)

func PrettyJSON(obj interface{}) string {
	jsonBytes, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
