package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"net/http/pprof"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	zlog "github.com/rs/zerolog/log"

	"github.com/reugn/async"
)

// PProfServer is the PProfServer object that manages an echo web server
type PProfServer struct {
	waitChannel chan os.Signal
	future      async.Future[string]
	e           *echo.Echo
}

// NewPProfServer creates a new PProf object
func NewPProfServer() *PProfServer {
	return &PProfServer{
		waitChannel: make(chan os.Signal),
	}
}

// Stop ...
func (s *PProfServer) Stop() {
	if s.future == nil {
		return
	}
	zlog.Info().Msg("Stopping PProf Web Server")

	if err := s.e.Shutdown(context.Background()); err != nil {
		s.e.Logger.Error(err)
	}
	s.future.Join()
	zlog.Info().Msg("PProf Web Server stopped")
}

// Start starts the echo web server using async and futures
func (s *PProfServer) Start() {
	pprofPort := os.Getenv("PPROF_PORT")
	if len(pprofPort) != 0 {
		// convert to int
		port, err := strconv.Atoi(pprofPort)
		if err != nil {
			zlog.Fatal().Err(err).Msg("Failed to convert pprof port to int")
		}

		s.e = echo.New()
		e := s.e
		e.Logger.SetLevel(log.DEBUG)
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello from PProf")
		})
		e.Any("/debug/pprof/", func(c echo.Context) error {
			// call pprof index
			pprof.Index(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof heap
		e.Any("/debug/pprof/heap", func(c echo.Context) error {
			// call pprof index specifying the gc
			pprof.Handler("heap").ServeHTTP(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof cmdline
		e.Any("/debug/pprof/cmdline", func(c echo.Context) error {
			// call pprof index
			pprof.Cmdline(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof profile
		e.Any("/debug/pprof/profile", func(c echo.Context) error {
			// call pprof index
			pprof.Profile(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof symbol
		e.Any("/debug/pprof/symbol", func(c echo.Context) error {
			// call pprof index
			pprof.Symbol(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof trace
		e.Any("/debug/pprof/trace", func(c echo.Context) error {
			// call pprof index
			pprof.Trace(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof goroutine
		e.Any("/debug/pprof/goroutine", func(c echo.Context) error {
			// call pprof index
			pprof.Handler("goroutine").ServeHTTP(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof threadcreate
		e.Any("/debug/pprof/threadcreate", func(c echo.Context) error {
			// call pprof index
			pprof.Handler("threadcreate").ServeHTTP(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof block
		e.Any("/debug/pprof/block", func(c echo.Context) error {
			// call pprof index
			pprof.Handler("block").ServeHTTP(c.Response().Writer, c.Request())
			return nil
		})

		// call pprof mutex
		e.Any("/debug/pprof/mutex", func(c echo.Context) error {
			// call pprof index
			pprof.Handler("mutex").ServeHTTP(c.Response().Writer, c.Request())
			return nil
		})
		// call pprof allocs
		e.Any("/debug/pprof/allocs", func(c echo.Context) error {
			// call pprof index
			pprof.Handler("allocs").ServeHTTP(c.Response().Writer, c.Request())
			return nil
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
		zlog.Info().Msg("Starting PProf Web Server")

	}
}
