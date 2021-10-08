package backgroundtasks

import (
	"time"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/robfig/cron/v3"
)

func NewScheduledJob(job cron.Job, schedule string) *backgroundtasksContracts.ScheduledJob {
	return &backgroundtasksContracts.ScheduledJob{
		Job:      job,
		Schedule: schedule,
	}
}

func NewScheduledJobs(jobs ...*backgroundtasksContracts.ScheduledJob) backgroundtasksContracts.ScheduledJobs {
	return jobs
}

func NewOneTimeJob(job cron.Job, delay time.Duration) *backgroundtasksContracts.OneTimeJob {
	return &backgroundtasksContracts.OneTimeJob{
		Job:   job,
		Delay: delay,
	}
}

func NewOneTimeJobs(jobs ...*backgroundtasksContracts.OneTimeJob) backgroundtasksContracts.OneTimeJobs {
	return jobs
}

var (
	TypeIBackgroundTasks = di.GetInterfaceReflectType((*backgroundtasksContracts.IBackgroundTasks)(nil))
	TypeIJobsProvider    = di.GetInterfaceReflectType((*backgroundtasksContracts.IJobsProvider)(nil))
)

type serviceBackgroundTasks struct {
	Logger servicesLogger.ILogger
}
