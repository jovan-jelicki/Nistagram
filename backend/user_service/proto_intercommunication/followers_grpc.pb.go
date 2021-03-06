// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto_intercommunication

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

// FollowersClient is the client API for Followers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowersClient interface {
	CreateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllFollowers(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteBiDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetAllFollowing(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error)
	GetFollowersConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error)
}

type followersClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowersClient(cc grpc.ClientConnInterface) FollowersClient {
	return &followersClient{cc}
}

func (c *followersClient) CreateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/CreateUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetAllFollowers(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/GetAllFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) DeleteDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/DeleteDirectedConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) DeleteBiDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/DeleteBiDirectedConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetAllFollowing(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/GetAllFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) UpdateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error) {
	out := new(CreateFollowerResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/UpdateUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetFollowersConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error) {
	out := new(CreateFollowerResponse)
	err := c.cc.Invoke(ctx, "/proto_intercommunication.Followers/GetFollowersConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowersServer is the server API for Followers service.
// All implementations must embed UnimplementedFollowersServer
// for forward compatibility
type FollowersServer interface {
	CreateUserConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error)
	GetAllFollowers(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*EmptyResponse, error)
	DeleteDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error)
	DeleteBiDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error)
	GetAllFollowing(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUserConnection(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error)
	GetFollowersConnection(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error)
	mustEmbedUnimplementedFollowersServer()
}

// UnimplementedFollowersServer must be embedded to have forward compatible implementations.
type UnimplementedFollowersServer struct {
}

func (UnimplementedFollowersServer) CreateUserConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserConnection not implemented")
}
func (UnimplementedFollowersServer) GetAllFollowers(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFollowers not implemented")
}
func (UnimplementedFollowersServer) CreateUser(context.Context, *CreateUserRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedFollowersServer) DeleteDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDirectedConnection not implemented")
}
func (UnimplementedFollowersServer) DeleteBiDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBiDirectedConnection not implemented")
}
func (UnimplementedFollowersServer) GetAllFollowing(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFollowing not implemented")
}
func (UnimplementedFollowersServer) UpdateUserConnection(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserConnection not implemented")
}
func (UnimplementedFollowersServer) GetFollowersConnection(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowersConnection not implemented")
}
func (UnimplementedFollowersServer) mustEmbedUnimplementedFollowersServer() {}

// UnsafeFollowersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowersServer will
// result in compilation errors.
type UnsafeFollowersServer interface {
	mustEmbedUnimplementedFollowersServer()
}

func RegisterFollowersServer(s grpc.ServiceRegistrar, srv FollowersServer) {
	s.RegisterService(&Followers_ServiceDesc, srv)
}

func _Followers_CreateUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).CreateUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/CreateUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).CreateUserConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetAllFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetAllFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/GetAllFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetAllFollowers(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_DeleteDirectedConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).DeleteDirectedConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/DeleteDirectedConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).DeleteDirectedConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_DeleteBiDirectedConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).DeleteBiDirectedConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/DeleteBiDirectedConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).DeleteBiDirectedConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetAllFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetAllFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/GetAllFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetAllFollowing(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_UpdateUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).UpdateUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/UpdateUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).UpdateUserConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetFollowersConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetFollowersConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_intercommunication.Followers/GetFollowersConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetFollowersConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Followers_ServiceDesc is the grpc.ServiceDesc for Followers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Followers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_intercommunication.Followers",
	HandlerType: (*FollowersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserConnection",
			Handler:    _Followers_CreateUserConnection_Handler,
		},
		{
			MethodName: "GetAllFollowers",
			Handler:    _Followers_GetAllFollowers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Followers_CreateUser_Handler,
		},
		{
			MethodName: "DeleteDirectedConnection",
			Handler:    _Followers_DeleteDirectedConnection_Handler,
		},
		{
			MethodName: "DeleteBiDirectedConnection",
			Handler:    _Followers_DeleteBiDirectedConnection_Handler,
		},
		{
			MethodName: "GetAllFollowing",
			Handler:    _Followers_GetAllFollowing_Handler,
		},
		{
			MethodName: "UpdateUserConnection",
			Handler:    _Followers_UpdateUserConnection_Handler,
		},
		{
			MethodName: "GetFollowersConnection",
			Handler:    _Followers_GetFollowersConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followers.proto",
}
