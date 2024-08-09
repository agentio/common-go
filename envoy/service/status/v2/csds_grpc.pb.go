// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: envoy/service/status/v2/csds.proto

package statusv2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ClientStatusDiscoveryService_StreamClientStatus_FullMethodName = "/envoy.service.status.v2.ClientStatusDiscoveryService/StreamClientStatus"
	ClientStatusDiscoveryService_FetchClientStatus_FullMethodName  = "/envoy.service.status.v2.ClientStatusDiscoveryService/FetchClientStatus"
)

// ClientStatusDiscoveryServiceClient is the client API for ClientStatusDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// CSDS is Client Status Discovery Service. It can be used to get the status of
// an xDS-compliant client from the management server's point of view. In the
// future, it can potentially be used as an interface to get the current
// state directly from the client.
type ClientStatusDiscoveryServiceClient interface {
	StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientStatusRequest, ClientStatusResponse], error)
	FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error)
}

type clientStatusDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStatusDiscoveryServiceClient(cc grpc.ClientConnInterface) ClientStatusDiscoveryServiceClient {
	return &clientStatusDiscoveryServiceClient{cc}
}

func (c *clientStatusDiscoveryServiceClient) StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientStatusRequest, ClientStatusResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ClientStatusDiscoveryService_ServiceDesc.Streams[0], ClientStatusDiscoveryService_StreamClientStatus_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ClientStatusRequest, ClientStatusResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ClientStatusDiscoveryService_StreamClientStatusClient = grpc.BidiStreamingClient[ClientStatusRequest, ClientStatusResponse]

func (c *clientStatusDiscoveryServiceClient) FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientStatusResponse)
	err := c.cc.Invoke(ctx, ClientStatusDiscoveryService_FetchClientStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientStatusDiscoveryServiceServer is the server API for ClientStatusDiscoveryService service.
// All implementations must embed UnimplementedClientStatusDiscoveryServiceServer
// for forward compatibility.
//
// CSDS is Client Status Discovery Service. It can be used to get the status of
// an xDS-compliant client from the management server's point of view. In the
// future, it can potentially be used as an interface to get the current
// state directly from the client.
type ClientStatusDiscoveryServiceServer interface {
	StreamClientStatus(grpc.BidiStreamingServer[ClientStatusRequest, ClientStatusResponse]) error
	FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error)
	mustEmbedUnimplementedClientStatusDiscoveryServiceServer()
}

// UnimplementedClientStatusDiscoveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientStatusDiscoveryServiceServer struct{}

func (UnimplementedClientStatusDiscoveryServiceServer) StreamClientStatus(grpc.BidiStreamingServer[ClientStatusRequest, ClientStatusResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientStatus not implemented")
}
func (UnimplementedClientStatusDiscoveryServiceServer) FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClientStatus not implemented")
}
func (UnimplementedClientStatusDiscoveryServiceServer) mustEmbedUnimplementedClientStatusDiscoveryServiceServer() {
}
func (UnimplementedClientStatusDiscoveryServiceServer) testEmbeddedByValue() {}

// UnsafeClientStatusDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStatusDiscoveryServiceServer will
// result in compilation errors.
type UnsafeClientStatusDiscoveryServiceServer interface {
	mustEmbedUnimplementedClientStatusDiscoveryServiceServer()
}

func RegisterClientStatusDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ClientStatusDiscoveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientStatusDiscoveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientStatusDiscoveryService_ServiceDesc, srv)
}

func _ClientStatusDiscoveryService_StreamClientStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStatusDiscoveryServiceServer).StreamClientStatus(&grpc.GenericServerStream[ClientStatusRequest, ClientStatusResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ClientStatusDiscoveryService_StreamClientStatusServer = grpc.BidiStreamingServer[ClientStatusRequest, ClientStatusResponse]

func _ClientStatusDiscoveryService_FetchClientStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientStatusDiscoveryService_FetchClientStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, req.(*ClientStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientStatusDiscoveryService_ServiceDesc is the grpc.ServiceDesc for ClientStatusDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStatusDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.status.v2.ClientStatusDiscoveryService",
	HandlerType: (*ClientStatusDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClientStatus",
			Handler:    _ClientStatusDiscoveryService_FetchClientStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientStatus",
			Handler:       _ClientStatusDiscoveryService_StreamClientStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/status/v2/csds.proto",
}
