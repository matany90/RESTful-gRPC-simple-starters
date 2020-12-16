package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/cmd/server"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/logging"
)

// Setup packages global instances
// for grpc service
func init() {
	// set logging global instace
	logging.Setup()

	// set server global instace
	server.Setup()
}

// grpc service entry point
func main() {
	// start server
	server.Svr.Start()

	// graceful exit
	sig := make(chan os.Signal, 1)

	// notify any signal to channel
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	// return os signal channel as exit sign
	<-sig

	// stopping the server
	defer server.Svr.Stop()
}