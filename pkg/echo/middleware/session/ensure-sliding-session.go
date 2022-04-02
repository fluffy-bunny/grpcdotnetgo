package session

import (
	"net/http"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	contracts_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/session"
)

// EnsureSlidingSession ...
func EnsureSlidingSession(_ di.Container, getSession contracts_session.GetSession) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					log.Error().Msg("Panic in EnsureSlidingSession")
					cookies := c.Cookies()
					for _, cookie := range cookies {
						cookie.MaxAge = -1
						c.SetCookie(cookie)
					}
					c.Redirect(http.StatusFound, "/")
				}
			}()
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
