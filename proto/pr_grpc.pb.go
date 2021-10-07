// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AirportClient is the client API for Airport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AirportClient interface {
	Status(ctx context.Context, opts ...grpc.CallOption) (Airport_StatusClient, error)
}

type airportClient struct {
	cc grpc.ClientConnInterface
}

func NewAirportClient(cc grpc.ClientConnInterface) AirportClient {
	return &airportClient{cc}
}

func (c *airportClient) Status(ctx context.Context, opts ...grpc.CallOption) (Airport_StatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Airport_ServiceDesc.Streams[0], "/proto.Airport/Status", opts...)
	if err != nil {
		return nil, err
	}
	x := &airportStatusClient{stream}
	return x, nil
}

type Airport_StatusClient interface {
	Send(*Req) error
	Recv() (*Res, error)
	grpc.ClientStream
}

type airportStatusClient struct {
	grpc.ClientStream
}

func (x *airportStatusClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *airportStatusClient) Recv() (*Res, error) {
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AirportServer is the server API for Airport service.
// All implementations must embed UnimplementedAirportServer
// for forward compatibility
type AirportServer interface {
	Status(Airport_StatusServer) error
	mustEmbedUnimplementedAirportServer()
}

// UnimplementedAirportServer must be embedded to have forward compatible implementations.
type UnimplementedAirportServer struct {
}

func (UnimplementedAirportServer) Status(Airport_StatusServer) error {
	return status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedAirportServer) mustEmbedUnimplementedAirportServer() {}

// UnsafeAirportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AirportServer will
// result in compilation errors.
type UnsafeAirportServer interface {
	mustEmbedUnimplementedAirportServer()
}

func RegisterAirportServer(s grpc.ServiceRegistrar, srv AirportServer) {
	s.RegisterService(&Airport_ServiceDesc, srv)
}

func _Airport_Status_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AirportServer).Status(&airportStatusServer{stream})
}

type Airport_StatusServer interface {
	Send(*Res) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type airportStatusServer struct {
	grpc.ServerStream
}

func (x *airportStatusServer) Send(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *airportStatusServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Airport_ServiceDesc is the grpc.ServiceDesc for Airport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Airport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Airport",
	HandlerType: (*AirportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Status",
			Handler:       _Airport_Status_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/pr.proto",
}