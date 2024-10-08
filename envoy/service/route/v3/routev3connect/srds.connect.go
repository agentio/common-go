// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/route/v3/srds.proto

package routev3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v31 "github.com/agentio/common-go/envoy/service/discovery/v3"
	v3 "github.com/agentio/common-go/envoy/service/route/v3"
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
	// ScopedRoutesDiscoveryServiceName is the fully-qualified name of the ScopedRoutesDiscoveryService
	// service.
	ScopedRoutesDiscoveryServiceName = "envoy.service.route.v3.ScopedRoutesDiscoveryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ScopedRoutesDiscoveryServiceStreamScopedRoutesProcedure is the fully-qualified name of the
	// ScopedRoutesDiscoveryService's StreamScopedRoutes RPC.
	ScopedRoutesDiscoveryServiceStreamScopedRoutesProcedure = "/envoy.service.route.v3.ScopedRoutesDiscoveryService/StreamScopedRoutes"
	// ScopedRoutesDiscoveryServiceDeltaScopedRoutesProcedure is the fully-qualified name of the
	// ScopedRoutesDiscoveryService's DeltaScopedRoutes RPC.
	ScopedRoutesDiscoveryServiceDeltaScopedRoutesProcedure = "/envoy.service.route.v3.ScopedRoutesDiscoveryService/DeltaScopedRoutes"
	// ScopedRoutesDiscoveryServiceFetchScopedRoutesProcedure is the fully-qualified name of the
	// ScopedRoutesDiscoveryService's FetchScopedRoutes RPC.
	ScopedRoutesDiscoveryServiceFetchScopedRoutesProcedure = "/envoy.service.route.v3.ScopedRoutesDiscoveryService/FetchScopedRoutes"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	scopedRoutesDiscoveryServiceServiceDescriptor                  = v3.File_envoy_service_route_v3_srds_proto.Services().ByName("ScopedRoutesDiscoveryService")
	scopedRoutesDiscoveryServiceStreamScopedRoutesMethodDescriptor = scopedRoutesDiscoveryServiceServiceDescriptor.Methods().ByName("StreamScopedRoutes")
	scopedRoutesDiscoveryServiceDeltaScopedRoutesMethodDescriptor  = scopedRoutesDiscoveryServiceServiceDescriptor.Methods().ByName("DeltaScopedRoutes")
	scopedRoutesDiscoveryServiceFetchScopedRoutesMethodDescriptor  = scopedRoutesDiscoveryServiceServiceDescriptor.Methods().ByName("FetchScopedRoutes")
)

