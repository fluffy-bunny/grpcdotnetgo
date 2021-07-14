package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	_ "github.com/fluffy-bunny/grpcdotnetgo/proto/error"
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
		log.Fatalf("did not connect: %v", err)
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
		Directive: pb.HelloDirectives_HELLO_DIRECTIVES_PANIC,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", internal.PrettyJSON(r))

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	r2, err := c2.SayHello(ctx2, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", internal.PrettyJSON(r2))

}
