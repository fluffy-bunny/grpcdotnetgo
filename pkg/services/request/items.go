package request

import (
	"reflect"

	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type serviceItems struct {
	items map[string]interface{}
}

func (s *serviceItems) Ctor() {
	s.items = make(map[string]interface{})
}

// AddScopedIItems adds service to the DI container
func AddScopedIItems(builder *di.Builder) {
	contracts_request.AddScopedIItems(builder, reflect.TypeOf(&serviceItems{}))
}
func (s *serviceItems) Set(key string, value interface{}) {
	s.items[key] = value
}
func (s *serviceItems) Get(key string) interface{} {
	obj, ok := s.items[key]
	if !ok {
		return nil
	}
	return obj
}
func (s *serviceItems) Delete(key string) {
	delete(s.items, key)
}
func (s *serviceItems) Clear() {
	s.items = make(map[string]interface{})
}
func (s *serviceItems) Keys() []string {
	keys := make([]string, 0)
	for key := range s.items {
		keys = append(keys, key)
	}
	return keys
}
func (s *serviceItems) GetItems() map[string]interface{} {
	return s.items
}
