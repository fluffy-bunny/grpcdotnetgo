package services

import di "github.com/fluffy-bunny/sarulabsdi"

var (
	ReflectTypeISomething = di.GetInterfaceReflectType((*ISomething)(nil))
)

type ISomething interface {
	GetName() string
	SetName(name string)
}
