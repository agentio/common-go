// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/rate_limit_quota/v3/rlqs.proto

package rate_limit_quotav3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v3 "github.com/agentio/common-go/envoy/service/rate_limit_quota/v3"
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
	// RateLimitQuotaServiceName is the fully-qualified name of the RateLimitQuotaService service.
	RateLimitQuotaServiceName = "envoy.service.rate_limit_quota.v3.RateLimitQuotaService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RateLimitQuotaServiceStreamRateLimitQuotasProcedure is the fully-qualified name of the
	// RateLimitQuotaService's StreamRateLimitQuotas RPC.
	RateLimitQuotaServiceStreamRateLimitQuotasProcedure = "/envoy.service.rate_limit_quota.v3.RateLimitQuotaService/StreamRateLimitQuotas"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	rateLimitQuotaServiceServiceDescriptor                     = v3.File_envoy_service_rate_limit_quota_v3_rlqs_proto.Services().ByName("RateLimitQuotaService")
	rateLimitQuotaServiceStreamRateLimitQuotasMethodDescriptor = rateLimitQuotaServiceServiceDescriptor.Methods().ByName("StreamRateLimitQuotas")
)

// RateLimitQuotaServiceClient is a client for the
// envoy.service.rate_limit_quota.v3.RateLimitQuotaService service.
type RateLimitQuotaServiceClient interface {
	// Main communication channel: the data plane sends usage reports to the RLQS server,
	// and the server asynchronously responding with the assignments.
	StreamRateLimitQuotas(context.Context) *connect.BidiStreamForClient[v3.RateLimitQuotaUsageReports, v3.RateLimitQuotaResponse]
}

// NewRateLimitQuotaServiceClient constructs a client for the
// envoy.service.rate_limit_quota.v3.RateLimitQuotaService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRateLimitQuotaServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RateLimitQuotaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rateLimitQuotaServiceClient{
		streamRateLimitQuotas: connect.NewClient[v3.RateLimitQuotaUsageReports, v3.RateLimitQuotaResponse](
			httpClient,
			baseURL+RateLimitQuotaServiceStreamRateLimitQuotasProcedure,
			connect.WithSchema(rateLimitQuotaServiceStreamRateLimitQuotasMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// rateLimitQuotaServiceClient implements RateLimitQuotaServiceClient.
type rateLimitQuotaServiceClient struct {
	streamRateLimitQuotas *connect.Client[v3.RateLimitQuotaUsageReports, v3.RateLimitQuotaResponse]
}

// StreamRateLimitQuotas calls
// envoy.service.rate_limit_quota.v3.RateLimitQuotaService.StreamRateLimitQuotas.
func (c *rateLimitQuotaServiceClient) StreamRateLimitQuotas(ctx context.Context) *connect.BidiStreamForClient[v3.RateLimitQuotaUsageReports, v3.RateLimitQuotaResponse] {
	return c.streamRateLimitQuotas.CallBidiStream(ctx)
}

// RateLimitQuotaServiceHandler is an implementation of the
// envoy.service.rate_limit_quota.v3.RateLimitQuotaService service.
type RateLimitQuotaServiceHandler interface {
	// Main communication channel: the data plane sends usage reports to the RLQS server,
	// and the server asynchronously responding with the assignments.
	StreamRateLimitQuotas(context.Context, *connect.BidiStream[v3.RateLimitQuotaUsageReports, v3.RateLimitQuotaResponse]) error
}

// NewRateLimitQuotaServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRateLimitQuotaServiceHandler(svc RateLimitQuotaServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	rateLimitQuotaServiceStreamRateLimitQuotasHandler := connect.NewBidiStreamHandler(
		RateLimitQuotaServiceStreamRateLimitQuotasProcedure,
		svc.StreamRateLimitQuotas,
		connect.WithSchema(rateLimitQuotaServiceStreamRateLimitQuotasMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.rate_limit_quota.v3.RateLimitQuotaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RateLimitQuotaServiceStreamRateLimitQuotasProcedure:
			rateLimitQuotaServiceStreamRateLimitQuotasHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRateLimitQuotaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRateLimitQuotaServiceHandler struct{}

func (UnimplementedRateLimitQuotaServiceHandler) StreamRateLimitQuotas(context.Context, *connect.BidiStream[v3.RateLimitQuotaUsageReports, v3.RateLimitQuotaResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.rate_limit_quota.v3.RateLimitQuotaService.StreamRateLimitQuotas is not implemented"))
}
