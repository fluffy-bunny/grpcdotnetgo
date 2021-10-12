package welcome

import (
	"time"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	servicesBackgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	"github.com/rs/zerolog/log"
)

// JOB
//------------------------------------------
type welcomeJob struct {
	message string
}

func newWelcomeJob(message string) *welcomeJob {
	return &welcomeJob{
		message: message,
	}
}
func (j *welcomeJob) Run() {
	log.Info().Str("message", j.message).
		Msg("Welcome Job")
}

// Job Provider
//------------------------------------------
type service struct {
	Logger loggerContracts.ILogger
}

func (s *service) GetOneTimeJobs() backgroundtasksContracts.OneTimeJobs {
	welcomeJob1 := newWelcomeJob("Hello From The Cron")
	cronJob1 := servicesBackgroundtasks.NewOneTimeJob(welcomeJob1, time.Millisecond)

	welcomeJob2 := newWelcomeJob("Hello From The Cron - Much Later")
	cronJob2 := servicesBackgroundtasks.NewOneTimeJob(welcomeJob2, time.Minute)

	return servicesBackgroundtasks.NewOneTimeJobs(cronJob1, cronJob2)
}
func (s *service) GetScheduledJobs() backgroundtasksContracts.ScheduledJobs {
	return nil
}
