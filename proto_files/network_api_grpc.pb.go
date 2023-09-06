// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: network_api.proto

package base

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkServiceClient interface {
	/// Subscribe to signals.
	SubscribeToSignals(ctx context.Context, in *SubscriberConfig, opts ...grpc.CallOption) (NetworkService_SubscribeToSignalsClient, error)
	/// Subscribe to signals with mapping. Not released yet for general availability. Available in limited environments.
	SubscribeToSignalsWithMapping(ctx context.Context, in *SubscriberWithMappingConfig, opts ...grpc.CallOption) (NetworkService_SubscribeToSignalsWithMappingClient, error)
	/// Publish signals with values.
	PublishSignals(ctx context.Context, in *PublisherConfig, opts ...grpc.CallOption) (*Empty, error)
	/// Read signals from cache.
	ReadSignals(ctx context.Context, in *SignalIds, opts ...grpc.CallOption) (*Signals, error)
}

type networkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServiceClient(cc grpc.ClientConnInterface) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) SubscribeToSignals(ctx context.Context, in *SubscriberConfig, opts ...grpc.CallOption) (NetworkService_SubscribeToSignalsClient, error) {
	stream, err := c.cc.NewStream(ctx, &NetworkService_ServiceDesc.Streams[0], "/base.NetworkService/SubscribeToSignals", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkServiceSubscribeToSignalsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkService_SubscribeToSignalsClient interface {
	Recv() (*Signals, error)
	grpc.ClientStream
}

type networkServiceSubscribeToSignalsClient struct {
	grpc.ClientStream
}

func (x *networkServiceSubscribeToSignalsClient) Recv() (*Signals, error) {
	m := new(Signals)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkServiceClient) SubscribeToSignalsWithMapping(ctx context.Context, in *SubscriberWithMappingConfig, opts ...grpc.CallOption) (NetworkService_SubscribeToSignalsWithMappingClient, error) {
	stream, err := c.cc.NewStream(ctx, &NetworkService_ServiceDesc.Streams[1], "/base.NetworkService/SubscribeToSignalsWithMapping", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkServiceSubscribeToSignalsWithMappingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkService_SubscribeToSignalsWithMappingClient interface {
	Recv() (*Signals, error)
	grpc.ClientStream
}

type networkServiceSubscribeToSignalsWithMappingClient struct {
	grpc.ClientStream
}

func (x *networkServiceSubscribeToSignalsWithMappingClient) Recv() (*Signals, error) {
	m := new(Signals)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkServiceClient) PublishSignals(ctx context.Context, in *PublisherConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/base.NetworkService/PublishSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ReadSignals(ctx context.Context, in *SignalIds, opts ...grpc.CallOption) (*Signals, error) {
	out := new(Signals)
	err := c.cc.Invoke(ctx, "/base.NetworkService/ReadSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
// All implementations must embed UnimplementedNetworkServiceServer
// for forward compatibility
type NetworkServiceServer interface {
	/// Subscribe to signals.
	SubscribeToSignals(*SubscriberConfig, NetworkService_SubscribeToSignalsServer) error
	/// Subscribe to signals with mapping. Not released yet for general availability. Available in limited environments.
	SubscribeToSignalsWithMapping(*SubscriberWithMappingConfig, NetworkService_SubscribeToSignalsWithMappingServer) error
	/// Publish signals with values.
	PublishSignals(context.Context, *PublisherConfig) (*Empty, error)
	/// Read signals from cache.
	ReadSignals(context.Context, *SignalIds) (*Signals, error)
	mustEmbedUnimplementedNetworkServiceServer()
}

// UnimplementedNetworkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (UnimplementedNetworkServiceServer) SubscribeToSignals(*SubscriberConfig, NetworkService_SubscribeToSignalsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToSignals not implemented")
}
func (UnimplementedNetworkServiceServer) SubscribeToSignalsWithMapping(*SubscriberWithMappingConfig, NetworkService_SubscribeToSignalsWithMappingServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToSignalsWithMapping not implemented")
}
func (UnimplementedNetworkServiceServer) PublishSignals(context.Context, *PublisherConfig) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishSignals not implemented")
}
func (UnimplementedNetworkServiceServer) ReadSignals(context.Context, *SignalIds) (*Signals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSignals not implemented")
}
func (UnimplementedNetworkServiceServer) mustEmbedUnimplementedNetworkServiceServer() {}

// UnsafeNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServiceServer will
// result in compilation errors.
type UnsafeNetworkServiceServer interface {
	mustEmbedUnimplementedNetworkServiceServer()
}

func RegisterNetworkServiceServer(s grpc.ServiceRegistrar, srv NetworkServiceServer) {
	s.RegisterService(&NetworkService_ServiceDesc, srv)
}

func _NetworkService_SubscribeToSignals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriberConfig)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkServiceServer).SubscribeToSignals(m, &networkServiceSubscribeToSignalsServer{stream})
}

type NetworkService_SubscribeToSignalsServer interface {
	Send(*Signals) error
	grpc.ServerStream
}

type networkServiceSubscribeToSignalsServer struct {
	grpc.ServerStream
}

func (x *networkServiceSubscribeToSignalsServer) Send(m *Signals) error {
	return x.ServerStream.SendMsg(m)
}

func _NetworkService_SubscribeToSignalsWithMapping_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriberWithMappingConfig)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkServiceServer).SubscribeToSignalsWithMapping(m, &networkServiceSubscribeToSignalsWithMappingServer{stream})
}

type NetworkService_SubscribeToSignalsWithMappingServer interface {
	Send(*Signals) error
	grpc.ServerStream
}

type networkServiceSubscribeToSignalsWithMappingServer struct {
	grpc.ServerStream
}

func (x *networkServiceSubscribeToSignalsWithMappingServer) Send(m *Signals) error {
	return x.ServerStream.SendMsg(m)
}

func _NetworkService_PublishSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublisherConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).PublishSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.NetworkService/PublishSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).PublishSignals(ctx, req.(*PublisherConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ReadSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ReadSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.NetworkService/ReadSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ReadSignals(ctx, req.(*SignalIds))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkService_ServiceDesc is the grpc.ServiceDesc for NetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishSignals",
			Handler:    _NetworkService_PublishSignals_Handler,
		},
		{
			MethodName: "ReadSignals",
			Handler:    _NetworkService_ReadSignals_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToSignals",
			Handler:       _NetworkService_SubscribeToSignals_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToSignalsWithMapping",
			Handler:       _NetworkService_SubscribeToSignalsWithMapping_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "network_api.proto",
}
