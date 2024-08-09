// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: envoy/service/tap/v2alpha/tap.proto

package v2alpha

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
	TapSinkService_StreamTaps_FullMethodName = "/envoy.service.tap.v2alpha.TapSinkService/StreamTaps"
)

// TapSinkServiceClient is the client API for TapSinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [#not-implemented-hide:] A tap service to receive incoming taps. Envoy will call
// StreamTaps to deliver captured taps to the server
type TapSinkServiceClient interface {
	// Envoy will connect and send StreamTapsRequest messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect.
	StreamTaps(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StreamTapsRequest, StreamTapsResponse], error)
}

type tapSinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTapSinkServiceClient(cc grpc.ClientConnInterface) TapSinkServiceClient {
	return &tapSinkServiceClient{cc}
}

func (c *tapSinkServiceClient) StreamTaps(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StreamTapsRequest, StreamTapsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TapSinkService_ServiceDesc.Streams[0], TapSinkService_StreamTaps_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamTapsRequest, StreamTapsResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TapSinkService_StreamTapsClient = grpc.ClientStreamingClient[StreamTapsRequest, StreamTapsResponse]

// TapSinkServiceServer is the server API for TapSinkService service.
// All implementations must embed UnimplementedTapSinkServiceServer
// for forward compatibility.
//
// [#not-implemented-hide:] A tap service to receive incoming taps. Envoy will call
// StreamTaps to deliver captured taps to the server
type TapSinkServiceServer interface {
	// Envoy will connect and send StreamTapsRequest messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect.
	StreamTaps(grpc.ClientStreamingServer[StreamTapsRequest, StreamTapsResponse]) error
	mustEmbedUnimplementedTapSinkServiceServer()
}

// UnimplementedTapSinkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTapSinkServiceServer struct{}

func (UnimplementedTapSinkServiceServer) StreamTaps(grpc.ClientStreamingServer[StreamTapsRequest, StreamTapsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamTaps not implemented")
}
func (UnimplementedTapSinkServiceServer) mustEmbedUnimplementedTapSinkServiceServer() {}
func (UnimplementedTapSinkServiceServer) testEmbeddedByValue()                        {}

// UnsafeTapSinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TapSinkServiceServer will
// result in compilation errors.
type UnsafeTapSinkServiceServer interface {
	mustEmbedUnimplementedTapSinkServiceServer()
}

func RegisterTapSinkServiceServer(s grpc.ServiceRegistrar, srv TapSinkServiceServer) {
	// If the following call pancis, it indicates UnimplementedTapSinkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TapSinkService_ServiceDesc, srv)
}

func _TapSinkService_StreamTaps_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapSinkServiceServer).StreamTaps(&grpc.GenericServerStream[StreamTapsRequest, StreamTapsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TapSinkService_StreamTapsServer = grpc.ClientStreamingServer[StreamTapsRequest, StreamTapsResponse]

// TapSinkService_ServiceDesc is the grpc.ServiceDesc for TapSinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TapSinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapSinkService",
	HandlerType: (*TapSinkServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTaps",
			Handler:       _TapSinkService_StreamTaps_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tap.proto",
}
