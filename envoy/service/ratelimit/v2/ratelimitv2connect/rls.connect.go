// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/ratelimit/v2/rls.proto

package ratelimitv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "github.com/agentio/common-go/envoy/service/ratelimit/v2"
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
	// RateLimitServiceName is the fully-qualified name of the RateLimitService service.
	RateLimitServiceName = "envoy.service.ratelimit.v2.RateLimitService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RateLimitServiceShouldRateLimitProcedure is the fully-qualified name of the RateLimitService's
	// ShouldRateLimit RPC.
	RateLimitServiceShouldRateLimitProcedure = "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	rateLimitServiceServiceDescriptor               = v2.File_envoy_service_ratelimit_v2_rls_proto.Services().ByName("RateLimitService")
	rateLimitServiceShouldRateLimitMethodDescriptor = rateLimitServiceServiceDescriptor.Methods().ByName("ShouldRateLimit")
)

// RateLimitServiceClient is a client for the envoy.service.ratelimit.v2.RateLimitService service.
type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *connect.Request[v2.RateLimitRequest]) (*connect.Response[v2.RateLimitResponse], error)
}

// NewRateLimitServiceClient constructs a client for the envoy.service.ratelimit.v2.RateLimitService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRateLimitServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RateLimitServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rateLimitServiceClient{
		shouldRateLimit: connect.NewClient[v2.RateLimitRequest, v2.RateLimitResponse](
			httpClient,
			baseURL+RateLimitServiceShouldRateLimitProcedure,
			connect.WithSchema(rateLimitServiceShouldRateLimitMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// rateLimitServiceClient implements RateLimitServiceClient.
type rateLimitServiceClient struct {
	shouldRateLimit *connect.Client[v2.RateLimitRequest, v2.RateLimitResponse]
}

// ShouldRateLimit calls envoy.service.ratelimit.v2.RateLimitService.ShouldRateLimit.
func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, req *connect.Request[v2.RateLimitRequest]) (*connect.Response[v2.RateLimitResponse], error) {
	return c.shouldRateLimit.CallUnary(ctx, req)
}

// RateLimitServiceHandler is an implementation of the envoy.service.ratelimit.v2.RateLimitService
// service.
type RateLimitServiceHandler interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *connect.Request[v2.RateLimitRequest]) (*connect.Response[v2.RateLimitResponse], error)
}

// NewRateLimitServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRateLimitServiceHandler(svc RateLimitServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	rateLimitServiceShouldRateLimitHandler := connect.NewUnaryHandler(
		RateLimitServiceShouldRateLimitProcedure,
		svc.ShouldRateLimit,
		connect.WithSchema(rateLimitServiceShouldRateLimitMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.ratelimit.v2.RateLimitService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RateLimitServiceShouldRateLimitProcedure:
			rateLimitServiceShouldRateLimitHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRateLimitServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRateLimitServiceHandler struct{}

func (UnimplementedRateLimitServiceHandler) ShouldRateLimit(context.Context, *connect.Request[v2.RateLimitRequest]) (*connect.Response[v2.RateLimitResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.ratelimit.v2.RateLimitService.ShouldRateLimit is not implemented"))
}
