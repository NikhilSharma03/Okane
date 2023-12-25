// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: protobuf/okane.proto

package okanepb

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
	OkaneUser_CreateUser_FullMethodName     = "/okanepb.OkaneUser/CreateUser"
	OkaneUser_GetUserByID_FullMethodName    = "/okanepb.OkaneUser/GetUserByID"
	OkaneUser_UpdateUserByID_FullMethodName = "/okanepb.OkaneUser/UpdateUserByID"
	OkaneUser_DeleteUserByID_FullMethodName = "/okanepb.OkaneUser/DeleteUserByID"
)

// OkaneUserClient is the client API for OkaneUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OkaneUserClient interface {
	// Create User creates a new User
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Get User By ID returns the user with provided ID
	GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDResponse, error)
	// Update User By ID updates the user with provided ID
	UpdateUserByID(ctx context.Context, in *UpdateUserByIDRequest, opts ...grpc.CallOption) (*UpdateUserByIDResponse, error)
	// Delete User By ID deletes the user with provided ID
	DeleteUserByID(ctx context.Context, in *DeleteUserByIDRequest, opts ...grpc.CallOption) (*DeleteUserByIDResponse, error)
}

type okaneUserClient struct {
	cc grpc.ClientConnInterface
}

func NewOkaneUserClient(cc grpc.ClientConnInterface) OkaneUserClient {
	return &okaneUserClient{cc}
}

