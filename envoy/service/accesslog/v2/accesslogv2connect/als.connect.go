// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/accesslog/v2/als.proto

package accesslogv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "github.com/agentio/common-go/envoy/service/accesslog/v2"
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
	// AccessLogServiceName is the fully-qualified name of the AccessLogService service.
	AccessLogServiceName = "envoy.service.accesslog.v2.AccessLogService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccessLogServiceStreamAccessLogsProcedure is the fully-qualified name of the AccessLogService's
	// StreamAccessLogs RPC.
	AccessLogServiceStreamAccessLogsProcedure = "/envoy.service.accesslog.v2.AccessLogService/StreamAccessLogs"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accessLogServiceServiceDescriptor                = v2.File_envoy_service_accesslog_v2_als_proto.Services().ByName("AccessLogService")
	accessLogServiceStreamAccessLogsMethodDescriptor = accessLogServiceServiceDescriptor.Methods().ByName("StreamAccessLogs")
)

// AccessLogServiceClient is a client for the envoy.service.accesslog.v2.AccessLogService service.
type AccessLogServiceClient interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(context.Context) *connect.ClientStreamForClient[v2.StreamAccessLogsMessage, v2.StreamAccessLogsResponse]
}

// NewAccessLogServiceClient constructs a client for the envoy.service.accesslog.v2.AccessLogService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccessLogServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccessLogServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accessLogServiceClient{
		streamAccessLogs: connect.NewClient[v2.StreamAccessLogsMessage, v2.StreamAccessLogsResponse](
			httpClient,
			baseURL+AccessLogServiceStreamAccessLogsProcedure,
			connect.WithSchema(accessLogServiceStreamAccessLogsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accessLogServiceClient implements AccessLogServiceClient.
type accessLogServiceClient struct {
	streamAccessLogs *connect.Client[v2.StreamAccessLogsMessage, v2.StreamAccessLogsResponse]
}

// StreamAccessLogs calls envoy.service.accesslog.v2.AccessLogService.StreamAccessLogs.
func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context) *connect.ClientStreamForClient[v2.StreamAccessLogsMessage, v2.StreamAccessLogsResponse] {
	return c.streamAccessLogs.CallClientStream(ctx)
}

// AccessLogServiceHandler is an implementation of the envoy.service.accesslog.v2.AccessLogService
// service.
type AccessLogServiceHandler interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(context.Context, *connect.ClientStream[v2.StreamAccessLogsMessage]) (*connect.Response[v2.StreamAccessLogsResponse], error)
}

// NewAccessLogServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccessLogServiceHandler(svc AccessLogServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accessLogServiceStreamAccessLogsHandler := connect.NewClientStreamHandler(
		AccessLogServiceStreamAccessLogsProcedure,
		svc.StreamAccessLogs,
		connect.WithSchema(accessLogServiceStreamAccessLogsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.accesslog.v2.AccessLogService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccessLogServiceStreamAccessLogsProcedure:
			accessLogServiceStreamAccessLogsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccessLogServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccessLogServiceHandler struct{}

func (UnimplementedAccessLogServiceHandler) StreamAccessLogs(context.Context, *connect.ClientStream[v2.StreamAccessLogsMessage]) (*connect.Response[v2.StreamAccessLogsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.accesslog.v2.AccessLogService.StreamAccessLogs is not implemented"))
}
