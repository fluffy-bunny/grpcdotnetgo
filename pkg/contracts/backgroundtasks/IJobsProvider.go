package backgroundtasks

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IJobsProvider"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IJobsProvider

// IJobsProvider ...
type IJobsProvider interface {
	GetScheduledJobs() ScheduledJobs
	GetOneTimeJobs() OneTimeJobs
}
