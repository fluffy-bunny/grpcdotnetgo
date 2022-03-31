package session

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"

	contracts_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/session"
)

// EnsureSlidingSession ...
func EnsureSlidingSession(_ di.Container, getSession contracts_session.GetSession) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess := getSession(c)
			if !sess.IsNew {
				// we don't want to create a new session if nobody every created one before
				// we are only here to ensure that the session is an old one and slide it out.
				// i.e. bump out the expiration time
				sess.Save(c.Request(), c.Response())
			}
			return next(c)
		}
	}
}
