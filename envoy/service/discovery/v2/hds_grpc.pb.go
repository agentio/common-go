// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: envoy/service/discovery/v2/hds.proto

package discoveryv2

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
	HealthDiscoveryService_StreamHealthCheck_FullMethodName = "/envoy.service.discovery.v2.HealthDiscoveryService/StreamHealthCheck"
	HealthDiscoveryService_FetchHealthCheck_FullMethodName  = "/envoy.service.discovery.v2.HealthDiscoveryService/FetchHealthCheck"
)

// HealthDiscoveryServiceClient is the client API for HealthDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// HDS is Health Discovery Service. It compliments Envoy’s health checking
// service by designating this Envoy to be a healthchecker for a subset of hosts
// in the cluster. The status of these health checks will be reported to the
// management server, where it can be aggregated etc and redistributed back to
// Envoy through EDS.
type HealthDiscoveryServiceClient interface {
	//  1. Envoy starts up and if its can_healthcheck option in the static
	//     bootstrap config is enabled, sends HealthCheckRequest to the management
	//     server. It supplies its capabilities (which protocol it can health check
	//     with, what zone it resides in, etc.).
	//  2. In response to (1), the management server designates this Envoy as a
	//     healthchecker to health check a subset of all upstream hosts for a given
	//     cluster (for example upstream Host 1 and Host 2). It streams
	//     HealthCheckSpecifier messages with cluster related configuration for all
	//     clusters this Envoy is designated to health check. Subsequent
	//     HealthCheckSpecifier message will be sent on changes to:
	//     a. Endpoints to health checks
	//     b. Per cluster configuration change
	//  3. Envoy creates a health probe based on the HealthCheck config and sends
	//     it to endpoint(ip:port) of Host 1 and 2. Based on the HealthCheck
	//     configuration Envoy waits upon the arrival of the probe response and
	//     looks at the content of the response to decide whether the endpoint is
	//     healthy or not. If a response hasn't been received within the timeout
	//     interval, the endpoint health status is considered TIMEOUT.
	//  4. Envoy reports results back in an EndpointHealthResponse message.
	//     Envoy streams responses as often as the interval configured by the
	//     management server in HealthCheckSpecifier.
	//  5. The management Server collects health statuses for all endpoints in the
	//     cluster (for all clusters) and uses this information to construct
	//     EndpointDiscoveryResponse messages.
	//  6. Once Envoy has a list of upstream endpoints to send traffic to, it load
	//     balances traffic to them without additional health checking. It may
	//     use inline healthcheck (i.e. consider endpoint UNHEALTHY if connection
	//     failed to a particular endpoint to account for health status propagation
	//     delay between HDS and EDS).
	//
	// By default, can_healthcheck is true. If can_healthcheck is false, Cluster
	// configuration may not contain HealthCheck message.
	// TODO(htuch): How is can_healthcheck communicated to CDS to ensure the above
	// invariant?
	// TODO(htuch): Add @amb67's diagram.
	StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier], error)
	// TODO(htuch): Unlike the gRPC version, there is no stream-based binding of
	// request/response. Should we add an identifier to the HealthCheckSpecifier
	// to bind with the response?
	FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error)
}

type healthDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthDiscoveryServiceClient(cc grpc.ClientConnInterface) HealthDiscoveryServiceClient {
	return &healthDiscoveryServiceClient{cc}
}

func (c *healthDiscoveryServiceClient) StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HealthDiscoveryService_ServiceDesc.Streams[0], HealthDiscoveryService_StreamHealthCheck_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HealthDiscoveryService_StreamHealthCheckClient = grpc.BidiStreamingClient[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier]

