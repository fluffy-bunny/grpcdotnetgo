package backgroundtasks

import (
	"context"
	"sync/atomic"
	"time"

	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/vmihailenco/taskq/v3"
)

type IBackgroundTasks interface {
}

type IConsumer interface {
	GetTaskMessages() []*taskq.Message
}

var (
	rtIBackgroundTasks = di.GetInterfaceReflectType((*IBackgroundTasks)(nil))
	rtIConsumer        = di.GetInterfaceReflectType((*IConsumer)(nil))
)

type serviceBackgroundTasks struct {
	QueueFactory taskq.Factory
	MainQueue    taskq.Queue
	Logger       servicesLogger.ILogger
}

type counterConsumer struct {
	Logger servicesLogger.ILogger
}

var counter int32

func IncrLocalCounter() {
	atomic.AddInt32(&counter, 1)
}
func GetLocalCounter() int32 {
	return atomic.LoadInt32(&counter)
}
func (s *counterConsumer) GetTaskMessages() []*taskq.Message {
	ctx := context.Background()

	var countTask = taskq.RegisterTask(&taskq.TaskOptions{
		Name: "counter",
		Handler: func() error {
			IncrLocalCounter()
			time.Sleep(time.Millisecond)
			return nil
		},
	})

	msg := countTask.WithArgs(ctx)
	return []*taskq.Message{
		msg,
	}
}
