// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/endpoint/v3/leds.proto

package endpointv3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v31 "github.com/agentio/common-go/envoy/service/discovery/v3"
	v3 "github.com/agentio/common-go/envoy/service/endpoint/v3"
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
	// LocalityEndpointDiscoveryServiceName is the fully-qualified name of the
	// LocalityEndpointDiscoveryService service.
	LocalityEndpointDiscoveryServiceName = "envoy.service.endpoint.v3.LocalityEndpointDiscoveryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LocalityEndpointDiscoveryServiceDeltaLocalityEndpointsProcedure is the fully-qualified name of
	// the LocalityEndpointDiscoveryService's DeltaLocalityEndpoints RPC.
	LocalityEndpointDiscoveryServiceDeltaLocalityEndpointsProcedure = "/envoy.service.endpoint.v3.LocalityEndpointDiscoveryService/DeltaLocalityEndpoints"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	localityEndpointDiscoveryServiceServiceDescriptor                      = v3.File_envoy_service_endpoint_v3_leds_proto.Services().ByName("LocalityEndpointDiscoveryService")
	localityEndpointDiscoveryServiceDeltaLocalityEndpointsMethodDescriptor = localityEndpointDiscoveryServiceServiceDescriptor.Methods().ByName("DeltaLocalityEndpoints")
)

// LocalityEndpointDiscoveryServiceClient is a client for the
// envoy.service.endpoint.v3.LocalityEndpointDiscoveryService service.
type LocalityEndpointDiscoveryServiceClient interface {
	// The resource_names_subscribe resource_names_unsubscribe fields in DeltaDiscoveryRequest
	// specify a list of glob collections to subscribe to updates for.
	DeltaLocalityEndpoints(context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
}

// NewLocalityEndpointDiscoveryServiceClient constructs a client for the
// envoy.service.endpoint.v3.LocalityEndpointDiscoveryService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLocalityEndpointDiscoveryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LocalityEndpointDiscoveryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &localityEndpointDiscoveryServiceClient{
		deltaLocalityEndpoints: connect.NewClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse](
			httpClient,
			baseURL+LocalityEndpointDiscoveryServiceDeltaLocalityEndpointsProcedure,
			connect.WithSchema(localityEndpointDiscoveryServiceDeltaLocalityEndpointsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// localityEndpointDiscoveryServiceClient implements LocalityEndpointDiscoveryServiceClient.
type localityEndpointDiscoveryServiceClient struct {
	deltaLocalityEndpoints *connect.Client[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
}

// DeltaLocalityEndpoints calls
// envoy.service.endpoint.v3.LocalityEndpointDiscoveryService.DeltaLocalityEndpoints.
func (c *localityEndpointDiscoveryServiceClient) DeltaLocalityEndpoints(ctx context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse] {
	return c.deltaLocalityEndpoints.CallBidiStream(ctx)
}

// LocalityEndpointDiscoveryServiceHandler is an implementation of the
// envoy.service.endpoint.v3.LocalityEndpointDiscoveryService service.
type LocalityEndpointDiscoveryServiceHandler interface {
	// The resource_names_subscribe resource_names_unsubscribe fields in DeltaDiscoveryRequest
	// specify a list of glob collections to subscribe to updates for.
	DeltaLocalityEndpoints(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error
}

// NewLocalityEndpointDiscoveryServiceHandler builds an HTTP handler from the service
// implementation. It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLocalityEndpointDiscoveryServiceHandler(svc LocalityEndpointDiscoveryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	localityEndpointDiscoveryServiceDeltaLocalityEndpointsHandler := connect.NewBidiStreamHandler(
		LocalityEndpointDiscoveryServiceDeltaLocalityEndpointsProcedure,
		svc.DeltaLocalityEndpoints,
		connect.WithSchema(localityEndpointDiscoveryServiceDeltaLocalityEndpointsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.endpoint.v3.LocalityEndpointDiscoveryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LocalityEndpointDiscoveryServiceDeltaLocalityEndpointsProcedure:
			localityEndpointDiscoveryServiceDeltaLocalityEndpointsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLocalityEndpointDiscoveryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLocalityEndpointDiscoveryServiceHandler struct{}

func (UnimplementedLocalityEndpointDiscoveryServiceHandler) DeltaLocalityEndpoints(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.endpoint.v3.LocalityEndpointDiscoveryService.DeltaLocalityEndpoints is not implemented"))
}
