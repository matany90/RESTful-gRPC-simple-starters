package server

import (
	"context"
	"net"
	"sync"

	// "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	// "github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/middleware"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/logging"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/utils"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/services"
	ms "github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/matan-service"
)

// Define the server struct
type Server struct {
	srv *grpc.Server
	wg 	sync.WaitGroup
}

// server instace
var Svr *Server

// Setup the server instance
func Setup() {
	// init server global instance
	Svr = &Server{
		srv: grpc.NewServer(),
	}
}

// Start running server instance
func (s *Server) Start() {
	// init context without timeout
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add to the waitgroup
	// for the listener goroutine
	s.wg.Add(1)

	// start the server
	go func() {
		// define port
		port := utils.GetEnv("PORT", "50004")

		// define net instance
		lis, err := net.Listen("tcp", ":"+port)
		if err != nil {
			logging.Log.Error("error while init service.")
		}

		// init service
		matanServerInstance, err := services.NewMatanService(&s.wg)
		if err != nil {
			logging.Log.Error("error while init service.")
		}

		// register matan gRPC service instance
		ms.RegisterMatanServiceServer(Svr.srv, matanServerInstance)

		// log to listen port
		logging.Log.Info("Listen on port " + port + "...")
		if err := Svr.srv.Serve(lis); err != nil {
			logging.Log.Error("failed to serve: %v", err)
		}

		// close wait group
		s.wg.Done()
	}()
}

// Stop the server
func (s *Server) Stop() {
	// handle any error while stopping the server
	Svr.srv.Stop()

	s.wg.Wait()
}