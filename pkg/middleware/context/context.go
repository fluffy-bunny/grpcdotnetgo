/*
There be a memory leak when calling WithTimeout and you don't call the cancel func.
Its really easy to forget to put the defer cancel() after the WithTimeout call.

Its easy to seach our code base for anyone calling WithTimeout and replacing it with NewContextWithTimeout.

Here we ensure that the cancel func is called when the request pipeline is complete.

https://developer.squareup.com/blog/always-be-closing/
*/
package context

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type (
	// WithTimeoutManager manages cancel funcs
	WithTimeoutManager struct {
		cancelFuncs []func()
	}
	contextKey struct{}
)

// NewContextWithTimeoutManager creates a new ContextWithTimeoutManager
func NewContextWithTimeoutManager() *WithTimeoutManager {
	return &WithTimeoutManager{
		cancelFuncs: make([]func(), 0),
	}
}

// addCancelFunc ContextWithTimeoutManager to a context
func (c *WithTimeoutManager) addCancelFunc(cancelFunc func()) {
	c.cancelFuncs = append(c.cancelFuncs, cancelFunc)
}

// CancelAll cancels all the cancel funcs
func (c *WithTimeoutManager) CancelAll() {
	for _, cancelFunc := range c.cancelFuncs {
		cancelFunc()
	}
}

var _contextKey = &contextKey{}

// SetContextWithTimeoutManager adds the ContextWithTimeoutManager to the context
func SetContextWithTimeoutManager(ctx context.Context, contextWithTimeoutManager *WithTimeoutManager) context.Context {
	ctx = context.WithValue(ctx, _contextKey, contextWithTimeoutManager)
	return ctx
}

// getContextWithTimeoutManager pulls the ContextWithTimeoutManager from the context
func getContextWithTimeoutManager(ctx context.Context) *WithTimeoutManager {
	val := ctx.Value(_contextKey)
	if val == nil {
		return nil
	}
	contextWithTimeoutManager, ok := val.(*WithTimeoutManager)
	if !ok || contextWithTimeoutManager == nil {
		log.Fatal().Msg("ContextWithTimeoutManager is not of in the context.  It needs to be put in there in an early middleware.  This is only meant for the Request Pipeline")
	}
	return contextWithTimeoutManager
}

// NewContextWithTimeout creates a new context with a timeout
func NewContextWithTimeout(ctx context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, duration)
	contextWithTimeoutManager := getContextWithTimeoutManager(ctx)

	contextWithTimeoutManager.addCancelFunc(cancel)
	return ctx, cancel
}
