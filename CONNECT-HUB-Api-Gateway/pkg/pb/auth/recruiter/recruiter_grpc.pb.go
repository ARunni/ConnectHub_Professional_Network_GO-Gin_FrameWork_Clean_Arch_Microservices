// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/pb/auth/recruiter/recruiter.proto

package recruiter

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
	Recruiter_RecruiterSignup_FullMethodName      = "/recruiterauth.Recruiter/RecruiterSignup"
	Recruiter_RecruiterLogin_FullMethodName       = "/recruiterauth.Recruiter/RecruiterLogin"
	Recruiter_GetUsers_FullMethodName             = "/recruiterauth.Recruiter/GetUsers"
	Recruiter_RecruiterGetProfile_FullMethodName  = "/recruiterauth.Recruiter/RecruiterGetProfile"
	Recruiter_RecruiterEditProfile_FullMethodName = "/recruiterauth.Recruiter/RecruiterEditProfile"
)

// RecruiterClient is the client API for Recruiter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecruiterClient interface {
	RecruiterSignup(ctx context.Context, in *RecruiterSignupRequest, opts ...grpc.CallOption) (*RecruiterSignupResponse, error)
	RecruiterLogin(ctx context.Context, in *RecruiterLoginInRequest, opts ...grpc.CallOption) (*RecruiterLoginResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	RecruiterGetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*RecruiterDetailsResponse, error)
	RecruiterEditProfile(ctx context.Context, in *RecruiterEditProfileRequest, opts ...grpc.CallOption) (*RecruiterEditProfileResponse, error)
}

type recruiterClient struct {
	cc grpc.ClientConnInterface
}

func NewRecruiterClient(cc grpc.ClientConnInterface) RecruiterClient {
	return &recruiterClient{cc}
}

func (c *recruiterClient) RecruiterSignup(ctx context.Context, in *RecruiterSignupRequest, opts ...grpc.CallOption) (*RecruiterSignupResponse, error) {
	out := new(RecruiterSignupResponse)
	err := c.cc.Invoke(ctx, Recruiter_RecruiterSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterClient) RecruiterLogin(ctx context.Context, in *RecruiterLoginInRequest, opts ...grpc.CallOption) (*RecruiterLoginResponse, error) {
	out := new(RecruiterLoginResponse)
	err := c.cc.Invoke(ctx, Recruiter_RecruiterLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, Recruiter_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterClient) RecruiterGetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*RecruiterDetailsResponse, error) {
	out := new(RecruiterDetailsResponse)
	err := c.cc.Invoke(ctx, Recruiter_RecruiterGetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterClient) RecruiterEditProfile(ctx context.Context, in *RecruiterEditProfileRequest, opts ...grpc.CallOption) (*RecruiterEditProfileResponse, error) {
	out := new(RecruiterEditProfileResponse)
	err := c.cc.Invoke(ctx, Recruiter_RecruiterEditProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecruiterServer is the server API for Recruiter service.
// All implementations must embed UnimplementedRecruiterServer
// for forward compatibility
type RecruiterServer interface {
	RecruiterSignup(context.Context, *RecruiterSignupRequest) (*RecruiterSignupResponse, error)
	RecruiterLogin(context.Context, *RecruiterLoginInRequest) (*RecruiterLoginResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	RecruiterGetProfile(context.Context, *GetProfileRequest) (*RecruiterDetailsResponse, error)
	RecruiterEditProfile(context.Context, *RecruiterEditProfileRequest) (*RecruiterEditProfileResponse, error)
	mustEmbedUnimplementedRecruiterServer()
}

// UnimplementedRecruiterServer must be embedded to have forward compatible implementations.
type UnimplementedRecruiterServer struct {
}

func (UnimplementedRecruiterServer) RecruiterSignup(context.Context, *RecruiterSignupRequest) (*RecruiterSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecruiterSignup not implemented")
}
func (UnimplementedRecruiterServer) RecruiterLogin(context.Context, *RecruiterLoginInRequest) (*RecruiterLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecruiterLogin not implemented")
}
func (UnimplementedRecruiterServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedRecruiterServer) RecruiterGetProfile(context.Context, *GetProfileRequest) (*RecruiterDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecruiterGetProfile not implemented")
}
func (UnimplementedRecruiterServer) RecruiterEditProfile(context.Context, *RecruiterEditProfileRequest) (*RecruiterEditProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecruiterEditProfile not implemented")
}
func (UnimplementedRecruiterServer) mustEmbedUnimplementedRecruiterServer() {}

// UnsafeRecruiterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecruiterServer will
// result in compilation errors.
type UnsafeRecruiterServer interface {
	mustEmbedUnimplementedRecruiterServer()
}

func RegisterRecruiterServer(s grpc.ServiceRegistrar, srv RecruiterServer) {
	s.RegisterService(&Recruiter_ServiceDesc, srv)
}

func _Recruiter_RecruiterSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecruiterSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterServer).RecruiterSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recruiter_RecruiterSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterServer).RecruiterSignup(ctx, req.(*RecruiterSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recruiter_RecruiterLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecruiterLoginInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterServer).RecruiterLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recruiter_RecruiterLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterServer).RecruiterLogin(ctx, req.(*RecruiterLoginInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recruiter_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recruiter_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recruiter_RecruiterGetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterServer).RecruiterGetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recruiter_RecruiterGetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterServer).RecruiterGetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recruiter_RecruiterEditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecruiterEditProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterServer).RecruiterEditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recruiter_RecruiterEditProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterServer).RecruiterEditProfile(ctx, req.(*RecruiterEditProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Recruiter_ServiceDesc is the grpc.ServiceDesc for Recruiter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Recruiter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recruiterauth.Recruiter",
	HandlerType: (*RecruiterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecruiterSignup",
			Handler:    _Recruiter_RecruiterSignup_Handler,
		},
		{
			MethodName: "RecruiterLogin",
			Handler:    _Recruiter_RecruiterLogin_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Recruiter_GetUsers_Handler,
		},
		{
			MethodName: "RecruiterGetProfile",
			Handler:    _Recruiter_RecruiterGetProfile_Handler,
		},
		{
			MethodName: "RecruiterEditProfile",
			Handler:    _Recruiter_RecruiterEditProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth/recruiter/recruiter.proto",
}
