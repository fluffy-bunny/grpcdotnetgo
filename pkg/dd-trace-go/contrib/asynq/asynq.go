package asynq

import (
	"context"
	"math"

	"github.com/hibiken/asynq"
	// add zerolog
	"github.com/rs/zerolog/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Client is a traced version of asynq.Client.
type Client struct {
	*asynq.Client
	cfg *config
}

// A messageCarrier implements TextMapReader/TextMapWriter for extracting/injecting traces on a kafka.Message
type messageCarrier struct {
	msg *asynq.Task
}

// WrapClient wraps a asynq.Client so that all requests are traced.
func WrapClient(c *asynq.Client, opts ...Option) *Client {
	client := &Client{
		Client: c,
		cfg:    newConfig(opts...),
	}
	log.Debug().Msgf("contrib/segmentio/kafka.go.v0: Wrapping Writer: %#v", client.cfg)
	return client
}
func (c *Client) startSpan(ctx context.Context, msg *asynq.Task, asynqOpts ...asynq.Option) ddtrace.Span {
	opts := []tracer.StartSpanOption{
		tracer.ServiceName(c.cfg.producerServiceName),
		tracer.SpanType(ext.SpanTypeMessageProducer),
		tracer.Tag(ext.Component, "asynq"),
		tracer.Tag(ext.SpanKind, ext.SpanKindProducer),
		tracer.Tag("messaging.system", "asynq"),
		tracer.Measured(),
	}
	for _, v := range asynqOpts {
		if v.Type() == asynq.QueueOpt {
			opts = append(opts, tracer.Tag("asynq.queue", v.String()))
		}
		if v.Type() == asynq.GroupOpt {
			opts = append(opts, tracer.Tag("asynq.group", v.String()))
		}
	}

	if !math.IsNaN(c.cfg.analyticsRate) {
		opts = append(opts, tracer.Tag(ext.EventSampleRate, c.cfg.analyticsRate))
	}
	carrier := messageCarrier{msg}
	span, _ := tracer.StartSpanFromContext(ctx, "asynq.enqueue", opts...)
	err := tracer.Inject(span.Context(), carrier)
	log.Debug().Msgf("contrib/segmentio/kafka.go.v0: Failed to inject span context into carrier, %v", err)
	return span
}

func finishSpan(span ddtrace.Span, taskInfo *asynq.TaskInfo, err error) {
	span.SetTag("task.id", taskInfo.ID)
	span.SetTag("task.type", taskInfo.Type)
	span.Finish(tracer.WithError(err))
}

// Enqueue enqueues a task and returns a task info.
func (c *Client) Enqueue(task *asynq.Task, opts ...asynq.Option) (taskInfo *asynq.TaskInfo, err error) {
	span := c.startSpan(context.Background(), task, opts...)
	defer func() {
		finishSpan(span, taskInfo, err)
	}()
	taskInfo, err = c.Client.Enqueue(task, opts...)
	return taskInfo, err
}
