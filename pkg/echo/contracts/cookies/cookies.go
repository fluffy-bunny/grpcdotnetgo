package cookies

import (
	"time"
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
		SetCookieValue(name string, value string, expires time.Time) error
		GetCookieValue(name string) (string, error)
		DeleteCookie(name string) error
		RefreshCookie(name string, durration time.Duration) error
	}
)
