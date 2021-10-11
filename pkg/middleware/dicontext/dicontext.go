package dicontext

import (
	"context"

	di "github.com/fluffy-bunny/sarulabsdi"
)

const ctxRequestContainer string = "ctx-request-container"

// GetRequestContainer pulls the request container from the context
func GetRequestContainer(ctx context.Context) di.Container {
	val := ctx.Value(ctxRequestContainer)
	if val == nil {
		return nil
	}
	requestContainer := val.(di.Container)
	return requestContainer
}

// SetRequestContainer adds the request container to the context
func SetRequestContainer(ctx context.Context, requestContainer di.Container) context.Context {
	ctx = context.WithValue(ctx, ctxRequestContainer, requestContainer)
	return ctx
}
