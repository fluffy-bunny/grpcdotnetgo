package cookies

import (
	"time"

	"github.com/labstack/echo/v4"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE      gen "InterfaceType=ISecureCookie"
//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/func-types.go      -out=gen-func-$GOFILE gen "FuncType=SecureCookieConfigAccessor"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/$GOPACKAGE ISecureCookie

type (
	// SecureCookieConfig ...
	SecureCookieConfig struct {
		SecureCookieHashKey       string
		SecureCookieEncryptionKey string
	}
	// SecureCookieConfigAccessor func in the DI
	SecureCookieConfigAccessor func() *SecureCookieConfig
	// ISecureCookie ...
	ISecureCookie interface {
		SetCookieValue(c echo.Context, name string, value string, expires time.Time) error
		GetCookieValue(c echo.Context, name string) (string, error)
		DeleteCookie(c echo.Context, name string) error
		RefreshCookie(c echo.Context, name string, durration time.Duration) error
	}
)
