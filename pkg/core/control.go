package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	go_runtime "runtime"
	"strconv"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/reugn/async"
	zlog "github.com/rs/zerolog/log"
)

// Control is the control object that manages an echo web server
type Control struct {
	waitChannel chan os.Signal
	future      async.Future[string]
	e           *echo.Echo
	runtime     *Runtime
}

// NewControl creates a new control object
func NewControl(runtime *Runtime) *Control {
	return &Control{
		waitChannel: make(chan os.Signal),
		runtime:     runtime,
	}
}

// Stop ...
func (s *Control) Stop() {
	if s.future == nil {
		return
	}
	zlog.Info().Msg("Stopping Control Web Server")
	if err := s.e.Shutdown(context.Background()); err != nil {
		s.e.Logger.Error(err)
	}
	s.future.Join()
	zlog.Info().Msg("Control Web Server stopped")
}

// Start starts the echo web server using async and futures
func (s *Control) Start() {
	controlPort := os.Getenv("CONTROL_PORT")
	if len(controlPort) != 0 {
		// convert to int
		port, err := strconv.Atoi(controlPort)
		if err != nil {
			zlog.Fatal().Err(err).Msg("Failed to convert Control port to int")
		}
		// start the control server
		zlog.Info().Int("port", port).Msg("Starting Control server")

		s.e = echo.New()
		e := s.e
		e.Logger.SetLevel(log.DEBUG)
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello from Control")
		})
		e.GET("/stop", func(c echo.Context) error {
			s.runtime.Stop()
			return c.String(http.StatusOK, "Signalled server to stop")
		})
		e.GET("/gc", func(c echo.Context) error {
			go_runtime.GC()
			return c.String(http.StatusOK, "Called GC")
		})
		asyncAction := func() async.Future[string] {
			promise := async.NewPromise[string]()
			go func() {
				port := fmt.Sprintf(":%d", port)
				if err := e.Start(port); err != nil {
					e.Logger.Info("shutting down the server")
					promise.Success(utils.Ptr("OK"))
				} else {
					promise.Failure(err)
				}
			}()

			return promise.Future()
		}
		s.future = asyncAction()
	}
}
