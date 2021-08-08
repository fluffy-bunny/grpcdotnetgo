package main

import (
	"context"
	"fmt"
	"reflect"

	"os"
	"time"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	_ "github.com/fluffy-bunny/grpcdotnetgo/proto/error"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:5105"
	defaultName = "world"
)

var version = "development"

func main() {
	fmt.Println("Version:\t", version)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal().Err(err).Msg("did not connect")
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c2 := pb.NewGreeter2Client(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name:      name,
		Directive: pb.HelloDirectives_HELLO_DIRECTIVES_ERROR,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not greet")
		st, ok := status.FromError(err)
		if ok {
			log.Error().Interface("statusError", st).Send()
			st = status.Convert(err)

			for _, detail := range st.Details() {
				fmt.Println(reflect.TypeOf(detail))
				log.Error().Interface("detail", detail).Send()
				switch t := detail.(type) {
				case *rpc.BadRequest:
					fmt.Println("Oops! Your request was rejected by the server.")

					for _, violation := range t.GetFieldViolations() {
						fmt.Printf("The %q field was wrong:\n", violation.GetField())
						fmt.Printf("\t%s\n", violation.GetDescription())
					}
				}
			}

		}
		return
	}
	log.Printf("Greeting: %s", internal.PrettyJSON(r))

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	r2, err := c2.SayHello(ctx2, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Error().Err(err).Msg("could not greet")
		st, ok := status.FromError(err)
		if ok {
			log.Error().Interface("statusError", st).Send()
		}
	}
	log.Printf("Greeting: %s", internal.PrettyJSON(r2))

}
