// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: event.proto

package proto

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

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventClient interface {
	GetEvent(ctx context.Context, in *Request, opts ...grpc.CallOption) (Event_GetEventClient, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) GetEvent(ctx context.Context, in *Request, opts ...grpc.CallOption) (Event_GetEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &Event_ServiceDesc.Streams[0], "/event.grpc.Event/GetEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventGetEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Event_GetEventClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type eventGetEventClient struct {
	grpc.ClientStream
}

func (x *eventGetEventClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServer is the server API for Event service.
// All implementations must embed UnimplementedEventServer
// for forward compatibility
type EventServer interface {
	GetEvent(*Request, Event_GetEventServer) error
	mustEmbedUnimplementedEventServer()
}

// UnimplementedEventServer must be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) GetEvent(*Request, Event_GetEventServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {}

// UnsafeEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServer will
// result in compilation errors.
type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	s.RegisterService(&Event_ServiceDesc, srv)
}

func _Event_GetEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServer).GetEvent(m, &eventGetEventServer{stream})
}

type Event_GetEventServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type eventGetEventServer struct {
	grpc.ServerStream
}

func (x *eventGetEventServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// Event_ServiceDesc is the grpc.ServiceDesc for Event service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.grpc.Event",
	HandlerType: (*EventServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvent",
			Handler:       _Event_GetEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}
