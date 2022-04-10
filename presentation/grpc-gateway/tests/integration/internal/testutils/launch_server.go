package testutils

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/pocket7878/go-ddd-learning/infra/ent"
	gw "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
	"github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/grpc_service"
)

const _PORT = 8081

type LaunchTestServerOpts struct {
	Client *ent.Client
}

type option interface {
	Apply(*LaunchTestServerOpts)
}

type entClient struct {
	client *ent.Client
}

func (c entClient) Apply(opts *LaunchTestServerOpts) {
	opts.Client = c.client
}

func WithEntClient(client *ent.Client) option {
	return entClient{client}
}

func LaunchTestServer(t *testing.T, options ...option) (func(), error) {
	launchOption := &LaunchTestServerOpts{
		Client: BuildDbConnection(t),
	}
	for _, opt := range options {
		opt.Apply(launchOption)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Register gRPC server endpoint
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	grpcService := grpc_service.NewGrpcService(launchOption.Client)
	gw.RegisterGoDDDLearningServiceServer(grpcServer, grpcService)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterGoDDDLearningServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", _PORT), opts)
	if err != nil {
		return func() { launchOption.Client.Close(); cancel() }, err
	}

	httpServer := http.Server{
		Handler: mux,
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", _PORT))
	if err != nil {
		return func() { launchOption.Client.Close(); cancel() }, err
	}
	m := cmux.New(l)
	httpListener := m.Match(cmux.HTTP1Fast())
	grpcListener := m.Match(cmux.HTTP2())

	go func() {
		err := grpcServer.Serve(grpcListener)
		if err != nil && err != grpc.ErrServerStopped && err != grpc.ErrClientConnClosing {
			t.Fatalf("failed to serve grpc: %v", err)
		}
	}()

	go func() {
		err := httpServer.Serve(httpListener)
		if err != nil && err != cmux.ErrListenerClosed && err != cmux.ErrServerClosed {
			t.Fatalf("failed to serve http: %v", err)
		}
	}()

	go func() {
		err := m.Serve()
		if !isAcceptableCmuxError(err) {
			t.Fatalf("failed to serve mux: %v", err)
		}
	}()

	return func() {
		m.Close()
		launchOption.Client.Close()
		cancel()
	}, nil
}

func isAcceptableCmuxError(err error) bool {
	if err == nil {
		return true
	}

	return !strings.Contains(err.Error(), "use of closed") && err != cmux.ErrServerClosed && err != cmux.ErrListenerClosed && err != http.ErrServerClosed
}
