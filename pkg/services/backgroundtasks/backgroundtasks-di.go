package backgroundtasks

import (
	"reflect"

	"github.com/bamzi/jobrunner"
	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddSingletonBackgroundTasks adds service to the DI container
func AddSingletonBackgroundTasks(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddBackgroundTasks")

	types := di.NewTypeSet()
	types.Add(backgroundtasksContracts.ReflectTypeIBackgroundTasks)

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&serviceBackgroundTasks{}),
		Build: func(ctn di.Container) (interface{}, error) {
			jobrunner.Start()
			obj := &serviceBackgroundTasks{
				Logger: contracts_logger.GetISingletonLoggerFromContainer(ctn),
			}
			//jobrunner.Schedule("@every 5s", ReminderEmails{})

			jobsProviders, err := ctn.SafeGetManyByType(backgroundtasksContracts.ReflectTypeIJobsProvider)
			if err == nil && jobsProviders != nil && len(jobsProviders) > 0 {
				for _, jp := range jobsProviders {
					jpi := jp.(backgroundtasksContracts.IJobsProvider)
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
