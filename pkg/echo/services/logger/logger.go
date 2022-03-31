package logger

import (
	"reflect"

	contracts_contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/contextaccessor"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type service struct {
	Logger *zerolog.Logger
}

type serviceScoped struct {
	Logger          *zerolog.Logger
	ContextAccessor contracts_contextaccessor.IEchoContextAccessor `inject:""`
}

func (s *serviceScoped) Ctor() {
	ctx := s.ContextAccessor.GetContext().Request().Context()
	s.Logger = zerolog.Ctx(ctx)
}
func assertImplementation() {
	var _ contracts_logger.ILogger = (*service)(nil)
	var _ contracts_logger.ISingletonLogger = (*service)(nil)
	var _ contracts_logger.ILogger = (*serviceScoped)(nil)
}

// AddILogger ...
func AddILogger(builder *di.Builder) {
	AddSingletonILogger(builder)
	AddScopedILogger(builder)
}

// AddSingletonILogger ...
func AddSingletonILogger(builder *di.Builder) {
	bldFunc := func(ctn di.Container) (interface{}, error) {
		logger := log.With().Logger()
		return &serviceScoped{
			Logger: &logger,
		}, nil
	}
	log.Info().Msg("DI: ISingletonLogger - singleton")
	contracts_logger.AddSingletonISingletonLoggerByFunc(builder, reflect.TypeOf(&service{}), bldFunc)
	log.Info().Msg("DI: ILogger - singleton")
	contracts_logger.AddSingletonILoggerByFunc(builder, reflect.TypeOf(&service{}), bldFunc)
}

// AddScopedILogger ...
func AddScopedILogger(builder *di.Builder) {
	log.Info().Msg("DI: ILogger - SCOPED")
	contracts_logger.AddScopedILogger(builder, reflect.TypeOf(&serviceScoped{}))
}

func (s *serviceScoped) GetLogger() *zerolog.Logger {
	return s.Logger
}
func (s *serviceScoped) Error() *zerolog.Event {
	return s.Logger.Error()
}

func (s *serviceScoped) Debug() *zerolog.Event {
	return s.Logger.Debug()
}

func (s *serviceScoped) Fatal() *zerolog.Event {
	return s.Logger.Fatal()
}

func (s *serviceScoped) Info() *zerolog.Event {
	return s.Logger.Info()
}

func (s *serviceScoped) Warn() *zerolog.Event {
	return s.Logger.Warn()
}

func (s *serviceScoped) Trace() *zerolog.Event {
	return s.Logger.Trace()
}

func (s *service) GetLogger() *zerolog.Logger {
	return s.Logger
}
func (s *service) Error() *zerolog.Event {
	return s.Logger.Error()
}

func (s *service) Debug() *zerolog.Event {
	return s.Logger.Debug()
}

func (s *service) Fatal() *zerolog.Event {
	return s.Logger.Fatal()
}

func (s *service) Info() *zerolog.Event {
	return s.Logger.Info()
}

func (s *service) Warn() *zerolog.Event {
	return s.Logger.Warn()
}

func (s *service) Trace() *zerolog.Event {
	return s.Logger.Trace()
}
