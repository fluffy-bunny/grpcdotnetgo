package pkg

var fullMethodNameToMap = make(map[string]interface{})

func AddFullMethodNameSliceToMap(slice []string) {
	for _, v := range slice {
		fullMethodNameToMap[v] = make(map[string]interface{})
	}
}
func NewFullMethodNameToMap(objMaker func(fullMethodName string) interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	for k := range fullMethodNameToMap {
		result[k] = objMaker(k)
	}
	return result
}
