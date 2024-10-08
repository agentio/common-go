// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: envoy/service/extension/v3/config_discovery.proto

package extensionv3

import (
	context "context"
	v3 "github.com/agentio/common-go/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExtensionConfigDiscoveryService_StreamExtensionConfigs_FullMethodName = "/envoy.service.extension.v3.ExtensionConfigDiscoveryService/StreamExtensionConfigs"
	ExtensionConfigDiscoveryService_DeltaExtensionConfigs_FullMethodName  = "/envoy.service.extension.v3.ExtensionConfigDiscoveryService/DeltaExtensionConfigs"
	ExtensionConfigDiscoveryService_FetchExtensionConfigs_FullMethodName  = "/envoy.service.extension.v3.ExtensionConfigDiscoveryService/FetchExtensionConfigs"
)

// ExtensionConfigDiscoveryServiceClient is the client API for ExtensionConfigDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Return extension configurations.
type ExtensionConfigDiscoveryServiceClient interface {
	StreamExtensionConfigs(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[v3.DiscoveryRequest, v3.DiscoveryResponse], error)
	DeltaExtensionConfigs(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse], error)
	FetchExtensionConfigs(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type extensionConfigDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionConfigDiscoveryServiceClient(cc grpc.ClientConnInterface) ExtensionConfigDiscoveryServiceClient {
	return &extensionConfigDiscoveryServiceClient{cc}
}

func (c *extensionConfigDiscoveryServiceClient) StreamExtensionConfigs(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[v3.DiscoveryRequest, v3.DiscoveryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExtensionConfigDiscoveryService_ServiceDesc.Streams[0], ExtensionConfigDiscoveryService_StreamExtensionConfigs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[v3.DiscoveryRequest, v3.DiscoveryResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExtensionConfigDiscoveryService_StreamExtensionConfigsClient = grpc.BidiStreamingClient[v3.DiscoveryRequest, v3.DiscoveryResponse]

func (c *extensionConfigDiscoveryServiceClient) DeltaExtensionConfigs(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExtensionConfigDiscoveryService_ServiceDesc.Streams[1], ExtensionConfigDiscoveryService_DeltaExtensionConfigs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExtensionConfigDiscoveryService_DeltaExtensionConfigsClient = grpc.BidiStreamingClient[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse]

func (c *extensionConfigDiscoveryServiceClient) FetchExtensionConfigs(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, ExtensionConfigDiscoveryService_FetchExtensionConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionConfigDiscoveryServiceServer is the server API for ExtensionConfigDiscoveryService service.
// All implementations must embed UnimplementedExtensionConfigDiscoveryServiceServer
// for forward compatibility.
//
// Return extension configurations.
type ExtensionConfigDiscoveryServiceServer interface {
	StreamExtensionConfigs(grpc.BidiStreamingServer[v3.DiscoveryRequest, v3.DiscoveryResponse]) error
	DeltaExtensionConfigs(grpc.BidiStreamingServer[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse]) error
	FetchExtensionConfigs(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
	mustEmbedUnimplementedExtensionConfigDiscoveryServiceServer()
}

// UnimplementedExtensionConfigDiscoveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExtensionConfigDiscoveryServiceServer struct{}

func (UnimplementedExtensionConfigDiscoveryServiceServer) StreamExtensionConfigs(grpc.BidiStreamingServer[v3.DiscoveryRequest, v3.DiscoveryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamExtensionConfigs not implemented")
}
func (UnimplementedExtensionConfigDiscoveryServiceServer) DeltaExtensionConfigs(grpc.BidiStreamingServer[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DeltaExtensionConfigs not implemented")
}
func (UnimplementedExtensionConfigDiscoveryServiceServer) FetchExtensionConfigs(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchExtensionConfigs not implemented")
}
func (UnimplementedExtensionConfigDiscoveryServiceServer) mustEmbedUnimplementedExtensionConfigDiscoveryServiceServer() {
}
func (UnimplementedExtensionConfigDiscoveryServiceServer) testEmbeddedByValue() {}

// UnsafeExtensionConfigDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionConfigDiscoveryServiceServer will
// result in compilation errors.
type UnsafeExtensionConfigDiscoveryServiceServer interface {
	mustEmbedUnimplementedExtensionConfigDiscoveryServiceServer()
}

func RegisterExtensionConfigDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ExtensionConfigDiscoveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedExtensionConfigDiscoveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExtensionConfigDiscoveryService_ServiceDesc, srv)
}

func _ExtensionConfigDiscoveryService_StreamExtensionConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExtensionConfigDiscoveryServiceServer).StreamExtensionConfigs(&grpc.GenericServerStream[v3.DiscoveryRequest, v3.DiscoveryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExtensionConfigDiscoveryService_StreamExtensionConfigsServer = grpc.BidiStreamingServer[v3.DiscoveryRequest, v3.DiscoveryResponse]

func _ExtensionConfigDiscoveryService_DeltaExtensionConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExtensionConfigDiscoveryServiceServer).DeltaExtensionConfigs(&grpc.GenericServerStream[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExtensionConfigDiscoveryService_DeltaExtensionConfigsServer = grpc.BidiStreamingServer[v3.DeltaDiscoveryRequest, v3.DeltaDiscoveryResponse]

func _ExtensionConfigDiscoveryService_FetchExtensionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionConfigDiscoveryServiceServer).FetchExtensionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionConfigDiscoveryService_FetchExtensionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionConfigDiscoveryServiceServer).FetchExtensionConfigs(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtensionConfigDiscoveryService_ServiceDesc is the grpc.ServiceDesc for ExtensionConfigDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionConfigDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.extension.v3.ExtensionConfigDiscoveryService",
	HandlerType: (*ExtensionConfigDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchExtensionConfigs",
			Handler:    _ExtensionConfigDiscoveryService_FetchExtensionConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamExtensionConfigs",
			Handler:       _ExtensionConfigDiscoveryService_StreamExtensionConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaExtensionConfigs",
			Handler:       _ExtensionConfigDiscoveryService_DeltaExtensionConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/extension/v3/config_discovery.proto",
}