func (c *healthDiscoveryServiceClient) FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckSpecifier)
	err := c.cc.Invoke(ctx, HealthDiscoveryService_FetchHealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthDiscoveryServiceServer is the server API for HealthDiscoveryService service.
// All implementations must embed UnimplementedHealthDiscoveryServiceServer
// for forward compatibility.
//
// HDS is Health Discovery Service. It compliments Envoy’s health checking
// service by designating this Envoy to be a healthchecker for a subset of hosts
// in the cluster. The status of these health checks will be reported to the
// management server, where it can be aggregated etc and redistributed back to
// Envoy through EDS.
type HealthDiscoveryServiceServer interface {
	//  1. Envoy starts up and if its can_healthcheck option in the static
	//     bootstrap config is enabled, sends HealthCheckRequest to the management
	//     server. It supplies its capabilities (which protocol it can health check
	//     with, what zone it resides in, etc.).
	//  2. In response to (1), the management server designates this Envoy as a
	//     healthchecker to health check a subset of all upstream hosts for a given
	//     cluster (for example upstream Host 1 and Host 2). It streams
	//     HealthCheckSpecifier messages with cluster related configuration for all
	//     clusters this Envoy is designated to health check. Subsequent
	//     HealthCheckSpecifier message will be sent on changes to:
	//     a. Endpoints to health checks
	//     b. Per cluster configuration change
	//  3. Envoy creates a health probe based on the HealthCheck config and sends
	//     it to endpoint(ip:port) of Host 1 and 2. Based on the HealthCheck
	//     configuration Envoy waits upon the arrival of the probe response and
	//     looks at the content of the response to decide whether the endpoint is
	//     healthy or not. If a response hasn't been received within the timeout
	//     interval, the endpoint health status is considered TIMEOUT.
	//  4. Envoy reports results back in an EndpointHealthResponse message.
	//     Envoy streams responses as often as the interval configured by the
	//     management server in HealthCheckSpecifier.
	//  5. The management Server collects health statuses for all endpoints in the
	//     cluster (for all clusters) and uses this information to construct
	//     EndpointDiscoveryResponse messages.
	//  6. Once Envoy has a list of upstream endpoints to send traffic to, it load
	//     balances traffic to them without additional health checking. It may
	//     use inline healthcheck (i.e. consider endpoint UNHEALTHY if connection
	//     failed to a particular endpoint to account for health status propagation
	//     delay between HDS and EDS).
	//
	// By default, can_healthcheck is true. If can_healthcheck is false, Cluster
	// configuration may not contain HealthCheck message.
	// TODO(htuch): How is can_healthcheck communicated to CDS to ensure the above
	// invariant?
	// TODO(htuch): Add @amb67's diagram.
	StreamHealthCheck(grpc.BidiStreamingServer[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier]) error
	// TODO(htuch): Unlike the gRPC version, there is no stream-based binding of
	// request/response. Should we add an identifier to the HealthCheckSpecifier
	// to bind with the response?
	FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error)
	mustEmbedUnimplementedHealthDiscoveryServiceServer()
}

// UnimplementedHealthDiscoveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHealthDiscoveryServiceServer struct{}

func (UnimplementedHealthDiscoveryServiceServer) StreamHealthCheck(grpc.BidiStreamingServer[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier]) error {
	return status.Errorf(codes.Unimplemented, "method StreamHealthCheck not implemented")
}
func (UnimplementedHealthDiscoveryServiceServer) FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHealthCheck not implemented")
}
func (UnimplementedHealthDiscoveryServiceServer) mustEmbedUnimplementedHealthDiscoveryServiceServer() {
}
func (UnimplementedHealthDiscoveryServiceServer) testEmbeddedByValue() {}

// UnsafeHealthDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthDiscoveryServiceServer will
// result in compilation errors.
type UnsafeHealthDiscoveryServiceServer interface {
	mustEmbedUnimplementedHealthDiscoveryServiceServer()
}

func RegisterHealthDiscoveryServiceServer(s grpc.ServiceRegistrar, srv HealthDiscoveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedHealthDiscoveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HealthDiscoveryService_ServiceDesc, srv)
}

func _HealthDiscoveryService_StreamHealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthDiscoveryServiceServer).StreamHealthCheck(&grpc.GenericServerStream[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HealthDiscoveryService_StreamHealthCheckServer = grpc.BidiStreamingServer[HealthCheckRequestOrEndpointHealthResponse, HealthCheckSpecifier]

func _HealthDiscoveryService_FetchHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthDiscoveryService_FetchHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, req.(*HealthCheckRequestOrEndpointHealthResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthDiscoveryService_ServiceDesc is the grpc.ServiceDesc for HealthDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.HealthDiscoveryService",
	HandlerType: (*HealthDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHealthCheck",
			Handler:    _HealthDiscoveryService_FetchHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHealthCheck",
			Handler:       _HealthDiscoveryService_StreamHealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/hds.proto",
}
