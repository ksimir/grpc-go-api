package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"
	"time"

	v1 "github.com/ksimir/grpc-go-api/pkg/api/v1/inventory"
	"github.com/ksimir/grpc-go-api/pkg/logger"
	"github.com/ksimir/grpc-go-api/pkg/protocol/grpc/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// RunServer runs gRPC service to publish Player service
func RunServer(ctx context.Context, v1API v1.InventoryServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)
	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}))

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterInventoryServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
