package handler

import (
	"fmt"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	ContextAccessor contextaccessor.IContextAccessor
	ClaimsPrincipal claimsprincipal.IClaimsPrincipal
	Logger          servicesLogger.ILogger
	ServiceProvider servicesServiceProvider.IServiceProvider
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.Logger.Info().Msg("Enter")

	fmt.Println(internal.PrettyJSON(s.ClaimsPrincipal))
	s.Logger.Info().Msgf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
