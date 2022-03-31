package handler

import (
	contracts_container "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/container"
	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	wellknown "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/wellknown"

	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type (
	service struct {
		ContainerAccessor contracts_container.ContainerAccessor `inject:""`
	}
)

func assertImplementation() {
	var _ contracts_handler.IHandlerFactory = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIHandlerFactory registers the *service as a singleton.
func AddSingletonIHandlerFactory(builder *di.Builder) {
	log.Info().Str("DI", "IHandlerFactory").Send()
	contracts_handler.AddSingletonIHandlerFactory(builder, reflectType)
}

func (s *service) RegisterHandlers(app *echo.Group) {
	rootContainer := s.ContainerAccessor()
	scopedContainer, _ := rootContainer.SubContainer()
	defs := contracts_handler.GetIHandlerDefinitions(scopedContainer)

	for _, def := range defs {
		path := def.MetaData["path"].(string)
		httpVerbs := def.MetaData["httpVerbs"].([]contracts_handler.HTTPVERB)
		defName := def.Name
		doFunc := func(c echo.Context) error {
			scopedContainer = c.Get(wellknown.SCOPED_CONTAINER_KEY).(di.Container)
			handlerInstance := di.Get(scopedContainer, defName).(contracts_handler.IHandler)
			return handlerInstance.Do(c)
		}

		for _, httpVerb := range httpVerbs {
			switch httpVerb {
			case contracts_handler.GET:
				app.GET(path, doFunc)
			case contracts_handler.POST:
				app.POST(path, doFunc)
			case contracts_handler.PUT:
				app.PUT(path, doFunc)
			case contracts_handler.DELETE:
				app.DELETE(path, doFunc)
			case contracts_handler.PATCH:
				app.PATCH(path, doFunc)
			case contracts_handler.HEAD:
				app.HEAD(path, doFunc)
			case contracts_handler.OPTIONS:
				app.OPTIONS(path, doFunc)
			case contracts_handler.CONNECT:
				app.CONNECT(path, doFunc)
			case contracts_handler.TRACE:
				app.TRACE(path, doFunc)
			}
		}
		log.Info().Str("echo", "RegisterHandlers").Str("path", path).Send()
	}
}
