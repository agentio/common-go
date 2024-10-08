// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/secret/v3/sds.proto

package secretv3connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v31 "github.com/agentio/common-go/envoy/service/discovery/v3"
	v3 "github.com/agentio/common-go/envoy/service/secret/v3"
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
	// SecretDiscoveryServiceName is the fully-qualified name of the SecretDiscoveryService service.
	SecretDiscoveryServiceName = "envoy.service.secret.v3.SecretDiscoveryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SecretDiscoveryServiceDeltaSecretsProcedure is the fully-qualified name of the
	// SecretDiscoveryService's DeltaSecrets RPC.
	SecretDiscoveryServiceDeltaSecretsProcedure = "/envoy.service.secret.v3.SecretDiscoveryService/DeltaSecrets"
	// SecretDiscoveryServiceStreamSecretsProcedure is the fully-qualified name of the
	// SecretDiscoveryService's StreamSecrets RPC.
	SecretDiscoveryServiceStreamSecretsProcedure = "/envoy.service.secret.v3.SecretDiscoveryService/StreamSecrets"
	// SecretDiscoveryServiceFetchSecretsProcedure is the fully-qualified name of the
	// SecretDiscoveryService's FetchSecrets RPC.
	SecretDiscoveryServiceFetchSecretsProcedure = "/envoy.service.secret.v3.SecretDiscoveryService/FetchSecrets"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	secretDiscoveryServiceServiceDescriptor             = v3.File_envoy_service_secret_v3_sds_proto.Services().ByName("SecretDiscoveryService")
	secretDiscoveryServiceDeltaSecretsMethodDescriptor  = secretDiscoveryServiceServiceDescriptor.Methods().ByName("DeltaSecrets")
	secretDiscoveryServiceStreamSecretsMethodDescriptor = secretDiscoveryServiceServiceDescriptor.Methods().ByName("StreamSecrets")
	secretDiscoveryServiceFetchSecretsMethodDescriptor  = secretDiscoveryServiceServiceDescriptor.Methods().ByName("FetchSecrets")
)

// SecretDiscoveryServiceClient is a client for the envoy.service.secret.v3.SecretDiscoveryService
// service.
type SecretDiscoveryServiceClient interface {
	DeltaSecrets(context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	StreamSecrets(context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse]
	FetchSecrets(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewSecretDiscoveryServiceClient constructs a client for the
// envoy.service.secret.v3.SecretDiscoveryService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSecretDiscoveryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SecretDiscoveryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &secretDiscoveryServiceClient{
		deltaSecrets: connect.NewClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse](
			httpClient,
			baseURL+SecretDiscoveryServiceDeltaSecretsProcedure,
			connect.WithSchema(secretDiscoveryServiceDeltaSecretsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		streamSecrets: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+SecretDiscoveryServiceStreamSecretsProcedure,
			connect.WithSchema(secretDiscoveryServiceStreamSecretsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		fetchSecrets: connect.NewClient[v31.DiscoveryRequest, v31.DiscoveryResponse](
			httpClient,
			baseURL+SecretDiscoveryServiceFetchSecretsProcedure,
			connect.WithSchema(secretDiscoveryServiceFetchSecretsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// secretDiscoveryServiceClient implements SecretDiscoveryServiceClient.
type secretDiscoveryServiceClient struct {
	deltaSecrets  *connect.Client[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]
	streamSecrets *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
	fetchSecrets  *connect.Client[v31.DiscoveryRequest, v31.DiscoveryResponse]
}

// DeltaSecrets calls envoy.service.secret.v3.SecretDiscoveryService.DeltaSecrets.
func (c *secretDiscoveryServiceClient) DeltaSecrets(ctx context.Context) *connect.BidiStreamForClient[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse] {
	return c.deltaSecrets.CallBidiStream(ctx)
}

// StreamSecrets calls envoy.service.secret.v3.SecretDiscoveryService.StreamSecrets.
func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context) *connect.BidiStreamForClient[v31.DiscoveryRequest, v31.DiscoveryResponse] {
	return c.streamSecrets.CallBidiStream(ctx)
}

// FetchSecrets calls envoy.service.secret.v3.SecretDiscoveryService.FetchSecrets.
func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, req *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return c.fetchSecrets.CallUnary(ctx, req)
}

// SecretDiscoveryServiceHandler is an implementation of the
// envoy.service.secret.v3.SecretDiscoveryService service.
type SecretDiscoveryServiceHandler interface {
	DeltaSecrets(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error
	StreamSecrets(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error
	FetchSecrets(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error)
}

// NewSecretDiscoveryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSecretDiscoveryServiceHandler(svc SecretDiscoveryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	secretDiscoveryServiceDeltaSecretsHandler := connect.NewBidiStreamHandler(
		SecretDiscoveryServiceDeltaSecretsProcedure,
		svc.DeltaSecrets,
		connect.WithSchema(secretDiscoveryServiceDeltaSecretsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretDiscoveryServiceStreamSecretsHandler := connect.NewBidiStreamHandler(
		SecretDiscoveryServiceStreamSecretsProcedure,
		svc.StreamSecrets,
		connect.WithSchema(secretDiscoveryServiceStreamSecretsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretDiscoveryServiceFetchSecretsHandler := connect.NewUnaryHandler(
		SecretDiscoveryServiceFetchSecretsProcedure,
		svc.FetchSecrets,
		connect.WithSchema(secretDiscoveryServiceFetchSecretsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/envoy.service.secret.v3.SecretDiscoveryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SecretDiscoveryServiceDeltaSecretsProcedure:
			secretDiscoveryServiceDeltaSecretsHandler.ServeHTTP(w, r)
		case SecretDiscoveryServiceStreamSecretsProcedure:
			secretDiscoveryServiceStreamSecretsHandler.ServeHTTP(w, r)
		case SecretDiscoveryServiceFetchSecretsProcedure:
			secretDiscoveryServiceFetchSecretsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSecretDiscoveryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSecretDiscoveryServiceHandler struct{}

func (UnimplementedSecretDiscoveryServiceHandler) DeltaSecrets(context.Context, *connect.BidiStream[v31.DeltaDiscoveryRequest, v31.DeltaDiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.secret.v3.SecretDiscoveryService.DeltaSecrets is not implemented"))
}

func (UnimplementedSecretDiscoveryServiceHandler) StreamSecrets(context.Context, *connect.BidiStream[v31.DiscoveryRequest, v31.DiscoveryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.secret.v3.SecretDiscoveryService.StreamSecrets is not implemented"))
}

func (UnimplementedSecretDiscoveryServiceHandler) FetchSecrets(context.Context, *connect.Request[v31.DiscoveryRequest]) (*connect.Response[v31.DiscoveryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("envoy.service.secret.v3.SecretDiscoveryService.FetchSecrets is not implemented"))
}
