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

type LoggerService struct {
	Logger *zerolog.Logger
}

func (s *LoggerService) Error() *zerolog.Event {
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Error()
}

func (s *LoggerService) Debug() *zerolog.Event {
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Debug()
}

func (s *LoggerService) Fatal() *zerolog.Event {
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Fatal()
}

func (s *LoggerService) Info() *zerolog.Event {
	e := s.Logger.Debug()
	if !e.Enabled() {
		return s.Logger.Info()
	}
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Info()
}
func (s *LoggerService) Warn() *zerolog.Event {
	e := s.Logger.Debug()
	if !e.Enabled() {
		return s.Logger.Warn()
	}
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Warn()
}
func (s *LoggerService) Trace() *zerolog.Event {
	e := s.Logger.Debug()
	if !e.Enabled() {
		return s.Logger.Trace()
	}
	sublogger := s.withFileNumber(s.Logger)
	return sublogger.Trace()
}

// withFileNumber is used internally to make sure our stack filenames are correct.
func (s *LoggerService) withFileNumber(log *zerolog.Logger) *zerolog.Logger {
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
func (s *LoggerService) getLoggerContext(ctx context.Context) *zerolog.Logger {
	log := zerolog.Ctx(ctx)
	return log
}
