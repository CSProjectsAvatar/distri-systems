// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1--rc1
// source: chord.proto

package pb_chord

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

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChordClient interface {
	// GetPredecessor returns the node believed to be the current predecessor.
	GetPredecessor(ctx context.Context, in *EmptyR, opts ...grpc.CallOption) (*Node, error)
	// GetSuccessor returns the node believed to be the current successor.
	GetSuccessor(ctx context.Context, in *EmptyR, opts ...grpc.CallOption) (*Node, error)
	// Notify notifies Chord that Node thinks it is our predecessor. This has
	// the potential to initiate the transferring of keys.
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyR, error)
	// FindSuccessor finds the node the succedes ID. May initiate RPC calls to
	// other nodes.
	FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	// CheckPredecessor checkes whether predecessor has failed.
	CheckPredecessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EmptyR, error)
	// SetPredecessor sets predecessor for a node.
	SetPredecessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyR, error)
	// SetPredecessor sets predecessor for a node.
	SetSuccessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyR, error)
	// Get returns the value in Chord ring for the given key.
	KGet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Set writes a key value pair to the Chord ring.
	KSet(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	// Delete returns the value in Chord ring for the given key.
	KDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Multiple delete returns the value in Chord ring between the given keys.
	KMultiDelete(ctx context.Context, in *MultiDeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// RequestKeys returns the keys between given range from the Chord ring.
	KRequestKeys(ctx context.Context, in *RequestKeysRequest, opts ...grpc.CallOption) (*RequestKeysResponse, error)
}

type chordClient struct {
	cc grpc.ClientConnInterface
}

func NewChordClient(cc grpc.ClientConnInterface) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) GetPredecessor(ctx context.Context, in *EmptyR, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/GetPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetSuccessor(ctx context.Context, in *EmptyR, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/GetSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyR, error) {
	out := new(EmptyR)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) CheckPredecessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EmptyR, error) {
	out := new(EmptyR)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/CheckPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetPredecessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyR, error) {
	out := new(EmptyR)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/SetPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) SetSuccessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyR, error) {
	out := new(EmptyR)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/SetSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) KGet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/KGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) KSet(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/KSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) KDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/KDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) KMultiDelete(ctx context.Context, in *MultiDeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/KMultiDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) KRequestKeys(ctx context.Context, in *RequestKeysRequest, opts ...grpc.CallOption) (*RequestKeysResponse, error) {
	out := new(RequestKeysResponse)
	err := c.cc.Invoke(ctx, "/pb_chord.Chord/KRequestKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordServer is the server API for Chord service.
// All implementations must embed UnimplementedChordServer
// for forward compatibility
type ChordServer interface {
	// GetPredecessor returns the node believed to be the current predecessor.
	GetPredecessor(context.Context, *EmptyR) (*Node, error)
	// GetSuccessor returns the node believed to be the current successor.
	GetSuccessor(context.Context, *EmptyR) (*Node, error)
	// Notify notifies Chord that Node thinks it is our predecessor. This has
	// the potential to initiate the transferring of keys.
	Notify(context.Context, *Node) (*EmptyR, error)
	// FindSuccessor finds the node the succedes ID. May initiate RPC calls to
	// other nodes.
	FindSuccessor(context.Context, *ID) (*Node, error)
	// CheckPredecessor checkes whether predecessor has failed.
	CheckPredecessor(context.Context, *ID) (*EmptyR, error)
	// SetPredecessor sets predecessor for a node.
	SetPredecessor(context.Context, *Node) (*EmptyR, error)
	// SetPredecessor sets predecessor for a node.
	SetSuccessor(context.Context, *Node) (*EmptyR, error)
	// Get returns the value in Chord ring for the given key.
	KGet(context.Context, *GetRequest) (*GetResponse, error)
	// Set writes a key value pair to the Chord ring.
	KSet(context.Context, *SetRequest) (*SetResponse, error)
	// Delete returns the value in Chord ring for the given key.
	KDelete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Multiple delete returns the value in Chord ring between the given keys.
	KMultiDelete(context.Context, *MultiDeleteRequest) (*DeleteResponse, error)
	// RequestKeys returns the keys between given range from the Chord ring.
	KRequestKeys(context.Context, *RequestKeysRequest) (*RequestKeysResponse, error)
	mustEmbedUnimplementedChordServer()
}

// UnimplementedChordServer must be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (UnimplementedChordServer) GetPredecessor(context.Context, *EmptyR) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessor not implemented")
}
func (UnimplementedChordServer) GetSuccessor(context.Context, *EmptyR) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuccessor not implemented")
}
func (UnimplementedChordServer) Notify(context.Context, *Node) (*EmptyR, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedChordServer) FindSuccessor(context.Context, *ID) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (UnimplementedChordServer) CheckPredecessor(context.Context, *ID) (*EmptyR, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPredecessor not implemented")
}
func (UnimplementedChordServer) SetPredecessor(context.Context, *Node) (*EmptyR, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPredecessor not implemented")
}
func (UnimplementedChordServer) SetSuccessor(context.Context, *Node) (*EmptyR, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSuccessor not implemented")
}
func (UnimplementedChordServer) KGet(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KGet not implemented")
}
func (UnimplementedChordServer) KSet(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KSet not implemented")
}
func (UnimplementedChordServer) KDelete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KDelete not implemented")
}
func (UnimplementedChordServer) KMultiDelete(context.Context, *MultiDeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KMultiDelete not implemented")
}
func (UnimplementedChordServer) KRequestKeys(context.Context, *RequestKeysRequest) (*RequestKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KRequestKeys not implemented")
}
func (UnimplementedChordServer) mustEmbedUnimplementedChordServer() {}

// UnsafeChordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChordServer will
// result in compilation errors.
type UnsafeChordServer interface {
	mustEmbedUnimplementedChordServer()
}

func RegisterChordServer(s grpc.ServiceRegistrar, srv ChordServer) {
	s.RegisterService(&Chord_ServiceDesc, srv)
}

func _Chord_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyR)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/GetPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetPredecessor(ctx, req.(*EmptyR))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyR)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/GetSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetSuccessor(ctx, req.(*EmptyR))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_CheckPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).CheckPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/CheckPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).CheckPredecessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/SetPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetPredecessor(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_SetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).SetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/SetSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).SetSuccessor(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_KGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).KGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/KGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).KGet(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_KSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).KSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/KSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).KSet(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_KDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).KDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/KDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).KDelete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_KMultiDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).KMultiDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/KMultiDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).KMultiDelete(ctx, req.(*MultiDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_KRequestKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).KRequestKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chord.Chord/KRequestKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).KRequestKeys(ctx, req.(*RequestKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chord_ServiceDesc is the grpc.ServiceDesc for Chord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_chord.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPredecessor",
			Handler:    _Chord_GetPredecessor_Handler,
		},
		{
			MethodName: "GetSuccessor",
			Handler:    _Chord_GetSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Chord_Notify_Handler,
		},
		{
			MethodName: "FindSuccessor",
			Handler:    _Chord_FindSuccessor_Handler,
		},
		{
			MethodName: "CheckPredecessor",
			Handler:    _Chord_CheckPredecessor_Handler,
		},
		{
			MethodName: "SetPredecessor",
			Handler:    _Chord_SetPredecessor_Handler,
		},
		{
			MethodName: "SetSuccessor",
			Handler:    _Chord_SetSuccessor_Handler,
		},
		{
			MethodName: "KGet",
			Handler:    _Chord_KGet_Handler,
		},
		{
			MethodName: "KSet",
			Handler:    _Chord_KSet_Handler,
		},
		{
			MethodName: "KDelete",
			Handler:    _Chord_KDelete_Handler,
		},
		{
			MethodName: "KMultiDelete",
			Handler:    _Chord_KMultiDelete_Handler,
		},
		{
			MethodName: "KRequestKeys",
			Handler:    _Chord_KRequestKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chord.proto",
}