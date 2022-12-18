package core

import (
	"os"
	"strings"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/pkg/profile"
	"github.com/rs/zerolog/log"
)

// PProf is the PProf object that manages an echo web server
type PProf struct {
	goProfilerStopper interface{ Stop() }
}

// NewPProf creates a new PProf object
func NewPProf() *PProf {
	return &PProf{}
}

// Stop ...
func (s *PProf) Stop() {
	if s.goProfilerStopper != nil {
		log.Info().Msg("Stopping Go Profiler")
		s.goProfilerStopper.Stop()
		log.Info().Msg("Go Profiler stopped")
	}
}

// Start starts the echo web server using async and futures
func (s *PProf) Start() {
	goProfiling := os.Getenv("GO_PROFILING")
	if !utils.IsEmptyOrNil(goProfiling) {
		parts := strings.Split(goProfiling, ";")
		if len(parts) == 2 {
			// Determine the profiling type (only one is supported per execution)
			var profType func(*profile.Profile)
			switch strings.ToUpper(parts[0]) {
			case "CPU":
				profType = profile.CPUProfile
			case "MEM":
				profType = profile.MemProfile
			case "BLOCK":
				profType = profile.BlockProfile
			case "GOROUTINE":
				profType = profile.GoroutineProfile
			case "MUTEX":
				profType = profile.MutexProfile
			case "THREAD":
				profType = profile.ThreadcreationProfile
			case "TRACE":
				profType = profile.TraceProfile
			}

			// Start the profiler
			if profType != nil {
				log.Info().Str("profile_type", parts[0]).Msg("Starting Go Profiler")
				s.goProfilerStopper = profile.Start(profType, profile.ProfilePath(parts[1]), profile.NoShutdownHook)
			}
		}
	}
}
