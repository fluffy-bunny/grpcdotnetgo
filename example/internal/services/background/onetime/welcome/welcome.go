package welcome

import (
	"time"

	servicesBackgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
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
	Logger servicesLogger.ILogger
}

func (s *service) GetOneTimeJobs() servicesBackgroundtasks.OneTimeJobs {
	welcomeJob1 := newWelcomeJob("Hello From The Cron")
	cronJob1 := servicesBackgroundtasks.NewOneTimeJob(welcomeJob1, time.Millisecond)

	welcomeJob2 := newWelcomeJob("Hello From The Cron - Much Later")
	cronJob2 := servicesBackgroundtasks.NewOneTimeJob(welcomeJob2, time.Minute)

	return servicesBackgroundtasks.NewOneTimeJobs(cronJob1, cronJob2)

}
func (s *service) GetScheduledJobs() servicesBackgroundtasks.ScheduledJobs {

	return nil
}
