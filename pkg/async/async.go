package async

import (
	"fmt"

	"github.com/reugn/async"
)

type (
	// AsyncResponse ...
	AsyncResponse struct {
		Message string
		Error   error
	}
	// PromiseResponse ...
	PromiseResponse[T any] struct {
		Future async.Future[T]
		done   chan struct{}
	}
)

// IsComplete checks to see if a promise has been completed
func (s *PromiseResponse[T]) IsComplete() bool {
	select {
	case <-s.done:
		return true
	default:
		return false
	}
}

// FutureResponse ...
type FutureResponse struct {
	Err   error
	Value interface{}
}

// ExecuteAsync returns a response that contains a future and a helper method to check if the future has been completed
func ExecuteAsync(f func() (interface{}, error)) *PromiseResponse[*FutureResponse] {
	p := async.NewPromise[*FutureResponse]()
	done := make(chan struct{})
	go func() {
		if err := recover(); err != nil {
			p.Success(&FutureResponse{Value: nil, Err: fmt.Errorf("%v", err)})
		}
		value, err := f()
		p.Success(&FutureResponse{Value: value, Err: err})
	}()
	future := p.Future().Map(func(v *FutureResponse) (*FutureResponse, error) {
		response := v
		done <- struct{}{}
		return response, response.Err
	})
	return &PromiseResponse[*FutureResponse]{
		Future: future,
		done:   done,
	}
}

// ExecuteWithPromiseAsync ...
func ExecuteWithPromiseAsync[T any](fn func(async.Promise[T])) async.Future[T] {
	promise := async.NewPromise[T]()
	go fn(promise)
	return promise.Future()
}
