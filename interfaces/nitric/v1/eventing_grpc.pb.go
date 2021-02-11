// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventingClient is the client API for Eventing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventingClient interface {
	// Publish a message to a given topic
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get a list of available topics
	GetTopics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTopicsReply, error)
}

type eventingClient struct {
	cc grpc.ClientConnInterface
}

func NewEventingClient(cc grpc.ClientConnInterface) EventingClient {
	return &eventingClient{cc}
}

func (c *eventingClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/nitric.v1.eventing.Eventing/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventingClient) GetTopics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTopicsReply, error) {
	out := new(GetTopicsReply)
	err := c.cc.Invoke(ctx, "/nitric.v1.eventing.Eventing/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventingServer is the server API for Eventing service.
// All implementations must embed UnimplementedEventingServer
// for forward compatibility
type EventingServer interface {
	// Publish a message to a given topic
	Publish(context.Context, *PublishRequest) (*emptypb.Empty, error)
	// Get a list of available topics
	GetTopics(context.Context, *emptypb.Empty) (*GetTopicsReply, error)
	mustEmbedUnimplementedEventingServer()
}

// UnimplementedEventingServer must be embedded to have forward compatible implementations.
type UnimplementedEventingServer struct {
}

func (UnimplementedEventingServer) Publish(context.Context, *PublishRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedEventingServer) GetTopics(context.Context, *emptypb.Empty) (*GetTopicsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (UnimplementedEventingServer) mustEmbedUnimplementedEventingServer() {}

// UnsafeEventingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventingServer will
// result in compilation errors.
type UnsafeEventingServer interface {
	mustEmbedUnimplementedEventingServer()
}

func RegisterEventingServer(s grpc.ServiceRegistrar, srv EventingServer) {
	s.RegisterService(&Eventing_ServiceDesc, srv)
}

func _Eventing_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventingServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.v1.eventing.Eventing/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventingServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Eventing_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventingServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.v1.eventing.Eventing/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventingServer).GetTopics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Eventing_ServiceDesc is the grpc.ServiceDesc for Eventing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Eventing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nitric.v1.eventing.Eventing",
	HandlerType: (*EventingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Eventing_Publish_Handler,
		},
		{
			MethodName: "GetTopics",
			Handler:    _Eventing_GetTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/eventing.proto",
}
