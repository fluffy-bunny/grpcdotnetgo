package logger

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/wellknown"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

// EnsureContextLoggerCorrelation ...
func EnsureContextLoggerCorrelation(_ di.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var loggerMap = make(map[string]string)
			headers := c.Request().Header

			// CORRELATION ID
			correlationID := headers.Get(wellknown.XCorrelationIDName)
			if core_utils.IsEmptyOrNil(correlationID) {
				correlationID = utils.GenerateUniqueID()
			}
			loggerMap["correlation_id"] = correlationID

			// SPANS
			span := headers.Get(wellknown.XSpanName)

			if !core_utils.IsEmptyOrNil(span) {
				loggerMap[wellknown.LogParentName] = span
				span = utils.GenerateUniqueID()
			}
			// generate a new span for this context
			newSpanID := utils.GenerateUniqueID()
			loggerMap[wellknown.LogSpanName] = newSpanID

			ctx := c.Request().Context()
			// add the correlation id to the context
			ctx = context.
				WithValue(ctx, wellknown.XCorrelationIDName, correlationID)
			ctx = context.
				WithValue(ctx, wellknown.XParentName, span)
			ctx = context.
				WithValue(ctx, wellknown.XSpanName, newSpanID)

			log := zerolog.Ctx(ctx)
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				for k, v := range loggerMap {
					c = c.Str(k, v)
				}
				return c
			})

			return next(c)
		}
	}
}
