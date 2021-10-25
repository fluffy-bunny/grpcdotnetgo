package grpcserver

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal/plugin"
	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/pkg/async"
	pluginContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
	"github.com/reugn/async"

	//	_ "github.com/fluffy-bunny/grpcdotnetgo/example/internal/plugin"
	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/pkg/core"
	"google.golang.org/grpc"
	bufconn "google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// TODO: fix this
func TestSayHello(t *testing.T) {
	var plugins []pluginContracts.IGRPCDotNetGoPlugin
	plugins = append(plugins, plugin.NewPlugin())
	lis := bufconn.Listen(bufSize)
	myRuntime := grpcdotnetgocore.NewRuntime()
	future := grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise) {
		var err error

		defer func() {
			promise.Success(&grpcdotnetgoasync.AsyncResponse{
				Message: "End Serve - grpc Server",
				Error:   err,
			})
		}()

		myRuntime.StartWithListenterAndPlugins(lis, plugins)
	})

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeter2Client(conn)

	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name:      "zep",
		Directive: pb.HelloDirectives_HELLO_DIRECTIVES_UNKNOWN,
	})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
	myRuntime.Stop()
	future.Get()
}
