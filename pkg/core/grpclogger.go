package core

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	grpclog "google.golang.org/grpc/grpclog"
)

type grpcLogger struct {
}

// NewGRPCLogger Creates a new instance of a gRPC Logger interface that redirects to zerolog
func NewGRPCLogger() grpclog.DepthLoggerV2 {
	return &grpcLogger{}
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) Info(args ...interface{}) {
	log.Info().Msg(fmt.Sprint(args...))
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (a *grpcLogger) Infoln(args ...interface{}) {
	log.Info().Msg(fmt.Sprintln(args...))
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (a *grpcLogger) Infof(format string, args ...interface{}) {
	// google.golang.org/grpc/internal/transport/controlbuf.go
	// sends this log line, so we want to ignore it.
	if strings.Contains(format, "transport: loopyWriter.run returning") ||
		strings.Contains(format, "transport: http2Server.HandleStreams") {
		log.Trace().Msgf(format, args...)
		return
	}

	log.Info().Msgf(format, args...)
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) Warning(args ...interface{}) {
	log.Warn().Msg(fmt.Sprint(args...))
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (a *grpcLogger) Warningln(args ...interface{}) {
	log.Warn().Msg(fmt.Sprintln(args...))
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (a *grpcLogger) Warningf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) Error(args ...interface{}) {
	log.Error().Msg(fmt.Sprint(args...))
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (a *grpcLogger) Errorln(args ...interface{}) {
	log.Error().Msg(fmt.Sprintln(args...))
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (a *grpcLogger) Errorf(format string, args ...interface{}) {

	// google.golang.org/grpc/internal/transport/(http2_client.go | http2_server.go)
	// sends this log line as an error when it's really not. So we want to
	// ignore it.
	if strings.Contains(format, "transport: loopyWriter.run returning") ||
		strings.Contains(format, "transport: http2Server.HandleStreams") {
		log.Trace().Msgf(format, args...)
		return
	}

	log.Error().Msgf(format, args...)
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (a *grpcLogger) Fatal(args ...interface{}) {
	log.Fatal().Msg(fmt.Sprint(args...))
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (a *grpcLogger) Fatalln(args ...interface{}) {
	log.Fatal().Msg(fmt.Sprintln(args...))
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (a *grpcLogger) Fatalf(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (a *grpcLogger) V(l int) bool {
	return true
}

// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) InfoDepth(depth int, args ...interface{}) {
	// gRPC passes a ton of stuff in to here that is really debug level, so we make it Trace()
	log.Trace().Int("depth", depth).Msg(fmt.Sprint(args...))
}

// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) WarningDepth(depth int, args ...interface{}) {
	// gRPC passes a ton of stuff in to here that is really debug level, so we make it Trace()
	log.Trace().Int("depth", depth).Msg(fmt.Sprint(args...))
}

// ErrorDetph logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) ErrorDepth(depth int, args ...interface{}) {
	log.Error().Int("depth", depth).Msg(fmt.Sprint(args...))
}

// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Print.
func (a *grpcLogger) FatalDepth(depth int, args ...interface{}) {
	log.Fatal().Int("depth", depth).Msg(fmt.Sprint(args...))
}
