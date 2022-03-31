package session

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"

	contracts_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/session"
)

// EnsureDevelopmentSession is a middleware that ensures that the session is
// wiped out when the app restarts
func EnsureDevelopmentSession(_ di.Container, getSession contracts_session.GetSession, appInstanceID string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess := getSession(c)
			appInstanceValue, ok := sess.Values["_appInstanceID"]
			if !ok {
				sess.Values["_appInstanceID"] = appInstanceID
				sess.Save(c.Request(), c.Response())
			} else {
				if appInstanceValue != appInstanceID {
					sess.Values = make(map[interface{}]interface{}) // wipe out the session
					sess.Values["_appInstanceID"] = appInstanceID
					sess.Save(c.Request(), c.Response())
				}
			}
			return next(c)
		}
	}
}
