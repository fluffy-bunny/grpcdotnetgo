package backgroundtasks

import (
	"reflect"
	"time"

	"github.com/bamzi/jobrunner"
	contracts_backgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// NewScheduledJob ...
func NewScheduledJob(job cron.Job, schedule string) *contracts_backgroundtasks.ScheduledJob {
	return &contracts_backgroundtasks.ScheduledJob{
		Job:      job,
		Schedule: schedule,
	}
}

// NewScheduledJobs ...
func NewScheduledJobs(jobs ...*contracts_backgroundtasks.ScheduledJob) contracts_backgroundtasks.ScheduledJobs {
	return jobs
}

// NewOneTimeJob ...
func NewOneTimeJob(job cron.Job, delay time.Duration) *contracts_backgroundtasks.OneTimeJob {
	return &contracts_backgroundtasks.OneTimeJob{
		Job:   job,
		Delay: delay,
	}
}

// NewOneTimeJobs ...
func NewOneTimeJobs(jobs ...*contracts_backgroundtasks.OneTimeJob) contracts_backgroundtasks.OneTimeJobs {
	return jobs
}

type serviceBackgroundTasks struct {
	Logger        loggerContracts.ILogger                   `inject:""`
	JobsProviders []contracts_backgroundtasks.IJobsProvider `inject:"optional"`
}

// BuildBreak ...
func BuildBreak() contracts_backgroundtasks.IBackgroundTasks {
	return &serviceBackgroundTasks{}
}

// AddSingletonBackgroundTasks ...
func AddSingletonBackgroundTasks(builder *di.Builder) {
	log.Info().Msg("DI: AddSingletonBackgroundTasks")
	contracts_backgroundtasks.AddSingletonIBackgroundTasks(builder, reflect.TypeOf(&serviceBackgroundTasks{}))
}

func (s *serviceBackgroundTasks) Ctor() {
	if !utils.IsEmptyOrNil(s.JobsProviders) {
		s.Logger.Info().Msg("Starting BackgroundTasks")
		jobrunner.Start()
		for _, jp := range s.JobsProviders {
			sjs := jp.GetScheduledJobs()
			for _, sj := range sjs {
				jobrunner.Schedule(sj.Schedule, sj.Job)
			}
			otjs := jp.GetOneTimeJobs()
			for _, otj := range otjs {
				jobrunner.In(otj.Delay, otj.Job)
			}
		}
	}
}
func (s *serviceBackgroundTasks) Close() {
	if !utils.IsEmptyOrNil(s.JobsProviders) {
		s.Logger.Info().Msg("Closing BackgroundTasks")
		jobrunner.Stop()
	}
}
