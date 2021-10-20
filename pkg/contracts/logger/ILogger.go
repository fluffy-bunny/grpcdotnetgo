package logger

import (
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
