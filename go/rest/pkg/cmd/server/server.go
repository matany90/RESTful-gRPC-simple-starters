package server

import (
	"net/http"
	"sync"
	"net"
	"context"
	"time"

	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/utils"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/routers"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/logging"
)

// Server struct that defines
// the internal server structure
type Server struct {
	srv *http.Server
	wg 	sync.WaitGroup
}

// define server global instance
var Svr *Server

// Setup function init server
// global instace
func Setup() {
	Svr = &Server{
		srv: &http.Server{
			Addr: net.JoinHostPort("0.0.0.0", utils.GetEnv("PORT", "50003")),
			Handler: routers.InitRouters(),
		},
	}
}

// Start running the server
func (s *Server) Start() {
	// initialize context without timeout
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add to waitgroup for listener goroutine
	s.wg.Add(1)

	// start the server
	go func() {
		logging.Log.Info("Running RESTful GO server on port " +  utils.GetEnv("PORT", "50003") + "...")

		// add to the waitgroup
		s.srv.ListenAndServe()

		// close waitgroup
		s.wg.Done()
	}()
}

// Stop the server
func (s *Server) Stop() {
	// init context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// handle any error while stopping the server
	if err := s.srv.Shutdown(ctx); err != nil {
		if err = s.srv.Close(); err != nil {
			logging.Log.Error(err.Error())
			return
		}
	}

	// waitgroup wait
	s.wg.Wait()
}
