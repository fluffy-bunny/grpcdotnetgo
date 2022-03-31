package startup

import (
	"net"

	core_contracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	// Options for echo apps
	Options struct {
		Listener net.Listener
		Port     int
	}

	// IStartup for echo apps
	IStartup interface {
		// In order of execution

		// 1. GetConfigOptions
		GetConfigOptions() *core_contracts.ConfigOptions
		// 2. ConfigureServices
		ConfigureServices(builder *di.Builder) error
		// 3. SetContainer
		SetContainer(container di.Container)
		// 4. Configure
		Configure(e *echo.Echo, root di.Container) error
		// 5. RegisterStaticRoutes
		// i.e. e.Static("/css", "./css")
		RegisterStaticRoutes(e *echo.Echo) error
		// 6. GetOptions
		GetOptions() *Options
	}
)
