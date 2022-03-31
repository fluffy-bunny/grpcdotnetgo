package runtime

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	core_contracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/core"
	contracts_container "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/container"
	contracts_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/handler"
	contracts_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/session"
	echo_contracts_startup "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/startup"
	middleware_container "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/middleware/container"
	middleware_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/middleware/logger"
	core_middleware_session "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/middleware/session"
	services_contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/services/contextaccessor"
	services_cookies "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/services/cookies"
	services_handler "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/services/handler"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/services/logger"
	core_echo_templates "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/templates"
	services_core_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	services_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/timeutils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/ziflex/lecho"
)

type (
	// Runtime ...
	Runtime struct {
		Startup       echo_contracts_startup.IStartup
		Container     di.Container
		e             *echo.Echo
		instanceID    string
		configOptions *core_contracts.ConfigOptions
	}
)

// New creates a new runtime
func New(startup echo_contracts_startup.IStartup) *Runtime {
	return &Runtime{
		Startup:    startup,
		instanceID: uuid.New().String(),
	}
}

// GetContainer returns the container
func (s *Runtime) GetContainer() di.Container {
	return s.Container
}
func (s *Runtime) phase1() error {
	s.configOptions = s.Startup.GetConfigOptions()
	err := core.LoadConfig(s.configOptions)
	if err != nil {
		return err
	}

	if s.configOptions.PrettyLog {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	switch strings.ToLower(s.configOptions.LogLevel) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
	return nil
}

func (s *Runtime) phase2() error {
	builder, _ := di.NewBuilder(di.App, di.Request, "transient")
	err := s.addDefaultServices(builder)
	if err != nil {
		log.Error().Err(err).Msg("Failed to add default services")
		return err
	}
	err = s.Startup.ConfigureServices(builder)
	if err != nil {
		log.Error().Err(err).Msg("Failed to configure services")
		return err
	}
	s.Container = builder.Build()
	s.Startup.SetContainer(s.Container)
	return nil
}
func (s *Runtime) phase3() error {
	s.e = echo.New()
	//use our own zerolog logger
	s.e.Logger = lecho.New(os.Stdout)
	//Set Renderer
	s.e.Renderer = core_echo_templates.GetTemplateRender("./templates")

	// MIDDELWARE
	//-------------------------------------------------------
	s.e.Use(middleware_logger.EnsureContextLogger(s.Container))
	s.e.Use(middleware_logger.EnsureContextLoggerCorrelation(s.Container))
	s.e.Use(middleware_container.EnsureScopedContainer(s.Container))
	sessionStore := contracts_session.GetGetSessionStoreFromContainer(s.Container)
	s.e.Use(session.Middleware(sessionStore()))
	mainSession := contracts_session.GetGetSessionFromContainer(s.Container)
	s.e.Use(core_middleware_session.EnsureSlidingSession(s.Container, mainSession))

	if s.configOptions.ApplicationEnvironment == "Development" {
		// this wipes out the session if we have a mismatch
		s.e.Use(core_middleware_session.EnsureDevelopmentSession(s.Container, mainSession, s.instanceID))
	}
	app := s.e.Group("")
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-Csrf-Token,form:csrf",
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	// we have all our required upfront middleware running
	// now we can add the optional startup ones.
	s.Startup.Configure(s.e, s.Container)

	// our middleware that runs at the end
	//-------------------------------------------------------
	s.e.Use(middleware.Recover())
	s.Startup.RegisterStaticRoutes(s.e)

	// register our handlers
	handlerFactory := contracts_handler.GetIHandlerFactoryFromContainer(s.Container)
	handlerFactory.RegisterHandlers(app)
	handlerDefinitions := contracts_handler.GetIHandlerDefinitions(s.Container)

	t := table.NewWriter()
	t.SetTitle("Handler Definitions")
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Verbs", "Path"})

	for _, handlerDefinition := range handlerDefinitions {
		metaData := handlerDefinition.MetaData
		httpVerbs, _ := metaData["httpVerbs"].([]contracts_handler.HTTPVERB)
		verbBldr := strings.Builder{}

		for idx, verb := range httpVerbs {
			verbBldr.WriteString(verb.String())
			if idx < len(httpVerbs)-1 {
				verbBldr.WriteString(",")
			}
		}
		path, _ := metaData["path"].(string)

		t.AppendRow([]interface{}{verbBldr.String(), string(path)})
	}
	t.Render()
	return nil
}
func (s *Runtime) finalPhase() error {
	// Finally start the server
	//----------------------------------------------------------------------------------
	startupOptions := s.Startup.GetOptions()

	address := fmt.Sprintf(":%v", startupOptions.Port)
	if startupOptions != nil && startupOptions.Listener != nil {
		// if we are here we are usually under test
		s.e.Listener = startupOptions.Listener
	}

	err := s.e.Start(address)
	if err != nil {
		log.Error().Err(err).Msg("failed to start server")
	}
	return err
}

// Run ...
func (s *Runtime) Run() error {
	// Phase 1
	// Load config
	// Setup Logger
	err := s.phase1()
	if err != nil {
		log.Fatal().Err(err).Msg("phase1")
	}
	// Phase 2
	// Setup our DI Container
	// Configure services
	err = s.phase2()
	if err != nil {
		log.Fatal().Err(err).Msg("phase2")
	}

	// Phase 2
	// Setup Echo
	// Configure middlewares
	err = s.phase3()
	if err != nil {
		log.Fatal().Err(err).Msg("phase3")
	}

	// Phase 2
	// Setup Echo
	// Configure middlewares
	err = s.finalPhase()
	if err != nil {
		log.Fatal().Err(err).Msg("finalPhase")
	}

	return err
}

func (s *Runtime) addDefaultServices(builder *di.Builder) error {
	contracts_container.AddContainerAccessorFunc(builder, s.GetContainer)
	services_timeutils.AddTimeNow(builder)
	services_timeutils.AddTimeParse(builder)
	services_cookies.AddSingletonISecureCookie(builder)
	services_contextaccessor.AddScopedIEchoContextAccessor(builder)
	services_logger.AddILogger(builder)
	services_core_claimsprincipal.AddScopedIClaimsPrincipal(builder)
	services_handler.AddSingletonIHandlerFactory(builder)
	return nil
}
