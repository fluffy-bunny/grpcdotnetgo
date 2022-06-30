package taskenginefactory

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	"sync"

	zLog "github.com/rs/zerolog/log"

	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/pkg/async"
	sync_atomic "github.com/fluffy-bunny/grpcdotnetgo/pkg/atomic"
	contracts_asynqengine "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/asynqengine"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	service_asynqengine "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/asynqengine"
	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/reugn/async"
	"go.mapped.dev/micro/job-worker/internal"
	contracts_background_tasks_targetsender "go.mapped.dev/micro/job-worker/internal/contracts/background/tasks/targetsender"
	contracts_config "go.mapped.dev/micro/job-worker/internal/contracts/config"
	contracts_jobs_targetsubscriptions "go.mapped.dev/webhooks-sdk-go/pkg/contracts/jobs/targetsubscriptions"
)

type (
	serverMuxContainer struct {
		config      contracts_asynqengine.TaskEngineConfig
		mux         *asynq.ServeMux
		srv         *asynq.Server
		future      async.Future
		waitChannel chan os.Signal
	}
	service struct {
		Logger                      contracts_logger.ILogger                                        `inject:""`
		Config                      *contracts_config.Config                                        `inject:""`
		Handlers                    []contracts_asynqengine.ISingletonTask                          `inject:""`
		FetchTargetSubscriptionsJob contracts_jobs_targetsubscriptions.IFetchTargetSubscriptionsJob `inject:""`
		taskEngineConfigs           []contracts_asynqengine.TaskEngineConfig
		serverMuxContainers         []*serverMuxContainer
		shuttingdown                sync_atomic.IAtomicBool
		subscriptionFetcher         *backgroundSubscriptionFetcher
		patternToHandler            map[string]contracts_asynqengine.ISingletonTask
		mtx                         sync.Mutex
	}
	backgroundSubscriptionFetcher struct {
		// quit channel is closed when the shutdown is requested.
		quit                        chan struct{}
		targetSubscriptions         *contracts_jobs_targetsubscriptions.TargetSubscriptions
		future                      async.Future
		FetchTargetSubscriptionsJob contracts_jobs_targetsubscriptions.IFetchTargetSubscriptionsJob
		Callback                    func(targetSubscriptions *contracts_jobs_targetsubscriptions.TargetSubscriptions)
	}
)

