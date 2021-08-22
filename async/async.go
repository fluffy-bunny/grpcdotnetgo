package async

import (
	"github.com/reugn/async"
)

type AsyncResponse struct {
	Message string
	Error   error
}

func ExecuteWithPromiseAsync(fn func(async.Promise)) async.Future {
	promise := async.NewPromise()
	go fn(promise)
	return promise.Future()
}
