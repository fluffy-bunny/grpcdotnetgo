package async

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsyncMap(t *testing.T) {
	pr := ExecuteAsync(func() (interface{}, error) {
		return true, nil
	})

	for {
		if pr.IsComplete() {
			break
		}
	}
	future := pr.Future
	v, err := future.Join()
	require.NoError(t, err)

	require.True(t, v.Value.(bool))
}
