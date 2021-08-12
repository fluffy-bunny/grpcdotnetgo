package dicontext

import (
	"context"

	di "github.com/fluffy-bunny/sarulabsdi"
)

const ctxRequestContainer string = "ctx-request-container"

func GetRequestContainer(ctx context.Context) di.Container {
	requestContainer := ctx.Value(ctxRequestContainer).(di.Container)
	return requestContainer
}
func SetRequestContainer(ctx context.Context, requestContainer di.Container) context.Context {
	ctx = context.WithValue(ctx, ctxRequestContainer, requestContainer)
	return ctx
}
