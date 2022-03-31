package utils

import "encoding/json"

// PrettyJSON returns a pretty-printed JSON string for the given object.
func PrettyJSON(obj interface{}) string {
	jsonBytes, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
