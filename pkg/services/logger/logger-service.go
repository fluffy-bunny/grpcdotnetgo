package logger

import (
	"context"

	"github.com/rs/zerolog"
)

type loggerService struct {
	Logger *zerolog.Logger
}

func (s *loggerService) GetLogger() *zerolog.Logger {
	return s.Logger
}
func (s *loggerService) Error() *zerolog.Event {
	return s.Logger.Error()
}

func (s *loggerService) Debug() *zerolog.Event {
	return s.Logger.Debug()
}

func (s *loggerService) Fatal() *zerolog.Event {
	return s.Logger.Fatal()
}

func (s *loggerService) Info() *zerolog.Event {
	return s.Logger.Info()
}

func (s *loggerService) Warn() *zerolog.Event {
	return s.Logger.Warn()
}

func (s *loggerService) Trace() *zerolog.Event {
	return s.Logger.Trace()
}

func (s *loggerService) getLoggerContext(ctx context.Context) *zerolog.Logger {
	log := zerolog.Ctx(ctx)
	return log
}
