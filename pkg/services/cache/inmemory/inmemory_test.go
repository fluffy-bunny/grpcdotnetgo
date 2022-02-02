package inmemory

import (
	"testing"
	"time"

	contracts_cache "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/cache"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/stretchr/testify/require"
)

func TestSameTypeAsScopedTransientSingleton(t *testing.T) {
	//var err error
	b, _ := di.NewBuilder()
	// order maters for Singleton and Transient, they are both app scoped and the last one wins
	AddSingletonIMemoryCache(b) // winner
	app := b.Build()

	meCache := contracts_cache.GetIMemoryCacheFromContainer(app)
	require.NotNil(t, meCache)

	val, err := meCache.Get("test")
	require.Error(t, err)
	require.Nil(t, val)

	meCache.SetWithTTL("test", "bob", time.Second)
	val, err = meCache.Get("test")
	require.NoError(t, err)
	require.Equal(t, "bob", val)
	time.Sleep(time.Second)

	val, err = meCache.Get("test")
	require.Error(t, err)
	require.Nil(t, val)

	val = meCache.GetOrInsert("dog", func() (interface{}, time.Duration, error) {
		return "Bowie", time.Second, nil
	})
	require.Equal(t, "Bowie", val)
	time.Sleep(time.Second)

	val, err = meCache.Get("dog")
	require.Error(t, err)
	require.Nil(t, val)
}
