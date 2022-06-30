package taskclient

import (
	"reflect"

	contracts_asynqengine "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/asynqengine"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/hibiken/asynq"
)

type (
	serviceRedis struct {
		Logger          contracts_logger.ILogger              `inject:""`
		GetRedisOptions contracts_asynqengine.GetRedisOptions `inject:""`
		client          *asynq.Client
	}
	serviceMiniRedis struct {
		Logger          contracts_logger.ILogger                  `inject:""`
		GetRedisOptions contracts_asynqengine.GetMiniRedisOptions `inject:""`
		client          *asynq.Client
	}
)

func assertImplementation() {
	var _ contracts_asynqengine.IRedisTaskClient = (*serviceRedis)(nil)
	var _ contracts_asynqengine.IMiniRedisTaskClient = (*serviceMiniRedis)(nil)
}

var reflectTypeServiceRedis = reflect.TypeOf((*serviceRedis)(nil))
var reflectTypeServiceMiniRedis = reflect.TypeOf((*serviceMiniRedis)(nil))

// AddSingletonITaskClients registers the *service as a singleton.
func AddSingletonITaskClients(builder *di.Builder) {
	contracts_asynqengine.AddSingletonIRedisTaskClient(builder, reflectTypeServiceRedis)
	contracts_asynqengine.AddSingletonIMiniRedisTaskClient(builder, reflectTypeServiceMiniRedis)
}

func (s *serviceRedis) Close() {
	s.client.Close()
}
func (s *serviceRedis) Ctor() {
	options := s.GetRedisOptions()
	s.client = asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     options.Addr,
			Network:  options.Network,
			Password: options.Password,
			Username: options.Username},
	)
}

func (s *serviceRedis) EnqueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	info, err := s.client.Enqueue(task, opts...)
	if err != nil {
		s.Logger.Error().Err(err).Msg("EnqueTask")
	}
	return info, err
}

func (s *serviceMiniRedis) Close() {
	s.client.Close()
}
func (s *serviceMiniRedis) Ctor() {
	options := s.GetRedisOptions()
	s.client = asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     options.Addr,
			Network:  options.Network,
			Password: options.Password,
			Username: options.Username},
	)
}

func (s *serviceMiniRedis) EnqueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	info, err := s.client.Enqueue(task, opts...)
	if err != nil {
		s.Logger.Error().Err(err).Msg("EnqueTask")
	}
	return info, err
}
