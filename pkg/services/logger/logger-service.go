package logger

import (
	"context"
	"runtime"
	"strconv"

	"github.com/rs/zerolog"
)

// ILogger interface
type ILogger interface {
	Error() *zerolog.Event
	Debug() *zerolog.Event
	Fatal() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Trace() *zerolog.Event

	GetLogger() *zerolog.Logger

	ErrorL(logger *zerolog.Logger) *zerolog.Event
	DebugL(logger *zerolog.Logger) *zerolog.Event
	FatalL(logger *zerolog.Logger) *zerolog.Event
	InfoL(logger *zerolog.Logger) *zerolog.Event
	WarnL(logger *zerolog.Logger) *zerolog.Event
	TraceL(logger *zerolog.Logger) *zerolog.Event
}

type loggerService struct {
	Logger *zerolog.Logger
}

func (s *loggerService) ErrorL(logger *zerolog.Logger) *zerolog.Event {
	sublogger := s.withFileNumber(logger)
	return sublogger.Error()
}
func (s *loggerService) DebugL(logger *zerolog.Logger) *zerolog.Event {
	sublogger := s.withFileNumber(logger)
	return sublogger.Debug()
}
func (s *loggerService) FatalL(logger *zerolog.Logger) *zerolog.Event {
	sublogger := s.withFileNumber(logger)
	return sublogger.Fatal()
}
func (s *loggerService) InfoL(logger *zerolog.Logger) *zerolog.Event {
	e := logger.Debug()
	if !e.Enabled() {
		return logger.Info()
	}
	sublogger := s.withFileNumber(logger)
	return sublogger.Info()
}
func (s *loggerService) WarnL(logger *zerolog.Logger) *zerolog.Event {
	e := logger.Debug()
	if !e.Enabled() {
		return logger.Warn()
	}
	sublogger := s.withFileNumber(logger)
	return sublogger.Warn()
}
func (s *loggerService) TraceL(logger *zerolog.Logger) *zerolog.Event {
	e := logger.Debug()
	if !e.Enabled() {
		return logger.Trace()
	}
	sublogger := s.withFileNumber(logger)
	return sublogger.Trace()
}

func (s *loggerService) GetLogger() *zerolog.Logger {
	return s.Logger
}
func (s *loggerService) Error() *zerolog.Event {
	return s.ErrorL(s.Logger)
}

func (s *loggerService) Debug() *zerolog.Event {
	return s.DebugL(s.Logger)
}

func (s *loggerService) Fatal() *zerolog.Event {
	return s.FatalL(s.Logger)
}

func (s *loggerService) Info() *zerolog.Event {
	return s.InfoL(s.Logger)
}

func (s *loggerService) Warn() *zerolog.Event {
	return s.WarnL(s.Logger)
}

func (s *loggerService) Trace() *zerolog.Event {
	return s.TraceL(s.Logger)
}

// withFileNumber is used internally to make sure our stack filenames are correct.
func (s *loggerService) withFileNumber(log *zerolog.Logger) *zerolog.Logger {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	subLogger := log.With().
		Str("stack.file", frame.File).
		Str("stack.line", strconv.Itoa(frame.Line)).
		Str("stack.function", frame.Function).Logger()
	return &subLogger
}
func (s *loggerService) getLoggerContext(ctx context.Context) *zerolog.Logger {
	log := zerolog.Ctx(ctx)
	return log
}
