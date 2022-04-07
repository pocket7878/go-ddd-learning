package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/pocket7878/go-ddd-learning/infra"
	gw "github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/gen/go"
	"github.com/pocket7878/go-ddd-learning/presentation/grpc-gateway/grpc_service"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8081", "gRPC server endpoint")
)

func run() error {
	client, cleanupFunc, err := infra.BuildEntClient()
	defer cleanupFunc()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	s := grpc.NewServer()
	grpcService := grpc_service.NewGrpcService(client)
	gw.RegisterGoDDDLearningServiceServer(s, grpcService)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = gw.RegisterGoDDDLearningServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler: withLogger(mux),
	}
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	m := cmux.New(l)
	httpListener := m.Match(cmux.HTTP1Fast())
	grpcListener := m.Match(cmux.HTTP2())

	go server.Serve(httpListener)
	go s.Serve(grpcListener)

	return m.Serve()
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
