package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/cmd/server"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/logging"
)

// Init runs the setup process
// for main package.
// The init function will initialize global
// go objects that will be "alive" while server is up.
func init() {
	// init logging object
	logging.Setup()

	// init server object
	server.Setup()
}

// main function will runs server instance.
func main() {
	// start the server
	server.Svr.Start()

	// graceful exit
	sig := make(chan os.Signal, 1)

	// notify signal to to channel
	// while the server instance making an exit.
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	// catch the signal input, when throw a notify call
	<-sig

	// stopping server
	defer server.Svr.Stop()
}