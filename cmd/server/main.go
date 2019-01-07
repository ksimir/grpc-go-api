package main

import (
	"fmt"
	"os"

	cmd "github.com/ksimir/grpc-go-api/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
