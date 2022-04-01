package echo

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/labstack/echo/v4"
)

// HasWellknownAuthHeaders checks for wellknown auth headers
// 1 - Authorization
// 2 - X-Api-Key
func HasWellknownAuthHeaders(c echo.Context) bool {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if !utils.IsEmptyOrNil(authorizationHeader) {
		return true
	}
	authorizationHeader = c.Request().Header.Get("X-Api-Key")
	if !utils.IsEmptyOrNil(authorizationHeader) {
		return true
	}
	return false
}
