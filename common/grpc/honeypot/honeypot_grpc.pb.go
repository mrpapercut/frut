// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: grpc/honeypot/honeypot.proto

package honeypot

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Honeypot_GetTask_FullMethodName = "/honeypot.Honeypot/GetTask"
)

// HoneypotClient is the client API for Honeypot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HoneypotClient interface {
	GetTask(ctx context.Context, in *HoneypotID, opts ...grpc.CallOption) (*Task, error)
}

type honeypotClient struct {
	cc grpc.ClientConnInterface
}

func NewHoneypotClient(cc grpc.ClientConnInterface) HoneypotClient {
	return &honeypotClient{cc}
}

func (c *honeypotClient) GetTask(ctx context.Context, in *HoneypotID, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, Honeypot_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoneypotServer is the server API for Honeypot service.
// All implementations must embed UnimplementedHoneypotServer
// for forward compatibility.
type HoneypotServer interface {
	GetTask(context.Context, *HoneypotID) (*Task, error)
	mustEmbedUnimplementedHoneypotServer()
}

// UnimplementedHoneypotServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHoneypotServer struct{}

func (UnimplementedHoneypotServer) GetTask(context.Context, *HoneypotID) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedHoneypotServer) mustEmbedUnimplementedHoneypotServer() {}
func (UnimplementedHoneypotServer) testEmbeddedByValue()                  {}

// UnsafeHoneypotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HoneypotServer will
// result in compilation errors.
type UnsafeHoneypotServer interface {
	mustEmbedUnimplementedHoneypotServer()
}

func RegisterHoneypotServer(s grpc.ServiceRegistrar, srv HoneypotServer) {
	// If the following call pancis, it indicates UnimplementedHoneypotServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Honeypot_ServiceDesc, srv)
}

func _Honeypot_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HoneypotID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoneypotServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Honeypot_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoneypotServer).GetTask(ctx, req.(*HoneypotID))
	}
	return interceptor(ctx, in, info, handler)
}

// Honeypot_ServiceDesc is the grpc.ServiceDesc for Honeypot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Honeypot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "honeypot.Honeypot",
	HandlerType: (*HoneypotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _Honeypot_GetTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/honeypot/honeypot.proto",
}
