// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package matanservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MatanServiceClient is the client API for MatanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatanServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type matanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatanServiceClient(cc grpc.ClientConnInterface) MatanServiceClient {
	return &matanServiceClient{cc}
}

func (c *matanServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/matanservice.MatanService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matanServiceClient) SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/matanservice.MatanService/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatanServiceServer is the server API for MatanService service.
// All implementations must embed UnimplementedMatanServiceServer
// for forward compatibility
type MatanServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHi(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedMatanServiceServer()
}

// UnimplementedMatanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatanServiceServer struct {
}

func (UnimplementedMatanServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMatanServiceServer) SayHi(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedMatanServiceServer) mustEmbedUnimplementedMatanServiceServer() {}

// UnsafeMatanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatanServiceServer will
// result in compilation errors.
type UnsafeMatanServiceServer interface {
	mustEmbedUnimplementedMatanServiceServer()
}

func RegisterMatanServiceServer(s grpc.ServiceRegistrar, srv MatanServiceServer) {
	s.RegisterService(&_MatanService_serviceDesc, srv)
}

func _MatanService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatanServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matanservice.MatanService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatanServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatanService_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatanServiceServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matanservice.MatanService/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatanServiceServer).SayHi(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MatanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matanservice.MatanService",
	HandlerType: (*MatanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _MatanService_SayHello_Handler,
		},
		{
			MethodName: "SayHi",
			Handler:    _MatanService_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matan-service.proto",
}
