package startup

import (
	core_contracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	// Options for echo apps
	Options struct {
		Port int
	}
	// Hooks into the startup pipeline
	Hooks struct {
		PrebuildHook    func(builder *di.Builder) error
		PostBuildHook   func(container di.Container) error
		PreStartHook    func(echo *echo.Echo) error
		PreShutdownHook func(echo *echo.Echo) error
	}
	// IStartup for echo apps
	IStartup interface {
		// Config
		// SetHooks lets us add services at that end of the main ConfigServices chain
		// Typically used for unit testing where mocks are swapped in.
		AddHooks(hooks ...*Hooks)
		GetHooks() []*Hooks

		GetContainer() di.Container

		// In order of execution

		// 1. GetConfigOptions
		GetConfigOptions() *core_contracts.ConfigOptions
		// 2.a ConfigureServices
		// 2.b Call PreBuildHook if it is present
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
	// CommonStartup ...
	CommonStartup struct {
		hooks     []*Hooks
		container di.Container
	}
)

// AddHooks setter
func (s *CommonStartup) AddHooks(hooks ...*Hooks) {
	s.hooks = append(s.hooks, hooks...)
}

// GetHooks getter
func (s *CommonStartup) GetHooks() []*Hooks {
	return s.hooks
}

// SetContainer setter
func (s *CommonStartup) SetContainer(container di.Container) {
	s.container = container
}

// GetContainer setter
func (s *CommonStartup) GetContainer() di.Container {
	return s.container
}
