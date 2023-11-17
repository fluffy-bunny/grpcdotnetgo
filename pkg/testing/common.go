package testing

import (
	"context"
	"net"
	"testing"

	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/pkg/async"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/pkg/core"
	"github.com/golang/mock/gomock"
	"github.com/reugn/async"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	bufconn "google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// ExecuteWithPromiseAsync ...
func ExecuteWithPromiseAsync(runtime *grpcdotnetgocore.Runtime, lis *bufconn.Listener, plugins []plugin.IGRPCDotNetGoPlugin) async.Future[grpcdotnetgoasync.AsyncResponse] {
	future := grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise[grpcdotnetgoasync.AsyncResponse]) {
		var err error
		defer func() {
			promise.Success(&grpcdotnetgoasync.AsyncResponse{
				Message: "End Serve - grpc Server",
				Error:   err,
			})
		}()
		runtime.StartWithListenterAndPlugins(lis, plugins)
	})
	return future
}

// CreateConnection ...
func CreateConnection(ctx context.Context, lis *bufconn.Listener) (*grpc.ClientConn, error) {

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	return conn, nil
}

// CreateForEach ...
func CreateForEach(setUp func(*testing.T), tearDown func()) func(*testing.T, func(*gomock.Controller)) {
	return func(t *testing.T, testFunc func(ctrl *gomock.Controller)) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		setUp(t)
		testFunc(ctrl)
		tearDown()
	}
}

// RunTest ...
var RunTest = CreateForEach(setUp, tearDown)

func setUp(t *testing.T) {
	// SETUP METHOD WHICH IS REQUIRED TO RUN FOR EACH TEST METHOD
	// your code here
	t.Setenv("APPLICATION_ENVIRONMENT", "Test")
}

func tearDown() {
	// TEAR DOWN METHOD WHICH IS REQUIRED TO RUN FOR EACH TEST METHOD
	// your code here
}