// ScopedRoutesDiscoveryServiceClient is a client for the
// envoy.service.route.v3.ScopedRoutesDiscoveryService service.
type ScopedRoutesDiscoveryServiceClient interface {
	StreamScopedRoutes(context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse]
	DeltaScopedRoutes(context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	FetchScopedRoutes(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewScopedRoutesDiscoveryServiceClient constructs a client for the
// envoy.service.route.v3.ScopedRoutesDiscoveryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewScopedRoutesDiscoveryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ScopedRoutesDiscoveryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &scopedRoutesDiscoveryServiceClient{
		streamScopedRoutes: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+ScopedRoutesDiscoveryServiceStreamScopedRoutesProcedure,
			connect.WithSchema(scopedRoutesDiscoveryServiceStreamScopedRoutesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deltaScopedRoutes: connect.NewClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse](
			httpClient,
			baseURL+ScopedRoutesDiscoveryServiceDeltaScopedRoutesProcedure,
			connect.WithSchema(scopedRoutesDiscoveryServiceDeltaScopedRoutesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		fetchScopedRoutes: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+ScopedRoutesDiscoveryServiceFetchScopedRoutesProcedure,
			connect.WithSchema(scopedRoutesDiscoveryServiceFetchScopedRoutesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// scopedRoutesDiscoveryServiceClient implements ScopedRoutesDiscoveryServiceClient.
type scopedRoutesDiscoveryServiceClient struct {
	streamScopedRoutes *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
	deltaScopedRoutes  *connect.Client[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	fetchScopedRoutes  *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
}

// StreamScopedRoutes calls envoy.service.route.v3.ScopedRoutesDiscoveryService.StreamScopedRoutes.
func (c *scopedRoutesDiscoveryServiceClient) StreamScopedRoutes(ctx context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse] {
	return c.streamScopedRoutes.CallBidiStream(ctx)
}

// DeltaScopedRoutes calls envoy.service.route.v3.ScopedRoutesDiscoveryService.DeltaScopedRoutes.
func (c *scopedRoutesDiscoveryServiceClient) DeltaScopedRoutes(ctx context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse] {
	return c.deltaScopedRoutes.CallBidiStream(ctx)
}

// FetchScopedRoutes calls envoy.service.route.v3.ScopedRoutesDiscoveryService.FetchScopedRoutes.
func (c *scopedRoutesDiscoveryServiceClient) FetchScopedRoutes(ctx context.Context, req *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return c.fetchScopedRoutes.CallUnary(ctx, req)
}

// ScopedRoutesDiscoveryServiceHandler is an implementation of the
// envoy.service.route.v3.ScopedRoutesDiscoveryService service.
type ScopedRoutesDiscoveryServiceHandler interface {
	StreamScopedRoutes(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error
	DeltaScopedRoutes(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error
	FetchScopedRoutes(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewScopedRoutesDiscoveryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewScopedRoutesDiscoveryServiceHandler(svc ScopedRoutesDiscoveryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	scopedRoutesDiscoveryServiceStreamScopedRoutesHandler := connect.NewBidiStreamHandler(
		ScopedRoutesDiscoveryServiceStreamScopedRoutesProcedure,
		svc.StreamScopedRoutes,
		connect.WithSchema(scopedRoutesDiscoveryServiceStreamScopedRoutesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	scopedRoutesDiscoveryServiceDeltaScopedRoutesHandler := connect.NewBidiStreamHandler(
		ScopedRoutesDiscoveryServiceDeltaScopedRoutesProcedure,
		svc.DeltaScopedRoutes,
		connect.WithSchema(scopedRoutesDiscoveryServiceDeltaScopedRoutesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	scopedRoutesDiscoveryServiceFetchScopedRoutesHandler := connect.NewUnaryHandler(
		ScopedRoutesDiscoveryServiceFetchScopedRoutesProcedure,
		svc.FetchScopedRoutes,
		connect.WithSchema(scopedRoutesDiscoveryServiceFetchScopedRoutesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.route.v3.ScopedRoutesDiscoveryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ScopedRoutesDiscoveryServiceStreamScopedRoutesProcedure:
			scopedRoutesDiscoveryServiceStreamScopedRoutesHandler.ServeHTTP(w, r)
		case ScopedRoutesDiscoveryServiceDeltaScopedRoutesProcedure:
			scopedRoutesDiscoveryServiceDeltaScopedRoutesHandler.ServeHTTP(w, r)
		case ScopedRoutesDiscoveryServiceFetchScopedRoutesProcedure:
			scopedRoutesDiscoveryServiceFetchScopedRoutesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedScopedRoutesDiscoveryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedScopedRoutesDiscoveryServiceHandler struct{}

func (UnimplementedScopedRoutesDiscoveryServiceHandler) StreamScopedRoutes(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.route.v3.ScopedRoutesDiscoveryService.StreamScopedRoutes is not implemented"))
}

func (UnimplementedScopedRoutesDiscoveryServiceHandler) DeltaScopedRoutes(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.route.v3.ScopedRoutesDiscoveryService.DeltaScopedRoutes is not implemented"))
}

func (UnimplementedScopedRoutesDiscoveryServiceHandler) FetchScopedRoutes(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.route.v3.ScopedRoutesDiscoveryService.FetchScopedRoutes is not implemented"))
}
