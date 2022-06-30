package asynqengine

//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE      gen "InterfaceType=ISingletonTask,ITaskClient,ITaskEngineFactory,IRedisTaskClient,IMiniRedisTaskClient"
//go:generate genny -pkg $GOPACKAGE -in=../../../genny/sarulabsdi/func-types.go      -out=gen-func-$GOFILE gen "FuncType=GetMiniRedisOptions,GetRedisOptions"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/background/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ISingletonTask,ITaskClient,ITaskEngineFactory,IRedisTaskClient,IMiniRedisTaskClient

import (
	"context"

	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	"github.com/hibiken/asynq"
)

type (
	// RedisOptions ...
	RedisOptions struct {
		// The network type, either tcp or unix.
		// Default is tcp.
		Network string `json:"network" mapstructure:"NETWORK"`
		// host:port address.
		Addr string `json:"addr" mapstructure:"ADDR"`
		// Use the specified Username to authenticate the current connection
		// with one of the connections defined in the ACL list when connecting
		// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
		Username string `json:"username" mapstructure:"USERNAME"`
		// Optional password. Must match the password specified in the
		// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
		// or the User Password when connecting to a Redis 6.0 instance, or greater,
		// that is using the Redis ACL system.
		Password string `json:"password" mapstructure:"PASSWORD"`

		Namespace []string `json:"namespace" mapstructure:"NAMESPACE"`
	}
	// GetMiniRedisOptions ...
	GetMiniRedisOptions func() RedisOptions
	// GetRedisOptions ...
	GetRedisOptions func() RedisOptions

	// TaskEngineConfig ...
	TaskEngineConfig struct {
		RedisClientOpt asynq.RedisClientOpt
		Config         asynq.Config
		Patterns       *core_hashset.StringSet
	}
	// ITaskEngineFactory ...
	ITaskEngineFactory interface {
		Start() error
		Stop() error
	}

	// ISingletonTask ...
	ISingletonTask interface {
		GetPatterns() *core_hashset.StringSet
		ProcessTask(ctx context.Context, t *asynq.Task) error
	}
	// ITaskClient ...
	ITaskClient interface {
		EnqueTask(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
	}
	// IRedisTaskClient ...
	IRedisTaskClient interface {
		ITaskClient
	}
	// IMiniRedisTaskClient ...
	IMiniRedisTaskClient interface {
		ITaskClient
	}
)
