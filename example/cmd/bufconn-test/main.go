package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	//	_ "github.com/fluffy-bunny/grpcdotnetgo/example/internal/plugin"
	_ "github.com/fluffy-bunny/grpcdotnetgo/example/internal/plugin"
	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/pkg/core"
	"google.golang.org/grpc"
	bufconn "google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func main() {
	lis := bufconn.Listen(bufSize)
	go func() {
		grpcdotnetgocore.Start(lis)
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGreeter2Client(conn)
	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name:      "zep",
		Directive: pb.HelloDirectives_HELLO_DIRECTIVES_UNKNOWN,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}
