// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: detection/detection.proto

package detection

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
	Detection_CheckIP_FullMethodName = "/detection.Detection/CheckIP"
)

// DetectionClient is the client API for Detection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DetectionClient interface {
	CheckIP(ctx context.Context, in *CheckIPRequest, opts ...grpc.CallOption) (*CheckIPResponse, error)
}

type detectionClient struct {
	cc grpc.ClientConnInterface
}

func NewDetectionClient(cc grpc.ClientConnInterface) DetectionClient {
	return &detectionClient{cc}
}

func (c *detectionClient) CheckIP(ctx context.Context, in *CheckIPRequest, opts ...grpc.CallOption) (*CheckIPResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckIPResponse)
	err := c.cc.Invoke(ctx, Detection_CheckIP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetectionServer is the server API for Detection service.
// All implementations must embed UnimplementedDetectionServer
// for forward compatibility.
type DetectionServer interface {
	CheckIP(context.Context, *CheckIPRequest) (*CheckIPResponse, error)
	mustEmbedUnimplementedDetectionServer()
}

// UnimplementedDetectionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDetectionServer struct{}

func (UnimplementedDetectionServer) CheckIP(context.Context, *CheckIPRequest) (*CheckIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIP not implemented")
}
func (UnimplementedDetectionServer) mustEmbedUnimplementedDetectionServer() {}
func (UnimplementedDetectionServer) testEmbeddedByValue()                   {}

// UnsafeDetectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DetectionServer will
// result in compilation errors.
type UnsafeDetectionServer interface {
	mustEmbedUnimplementedDetectionServer()
}

func RegisterDetectionServer(s grpc.ServiceRegistrar, srv DetectionServer) {
	// If the following call pancis, it indicates UnimplementedDetectionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Detection_ServiceDesc, srv)
}

func _Detection_CheckIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectionServer).CheckIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Detection_CheckIP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectionServer).CheckIP(ctx, req.(*CheckIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Detection_ServiceDesc is the grpc.ServiceDesc for Detection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Detection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "detection.Detection",
	HandlerType: (*DetectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckIP",
			Handler:    _Detection_CheckIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "detection/detection.proto",
}