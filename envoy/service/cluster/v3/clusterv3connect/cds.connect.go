// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/cluster/v3/cds.proto

package clusterv3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v3 "github.com/agentio/common-go/envoy/service/cluster/v3"
	v31 "github.com/agentio/common-go/envoy/service/discovery/v3"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ClusterDiscoveryServiceName is the fully-qualified name of the ClusterDiscoveryService service.
	ClusterDiscoveryServiceName = "envoy.service.cluster.v3.ClusterDiscoveryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClusterDiscoveryServiceStreamClustersProcedure is the fully-qualified name of the
	// ClusterDiscoveryService's StreamClusters RPC.
	ClusterDiscoveryServiceStreamClustersProcedure = "/envoy.service.cluster.v3.ClusterDiscoveryService/StreamClusters"
	// ClusterDiscoveryServiceDeltaClustersProcedure is the fully-qualified name of the
	// ClusterDiscoveryService's DeltaClusters RPC.
	ClusterDiscoveryServiceDeltaClustersProcedure = "/envoy.service.cluster.v3.ClusterDiscoveryService/DeltaClusters"
	// ClusterDiscoveryServiceFetchClustersProcedure is the fully-qualified name of the
	// ClusterDiscoveryService's FetchClusters RPC.
	ClusterDiscoveryServiceFetchClustersProcedure = "/envoy.service.cluster.v3.ClusterDiscoveryService/FetchClusters"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	clusterDiscoveryServiceServiceDescriptor              = v3.File_envoy_service_cluster_v3_cds_proto.Services().ByName("ClusterDiscoveryService")
	clusterDiscoveryServiceStreamClustersMethodDescriptor = clusterDiscoveryServiceServiceDescriptor.Methods().ByName("StreamClusters")
	clusterDiscoveryServiceDeltaClustersMethodDescriptor  = clusterDiscoveryServiceServiceDescriptor.Methods().ByName("DeltaClusters")
	clusterDiscoveryServiceFetchClustersMethodDescriptor  = clusterDiscoveryServiceServiceDescriptor.Methods().ByName("FetchClusters")
)

// ClusterDiscoveryServiceClient is a client for the
// envoy.service.cluster.v3.ClusterDiscoveryService service.
type ClusterDiscoveryServiceClient interface {
	StreamClusters(context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse]
	DeltaClusters(context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	FetchClusters(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewClusterDiscoveryServiceClient constructs a client for the
// envoy.service.cluster.v3.ClusterDiscoveryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClusterDiscoveryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ClusterDiscoveryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clusterDiscoveryServiceClient{
		streamClusters: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+ClusterDiscoveryServiceStreamClustersProcedure,
			connect.WithSchema(clusterDiscoveryServiceStreamClustersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deltaClusters: connect.NewClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse](
			httpClient,
			baseURL+ClusterDiscoveryServiceDeltaClustersProcedure,
			connect.WithSchema(clusterDiscoveryServiceDeltaClustersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		fetchClusters: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+ClusterDiscoveryServiceFetchClustersProcedure,
			connect.WithSchema(clusterDiscoveryServiceFetchClustersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// clusterDiscoveryServiceClient implements ClusterDiscoveryServiceClient.
type clusterDiscoveryServiceClient struct {
	streamClusters *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
	deltaClusters  *connect.Client[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	fetchClusters  *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
}

// StreamClusters calls envoy.service.cluster.v3.ClusterDiscoveryService.StreamClusters.
func (c *clusterDiscoveryServiceClient) StreamClusters(ctx context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse] {
	return c.streamClusters.CallBidiStream(ctx)
}

// DeltaClusters calls envoy.service.cluster.v3.ClusterDiscoveryService.DeltaClusters.
func (c *clusterDiscoveryServiceClient) DeltaClusters(ctx context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse] {
	return c.deltaClusters.CallBidiStream(ctx)
}

// FetchClusters calls envoy.service.cluster.v3.ClusterDiscoveryService.FetchClusters.
func (c *clusterDiscoveryServiceClient) FetchClusters(ctx context.Context, req *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return c.fetchClusters.CallUnary(ctx, req)
}

// ClusterDiscoveryServiceHandler is an implementation of the
// envoy.service.cluster.v3.ClusterDiscoveryService service.
type ClusterDiscoveryServiceHandler interface {
	StreamClusters(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error
	DeltaClusters(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error
	FetchClusters(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewClusterDiscoveryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClusterDiscoveryServiceHandler(svc ClusterDiscoveryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	clusterDiscoveryServiceStreamClustersHandler := connect.NewBidiStreamHandler(
		ClusterDiscoveryServiceStreamClustersProcedure,
		svc.StreamClusters,
		connect.WithSchema(clusterDiscoveryServiceStreamClustersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	clusterDiscoveryServiceDeltaClustersHandler := connect.NewBidiStreamHandler(
		ClusterDiscoveryServiceDeltaClustersProcedure,
		svc.DeltaClusters,
		connect.WithSchema(clusterDiscoveryServiceDeltaClustersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	clusterDiscoveryServiceFetchClustersHandler := connect.NewUnaryHandler(
		ClusterDiscoveryServiceFetchClustersProcedure,
		svc.FetchClusters,
		connect.WithSchema(clusterDiscoveryServiceFetchClustersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.cluster.v3.ClusterDiscoveryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ClusterDiscoveryServiceStreamClustersProcedure:
			clusterDiscoveryServiceStreamClustersHandler.ServeHTTP(w, r)
		case ClusterDiscoveryServiceDeltaClustersProcedure:
			clusterDiscoveryServiceDeltaClustersHandler.ServeHTTP(w, r)
		case ClusterDiscoveryServiceFetchClustersProcedure:
			clusterDiscoveryServiceFetchClustersHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedClusterDiscoveryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClusterDiscoveryServiceHandler struct{}

func (UnimplementedClusterDiscoveryServiceHandler) StreamClusters(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.cluster.v3.ClusterDiscoveryService.StreamClusters is not implemented"))
}

func (UnimplementedClusterDiscoveryServiceHandler) DeltaClusters(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.cluster.v3.ClusterDiscoveryService.DeltaClusters is not implemented"))
}

func (UnimplementedClusterDiscoveryServiceHandler) FetchClusters(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.cluster.v3.ClusterDiscoveryService.FetchClusters is not implemented"))
}