func (c *okaneUserClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, OkaneUser_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneUserClient) GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDResponse, error) {
	out := new(GetUserByIDResponse)
	err := c.cc.Invoke(ctx, OkaneUser_GetUserByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneUserClient) UpdateUserByID(ctx context.Context, in *UpdateUserByIDRequest, opts ...grpc.CallOption) (*UpdateUserByIDResponse, error) {
	out := new(UpdateUserByIDResponse)
	err := c.cc.Invoke(ctx, OkaneUser_UpdateUserByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneUserClient) DeleteUserByID(ctx context.Context, in *DeleteUserByIDRequest, opts ...grpc.CallOption) (*DeleteUserByIDResponse, error) {
	out := new(DeleteUserByIDResponse)
	err := c.cc.Invoke(ctx, OkaneUser_DeleteUserByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OkaneUserServer is the server API for OkaneUser service.
// All implementations must embed UnimplementedOkaneUserServer
// for forward compatibility
type OkaneUserServer interface {
	// Create User creates a new User
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Get User By ID returns the user with provided ID
	GetUserByID(context.Context, *GetUserByIDRequest) (*GetUserByIDResponse, error)
	// Update User By ID updates the user with provided ID
	UpdateUserByID(context.Context, *UpdateUserByIDRequest) (*UpdateUserByIDResponse, error)
	// Delete User By ID deletes the user with provided ID
	DeleteUserByID(context.Context, *DeleteUserByIDRequest) (*DeleteUserByIDResponse, error)
	mustEmbedUnimplementedOkaneUserServer()
}

// UnimplementedOkaneUserServer must be embedded to have forward compatible implementations.
type UnimplementedOkaneUserServer struct {
}

func (UnimplementedOkaneUserServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedOkaneUserServer) GetUserByID(context.Context, *GetUserByIDRequest) (*GetUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedOkaneUserServer) UpdateUserByID(context.Context, *UpdateUserByIDRequest) (*UpdateUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserByID not implemented")
}
func (UnimplementedOkaneUserServer) DeleteUserByID(context.Context, *DeleteUserByIDRequest) (*DeleteUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByID not implemented")
}
func (UnimplementedOkaneUserServer) mustEmbedUnimplementedOkaneUserServer() {}

// UnsafeOkaneUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OkaneUserServer will
// result in compilation errors.
type UnsafeOkaneUserServer interface {
	mustEmbedUnimplementedOkaneUserServer()
}

func RegisterOkaneUserServer(s grpc.ServiceRegistrar, srv OkaneUserServer) {
	s.RegisterService(&OkaneUser_ServiceDesc, srv)
}

func _OkaneUser_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneUserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneUser_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneUserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneUser_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneUserServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneUser_GetUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneUserServer).GetUserByID(ctx, req.(*GetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneUser_UpdateUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneUserServer).UpdateUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneUser_UpdateUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneUserServer).UpdateUserByID(ctx, req.(*UpdateUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneUser_DeleteUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneUserServer).DeleteUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneUser_DeleteUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneUserServer).DeleteUserByID(ctx, req.(*DeleteUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OkaneUser_ServiceDesc is the grpc.ServiceDesc for OkaneUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OkaneUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "okanepb.OkaneUser",
	HandlerType: (*OkaneUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _OkaneUser_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _OkaneUser_GetUserByID_Handler,
		},
		{
			MethodName: "UpdateUserByID",
			Handler:    _OkaneUser_UpdateUserByID_Handler,
		},
		{
			MethodName: "DeleteUserByID",
			Handler:    _OkaneUser_DeleteUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/okane.proto",
}

const (
	OkaneExpense_CreateExpense_FullMethodName     = "/okanepb.OkaneExpense/CreateExpense"
	OkaneExpense_GetExpense_FullMethodName        = "/okanepb.OkaneExpense/GetExpense"
	OkaneExpense_GetExpenseByID_FullMethodName    = "/okanepb.OkaneExpense/GetExpenseByID"
	OkaneExpense_UpdateExpenseByID_FullMethodName = "/okanepb.OkaneExpense/UpdateExpenseByID"
	OkaneExpense_DeleteExpenseByID_FullMethodName = "/okanepb.OkaneExpense/DeleteExpenseByID"
)

// OkaneExpenseClient is the client API for OkaneExpense service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OkaneExpenseClient interface {
	// Create Expense creates a new Expense
	CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error)
	// Get Expense returns all the expenses by the provided user id
	GetExpense(ctx context.Context, in *GetExpenseRequest, opts ...grpc.CallOption) (*GetExpenseResponse, error)
	// Get Expense By ID returns the expense by the provided expense id
	GetExpenseByID(ctx context.Context, in *GetExpenseByIDRequest, opts ...grpc.CallOption) (*GetExpenseByIDResponse, error)
	// Update Expense updates the expense by the provided expense id
	UpdateExpenseByID(ctx context.Context, in *UpdateExpenseByIDRequest, opts ...grpc.CallOption) (*UpdateExpenseByIDResponse, error)
	// Delete Expense deletes the expense by the provided expense id
	DeleteExpenseByID(ctx context.Context, in *DeleteExpenseByIDRequest, opts ...grpc.CallOption) (*DeleteExpenseByIDResponse, error)
}

type okaneExpenseClient struct {
	cc grpc.ClientConnInterface
}

func NewOkaneExpenseClient(cc grpc.ClientConnInterface) OkaneExpenseClient {
	return &okaneExpenseClient{cc}
}

func (c *okaneExpenseClient) CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error) {
	out := new(CreateExpenseResponse)
	err := c.cc.Invoke(ctx, OkaneExpense_CreateExpense_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneExpenseClient) GetExpense(ctx context.Context, in *GetExpenseRequest, opts ...grpc.CallOption) (*GetExpenseResponse, error) {
	out := new(GetExpenseResponse)
	err := c.cc.Invoke(ctx, OkaneExpense_GetExpense_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneExpenseClient) GetExpenseByID(ctx context.Context, in *GetExpenseByIDRequest, opts ...grpc.CallOption) (*GetExpenseByIDResponse, error) {
	out := new(GetExpenseByIDResponse)
	err := c.cc.Invoke(ctx, OkaneExpense_GetExpenseByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneExpenseClient) UpdateExpenseByID(ctx context.Context, in *UpdateExpenseByIDRequest, opts ...grpc.CallOption) (*UpdateExpenseByIDResponse, error) {
	out := new(UpdateExpenseByIDResponse)
	err := c.cc.Invoke(ctx, OkaneExpense_UpdateExpenseByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okaneExpenseClient) DeleteExpenseByID(ctx context.Context, in *DeleteExpenseByIDRequest, opts ...grpc.CallOption) (*DeleteExpenseByIDResponse, error) {
	out := new(DeleteExpenseByIDResponse)
	err := c.cc.Invoke(ctx, OkaneExpense_DeleteExpenseByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OkaneExpenseServer is the server API for OkaneExpense service.
// All implementations must embed UnimplementedOkaneExpenseServer
// for forward compatibility
type OkaneExpenseServer interface {
	// Create Expense creates a new Expense
	CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error)
	// Get Expense returns all the expenses by the provided user id
	GetExpense(context.Context, *GetExpenseRequest) (*GetExpenseResponse, error)
	// Get Expense By ID returns the expense by the provided expense id
	GetExpenseByID(context.Context, *GetExpenseByIDRequest) (*GetExpenseByIDResponse, error)
	// Update Expense updates the expense by the provided expense id
	UpdateExpenseByID(context.Context, *UpdateExpenseByIDRequest) (*UpdateExpenseByIDResponse, error)
	// Delete Expense deletes the expense by the provided expense id
	DeleteExpenseByID(context.Context, *DeleteExpenseByIDRequest) (*DeleteExpenseByIDResponse, error)
	mustEmbedUnimplementedOkaneExpenseServer()
}

// UnimplementedOkaneExpenseServer must be embedded to have forward compatible implementations.
type UnimplementedOkaneExpenseServer struct {
}

func (UnimplementedOkaneExpenseServer) CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedOkaneExpenseServer) GetExpense(context.Context, *GetExpenseRequest) (*GetExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpense not implemented")
}
func (UnimplementedOkaneExpenseServer) GetExpenseByID(context.Context, *GetExpenseByIDRequest) (*GetExpenseByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpenseByID not implemented")
}
func (UnimplementedOkaneExpenseServer) UpdateExpenseByID(context.Context, *UpdateExpenseByIDRequest) (*UpdateExpenseByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExpenseByID not implemented")
}
func (UnimplementedOkaneExpenseServer) DeleteExpenseByID(context.Context, *DeleteExpenseByIDRequest) (*DeleteExpenseByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExpenseByID not implemented")
}
func (UnimplementedOkaneExpenseServer) mustEmbedUnimplementedOkaneExpenseServer() {}

// UnsafeOkaneExpenseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OkaneExpenseServer will
// result in compilation errors.
type UnsafeOkaneExpenseServer interface {
	mustEmbedUnimplementedOkaneExpenseServer()
}

func RegisterOkaneExpenseServer(s grpc.ServiceRegistrar, srv OkaneExpenseServer) {
	s.RegisterService(&OkaneExpense_ServiceDesc, srv)
}

func _OkaneExpense_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneExpenseServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneExpense_CreateExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneExpenseServer).CreateExpense(ctx, req.(*CreateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneExpense_GetExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneExpenseServer).GetExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneExpense_GetExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneExpenseServer).GetExpense(ctx, req.(*GetExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneExpense_GetExpenseByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpenseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneExpenseServer).GetExpenseByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneExpense_GetExpenseByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneExpenseServer).GetExpenseByID(ctx, req.(*GetExpenseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneExpense_UpdateExpenseByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExpenseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneExpenseServer).UpdateExpenseByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneExpense_UpdateExpenseByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneExpenseServer).UpdateExpenseByID(ctx, req.(*UpdateExpenseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkaneExpense_DeleteExpenseByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExpenseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkaneExpenseServer).DeleteExpenseByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OkaneExpense_DeleteExpenseByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkaneExpenseServer).DeleteExpenseByID(ctx, req.(*DeleteExpenseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OkaneExpense_ServiceDesc is the grpc.ServiceDesc for OkaneExpense service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OkaneExpense_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "okanepb.OkaneExpense",
	HandlerType: (*OkaneExpenseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpense",
			Handler:    _OkaneExpense_CreateExpense_Handler,
		},
		{
			MethodName: "GetExpense",
			Handler:    _OkaneExpense_GetExpense_Handler,
		},
		{
			MethodName: "GetExpenseByID",
			Handler:    _OkaneExpense_GetExpenseByID_Handler,
		},
		{
			MethodName: "UpdateExpenseByID",
			Handler:    _OkaneExpense_UpdateExpenseByID_Handler,
		},
		{
			MethodName: "DeleteExpenseByID",
			Handler:    _OkaneExpense_DeleteExpenseByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/okane.proto",
}
