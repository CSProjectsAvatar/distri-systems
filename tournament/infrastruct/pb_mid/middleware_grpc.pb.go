// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1--rc1
// source: middleware.proto

package pb_mid

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	UploadTournament(ctx context.Context, in *TournamentReq, opts ...grpc.CallOption) (*TournamentResp, error)
	RunTournament(ctx context.Context, in *RunReq, opts ...grpc.CallOption) (Middleware_RunTournamentClient, error)
	GetStats(ctx context.Context, in *StatsReq, opts ...grpc.CallOption) (*StatsResp, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) UploadTournament(ctx context.Context, in *TournamentReq, opts ...grpc.CallOption) (*TournamentResp, error) {
	out := new(TournamentResp)
	err := c.cc.Invoke(ctx, "/pb.Middleware/UploadTournament", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) RunTournament(ctx context.Context, in *RunReq, opts ...grpc.CallOption) (Middleware_RunTournamentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Middleware_ServiceDesc.Streams[0], "/pb.Middleware/RunTournament", opts...)
	if err != nil {
		return nil, err
	}
	x := &middlewareRunTournamentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Middleware_RunTournamentClient interface {
	Recv() (*RunResp, error)
	grpc.ClientStream
}

type middlewareRunTournamentClient struct {
	grpc.ClientStream
}

func (x *middlewareRunTournamentClient) Recv() (*RunResp, error) {
	m := new(RunResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *middlewareClient) GetStats(ctx context.Context, in *StatsReq, opts ...grpc.CallOption) (*StatsResp, error) {
	out := new(StatsResp)
	err := c.cc.Invoke(ctx, "/pb.Middleware/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	UploadTournament(context.Context, *TournamentReq) (*TournamentResp, error)
	RunTournament(*RunReq, Middleware_RunTournamentServer) error
	GetStats(context.Context, *StatsReq) (*StatsResp, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) UploadTournament(context.Context, *TournamentReq) (*TournamentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadTournament not implemented")
}
func (UnimplementedMiddlewareServer) RunTournament(*RunReq, Middleware_RunTournamentServer) error {
	return status.Errorf(codes.Unimplemented, "method RunTournament not implemented")
}
func (UnimplementedMiddlewareServer) GetStats(context.Context, *StatsReq) (*StatsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

func _Middleware_UploadTournament_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TournamentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UploadTournament(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middleware/UploadTournament",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UploadTournament(ctx, req.(*TournamentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_RunTournament_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MiddlewareServer).RunTournament(m, &middlewareRunTournamentServer{stream})
}

type Middleware_RunTournamentServer interface {
	Send(*RunResp) error
	grpc.ServerStream
}

type middlewareRunTournamentServer struct {
	grpc.ServerStream
}

func (x *middlewareRunTournamentServer) Send(m *RunResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Middleware_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middleware/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetStats(ctx, req.(*StatsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadTournament",
			Handler:    _Middleware_UploadTournament_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Middleware_GetStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunTournament",
			Handler:       _Middleware_RunTournament_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "middleware.proto",
}
