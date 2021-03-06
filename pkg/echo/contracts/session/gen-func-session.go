// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package session

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeGetSession used when your service claims to implement GetSession
var ReflectTypeGetSession = reflect.TypeOf(GetSession(nil))

// AddSingletonGetSessionFunc adds a func to the DI
func AddGetSessionFunc(builder *di.Builder, fnc GetSession) {
	di.AddFunc(builder, fnc)
}

// RemoveAllGetSessionFunc removes all GetSession functions from the DI
func RemoveAllGetSessionFunc(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeGetSession)
}

// GetGetSessionFromContainer alternative to SafeGetGetSessionFromContainer but panics of object is not present
func GetGetSessionFromContainer(ctn di.Container) GetSession {
	return ctn.GetByType(ReflectTypeGetSession).(GetSession)
}

// GetManyGetSessionFromContainer alternative to SafeGetManyGetSessionFromContainer but panics of object is not present
func GetManyGetSessionFromContainer(ctn di.Container) []GetSession {
	objs := ctn.GetManyByType(ReflectTypeGetSession)
	var results []GetSession
	for _, obj := range objs {
		results = append(results, obj.(GetSession))
	}
	return results
}

// SafeGetGetSessionFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetGetSessionFromContainer(ctn di.Container) (GetSession, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeGetSession)
	if err != nil {
		return nil, err
	}
	return obj.(GetSession), nil
}

// SafeGetManyGetSessionFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyGetSessionFromContainer(ctn di.Container) ([]GetSession, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeGetSession)
	if err != nil {
		return nil, err
	}
	var results []GetSession
	for _, obj := range objs {
		results = append(results, obj.(GetSession))
	}
	return results, nil
}

// ReflectTypeGetSessionStore used when your service claims to implement GetSessionStore
var ReflectTypeGetSessionStore = reflect.TypeOf(GetSessionStore(nil))

// AddSingletonGetSessionStoreFunc adds a func to the DI
func AddGetSessionStoreFunc(builder *di.Builder, fnc GetSessionStore) {
	di.AddFunc(builder, fnc)
}

// RemoveAllGetSessionStoreFunc removes all GetSessionStore functions from the DI
func RemoveAllGetSessionStoreFunc(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeGetSessionStore)
}

// GetGetSessionStoreFromContainer alternative to SafeGetGetSessionStoreFromContainer but panics of object is not present
func GetGetSessionStoreFromContainer(ctn di.Container) GetSessionStore {
	return ctn.GetByType(ReflectTypeGetSessionStore).(GetSessionStore)
}

// GetManyGetSessionStoreFromContainer alternative to SafeGetManyGetSessionStoreFromContainer but panics of object is not present
func GetManyGetSessionStoreFromContainer(ctn di.Container) []GetSessionStore {
	objs := ctn.GetManyByType(ReflectTypeGetSessionStore)
	var results []GetSessionStore
	for _, obj := range objs {
		results = append(results, obj.(GetSessionStore))
	}
	return results
}

// SafeGetGetSessionStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetGetSessionStoreFromContainer(ctn di.Container) (GetSessionStore, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeGetSessionStore)
	if err != nil {
		return nil, err
	}
	return obj.(GetSessionStore), nil
}

// SafeGetManyGetSessionStoreFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyGetSessionStoreFromContainer(ctn di.Container) ([]GetSessionStore, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeGetSessionStore)
	if err != nil {
		return nil, err
	}
	var results []GetSessionStore
	for _, obj := range objs {
		results = append(results, obj.(GetSessionStore))
	}
	return results, nil
}
