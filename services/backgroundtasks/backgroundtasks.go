package backgroundtasks

import (
	"fmt"
	"sync/atomic"
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
type OneTimeJob struct {
	// Job must support Run() func
	Job   cron.Job
	Delay time.Duration
}
type IJobsProvider interface {
	GetScheduledJobs() []*ScheduledJob
	GetOneTimeJobs() []*OneTimeJob
}

var (
	rtIBackgroundTasks = di.GetInterfaceReflectType((*IBackgroundTasks)(nil))
	rtIJobsProvider    = di.GetInterfaceReflectType((*IJobsProvider)(nil))
)

type serviceBackgroundTasks struct {
	Logger servicesLogger.ILogger
}

type counterConsumer struct {
	Logger servicesLogger.ILogger
}

type counterJob struct {
	counter int32
}

func newCounterJob() *counterJob {
	return &counterJob{}
}
func (j *counterJob) Run() {
	j.incrLocalCounter()
	log.Info().Str("count",
		fmt.Sprintf("%v", j.getLocalCounter())).
		Msg("Background Counter")
}
func (j *counterJob) incrLocalCounter() {
	atomic.AddInt32(&j.counter, 1)
}
func (j *counterJob) getLocalCounter() int32 {
	return atomic.LoadInt32(&j.counter)
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

func (s *counterConsumer) GetOneTimeJobs() []*OneTimeJob {
	return []*OneTimeJob{
		{
			Job:   newWelcomeJob("Hi Bunny"),
			Delay: time.Microsecond,
		},
		{
			Job:   newWelcomeJob("Hi Porky"),
			Delay: time.Minute,
		},
	}

}
func (s *counterConsumer) GetScheduledJobs() []*ScheduledJob {
	counterJob := newCounterJob()
	cronJob := &ScheduledJob{
		Job:      counterJob,
		Schedule: "@every 0h0m5s",
	}

	return []*ScheduledJob{
		cronJob,
	}
}
