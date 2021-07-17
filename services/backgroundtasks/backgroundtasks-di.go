package backgroundtasks

import (
	"fmt"
	"reflect"

	"github.com/bamzi/jobrunner"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// GetContextAccessorFromContainer from the Container

// AddCounterTaskConsumer adds service to the DI container
func AddCounterTaskConsumer(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddCounterTaskConsumer")
	types := di.NewTypeSet()
	types.Add(rtIJobsProvider)

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&counterConsumer{}),
		Build: func(ctn di.Container) (interface{}, error) {
			obj := &counterConsumer{
				Logger: servicesLogger.GetSingletonLoggerFromContainer(ctn),
			}

			return obj, nil
		},
		Close: func(obj interface{}) error {

			return nil
		},
	})
}
func GetBackgroundTasksFromContainer(ctn di.Container) IBackgroundTasks {
	obj := ctn.GetByType(rtIBackgroundTasks).(IBackgroundTasks)
	return obj
}

type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}

// AddBackgroundTasks adds service to the DI container
func AddBackgroundTasks(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddBackgroundTasks")
	types := di.NewTypeSet()

	types.Add(rtIBackgroundTasks)

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&serviceBackgroundTasks{}),
		Build: func(ctn di.Container) (interface{}, error) {
			jobrunner.Start()
			obj := &serviceBackgroundTasks{
				Logger: servicesLogger.GetSingletonLoggerFromContainer(ctn),
			}
			//jobrunner.Schedule("@every 5s", ReminderEmails{})

			jobsProviders, err := ctn.SafeGetManyByType(rtIJobsProvider)
			if err == nil && jobsProviders != nil && len(jobsProviders) > 0 {
				for _, jp := range jobsProviders {
					jpi := jp.(IJobsProvider)
					sjs := jpi.GetScheduledJobs()
					for _, sj := range sjs {
						jobrunner.Schedule(sj.Schedule, sj.Job)
					}

					otjs := jpi.GetOneTimeJobs()
					for _, otj := range otjs {
						jobrunner.In(otj.Delay, otj.Job)
					}
				}

			}

			return obj, nil
		},
		Close: func(obj interface{}) error {
			log.Info().Msg("Closing BackgroundTasks")

			jobrunner.Stop()
			return nil
		},
	})
}
