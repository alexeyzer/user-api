// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// UserApiServiceClient is the client API for UserApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserApiServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SessionCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SessionCheckResponse, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type userApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserApiServiceClient(cc grpc.ClientConnInterface) UserApiServiceClient {
	return &userApiServiceClient{cc}
}

func (c *userApiServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.api.userApiService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.api.userApiService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiServiceClient) Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.api.userApiService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiServiceClient) SessionCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SessionCheckResponse, error) {
	out := new(SessionCheckResponse)
	err := c.cc.Invoke(ctx, "/user.api.userApiService/SessionCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/user.api.userApiService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServiceServer is the server API for UserApiService service.
// All implementations must embed UnimplementedUserApiServiceServer
// for forward compatibility
type UserApiServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	SessionCheck(context.Context, *emptypb.Empty) (*SessionCheckResponse, error)
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedUserApiServiceServer()
}

// UnimplementedUserApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserApiServiceServer struct {
}

func (UnimplementedUserApiServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserApiServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserApiServiceServer) Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserApiServiceServer) SessionCheck(context.Context, *emptypb.Empty) (*SessionCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionCheck not implemented")
}
func (UnimplementedUserApiServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedUserApiServiceServer) mustEmbedUnimplementedUserApiServiceServer() {}

// UnsafeUserApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserApiServiceServer will
// result in compilation errors.
type UnsafeUserApiServiceServer interface {
	mustEmbedUnimplementedUserApiServiceServer()
}

func RegisterUserApiServiceServer(s grpc.ServiceRegistrar, srv UserApiServiceServer) {
	s.RegisterService(&UserApiService_ServiceDesc, srv)
}

func _UserApiService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.userApiService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApiService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.userApiService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApiService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.userApiService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServiceServer).Logout(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApiService_SessionCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServiceServer).SessionCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.userApiService/SessionCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServiceServer).SessionCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApiService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.userApiService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserApiService_ServiceDesc is the grpc.ServiceDesc for UserApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.userApiService",
	HandlerType: (*UserApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserApiService_CreateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserApiService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserApiService_Logout_Handler,
		},
		{
			MethodName: "SessionCheck",
			Handler:    _UserApiService_SessionCheck_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _UserApiService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user-api.proto",
}
