// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: envoy/service/load_stats/v2/lrs.proto

package load_statsv2

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
	LoadReportingService_StreamLoadStats_FullMethodName = "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats"
)

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//
	//	capacity configuration. The capacity configuration definition is
	//	outside of the scope of this document.
	//  2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//     to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	//  1. Once a connection establishes, the management server publishes a
	//     LoadStatsResponse for all clusters it is interested in learning load
	//     stats about.
	//  2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//     based on per-zone weights and/or per-instance weights (if specified)
	//     based on intra-zone LbPolicy. This information comes from the above
	//     {Stream,Fetch}Endpoints.
	//  3. When upstream hosts reply, they optionally add header <define header
	//     name> with ASCII representation of EndpointLoadMetricStats.
	//  4. Envoy aggregates load reports over the period of time given to it in
	//     LoadStatsResponse.load_reporting_interval. This includes aggregation
	//     stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//     well as load metrics from upstream hosts.
	//  5. When the timer of load_reporting_interval expires, Envoy sends new
	//     LoadStatsRequest filled with load reports for each cluster.
	//  6. The management server uses the load reports from all reported Envoys
	//     from around the world, computes global assignment and prepares traffic
	//     assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[LoadStatsRequest, LoadStatsResponse], error)
}

type loadReportingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadReportingServiceClient(cc grpc.ClientConnInterface) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[LoadStatsRequest, LoadStatsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LoadReportingService_ServiceDesc.Streams[0], LoadReportingService_StreamLoadStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LoadStatsRequest, LoadStatsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LoadReportingService_StreamLoadStatsClient = grpc.BidiStreamingClient[LoadStatsRequest, LoadStatsResponse]

// LoadReportingServiceServer is the server API for LoadReportingService service.
// All implementations must embed UnimplementedLoadReportingServiceServer
// for forward compatibility.
type LoadReportingServiceServer interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//
	//	capacity configuration. The capacity configuration definition is
	//	outside of the scope of this document.
	//  2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//     to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	//  1. Once a connection establishes, the management server publishes a
	//     LoadStatsResponse for all clusters it is interested in learning load
	//     stats about.
	//  2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//     based on per-zone weights and/or per-instance weights (if specified)
	//     based on intra-zone LbPolicy. This information comes from the above
	//     {Stream,Fetch}Endpoints.
	//  3. When upstream hosts reply, they optionally add header <define header
	//     name> with ASCII representation of EndpointLoadMetricStats.
	//  4. Envoy aggregates load reports over the period of time given to it in
	//     LoadStatsResponse.load_reporting_interval. This includes aggregation
	//     stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//     well as load metrics from upstream hosts.
	//  5. When the timer of load_reporting_interval expires, Envoy sends new
	//     LoadStatsRequest filled with load reports for each cluster.
	//  6. The management server uses the load reports from all reported Envoys
	//     from around the world, computes global assignment and prepares traffic
	//     assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(grpc.BidiStreamingServer[LoadStatsRequest, LoadStatsResponse]) error
	mustEmbedUnimplementedLoadReportingServiceServer()
}

// UnimplementedLoadReportingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLoadReportingServiceServer struct{}

func (UnimplementedLoadReportingServiceServer) StreamLoadStats(grpc.BidiStreamingServer[LoadStatsRequest, LoadStatsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}
func (UnimplementedLoadReportingServiceServer) mustEmbedUnimplementedLoadReportingServiceServer() {}
func (UnimplementedLoadReportingServiceServer) testEmbeddedByValue()                              {}

// UnsafeLoadReportingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoadReportingServiceServer will
// result in compilation errors.
type UnsafeLoadReportingServiceServer interface {
	mustEmbedUnimplementedLoadReportingServiceServer()
}

func RegisterLoadReportingServiceServer(s grpc.ServiceRegistrar, srv LoadReportingServiceServer) {
	// If the following call pancis, it indicates UnimplementedLoadReportingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LoadReportingService_ServiceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&grpc.GenericServerStream[LoadStatsRequest, LoadStatsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LoadReportingService_StreamLoadStatsServer = grpc.BidiStreamingServer[LoadStatsRequest, LoadStatsResponse]

// LoadReportingService_ServiceDesc is the grpc.ServiceDesc for LoadReportingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoadReportingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}
