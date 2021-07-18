package backgroundtasks

import (
	"time"

	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type IBackgroundTasks interface {
}

// ScheduledJob cron
type ScheduledJob struct {
	// Job must support Run() func
	Job cron.Job
	// Schedule "* */5 * * * *","@every 1h30m10s","@midnight"
	Schedule string
}
type ScheduledJobs []*ScheduledJob

func NewScheduledJob(job cron.Job, schedule string) *ScheduledJob {
	return &ScheduledJob{
		Job:      job,
		Schedule: schedule,
	}
}

func NewScheduledJobs(jobs ...*ScheduledJob) ScheduledJobs {
	return jobs
}

type OneTimeJob struct {
	// Job must support Run() func
	Job   cron.Job
	Delay time.Duration
}
type OneTimeJobs []*OneTimeJob

func NewOneTimeJob(job cron.Job, delay time.Duration) *OneTimeJob {
	return &OneTimeJob{
		Job:   job,
		Delay: delay,
	}
}

func NewOneTimeJobs(jobs ...*OneTimeJob) OneTimeJobs {
	return jobs
}

type IJobsProvider interface {
	GetScheduledJobs() ScheduledJobs
	GetOneTimeJobs() OneTimeJobs
}

var (
	TypeIBackgroundTasks = di.GetInterfaceReflectType((*IBackgroundTasks)(nil))
	TypeIJobsProvider    = di.GetInterfaceReflectType((*IJobsProvider)(nil))
)

type serviceBackgroundTasks struct {
	Logger servicesLogger.ILogger
}

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
