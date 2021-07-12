package backgroundtasks

import (
	"reflect"

	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/memqueue"
)

// GetContextAccessorFromContainer from the Container

// AddCounterTaskConsumer adds service to the DI container
func AddCounterTaskConsumer(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddCounterTaskConsumer")
	types := di.NewTypeSet()
	types.Add(rtIConsumer)

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
			obj := &serviceBackgroundTasks{
				Logger: servicesLogger.GetSingletonLoggerFromContainer(ctn),
			}

			obj.QueueFactory = memqueue.NewFactory()
			obj.MainQueue = obj.QueueFactory.RegisterQueue(&taskq.QueueOptions{
				Name: "api-worker",
			})
			consumers, err := ctn.SafeGetManyByType(rtIConsumer)
			if err == nil && consumers != nil && len(consumers) > 0 {
				for _, c := range consumers {
					co := c.(IConsumer)
					taskMessages := co.GetTaskMessages()
					for _, to := range taskMessages {
						obj.MainQueue.Add(to)
					}

				}

			}

			return obj, nil
		},
		Close: func(obj interface{}) error {
			log.Info().Msg("Closing BackgroundTasks")
			o := obj.(*serviceBackgroundTasks)
			o.QueueFactory.StopConsumers()
			o.QueueFactory.Close()

			return nil
		},
	})
}
