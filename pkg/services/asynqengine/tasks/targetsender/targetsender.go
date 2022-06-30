package targetsender

import (
	"context"
	"fmt"
	"reflect"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	core_hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/hibiken/asynq"
	contracts_background_tasks_targetsender "go.mapped.dev/micro/job-worker/internal/contracts/background/tasks/targetsender"
	cloud_api_jobworker "go.mapped.dev/proto/cloud/api/jobworker"
	cloud_api_webhooks_senderv2 "go.mapped.dev/proto/cloud/api/webhooks/senderv2"
	contracts_background_tasks "go.mapped.dev/webhooks-sdk-go/pkg/contracts/background/tasks"
	contracts_jobs_targetsubscriptions "go.mapped.dev/webhooks-sdk-go/pkg/contracts/jobs/targetsubscriptions"
	contracts_webhookssenderclient "go.mapped.dev/webhooks-sdk-go/pkg/contracts/webhookssenderclient"
	"google.golang.org/protobuf/proto"
)

type (
	service struct {
		Logger                      contracts_logger.ILogger                                        `inject:""`
		MiniRedisTaskClient         contracts_background_tasks.IMiniRedisTaskClient                 `inject:""`
		RedisTaskClient             contracts_background_tasks.IRedisTaskClient                     `inject:""`
		FetchTargetSubscriptionsJob contracts_jobs_targetsubscriptions.IFetchTargetSubscriptionsJob `inject:""`
		WebhooksSenderClient        contracts_webhookssenderclient.IWebhooksSenderClient            `inject:""`
		client                      contracts_webhookssenderclient.IWebhooksSenderServiceClient
	}
)

func assertImplementation() {
	var _ contracts_background_tasks.ISingletonTask = (*service)(nil)
	var _ contracts_background_tasks_targetsender.ITargetSenderSingletonTask = (*service)(nil)

}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonISingletonTask registers the *service as a singleton.
func AddSingletonISingletonTask(builder *di.Builder) {
	contracts_background_tasks.AddSingletonISingletonTask(builder, reflectType,
		contracts_background_tasks_targetsender.ReflectTypeITargetSenderSingletonTask)
}
func (s *service) Ctor() {}

func (s *service) tryGetWebhooksSenderGRPCClient() (contracts_webhookssenderclient.IWebhooksSenderServiceClient, error) {
	var err error
	if s.client == nil {
		s.client, err = s.WebhooksSenderClient.GetGRPCClient()
	}
	return s.client, err
}
func (s *service) GetPatterns() *core_hashset.StringSet {
	return core_hashset.NewStringSet(contracts_background_tasks_targetsender.TypeWebhooksMessage)
}

func (s *service) ProcessTask(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case contracts_background_tasks_targetsender.TypeWebhooksMessage:
		return s.processWebhooksMesage(ctx, t)
	default:
		return fmt.Errorf("unknown task type: %s", t.Type())
	}
}
func (s *service) processWebhooksMesage(ctx context.Context, t *asynq.Task) error {
	var p cloud_api_jobworker.WebhookMessage
	if err := proto.Unmarshal(t.Payload(), &p); err != nil {
		s.Logger.Error().Err(err).Msg("failed to unmarshal task payload")
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	client, err := s.tryGetWebhooksSenderGRPCClient()
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to get grpc client")
		return fmt.Errorf("failed to get grpc client: %w", err)
	}
	s.Logger.Trace().Interface("webhookMessage", &p).Msg("processing task")

	response, err := client.SendEvent(ctx, &cloud_api_webhooks_senderv2.SendEventRequest{
		TargetId: p.TargetId,
		OrgId:    p.OrgId,
		Event:    p.Event,
	})
	if err != nil {
		s.Logger.Error().Err(err).Msg("failed to send event")
		return fmt.Errorf("failed to send event: %w", err)
	}
	s.Logger.Trace().Msgf("response: %v", response)
	return nil
}

func (s *service) EnqueTaskWebhookMessage(webhookMessage *cloud_api_jobworker.WebhookMessage, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	targetConfig, ok := s.FetchTargetSubscriptionsJob.GetCurrentTargetSubscriptions().ByTargetMap[webhookMessage.TargetId]
	if !ok {
		return nil, fmt.Errorf("target not found: %s", webhookMessage.TargetId)
	}
	name := contracts_background_tasks_targetsender.TypeWebhooksMessage

	payloadJson, err := proto.Marshal(webhookMessage)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(name, payloadJson)
	// retry is message specific
	if targetConfig.WebhookTarget.Durable {
		opts = append(opts, asynq.MaxRetry(int(targetConfig.WebhookTarget.MaxRetry)))
		return s.RedisTaskClient.EnqueTask(task, opts...)
	}
	opts = append(opts, asynq.MaxRetry(0)) // non durable
	return s.MiniRedisTaskClient.EnqueTask(task, opts...)
}
