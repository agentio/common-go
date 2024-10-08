// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/event_reporting/v2alpha/event_reporting_service.proto

package v2alphaconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2alpha "github.com/agentio/common-go/envoy/service/event_reporting/v2alpha"
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
	// EventReportingServiceName is the fully-qualified name of the EventReportingService service.
	EventReportingServiceName = "envoy.service.event_reporting.v2alpha.EventReportingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EventReportingServiceStreamEventsProcedure is the fully-qualified name of the
	// EventReportingService's StreamEvents RPC.
	EventReportingServiceStreamEventsProcedure = "/envoy.service.event_reporting.v2alpha.EventReportingService/StreamEvents"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	eventReportingServiceServiceDescriptor            = v2alpha.File_envoy_service_event_reporting_v2alpha_event_reporting_service_proto.Services().ByName("EventReportingService")
	eventReportingServiceStreamEventsMethodDescriptor = eventReportingServiceServiceDescriptor.Methods().ByName("StreamEvents")
)

// EventReportingServiceClient is a client for the
// envoy.service.event_reporting.v2alpha.EventReportingService service.
type EventReportingServiceClient interface {
	// Envoy will connect and send StreamEventsRequest messages forever.
	// The management server may send StreamEventsResponse to configure event stream. See below.
	// This API is designed for high throughput with the expectation that it might be lossy.
	StreamEvents(context.Context) *connect.BidiStreamForClient[v2alpha.StreamEventsRequest, v2alpha.StreamEventsResponse]
}

// NewEventReportingServiceClient constructs a client for the
// envoy.service.event_reporting.v2alpha.EventReportingService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEventReportingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EventReportingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eventReportingServiceClient{
		streamEvents: connect.NewClient[v2alpha.StreamEventsRequest, v2alpha.StreamEventsResponse](
			httpClient,
			baseURL+EventReportingServiceStreamEventsProcedure,
			connect.WithSchema(eventReportingServiceStreamEventsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// eventReportingServiceClient implements EventReportingServiceClient.
type eventReportingServiceClient struct {
	streamEvents *connect.Client[v2alpha.StreamEventsRequest, v2alpha.StreamEventsResponse]
}

// StreamEvents calls envoy.service.event_reporting.v2alpha.EventReportingService.StreamEvents.
func (c *eventReportingServiceClient) StreamEvents(ctx context.Context) *connect.BidiStreamForClient[v2alpha.StreamEventsRequest, v2alpha.StreamEventsResponse] {
	return c.streamEvents.CallBidiStream(ctx)
}

// EventReportingServiceHandler is an implementation of the
// envoy.service.event_reporting.v2alpha.EventReportingService service.
type EventReportingServiceHandler interface {
	// Envoy will connect and send StreamEventsRequest messages forever.
	// The management server may send StreamEventsResponse to configure event stream. See below.
	// This API is designed for high throughput with the expectation that it might be lossy.
	StreamEvents(context.Context, *connect.BidiStream[v2alpha.StreamEventsRequest, v2alpha.StreamEventsResponse]) error
}

// NewEventReportingServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEventReportingServiceHandler(svc EventReportingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	eventReportingServiceStreamEventsHandler := connect.NewBidiStreamHandler(
		EventReportingServiceStreamEventsProcedure,
		svc.StreamEvents,
		connect.WithSchema(eventReportingServiceStreamEventsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.event_reporting.v2alpha.EventReportingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EventReportingServiceStreamEventsProcedure:
			eventReportingServiceStreamEventsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEventReportingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEventReportingServiceHandler struct{}

func (UnimplementedEventReportingServiceHandler) StreamEvents(context.Context, *connect.BidiStream[v2alpha.StreamEventsRequest, v2alpha.StreamEventsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.event_reporting.v2alpha.EventReportingService.StreamEvents is not implemented"))
}
