// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: allocator.proto

package allocator

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

const (
	DeliveryAgentAllocatorService_Allocate_FullMethodName = "/deliveryAgentAllocator.DeliveryAgentAllocatorService/allocate"
)

// DeliveryAgentAllocatorServiceClient is the client API for DeliveryAgentAllocatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryAgentAllocatorServiceClient interface {
	Allocate(ctx context.Context, in *AllocatorRequest, opts ...grpc.CallOption) (*AllocatorResponse, error)
}

type deliveryAgentAllocatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryAgentAllocatorServiceClient(cc grpc.ClientConnInterface) DeliveryAgentAllocatorServiceClient {
	return &deliveryAgentAllocatorServiceClient{cc}
}

func (c *deliveryAgentAllocatorServiceClient) Allocate(ctx context.Context, in *AllocatorRequest, opts ...grpc.CallOption) (*AllocatorResponse, error) {
	out := new(AllocatorResponse)
	err := c.cc.Invoke(ctx, DeliveryAgentAllocatorService_Allocate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryAgentAllocatorServiceServer is the server API for DeliveryAgentAllocatorService service.
// All implementations must embed UnimplementedDeliveryAgentAllocatorServiceServer
// for forward compatibility
type DeliveryAgentAllocatorServiceServer interface {
	Allocate(context.Context, *AllocatorRequest) (*AllocatorResponse, error)
	mustEmbedUnimplementedDeliveryAgentAllocatorServiceServer()
}

// UnimplementedDeliveryAgentAllocatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveryAgentAllocatorServiceServer struct {
}

func (UnimplementedDeliveryAgentAllocatorServiceServer) Allocate(context.Context, *AllocatorRequest) (*AllocatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocate not implemented")
}
func (UnimplementedDeliveryAgentAllocatorServiceServer) mustEmbedUnimplementedDeliveryAgentAllocatorServiceServer() {
}

// UnsafeDeliveryAgentAllocatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryAgentAllocatorServiceServer will
// result in compilation errors.
type UnsafeDeliveryAgentAllocatorServiceServer interface {
	mustEmbedUnimplementedDeliveryAgentAllocatorServiceServer()
}

func RegisterDeliveryAgentAllocatorServiceServer(s grpc.ServiceRegistrar, srv DeliveryAgentAllocatorServiceServer) {
	s.RegisterService(&DeliveryAgentAllocatorService_ServiceDesc, srv)
}

func _DeliveryAgentAllocatorService_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryAgentAllocatorServiceServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryAgentAllocatorService_Allocate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryAgentAllocatorServiceServer).Allocate(ctx, req.(*AllocatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryAgentAllocatorService_ServiceDesc is the grpc.ServiceDesc for DeliveryAgentAllocatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryAgentAllocatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deliveryAgentAllocator.DeliveryAgentAllocatorService",
	HandlerType: (*DeliveryAgentAllocatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "allocate",
			Handler:    _DeliveryAgentAllocatorService_Allocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "allocator.proto",
}
