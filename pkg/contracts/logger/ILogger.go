package logger

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ILogger,ISingletonLogger"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ILogger,ISingletonLogger

import (
	"github.com/rs/zerolog"
)

type (
	// ILogger interface
	ILogger interface {
		Error() *zerolog.Event
		Debug() *zerolog.Event
		Fatal() *zerolog.Event
		Info() *zerolog.Event
		Warn() *zerolog.Event
		Trace() *zerolog.Event

		GetLogger() *zerolog.Logger
	}
	// ISingletonLogger when you absolutely need a singleton logger
	ISingletonLogger interface {
		ILogger
	}
)
