// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: pkg/proto/caseItems.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateCaseItem(ctx context.Context, in *CaseItem, opts ...grpc.CallOption) (*Confirm, error)
	DeleteCaseItem(ctx context.Context, in *CaseItem, opts ...grpc.CallOption) (*Confirm, error)
	ShowCaseItem(ctx context.Context, in *CaseItemRequest, opts ...grpc.CallOption) (*CaseItem, error)
	GetAllCaseItems(ctx context.Context, in *Confirm, opts ...grpc.CallOption) (*CaseItems, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateCaseItem(ctx context.Context, in *CaseItem, opts ...grpc.CallOption) (*Confirm, error) {
	out := new(Confirm)
	err := c.cc.Invoke(ctx, "/caseItem.UserService/CreateCaseItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteCaseItem(ctx context.Context, in *CaseItem, opts ...grpc.CallOption) (*Confirm, error) {
	out := new(Confirm)
	err := c.cc.Invoke(ctx, "/caseItem.UserService/DeleteCaseItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ShowCaseItem(ctx context.Context, in *CaseItemRequest, opts ...grpc.CallOption) (*CaseItem, error) {
	out := new(CaseItem)
	err := c.cc.Invoke(ctx, "/caseItem.UserService/ShowCaseItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllCaseItems(ctx context.Context, in *Confirm, opts ...grpc.CallOption) (*CaseItems, error) {
	out := new(CaseItems)
	err := c.cc.Invoke(ctx, "/caseItem.UserService/GetAllCaseItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateCaseItem(context.Context, *CaseItem) (*Confirm, error)
	DeleteCaseItem(context.Context, *CaseItem) (*Confirm, error)
	ShowCaseItem(context.Context, *CaseItemRequest) (*CaseItem, error)
	GetAllCaseItems(context.Context, *Confirm) (*CaseItems, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateCaseItem(context.Context, *CaseItem) (*Confirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCaseItem not implemented")
}
func (UnimplementedUserServiceServer) DeleteCaseItem(context.Context, *CaseItem) (*Confirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCaseItem not implemented")
}
func (UnimplementedUserServiceServer) ShowCaseItem(context.Context, *CaseItemRequest) (*CaseItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCaseItem not implemented")
}
func (UnimplementedUserServiceServer) GetAllCaseItems(context.Context, *Confirm) (*CaseItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCaseItems not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateCaseItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateCaseItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caseItem.UserService/CreateCaseItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateCaseItem(ctx, req.(*CaseItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteCaseItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteCaseItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caseItem.UserService/DeleteCaseItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteCaseItem(ctx, req.(*CaseItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ShowCaseItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ShowCaseItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caseItem.UserService/ShowCaseItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ShowCaseItem(ctx, req.(*CaseItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllCaseItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Confirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllCaseItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caseItem.UserService/GetAllCaseItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllCaseItems(ctx, req.(*Confirm))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "caseItem.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCaseItem",
			Handler:    _UserService_CreateCaseItem_Handler,
		},
		{
			MethodName: "DeleteCaseItem",
			Handler:    _UserService_DeleteCaseItem_Handler,
		},
		{
			MethodName: "ShowCaseItem",
			Handler:    _UserService_ShowCaseItem_Handler,
		},
		{
			MethodName: "GetAllCaseItems",
			Handler:    _UserService_GetAllCaseItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/caseItems.proto",
}
