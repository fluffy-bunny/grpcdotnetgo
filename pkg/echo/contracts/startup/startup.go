package startup

import (
	"net"

	core_contracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type (
	// Options for echo apps
	Options struct {
		Listener net.Listener
	}

	// IStartup for echo apps
	IStartup interface {
		SetContainer(container di.Container)
		GetOptions() *Options
		GetConfigOptions() *core_contracts.ConfigOptions
		GetPort() int
		ConfigureServices(builder *di.Builder) error
		Configure(e *echo.Echo, root di.Container) error
		GetSessionStore() sessions.Store
		// RegisterStaticRoutes
		// i.e. e.Static("/css", "./css")
		RegisterStaticRoutes(e *echo.Echo) error
	}
)
