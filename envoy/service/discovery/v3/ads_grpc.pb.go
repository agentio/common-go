// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: envoy/service/discovery/v3/ads.proto

package discoveryv3

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
	AggregatedDiscoveryService_StreamAggregatedResources_FullMethodName = "/envoy.service.discovery.v3.AggregatedDiscoveryService/StreamAggregatedResources"
	AggregatedDiscoveryService_DeltaAggregatedResources_FullMethodName  = "/envoy.service.discovery.v3.AggregatedDiscoveryService/DeltaAggregatedResources"
)

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// See https://github.com/envoyproxy/envoy-api#apis for a description of the role of
// ADS and how it is intended to be used by a management server. ADS requests
// have the same structure as their singleton xDS counterparts, but can
// multiplex many resource types on a single stream. The type_url in the
// DiscoveryRequest/DiscoveryResponse provides sufficient information to recover
// the multiplexed singleton APIs at the Envoy instance and management server.
type AggregatedDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DiscoveryRequest, DiscoveryResponse], error)
	DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DeltaDiscoveryRequest, DeltaDiscoveryResponse], error)
}

type aggregatedDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatedDiscoveryServiceClient(cc grpc.ClientConnInterface) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DiscoveryRequest, DiscoveryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AggregatedDiscoveryService_ServiceDesc.Streams[0], AggregatedDiscoveryService_StreamAggregatedResources_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DiscoveryRequest, DiscoveryResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AggregatedDiscoveryService_StreamAggregatedResourcesClient = grpc.BidiStreamingClient[DiscoveryRequest, DiscoveryResponse]

func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[DeltaDiscoveryRequest, DeltaDiscoveryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AggregatedDiscoveryService_ServiceDesc.Streams[1], AggregatedDiscoveryService_DeltaAggregatedResources_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DeltaDiscoveryRequest, DeltaDiscoveryResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AggregatedDiscoveryService_DeltaAggregatedResourcesClient = grpc.BidiStreamingClient[DeltaDiscoveryRequest, DeltaDiscoveryResponse]

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
// All implementations must embed UnimplementedAggregatedDiscoveryServiceServer
// for forward compatibility.
//
// See https://github.com/envoyproxy/envoy-api#apis for a description of the role of
// ADS and how it is intended to be used by a management server. ADS requests
// have the same structure as their singleton xDS counterparts, but can
// multiplex many resource types on a single stream. The type_url in the
// DiscoveryRequest/DiscoveryResponse provides sufficient information to recover
// the multiplexed singleton APIs at the Envoy instance and management server.
type AggregatedDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(grpc.BidiStreamingServer[DiscoveryRequest, DiscoveryResponse]) error
	DeltaAggregatedResources(grpc.BidiStreamingServer[DeltaDiscoveryRequest, DeltaDiscoveryResponse]) error
	mustEmbedUnimplementedAggregatedDiscoveryServiceServer()
}

// UnimplementedAggregatedDiscoveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAggregatedDiscoveryServiceServer struct{}

func (UnimplementedAggregatedDiscoveryServiceServer) StreamAggregatedResources(grpc.BidiStreamingServer[DiscoveryRequest, DiscoveryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamAggregatedResources not implemented")
}
func (UnimplementedAggregatedDiscoveryServiceServer) DeltaAggregatedResources(grpc.BidiStreamingServer[DeltaDiscoveryRequest, DeltaDiscoveryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DeltaAggregatedResources not implemented")
}
func (UnimplementedAggregatedDiscoveryServiceServer) mustEmbedUnimplementedAggregatedDiscoveryServiceServer() {
}
func (UnimplementedAggregatedDiscoveryServiceServer) testEmbeddedByValue() {}

// UnsafeAggregatedDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatedDiscoveryServiceServer will
// result in compilation errors.
type UnsafeAggregatedDiscoveryServiceServer interface {
	mustEmbedUnimplementedAggregatedDiscoveryServiceServer()
}

func RegisterAggregatedDiscoveryServiceServer(s grpc.ServiceRegistrar, srv AggregatedDiscoveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedAggregatedDiscoveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AggregatedDiscoveryService_ServiceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&grpc.GenericServerStream[DiscoveryRequest, DiscoveryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AggregatedDiscoveryService_StreamAggregatedResourcesServer = grpc.BidiStreamingServer[DiscoveryRequest, DiscoveryResponse]

func _AggregatedDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).DeltaAggregatedResources(&grpc.GenericServerStream[DeltaDiscoveryRequest, DeltaDiscoveryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AggregatedDiscoveryService_DeltaAggregatedResourcesServer = grpc.BidiStreamingServer[DeltaDiscoveryRequest, DeltaDiscoveryResponse]

// AggregatedDiscoveryService_ServiceDesc is the grpc.ServiceDesc for AggregatedDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregatedDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v3.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaAggregatedResources",
			Handler:       _AggregatedDiscoveryService_DeltaAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v3/ads.proto",
}