func assertImplementation() {
	var _ contracts_asynqengine.ITaskEngineFactory = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonITaskEngineFactory registers the *service as a singleton.
func AddSingletonITaskEngineFactory(builder *di.Builder) {
	contracts_asynqengine.AddSingletonITaskEngineFactory(builder, reflectType)
}
func NewBackgroundSubscriptionFetcher(fetchTargetSubscriptionsJob contracts_jobs_targetsubscriptions.IFetchTargetSubscriptionsJob) *backgroundSubscriptionFetcher {
	return &backgroundSubscriptionFetcher{
		FetchTargetSubscriptionsJob: fetchTargetSubscriptionsJob,
		quit:                        make(chan struct{}),
	}
}
func (s *serverMuxContainer) Wait() {
	signal.Notify(
		s.waitChannel,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-s.waitChannel
}
func (s *serverMuxContainer) SignalStop() {
	s.waitChannel <- os.Interrupt
}
func (s *backgroundSubscriptionFetcher) stop() {
	close(s.quit)
}
func (s *backgroundSubscriptionFetcher) start() {
	s.future = grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise) {
		var err error
		defer func() {
			promise.Success(&grpcdotnetgoasync.AsyncResponse{
				Message: "End BackgroundSubscriptionWorker",
				Error:   err,
			})
		}()
		s.backgroundSubscriptionWorker()
	})

}
func (s *backgroundSubscriptionFetcher) backgroundSubscriptionWorker() {
	log := zLog.With().Caller().Str("function", internal.GetFunctionName(s.backgroundSubscriptionWorker)).Logger()
	defer func() {
		if r := recover(); r != nil {
			err := errors.Wrap(fmt.Errorf("%v", r), "panic")
			log.Error().Stack().Err(err).Send()
		}
	}()
	var workingVersion int64 = 0
	var fetchAndCallback = func() bool {
		s.targetSubscriptions = s.FetchTargetSubscriptionsJob.GetCurrentTargetSubscriptions()
		if s.targetSubscriptions == nil {
			return false
		}
		workingVersion = s.targetSubscriptions.Version
		s.Callback(s.targetSubscriptions)
		return true
	}
	for {
		select {
		case <-s.quit:
			return
		case <-time.After(time.Second * 5):

			if s.targetSubscriptions == nil {
				if !fetchAndCallback() {
					continue
				}
			}

			latestVersion := s.FetchTargetSubscriptionsJob.GetCurrentTargetSubscriptionsVersion()
			if latestVersion != workingVersion {
				// we need to fetch all the new subscriptions, stop the engines and restart them all.
				fetchAndCallback()
			}
		}
	}
}
func (s *service) OnNewSubscriptions(targetSubscriptions *contracts_jobs_targetsubscriptions.TargetSubscriptions) {
	log := s.Logger.GetLogger().With().Caller().Str("function", internal.GetFunctionName(s.OnNewSubscriptions)).Logger()
	defer func() {
		if r := recover(); r != nil {
			err := errors.Wrap(fmt.Errorf("%v", r), "panic")
			log.Error().Stack().Err(err).Send()
		}
	}()
	//--~--~--~--~-- BARBED WIRE --~--~--~--~--~--//
	s.mtx.Lock()
	defer s.mtx.Unlock()
	//--~--~--~--~-- BARBED WIRE --~--~--~--~--~--//
	shuttingDown := s.shuttingdown.Load()
	if shuttingDown {
		return
	}
	var taskEngineConfigs []contracts_asynqengine.TaskEngineConfig
	for targetID, targetConfig := range targetSubscriptions.ByTargetMap {
		// what redis cluster is this target using?
		// the live REDIS cluster or the inmemory miniredis one?
		redisAddr := s.Config.RedisOptions.Addr
		var maxArchiveSize *int

		if !targetConfig.WebhookTarget.Durable {
			value := 0
			maxArchiveSize = &value // durable means we dont keep any archived messages.
			redisAddr = s.Config.MiniRedisOptions.Addr
		}

		taskEngineConfig := contracts_asynqengine.TaskEngineConfig{
			RedisClientOpt: asynq.RedisClientOpt{
				Addr:     redisAddr,
				Network:  s.Config.RedisOptions.Network,
				Password: s.Config.RedisOptions.Password,
				Username: s.Config.RedisOptions.Username,
			},
			Config: asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: 10,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{
					contracts_background_tasks_targetsender.GenerateWebhooksQueueNameCritical(targetID): 6,
					contracts_background_tasks_targetsender.GenerateWebhooksQueueNameNormal(targetID):   3,
					contracts_background_tasks_targetsender.GenerateWebhooksQueueNameLow(targetID):      1,
				},
				MaxArchiveSize: maxArchiveSize,
				Logger: service_asynqengine.NewLogger(service_asynqengine.StrOption{
					Key:   "targetID",
					Value: targetID,
				}),
			},
			Patterns: core_hashset.NewStringSet(contracts_background_tasks_targetsender.TypeWebhooksMessage),
		}
		taskEngineConfigs = append(taskEngineConfigs, taskEngineConfig)
	}

	var serverMuxContainers []*serverMuxContainer
	for idx := range taskEngineConfigs {
		config := taskEngineConfigs[idx]
		// swap out the asynq logger for or own zerolog logger

		srv := asynq.NewServer(config.RedisClientOpt, config.Config)
		serverMuxContainers = append(serverMuxContainers, &serverMuxContainer{
			config:      config,
			mux:         asynq.NewServeMux(),
			srv:         srv,
			waitChannel: make(chan os.Signal),
		})
	}
	for idx := range serverMuxContainers {
		container := serverMuxContainers[idx]
		for idx2 := range container.config.Patterns.Values() {
			ctnPattern := container.config.Patterns.Values()[idx2]
			handler, ok := s.patternToHandler[ctnPattern]
			if ok {
				container.mux.Handle(ctnPattern, handler)
			}
		}
	}
	s.Logger.Info().Msg("Received new subscriptions, restarting AsynqServers")
	s.stopAsynqServers()
	s.taskEngineConfigs = taskEngineConfigs
	s.serverMuxContainers = serverMuxContainers
	s.startAsynqServers()

}

