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

// ExecuteAsync returns a response that contains a future and a helper method to check if the future has been completed
func ExecuteAsync(f func() (interface{}, error)) *PromiseResponse[interface{}] {
	type internalFutureResponse struct {
		Err   error
		Value interface{}
	}
	p := async.NewPromise[interface{}]()
	done := make(chan struct{})
	go func() {
		if err := recover(); err != nil {
			p.Success(&internalFutureResponse{Value: nil, Err: fmt.Errorf("%v", err)})
		}
		value, err := f()
		p.Success(&internalFutureResponse{Value: value, Err: err})
	}()
	future := p.Future().Map(func(v interface{}) (interface{}, error) {
		response := v.(*internalFutureResponse)
		done <- struct{}{}
		return response.Value, response.Err
	})
	return &PromiseResponse[interface{}]{
		Future: future,
		done:   done,
	}
}

// ExecuteWithPromiseAsync ...
func ExecuteWithPromiseAsync(fn func(async.Promise[interface{}])) async.Future[interface{}] {
	promise := async.NewPromise[interface{}]()
	go fn(promise)
	return promise.Future()
}
