package services

import (
	"context"
	"sync"

	ms "github.com/matany90/RESTful-gRPC-simple-starters/go/grpc/pkg/matan-service"
)

// MatanServer struct.
// handlers will be implemented and attach to this struct.
type MatanService struct {
	ms.UnimplementedMatanServiceServer
	wg *sync.WaitGroup
}

// init Matan service instance
func NewMatanService(wg *sync.WaitGroup) (*MatanService, error) {
	return &MatanService{
		wg:       wg,
	}, nil
}

// SayHello handler implement.
func (s *MatanService) SayHello(ctx context.Context, in *ms.HelloRequest) (*ms.HelloReply, error) {
	// returns an hello response
	return &ms.HelloReply{
		Message: "Hello from matan's service.",
	}, nil
}

// SayHi handler implement.
func (s *MatanService) SayHi(ctx context.Context, in *ms.HelloRequest) (*ms.HelloReply, error) {
	// returns an hello response
	return &ms.HelloReply{
		Message: "Hi from matan's service.",
	}, nil
}