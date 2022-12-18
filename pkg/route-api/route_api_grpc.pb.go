// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: route-api/route_api.proto

package route_api

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

// RouteServiceClient is the client API for RouteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteServiceClient interface {
	// Create
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Delete
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type routeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteServiceClient(cc grpc.ClientConnInterface) RouteServiceClient {
	return &routeServiceClient{cc}
}

func (c *routeServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/statusmok.server.RouteService.RouteService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/statusmok.server.RouteService.RouteService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServiceServer is the server API for RouteService service.
// All implementations must embed UnimplementedRouteServiceServer
// for forward compatibility
type RouteServiceServer interface {
	// Create
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Delete
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedRouteServiceServer()
}

// UnimplementedRouteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServiceServer struct {
}

func (UnimplementedRouteServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRouteServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRouteServiceServer) mustEmbedUnimplementedRouteServiceServer() {}

// UnsafeRouteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServiceServer will
// result in compilation errors.
type UnsafeRouteServiceServer interface {
	mustEmbedUnimplementedRouteServiceServer()
}

func RegisterRouteServiceServer(s grpc.ServiceRegistrar, srv RouteServiceServer) {
	s.RegisterService(&RouteService_ServiceDesc, srv)
}

func _RouteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statusmok.server.RouteService.RouteService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statusmok.server.RouteService.RouteService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteService_ServiceDesc is the grpc.ServiceDesc for RouteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statusmok.server.RouteService.RouteService",
	HandlerType: (*RouteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RouteService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RouteService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route-api/route_api.proto",
}