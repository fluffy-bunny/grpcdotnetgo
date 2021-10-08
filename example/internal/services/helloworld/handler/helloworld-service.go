package handler

import (
	"fmt"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	grpcError "github.com/fluffy-bunny/grpcdotnetgo/pkg/grpc/error"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	"google.golang.org/grpc/codes"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	ContextAccessor contextaccessor.IContextAccessor
	ClaimsPrincipal claimsprincipalContracts.IClaimsPrincipal
	Logger          servicesLogger.ILogger
	config          *internal.Config
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.Logger.Info().Msg("Enter")
	fmt.Println(internal.PrettyJSON(s.ClaimsPrincipal))

	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_PANIC {
		panic("shits on fire, yo")
	}
	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_ERROR {
		br := grpcError.NewBadRequest()
		desc := "The username must only contain alphanumeric characters"
		br.AddViolation("username", desc)
		errst := br.GetStatusError(codes.InvalidArgument, "HelloDirectives_HELLO_DIRECTIVES_ERROR")
		//	err := status.Errorf(codes.Internal, "%v", pb.HelloDirectives_HELLO_DIRECTIVES_ERROR)
		return nil, errst
	}
	s.Logger.Info().Msgf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type Service2 struct {
	ContextAccessor contextaccessor.IContextAccessor
	ClaimsPrincipal claimsprincipalContracts.IClaimsPrincipal
	Logger          servicesLogger.ILogger
	config          *internal.Config
}

// SayHello implements helloworld.GreeterServer
func (s *Service2) SayHello(in *pb.HelloRequest) (*pb.HelloReply2, error) {
	s.Logger.Info().Msg("Enter")
	fmt.Println(internal.PrettyJSON(s.ClaimsPrincipal))

	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_PANIC {
		panic("shits on fire, yo")
	}
	if in.Directive == pb.HelloDirectives_HELLO_DIRECTIVES_ERROR {
		reply := &pb.HelloReply2{}
		err := fmt.Errorf("Ermaghd")
		return reply, err
	}
	s.Logger.Info().Msgf("Received: %v", in.GetName())
	return &pb.HelloReply2{Message: "Hello " + in.GetName()}, nil
}