func (s *service) Ctor() {
	s.subscriptionFetcher = NewBackgroundSubscriptionFetcher(s.FetchTargetSubscriptionsJob)
	s.subscriptionFetcher.Callback = s.OnNewSubscriptions
	s.patternToHandler = make(map[string]contracts_asynqengine.ISingletonTask)
	s.sanitizeHandlers()
	s.shuttingdown = sync_atomic.NewAtomicBool(false)
}

func (s *service) sanitizeHandlers() {
	for idx := range s.Handlers {
		handler := s.Handlers[idx]
		for idx2 := range handler.GetPatterns().Values() {
			pattern := handler.GetPatterns().Values()[idx2]
			if _, ok := s.patternToHandler[pattern]; ok {
				panic(fmt.Sprintf("duplicate pattern %s", pattern))
			}
			s.patternToHandler[pattern] = handler
		}
	}
}

func (s *service) startAsynqServers() error {
	if core_utils.IsEmptyOrNil(s.serverMuxContainers) {
		return nil
	}
	for idx := range s.serverMuxContainers {
		container := s.serverMuxContainers[idx]
		if container.future != nil {
			panic("task engine already started")
		}
	}
	for idx := range s.serverMuxContainers {
		container := s.serverMuxContainers[idx]
		container.future = grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise) {
			var err error
			defer func() {
				promise.Success(&grpcdotnetgoasync.AsyncResponse{
					Message: "End Serve - echo Server",
					Error:   err,
				})
			}()
			err = container.srv.Start(container.mux)
			if err != nil {
				s.Logger.Fatal().Err(err).Msg("Failed to start asynq server")
			} else {
				container.Wait()
				container.srv.Shutdown()
			}
		})
	}

	return nil
}

func (s *service) stopAsynqServers() error {
	if core_utils.IsEmptyOrNil(s.serverMuxContainers) {
		return nil
	}
	// tell all to stop and shutdown
	for idx := range s.serverMuxContainers {
		container := s.serverMuxContainers[idx]
		container.SignalStop()
	}

	// wait for all to return the promise
	for idx := range s.serverMuxContainers {
		container := s.serverMuxContainers[idx]
		promise, err := container.future.Get()
		if err != nil {
			s.Logger.Error().Err(err).Msg("Failed to get server shutdown promise")
		}
		response := promise.(*grpcdotnetgoasync.AsyncResponse)
		s.Logger.Info().Msg(response.Message)
	}
	return nil
}
func (s *service) Start() error {
	s.subscriptionFetcher.start()
	return nil
}

func (s *service) Stop() error {
	log := s.Logger.GetLogger().With().Caller().Str("function", internal.GetFunctionName(s.Stop)).Logger()
	defer func() {
		if r := recover(); r != nil {
			err := errors.Wrap(fmt.Errorf("%v", r), "panic")
			log.Error().Stack().Err(err).Send()
		}
	}()
	// the ones up front are signalers that we are shutting down
	s.shuttingdown.Store(true)   // do this first, because there may be some outstanding new subscriptions comming in right now
	s.subscriptionFetcher.stop() // stop the subscription fetcher.  Beware one of these may be in the middle of a new subscription

	//--~--~--~--~-- BARBED WIRE --~--~--~--~--~--//
	s.mtx.Lock()
	defer s.mtx.Unlock()
	//--~--~--~--~-- BARBED WIRE --~--~--~--~--~--//
	s.stopAsynqServers()
	// wait for subscription background worker to finish
	s.subscriptionFetcher.future.Get()
	return nil
}
