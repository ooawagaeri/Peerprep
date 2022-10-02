// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: collab-service.proto

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

// CollabTunnelServiceClient is the client API for CollabTunnelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollabTunnelServiceClient interface {
	OpenStream(ctx context.Context, opts ...grpc.CallOption) (CollabTunnelService_OpenStreamClient, error)
}

type collabTunnelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollabTunnelServiceClient(cc grpc.ClientConnInterface) CollabTunnelServiceClient {
	return &collabTunnelServiceClient{cc}
}

func (c *collabTunnelServiceClient) OpenStream(ctx context.Context, opts ...grpc.CallOption) (CollabTunnelService_OpenStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CollabTunnelService_ServiceDesc.Streams[0], "/collaboration_service.CollabTunnelService/OpenStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &collabTunnelServiceOpenStreamClient{stream}
	return x, nil
}

type CollabTunnelService_OpenStreamClient interface {
	Send(*CollabTunnelRequest) error
	Recv() (*CollabTunnelResponse, error)
	grpc.ClientStream
}

type collabTunnelServiceOpenStreamClient struct {
	grpc.ClientStream
}

func (x *collabTunnelServiceOpenStreamClient) Send(m *CollabTunnelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *collabTunnelServiceOpenStreamClient) Recv() (*CollabTunnelResponse, error) {
	m := new(CollabTunnelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CollabTunnelServiceServer is the server API for CollabTunnelService service.
// All implementations must embed UnimplementedCollabTunnelServiceServer
// for forward compatibility
type CollabTunnelServiceServer interface {
	OpenStream(CollabTunnelService_OpenStreamServer) error
	mustEmbedUnimplementedCollabTunnelServiceServer()
}

// UnimplementedCollabTunnelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollabTunnelServiceServer struct {
}

func (UnimplementedCollabTunnelServiceServer) OpenStream(CollabTunnelService_OpenStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenStream not implemented")
}
func (UnimplementedCollabTunnelServiceServer) mustEmbedUnimplementedCollabTunnelServiceServer() {}

// UnsafeCollabTunnelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollabTunnelServiceServer will
// result in compilation errors.
type UnsafeCollabTunnelServiceServer interface {
	mustEmbedUnimplementedCollabTunnelServiceServer()
}

func RegisterCollabTunnelServiceServer(s grpc.ServiceRegistrar, srv CollabTunnelServiceServer) {
	s.RegisterService(&CollabTunnelService_ServiceDesc, srv)
}

func _CollabTunnelService_OpenStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CollabTunnelServiceServer).OpenStream(&collabTunnelServiceOpenStreamServer{stream})
}

type CollabTunnelService_OpenStreamServer interface {
	Send(*CollabTunnelResponse) error
	Recv() (*CollabTunnelRequest, error)
	grpc.ServerStream
}

type collabTunnelServiceOpenStreamServer struct {
	grpc.ServerStream
}

func (x *collabTunnelServiceOpenStreamServer) Send(m *CollabTunnelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *collabTunnelServiceOpenStreamServer) Recv() (*CollabTunnelRequest, error) {
	m := new(CollabTunnelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CollabTunnelService_ServiceDesc is the grpc.ServiceDesc for CollabTunnelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollabTunnelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collaboration_service.CollabTunnelService",
	HandlerType: (*CollabTunnelServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenStream",
			Handler:       _CollabTunnelService_OpenStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "collab-service.proto",
}
