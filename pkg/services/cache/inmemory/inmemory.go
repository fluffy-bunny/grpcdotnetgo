package inmemory

import (
	"reflect"
	"time"

	ttlcache "github.com/ReneKroon/ttlcache/v2"
	contracts_cache "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/cache"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type (
	service struct {
		ttlCache ttlcache.SimpleCache
	}
)

func assertImplementation() {
	var _ contracts_cache.IMemoryCache = (*service)(nil)
}

// ReflectTypeService returns the service type
var ReflectTypeService = reflect.TypeOf(&service{})

func (s *service) Ctor() {
	s.ttlCache = ttlcache.NewCache()
	s.ttlCache.SetTTL(contracts_cache.Forever)
}

// AddSingletonIMemoryCache adds service to the DI container
func AddSingletonIMemoryCache(builder *di.Builder) {
	log.Info().Msg("IoC: AddSingletonIMemoryCache")
	contracts_cache.AddSingletonIMemoryCache(builder, ReflectTypeService)
}

func (s *service) Get(key string) (interface{}, error) {
	return s.ttlCache.Get(key)
}

func (s *service) GetWithTTL(key string) (interface{}, time.Duration, error) {
	return s.ttlCache.GetWithTTL(key)
}

func (s *service) Set(key string, data interface{}) error {
	return s.ttlCache.Set(key, data)
}

func (s *service) SetTTL(ttl time.Duration) error {
	return s.ttlCache.SetTTL(ttl)
}

func (s *service) SetWithTTL(key string, data interface{}, ttl time.Duration) error {
	return s.ttlCache.SetWithTTL(key, data, ttl)
}

func (s *service) Remove(key string) error {
	return s.ttlCache.Remove(key)
}

func (s *service) Close() error {
	return s.ttlCache.Close()
}

func (s *service) Purge() error {
	return s.ttlCache.Purge()
}

func (s *service) GetOrInsert(k string, adder func() (interface{}, time.Duration, error)) interface{} {
	result, err := s.Get(k)
	if err != nil || result == nil {
		obj, ttl, err := adder()
		if err != nil {
			return nil
		}
		s.SetWithTTL(k, obj, ttl)
		result = obj
	}
	return result
}
