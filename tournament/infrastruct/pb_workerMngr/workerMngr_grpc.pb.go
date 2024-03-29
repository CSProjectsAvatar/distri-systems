// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1--rc1
// source: workerMngr.proto

package pb_workerMngr

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

// WorkerMngrClient is the client API for WorkerMngr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerMngrClient interface {
	// Returns a Match for the worker to run
	GiveMeWork(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchResp, error)
	// Receive the result of running a Match
	CatchResult(ctx context.Context, in *ResultReq, opts ...grpc.CallOption) (*ResultResp, error)
}

type workerMngrClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerMngrClient(cc grpc.ClientConnInterface) WorkerMngrClient {
	return &workerMngrClient{cc}
}

func (c *workerMngrClient) GiveMeWork(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchResp, error) {
	out := new(MatchResp)
	err := c.cc.Invoke(ctx, "/pb_workerMngr.WorkerMngr/GiveMeWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerMngrClient) CatchResult(ctx context.Context, in *ResultReq, opts ...grpc.CallOption) (*ResultResp, error) {
	out := new(ResultResp)
	err := c.cc.Invoke(ctx, "/pb_workerMngr.WorkerMngr/CatchResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerMngrServer is the server API for WorkerMngr service.
// All implementations must embed UnimplementedWorkerMngrServer
// for forward compatibility
type WorkerMngrServer interface {
	// Returns a Match for the worker to run
	GiveMeWork(context.Context, *MatchReq) (*MatchResp, error)
	// Receive the result of running a Match
	CatchResult(context.Context, *ResultReq) (*ResultResp, error)
	mustEmbedUnimplementedWorkerMngrServer()
}

// UnimplementedWorkerMngrServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerMngrServer struct {
}

func (UnimplementedWorkerMngrServer) GiveMeWork(context.Context, *MatchReq) (*MatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveMeWork not implemented")
}
func (UnimplementedWorkerMngrServer) CatchResult(context.Context, *ResultReq) (*ResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CatchResult not implemented")
}
func (UnimplementedWorkerMngrServer) mustEmbedUnimplementedWorkerMngrServer() {}

// UnsafeWorkerMngrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerMngrServer will
// result in compilation errors.
type UnsafeWorkerMngrServer interface {
	mustEmbedUnimplementedWorkerMngrServer()
}

func RegisterWorkerMngrServer(s grpc.ServiceRegistrar, srv WorkerMngrServer) {
	s.RegisterService(&WorkerMngr_ServiceDesc, srv)
}

func _WorkerMngr_GiveMeWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerMngrServer).GiveMeWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_workerMngr.WorkerMngr/GiveMeWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerMngrServer).GiveMeWork(ctx, req.(*MatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerMngr_CatchResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerMngrServer).CatchResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_workerMngr.WorkerMngr/CatchResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerMngrServer).CatchResult(ctx, req.(*ResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerMngr_ServiceDesc is the grpc.ServiceDesc for WorkerMngr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerMngr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_workerMngr.WorkerMngr",
	HandlerType: (*WorkerMngrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GiveMeWork",
			Handler:    _WorkerMngr_GiveMeWork_Handler,
		},
		{
			MethodName: "CatchResult",
			Handler:    _WorkerMngr_CatchResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workerMngr.proto",
}
