// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// UserApiServiceClient is the client API for UserApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserApiServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type userApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserApiServiceClient(cc grpc.ClientConnInterface) UserApiServiceClient {
	return &userApiServiceClient{cc}
}

func (c *userApiServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
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
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedUserApiServiceServer()
}

// UnimplementedUserApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserApiServiceServer struct {
}

func (UnimplementedUserApiServiceServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
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

func _UserApiService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
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
		return srv.(UserApiServiceServer).Echo(ctx, req.(*StringMessage))
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
			MethodName: "Echo",
			Handler:    _UserApiService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user-api.proto",
}
