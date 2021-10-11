package backgroundtasks

import (
	"time"

	"github.com/robfig/cron/v3"
)

// IBackgroundTasks ...
type IBackgroundTasks interface {
}

// ScheduledJob cron
type ScheduledJob struct {
	// Job must support Run() func
	Job cron.Job
	// Schedule "* */5 * * * *","@every 1h30m10s","@midnight"
	Schedule string
}

// ScheduledJobs list
type ScheduledJobs []*ScheduledJob

// OneTimeJob type
type OneTimeJob struct {
	// Job must support Run() func
	Job   cron.Job
	Delay time.Duration
}

// OneTimeJobs list
type OneTimeJobs []*OneTimeJob
