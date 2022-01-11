package inmemory

import (
	"reflect"
	"time"

	contracts_cache "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/cache"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gookit/cache"
	"github.com/gookit/cache/gocache"
)

type (
	service struct {
		theCache *gocache.GoCache
	}
)

// ReflectTypeService returns the service type
var ReflectTypeService = reflect.TypeOf(&service{})

func (s *service) Ctor() {
	s.theCache = gocache.NewGoCache(cache.OneDay, cache.FiveMinutes)
}

// AddSingletonIMemoryCache adds service to the DI container
func AddSingletonIMemoryCache(builder *di.Builder) {
	contracts_cache.AddSingletonIMemoryCache(builder, ReflectTypeService)
}
func (s *service) Clear() error {
	return s.theCache.Clear()
}

// Has basic operation
func (s *service) Has(key string) bool {
	return s.theCache.Has(key)
}
func (s *service) Del(key string) error {
	return s.theCache.Del(key)
}
func (s *service) Get(key string) interface{} {
	return s.theCache.Get(key)
}
func (s *service) Set(key string, val interface{}, ttl time.Duration) error {
	return s.theCache.Set(key, val, ttl)
}

// GetMulti multi operation
func (s *service) GetMulti(keys []string) map[string]interface{} {
	return s.theCache.GetMulti(keys)
}
func (s *service) SetMulti(values map[string]interface{}, ttl time.Duration) error {
	return s.theCache.SetMulti(values, ttl)
}
func (s *service) DelMulti(keys []string) error {
	return s.theCache.DelMulti(keys)
}
func (s *service) Close() error {
	return s.theCache.Close()
}
func (s *service) GetOrInsert(k string, adder func() (interface{}, time.Duration, error)) interface{} {
	result := s.Get(k)
	if result == nil {
		obj, ttl, err := adder()
		if err != nil {
			return nil
		}
		s.Set(k, obj, ttl)
		result = obj
	}
	return result
}
