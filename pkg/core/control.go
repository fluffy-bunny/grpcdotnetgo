package core

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/reugn/async"
)

// Control is the control object that manages an echo web server
type Control struct {
	waitChannel chan os.Signal
	future      async.Future[string]
	e           *echo.Echo
	runtime     *Runtime
	port        int
}

// NewControl creates a new control object
func NewControl(runtime *Runtime, port int) *Control {
	return &Control{
		waitChannel: make(chan os.Signal),
		runtime:     runtime,
		port:        port,
	}
}

// Stop ...
func (s *Control) Stop() {
	if err := s.e.Shutdown(context.Background()); err != nil {
		s.e.Logger.Error(err)
	}
	s.future.Join()
}

// Start starts the echo web server using async and futures
func (s *Control) Start() {
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
	asyncAction := func() async.Future[string] {
		promise := async.NewPromise[string]()
		go func() {
			port := fmt.Sprintf(":%d", s.port)
			if err := e.Start(port); err != nil {
				e.Logger.Info("shutting down the server")
				promise.Success("OK")
			} else {
				promise.Failure(err)
			}
		}()

		return promise.Future()
	}
	s.future = asyncAction()
}
