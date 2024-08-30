// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/ext_proc/v3/external_processor.proto

package ext_procv3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v3 "github.com/agentio/common-go/envoy/service/ext_proc/v3"
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
	// ExternalProcessorName is the fully-qualified name of the ExternalProcessor service.
	ExternalProcessorName = "envoy.service.ext_proc.v3.ExternalProcessor"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExternalProcessorProcessProcedure is the fully-qualified name of the ExternalProcessor's Process
	// RPC.
	ExternalProcessorProcessProcedure = "/envoy.service.ext_proc.v3.ExternalProcessor/Process"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	externalProcessorServiceDescriptor       = v3.File_envoy_service_ext_proc_v3_external_processor_proto.Services().ByName("ExternalProcessor")
	externalProcessorProcessMethodDescriptor = externalProcessorServiceDescriptor.Methods().ByName("Process")
)

// ExternalProcessorClient is a client for the envoy.service.ext_proc.v3.ExternalProcessor service.
type ExternalProcessorClient interface {
	// This begins the bidirectional stream that Envoy will use to
	// give the server control over what the filter does. The actual
	// protocol is described by the ProcessingRequest and ProcessingResponse
	// messages below.
	Process(context.Context) *connect.BidiStreamForClient[v3.ProcessingRequest, v3.ProcessingResponse]
}

// NewExternalProcessorClient constructs a client for the
// envoy.service.ext_proc.v3.ExternalProcessor service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExternalProcessorClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ExternalProcessorClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &externalProcessorClient{
		process: connect.NewClient[v3.ProcessingRequest, v3.ProcessingResponse](
			httpClient,
			baseURL+ExternalProcessorProcessProcedure,
			connect.WithSchema(externalProcessorProcessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// externalProcessorClient implements ExternalProcessorClient.
type externalProcessorClient struct {
	process *connect.Client[v3.ProcessingRequest, v3.ProcessingResponse]
}

// Process calls envoy.service.ext_proc.v3.ExternalProcessor.Process.
func (c *externalProcessorClient) Process(ctx context.Context) *connect.BidiStreamForClient[v3.ProcessingRequest, v3.ProcessingResponse] {
	return c.process.CallBidiStream(ctx)
}

// ExternalProcessorHandler is an implementation of the envoy.service.ext_proc.v3.ExternalProcessor
// service.
type ExternalProcessorHandler interface {
	// This begins the bidirectional stream that Envoy will use to
	// give the server control over what the filter does. The actual
	// protocol is described by the ProcessingRequest and ProcessingResponse
	// messages below.
	Process(context.Context, *connect.BidiStream[v3.ProcessingRequest, v3.ProcessingResponse]) error
}

// NewExternalProcessorHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExternalProcessorHandler(svc ExternalProcessorHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	externalProcessorProcessHandler := connect.NewBidiStreamHandler(
		ExternalProcessorProcessProcedure,
		svc.Process,
		connect.WithSchema(externalProcessorProcessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.ext_proc.v3.ExternalProcessor/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExternalProcessorProcessProcedure:
			externalProcessorProcessHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExternalProcessorHandler returns CodeUnimplemented from all methods.
type UnimplementedExternalProcessorHandler struct{}

func (UnimplementedExternalProcessorHandler) Process(context.Context, *connect.BidiStream[v3.ProcessingRequest, v3.ProcessingResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.ext_proc.v3.ExternalProcessor.Process is not implemented"))
}
