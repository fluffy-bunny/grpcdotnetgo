package services

import di "github.com/fluffy-bunny/sarulabsdi"

// GetContextAccessorFromContainer from the Container
func GetISomethingFromContainer(ctn di.Container) ISomething {
	obj := ctn.GetByType(ReflectTypeISomething).(ISomething)
	return obj
}

// GetContextAccessorFromContainer from the Container
func GetISomethingsFromContainer(ctn di.Container) []ISomething {
	objs := ctn.GetManyByType(ReflectTypeISomething)
	var result []ISomething
	for _, i := range objs {
		result = append(result, i.(ISomething))
	}
	return result
}
