// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: pkg/proto/case/case.proto

package _case

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

// CaseServiceClient is the client API for CaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaseServiceClient interface {
	CreateCase(ctx context.Context, in *Case, opts ...grpc.CallOption) (*Confirm, error)
	ViewCase(ctx context.Context, in *CaseRequest, opts ...grpc.CallOption) (*Case, error)
	DeleteCase(ctx context.Context, in *CaseRequest, opts ...grpc.CallOption) (*Confirm, error)
	ShowAllCases(ctx context.Context, in *Confirm, opts ...grpc.CallOption) (*Cases, error)
	GetCaseItem(ctx context.Context, in *CaseItemRequest, opts ...grpc.CallOption) (*Confirm, error)
}

type caseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseServiceClient(cc grpc.ClientConnInterface) CaseServiceClient {
	return &caseServiceClient{cc}
}

func (c *caseServiceClient) CreateCase(ctx context.Context, in *Case, opts ...grpc.CallOption) (*Confirm, error) {
	out := new(Confirm)
	err := c.cc.Invoke(ctx, "/case.CaseService/CreateCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) ViewCase(ctx context.Context, in *CaseRequest, opts ...grpc.CallOption) (*Case, error) {
	out := new(Case)
	err := c.cc.Invoke(ctx, "/case.CaseService/ViewCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) DeleteCase(ctx context.Context, in *CaseRequest, opts ...grpc.CallOption) (*Confirm, error) {
	out := new(Confirm)
	err := c.cc.Invoke(ctx, "/case.CaseService/DeleteCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) ShowAllCases(ctx context.Context, in *Confirm, opts ...grpc.CallOption) (*Cases, error) {
	out := new(Cases)
	err := c.cc.Invoke(ctx, "/case.CaseService/ShowAllCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) GetCaseItem(ctx context.Context, in *CaseItemRequest, opts ...grpc.CallOption) (*Confirm, error) {
	out := new(Confirm)
	err := c.cc.Invoke(ctx, "/case.CaseService/GetCaseItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseServiceServer is the server API for CaseService service.
// All implementations must embed UnimplementedCaseServiceServer
// for forward compatibility
type CaseServiceServer interface {
	CreateCase(context.Context, *Case) (*Confirm, error)
	ViewCase(context.Context, *CaseRequest) (*Case, error)
	DeleteCase(context.Context, *CaseRequest) (*Confirm, error)
	ShowAllCases(context.Context, *Confirm) (*Cases, error)
	GetCaseItem(context.Context, *CaseItemRequest) (*Confirm, error)
	mustEmbedUnimplementedCaseServiceServer()
}

// UnimplementedCaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCaseServiceServer struct {
}

func (UnimplementedCaseServiceServer) CreateCase(context.Context, *Case) (*Confirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCase not implemented")
}
func (UnimplementedCaseServiceServer) ViewCase(context.Context, *CaseRequest) (*Case, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewCase not implemented")
}
func (UnimplementedCaseServiceServer) DeleteCase(context.Context, *CaseRequest) (*Confirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCase not implemented")
}
func (UnimplementedCaseServiceServer) ShowAllCases(context.Context, *Confirm) (*Cases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllCases not implemented")
}
func (UnimplementedCaseServiceServer) GetCaseItem(context.Context, *CaseItemRequest) (*Confirm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaseItem not implemented")
}
func (UnimplementedCaseServiceServer) mustEmbedUnimplementedCaseServiceServer() {}

// UnsafeCaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaseServiceServer will
// result in compilation errors.
type UnsafeCaseServiceServer interface {
	mustEmbedUnimplementedCaseServiceServer()
}

func RegisterCaseServiceServer(s grpc.ServiceRegistrar, srv CaseServiceServer) {
	s.RegisterService(&CaseService_ServiceDesc, srv)
}

func _CaseService_CreateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Case)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).CreateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.CaseService/CreateCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).CreateCase(ctx, req.(*Case))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_ViewCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).ViewCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.CaseService/ViewCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).ViewCase(ctx, req.(*CaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_DeleteCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).DeleteCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.CaseService/DeleteCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).DeleteCase(ctx, req.(*CaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_ShowAllCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Confirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).ShowAllCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.CaseService/ShowAllCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).ShowAllCases(ctx, req.(*Confirm))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_GetCaseItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaseItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).GetCaseItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/case.CaseService/GetCaseItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).GetCaseItem(ctx, req.(*CaseItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaseService_ServiceDesc is the grpc.ServiceDesc for CaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "case.CaseService",
	HandlerType: (*CaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCase",
			Handler:    _CaseService_CreateCase_Handler,
		},
		{
			MethodName: "ViewCase",
			Handler:    _CaseService_ViewCase_Handler,
		},
		{
			MethodName: "DeleteCase",
			Handler:    _CaseService_DeleteCase_Handler,
		},
		{
			MethodName: "ShowAllCases",
			Handler:    _CaseService_ShowAllCases_Handler,
		},
		{
			MethodName: "GetCaseItem",
			Handler:    _CaseService_GetCaseItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/case/case.proto",
}
