// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: waf/waf.proto

package waf

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
	WAF_Analyze_FullMethodName = "/waf.WAF/Analyze"
)

// WAFClient is the client API for WAF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WAFClient interface {
	// Analyze analyzes http traffic and gives recommendations to block it
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type wAFClient struct {
	cc grpc.ClientConnInterface
}

func NewWAFClient(cc grpc.ClientConnInterface) WAFClient {
	return &wAFClient{cc}
}

func (c *wAFClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnalyzeResponse)
	err := c.cc.Invoke(ctx, WAF_Analyze_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WAFServer is the server API for WAF service.
// All implementations must embed UnimplementedWAFServer
// for forward compatibility.
type WAFServer interface {
	// Analyze analyzes http traffic and gives recommendations to block it
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
	mustEmbedUnimplementedWAFServer()
}

// UnimplementedWAFServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWAFServer struct{}

func (UnimplementedWAFServer) Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedWAFServer) mustEmbedUnimplementedWAFServer() {}
func (UnimplementedWAFServer) testEmbeddedByValue()             {}

// UnsafeWAFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WAFServer will
// result in compilation errors.
type UnsafeWAFServer interface {
	mustEmbedUnimplementedWAFServer()
}

func RegisterWAFServer(s grpc.ServiceRegistrar, srv WAFServer) {
	// If the following call pancis, it indicates UnimplementedWAFServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WAF_ServiceDesc, srv)
}

func _WAF_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WAFServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WAF_Analyze_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WAFServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WAF_ServiceDesc is the grpc.ServiceDesc for WAF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WAF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "waf.WAF",
	HandlerType: (*WAFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _WAF_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "waf/waf.proto",
}
