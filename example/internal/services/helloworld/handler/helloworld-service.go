package handler

import (
	"fmt"

	"strings"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	contracts_config "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/config"
	contracts_lambda "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/lambda"
	contracts_scoped "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/scoped"
	contracts_singleton "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/singleton"
	contracts_transient "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/transient"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
	contracts_serviceprovider "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/serviceprovider"
	contracts_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/timeutils"
	contracts_uuid "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/uuid"
	grpc_error "github.com/fluffy-bunny/grpcdotnetgo/pkg/grpc/error"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	Container       di.Container                               `inject:""`
	ServiceProvider contracts_serviceprovider.IServiceProvider `inject:""`
	Request         contracts_request.IRequest                 `inject:""`
	ScopedItems     contracts_request.IItems                   `inject:""`
	ClaimsPrincipal contracts_claimsprincipal.IClaimsPrincipal `inject:""`
	Config          *contracts_config.Config                   `inject:""`
	Singleton       contracts_singleton.ISingleton             `inject:""`
	Scoped          contracts_scoped.IScoped                   `inject:""`
	Transients      []contracts_transient.ITransient           `inject:""`
	Transient       contracts_transient.ITransient             `inject:""`
	TimeNow         contracts_timeutils.TimeNow                `inject:""`
	TimeParse       contracts_timeutils.TimeParse              `inject:""`
	Time            contracts_timeutils.ITime                  `inject:""`
	TimeUtils       contracts_timeutils.ITimeUtils             `inject:""`
	KSUID           contracts_uuid.IKSUID                      `inject:""`
	GenerateUUID    contracts_lambda.GenerateUUID              `inject:""`
	GenerateUUIDs   []contracts_lambda.GenerateUUID            `inject:""`
	instanceID      string
	multiUUIDs      []string
}

// Ctor if it exists is called when the service is created
func (s *Service) Ctor() {
	s.instanceID = s.GenerateUUID()
	builder := strings.Builder{}
	for _, t := range s.GenerateUUIDs {
		builder.WriteString(t())
		builder.WriteString(":")
		s.multiUUIDs = append(s.multiUUIDs, t())
	}

}
func (s *Service) getLogger() *zerolog.Logger {
	ctx := s.Request.GetContext()
	l := zerolog.Ctx(ctx)
	return l
}

// Close if it exists is called when the container is torn down
func (s *Service) Close() {
	logger := s.getLogger()
	logger.Info().Msg("Close")
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(in *pb.HelloRequest) (*pb.HelloReply, error) {
	logger := s.getLogger()
	logger.Info().Msg("Enter")
	fmt.Println(internal.PrettyJSON(s.ClaimsPrincipal))

	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_PANIC {
		panic("shits on fire, yo")
	}
	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_ERROR {
		br := grpc_error.NewBadRequest()
		desc := "The username must only contain alphanumeric characters"
		br.AddViolation("username", desc)
		errst := br.GetStatusError(codes.InvalidArgument, "HelloDirectives_HELLO_DIRECTIVES_ERROR")
		//	err := status.Errorf(codes.Internal, "%v", pb.HelloDirectives_HELLO_DIRECTIVES_ERROR)
		return nil, errst
	}
	logger.Info().Msgf("Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// Service2 ...
type Service2 struct {
	Request         contracts_request.IRequest                 `inject:""`
	ClaimsPrincipal contracts_claimsprincipal.IClaimsPrincipal `inject:""`
	Config          *contracts_config.Config                   `inject:""`
}

func (s *Service2) getLogger() *zerolog.Logger {
	ctx := s.Request.GetContext()
	l := zerolog.Ctx(ctx)
	return l
}

// SayHello implements helloworld.GreeterServer
func (s *Service2) SayHello(in *pb.HelloRequest) (*pb.HelloReply2, error) {
	logger := s.getLogger()
	logger.Info().Msg("Enter")
	fmt.Println(internal.PrettyJSON(s.ClaimsPrincipal))

	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_PANIC {
		panic("shits on fire, yo")
	}
	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_ERROR {
		reply := &pb.HelloReply2{}
		err := fmt.Errorf("Ermaghd")
		return reply, err
	}
	logger.Info().Msgf("Received: %v", in.GetName())
	return &pb.HelloReply2{Message: "Hello " + in.GetName()}, nil
}
