package backgroundtasks

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IBackgroundTasks"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IBackgroundTasks

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
