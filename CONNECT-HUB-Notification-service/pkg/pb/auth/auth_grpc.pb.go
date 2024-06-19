// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/pb/auth/auth.proto

package auth

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
	NotificationAuthService_UserData_FullMethodName = "/notification_auth.NotificationAuthService/UserData"
)

// NotificationAuthServiceClient is the client API for NotificationAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationAuthServiceClient interface {
	UserData(ctx context.Context, in *UserDataRequest, opts ...grpc.CallOption) (*UserDataResponse, error)
}

type notificationAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationAuthServiceClient(cc grpc.ClientConnInterface) NotificationAuthServiceClient {
	return &notificationAuthServiceClient{cc}
}

func (c *notificationAuthServiceClient) UserData(ctx context.Context, in *UserDataRequest, opts ...grpc.CallOption) (*UserDataResponse, error) {
	out := new(UserDataResponse)
	err := c.cc.Invoke(ctx, NotificationAuthService_UserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationAuthServiceServer is the server API for NotificationAuthService service.
// All implementations must embed UnimplementedNotificationAuthServiceServer
// for forward compatibility
type NotificationAuthServiceServer interface {
	UserData(context.Context, *UserDataRequest) (*UserDataResponse, error)
	mustEmbedUnimplementedNotificationAuthServiceServer()
}

// UnimplementedNotificationAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationAuthServiceServer struct {
}

func (UnimplementedNotificationAuthServiceServer) UserData(context.Context, *UserDataRequest) (*UserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserData not implemented")
}
func (UnimplementedNotificationAuthServiceServer) mustEmbedUnimplementedNotificationAuthServiceServer() {
}

// UnsafeNotificationAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationAuthServiceServer will
// result in compilation errors.
type UnsafeNotificationAuthServiceServer interface {
	mustEmbedUnimplementedNotificationAuthServiceServer()
}

func RegisterNotificationAuthServiceServer(s grpc.ServiceRegistrar, srv NotificationAuthServiceServer) {
	s.RegisterService(&NotificationAuthService_ServiceDesc, srv)
}

func _NotificationAuthService_UserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationAuthServiceServer).UserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationAuthService_UserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationAuthServiceServer).UserData(ctx, req.(*UserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationAuthService_ServiceDesc is the grpc.ServiceDesc for NotificationAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification_auth.NotificationAuthService",
	HandlerType: (*NotificationAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserData",
			Handler:    _NotificationAuthService_UserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth/auth.proto",
}
