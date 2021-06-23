package logger

import (
	"context"
	"runtime"
	"strconv"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Error() *zerolog.Event
	Debug() *zerolog.Event
	Fatal() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Trace() *zerolog.Event
}

func (s *loggerService) Error() *zerolog.Event {
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Error()
}

func (s *loggerService) Debug() *zerolog.Event {
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Debug()
}

func (s *loggerService) Fatal() *zerolog.Event {
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Fatal()
}

func (s *loggerService) Info() *zerolog.Event {
	e := s.Logger.Debug()
	if !e.Enabled() {
		return s.Logger.Info()
	}
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Info()
}
func (s *loggerService) Warn() *zerolog.Event {
	e := s.Logger.Debug()
	if !e.Enabled() {
		return s.Logger.Warn()
	}
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Warn()
}
func (s *loggerService) Trace() *zerolog.Event {
	e := s.Logger.Debug()
	if !e.Enabled() {
		return s.Logger.Trace()
	}
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Trace()
}

// withFileNumber is used internally to make sure our stack filenames are correct.
func (s *loggerService) withFileNumber(log *zerolog.Logger) *zerolog.Logger {
	e := log.Debug()
	if !e.Enabled() {
		return log
	}

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
