package backgroundtasks

// IJobsProvider ...
type IJobsProvider interface {
	GetScheduledJobs() ScheduledJobs
	GetOneTimeJobs() OneTimeJobs
}
