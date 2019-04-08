package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	// spanner client lib
	"cloud.google.com/go/spanner"
	"github.com/ksimir/grpc-go-api/pkg/logger"
	grpc "github.com/ksimir/grpc-go-api/pkg/protocol/grpc/inventory-server"
	v1 "github.com/ksimir/grpc-go-api/pkg/service/v1/inventory-service"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC serve
	GRPCPort string

	// Spanner DB parameters section
	// SpannerProjectName is the project name that host spanner
	SpannerProjectName string
	// SpannerInstanceName is the Spanner instance
	SpannerInstanceName string
	// SpannerDatabaseName is the Spanner database
	SpannerDatabaseName string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.SpannerProjectName, "project", "", "Project name")
	flag.StringVar(&cfg.SpannerInstanceName, "instance", "", "Instance name")
	flag.StringVar(&cfg.SpannerDatabaseName, "database", "", "Database name")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	dsn := fmt.Sprintf("projects/%s/instances/%s/databases/%s",
		cfg.SpannerProjectName,
		cfg.SpannerInstanceName,
		cfg.SpannerDatabaseName)
	dataClient, err := spanner.NewClient(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dataClient.Close()

	v1API := v1.NewInventoryServiceServer(dataClient)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
