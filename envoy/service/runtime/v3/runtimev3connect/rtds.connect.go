// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/runtime/v3/rtds.proto

package runtimev3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v31 "github.com/agentio/common-go/envoy/service/discovery/v3"
	v3 "github.com/agentio/common-go/envoy/service/runtime/v3"
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
	// RuntimeDiscoveryServiceName is the fully-qualified name of the RuntimeDiscoveryService service.
	RuntimeDiscoveryServiceName = "envoy.service.runtime.v3.RuntimeDiscoveryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RuntimeDiscoveryServiceStreamRuntimeProcedure is the fully-qualified name of the
	// RuntimeDiscoveryService's StreamRuntime RPC.
	RuntimeDiscoveryServiceStreamRuntimeProcedure = "/envoy.service.runtime.v3.RuntimeDiscoveryService/StreamRuntime"
	// RuntimeDiscoveryServiceDeltaRuntimeProcedure is the fully-qualified name of the
	// RuntimeDiscoveryService's DeltaRuntime RPC.
	RuntimeDiscoveryServiceDeltaRuntimeProcedure = "/envoy.service.runtime.v3.RuntimeDiscoveryService/DeltaRuntime"
	// RuntimeDiscoveryServiceFetchRuntimeProcedure is the fully-qualified name of the
	// RuntimeDiscoveryService's FetchRuntime RPC.
	RuntimeDiscoveryServiceFetchRuntimeProcedure = "/envoy.service.runtime.v3.RuntimeDiscoveryService/FetchRuntime"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	runtimeDiscoveryServiceServiceDescriptor             = v3.File_envoy_service_runtime_v3_rtds_proto.Services().ByName("RuntimeDiscoveryService")
	runtimeDiscoveryServiceStreamRuntimeMethodDescriptor = runtimeDiscoveryServiceServiceDescriptor.Methods().ByName("StreamRuntime")
	runtimeDiscoveryServiceDeltaRuntimeMethodDescriptor  = runtimeDiscoveryServiceServiceDescriptor.Methods().ByName("DeltaRuntime")
	runtimeDiscoveryServiceFetchRuntimeMethodDescriptor  = runtimeDiscoveryServiceServiceDescriptor.Methods().ByName("FetchRuntime")
)

// RuntimeDiscoveryServiceClient is a client for the
// envoy.service.runtime.v3.RuntimeDiscoveryService service.
type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse]
	DeltaRuntime(context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	FetchRuntime(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewRuntimeDiscoveryServiceClient constructs a client for the
// envoy.service.runtime.v3.RuntimeDiscoveryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRuntimeDiscoveryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RuntimeDiscoveryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &runtimeDiscoveryServiceClient{
		streamRuntime: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+RuntimeDiscoveryServiceStreamRuntimeProcedure,
			connect.WithSchema(runtimeDiscoveryServiceStreamRuntimeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deltaRuntime: connect.NewClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse](
			httpClient,
			baseURL+RuntimeDiscoveryServiceDeltaRuntimeProcedure,
			connect.WithSchema(runtimeDiscoveryServiceDeltaRuntimeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		fetchRuntime: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+RuntimeDiscoveryServiceFetchRuntimeProcedure,
			connect.WithSchema(runtimeDiscoveryServiceFetchRuntimeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// runtimeDiscoveryServiceClient implements RuntimeDiscoveryServiceClient.
type runtimeDiscoveryServiceClient struct {
	streamRuntime *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
	deltaRuntime  *connect.Client[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	fetchRuntime  *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
}

// StreamRuntime calls envoy.service.runtime.v3.RuntimeDiscoveryService.StreamRuntime.
func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse] {
	return c.streamRuntime.CallBidiStream(ctx)
}

// DeltaRuntime calls envoy.service.runtime.v3.RuntimeDiscoveryService.DeltaRuntime.
func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse] {
	return c.deltaRuntime.CallBidiStream(ctx)
}

// FetchRuntime calls envoy.service.runtime.v3.RuntimeDiscoveryService.FetchRuntime.
func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, req *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return c.fetchRuntime.CallUnary(ctx, req)
}

// RuntimeDiscoveryServiceHandler is an implementation of the
// envoy.service.runtime.v3.RuntimeDiscoveryService service.
type RuntimeDiscoveryServiceHandler interface {
	StreamRuntime(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error
	DeltaRuntime(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error
	FetchRuntime(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewRuntimeDiscoveryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRuntimeDiscoveryServiceHandler(svc RuntimeDiscoveryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	runtimeDiscoveryServiceStreamRuntimeHandler := connect.NewBidiStreamHandler(
		RuntimeDiscoveryServiceStreamRuntimeProcedure,
		svc.StreamRuntime,
		connect.WithSchema(runtimeDiscoveryServiceStreamRuntimeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	runtimeDiscoveryServiceDeltaRuntimeHandler := connect.NewBidiStreamHandler(
		RuntimeDiscoveryServiceDeltaRuntimeProcedure,
		svc.DeltaRuntime,
		connect.WithSchema(runtimeDiscoveryServiceDeltaRuntimeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	runtimeDiscoveryServiceFetchRuntimeHandler := connect.NewUnaryHandler(
		RuntimeDiscoveryServiceFetchRuntimeProcedure,
		svc.FetchRuntime,
		connect.WithSchema(runtimeDiscoveryServiceFetchRuntimeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.runtime.v3.RuntimeDiscoveryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RuntimeDiscoveryServiceStreamRuntimeProcedure:
			runtimeDiscoveryServiceStreamRuntimeHandler.ServeHTTP(w, r)
		case RuntimeDiscoveryServiceDeltaRuntimeProcedure:
			runtimeDiscoveryServiceDeltaRuntimeHandler.ServeHTTP(w, r)
		case RuntimeDiscoveryServiceFetchRuntimeProcedure:
			runtimeDiscoveryServiceFetchRuntimeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRuntimeDiscoveryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRuntimeDiscoveryServiceHandler struct{}

func (UnimplementedRuntimeDiscoveryServiceHandler) StreamRuntime(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.runtime.v3.RuntimeDiscoveryService.StreamRuntime is not implemented"))
}

func (UnimplementedRuntimeDiscoveryServiceHandler) DeltaRuntime(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.runtime.v3.RuntimeDiscoveryService.DeltaRuntime is not implemented"))
}

func (UnimplementedRuntimeDiscoveryServiceHandler) FetchRuntime(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.runtime.v3.RuntimeDiscoveryService.FetchRuntime is not implemented"))
}
