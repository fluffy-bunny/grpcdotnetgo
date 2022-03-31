package handler

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IHandler,IHandlerFactory"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/$GOPACKAGE IHandler,IHandlerFactory
// HTTPVERB is a list of HTTP verbs
type HTTPVERB uint

const (
	GET     HTTPVERB = 0
	POST             = 1
	PUT              = 2
	DELETE           = 3
	PATCH            = 4
	OPTIONS          = 5
	HEAD             = 6
	CONNECT          = 7
	TRACE            = 8
)

func (s HTTPVERB) String() string {
	switch s {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case PATCH:
		return "PATCH"
	case OPTIONS:
		return "OPTIONS"
	case HEAD:
		return "HEAD"
	case CONNECT:
		return "CONNECT"
	case TRACE:
		return "TRACE"
	}
	return "unknown"
}

type (
	// IHandler ...
	IHandler interface {
		GetMiddleware() []echo.MiddlewareFunc
		Do(c echo.Context) error
	}
	// IHandlerFactory ...
	IHandlerFactory interface {
		RegisterHandlers(app *echo.Group)
	}
)

// AddScopedIHandlerEx ...
func AddScopedIHandlerEx(builder *di.Builder, reflectType reflect.Type, httpVerbs []HTTPVERB, path string) {
	httpVerbS := []string{}
	for _, httpVerb := range httpVerbs {
		httpVerbS = append(httpVerbS, httpVerb.String())
	}
	metadata := map[string]interface{}{
		"path":      path,
		"httpVerbs": httpVerbs,
	}
	log.Info().
		Str("DI", "IHandler").
		Str("path", path).
		Str("httpVerbs", strings.Join(httpVerbS, "|")).Send()
	AddScopedIHandlerWithMetadata(builder, reflectType, metadata)
}
