// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/pb/auth/jobseeker/jobseeker.proto

package jobseeker

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
	Jobseeker_JobSeekerSignup_FullMethodName      = "/jobseekerauth.Jobseeker/JobSeekerSignup"
	Jobseeker_JobSeekerLogin_FullMethodName       = "/jobseekerauth.Jobseeker/JobSeekerLogin"
	Jobseeker_AddProfile_FullMethodName           = "/jobseekerauth.Jobseeker/AddProfile"
	Jobseeker_JobSeekerGetProfile_FullMethodName  = "/jobseekerauth.Jobseeker/JobSeekerGetProfile"
	Jobseeker_JobSeekerEditProfile_FullMethodName = "/jobseekerauth.Jobseeker/JobSeekerEditProfile"
	Jobseeker_JobSeekerOTPLogin_FullMethodName    = "/jobseekerauth.Jobseeker/JobSeekerOTPLogin"
	Jobseeker_OtpVerification_FullMethodName      = "/jobseekerauth.Jobseeker/OtpVerification"
	Jobseeker_ChangePassword_FullMethodName       = "/jobseekerauth.Jobseeker/ChangePassword"
)

// JobseekerClient is the client API for Jobseeker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobseekerClient interface {
	JobSeekerSignup(ctx context.Context, in *JobSeekerSignupRequest, opts ...grpc.CallOption) (*JobSeekerSignupResponse, error)
	JobSeekerLogin(ctx context.Context, in *JobSeekerLoginRequest, opts ...grpc.CallOption) (*JobSeekerLoginResponse, error)
	AddProfile(ctx context.Context, in *AddProfileRequest, opts ...grpc.CallOption) (*AddProfileResponse, error)
	JobSeekerGetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	JobSeekerEditProfile(ctx context.Context, in *JobSeekerEditProfileRequest, opts ...grpc.CallOption) (*JobSeekerEditProfileResponse, error)
	JobSeekerOTPLogin(ctx context.Context, in *JobSeekerOTPLoginRequest, opts ...grpc.CallOption) (*JobSeekerOTPLoginResponse, error)
	OtpVerification(ctx context.Context, in *OtpVerificationRequest, opts ...grpc.CallOption) (*OtpVerificationResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
}

type jobseekerClient struct {
	cc grpc.ClientConnInterface
}

func NewJobseekerClient(cc grpc.ClientConnInterface) JobseekerClient {
	return &jobseekerClient{cc}
}

func (c *jobseekerClient) JobSeekerSignup(ctx context.Context, in *JobSeekerSignupRequest, opts ...grpc.CallOption) (*JobSeekerSignupResponse, error) {
	out := new(JobSeekerSignupResponse)
	err := c.cc.Invoke(ctx, Jobseeker_JobSeekerSignup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) JobSeekerLogin(ctx context.Context, in *JobSeekerLoginRequest, opts ...grpc.CallOption) (*JobSeekerLoginResponse, error) {
	out := new(JobSeekerLoginResponse)
	err := c.cc.Invoke(ctx, Jobseeker_JobSeekerLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) AddProfile(ctx context.Context, in *AddProfileRequest, opts ...grpc.CallOption) (*AddProfileResponse, error) {
	out := new(AddProfileResponse)
	err := c.cc.Invoke(ctx, Jobseeker_AddProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) JobSeekerGetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, Jobseeker_JobSeekerGetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) JobSeekerEditProfile(ctx context.Context, in *JobSeekerEditProfileRequest, opts ...grpc.CallOption) (*JobSeekerEditProfileResponse, error) {
	out := new(JobSeekerEditProfileResponse)
	err := c.cc.Invoke(ctx, Jobseeker_JobSeekerEditProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) JobSeekerOTPLogin(ctx context.Context, in *JobSeekerOTPLoginRequest, opts ...grpc.CallOption) (*JobSeekerOTPLoginResponse, error) {
	out := new(JobSeekerOTPLoginResponse)
	err := c.cc.Invoke(ctx, Jobseeker_JobSeekerOTPLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) OtpVerification(ctx context.Context, in *OtpVerificationRequest, opts ...grpc.CallOption) (*OtpVerificationResponse, error) {
	out := new(OtpVerificationResponse)
	err := c.cc.Invoke(ctx, Jobseeker_OtpVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobseekerClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, Jobseeker_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobseekerServer is the server API for Jobseeker service.
// All implementations must embed UnimplementedJobseekerServer
// for forward compatibility
type JobseekerServer interface {
	JobSeekerSignup(context.Context, *JobSeekerSignupRequest) (*JobSeekerSignupResponse, error)
	JobSeekerLogin(context.Context, *JobSeekerLoginRequest) (*JobSeekerLoginResponse, error)
	AddProfile(context.Context, *AddProfileRequest) (*AddProfileResponse, error)
	JobSeekerGetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	JobSeekerEditProfile(context.Context, *JobSeekerEditProfileRequest) (*JobSeekerEditProfileResponse, error)
	JobSeekerOTPLogin(context.Context, *JobSeekerOTPLoginRequest) (*JobSeekerOTPLoginResponse, error)
	OtpVerification(context.Context, *OtpVerificationRequest) (*OtpVerificationResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	mustEmbedUnimplementedJobseekerServer()
}

// UnimplementedJobseekerServer must be embedded to have forward compatible implementations.
type UnimplementedJobseekerServer struct {
}

func (UnimplementedJobseekerServer) JobSeekerSignup(context.Context, *JobSeekerSignupRequest) (*JobSeekerSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerSignup not implemented")
}
func (UnimplementedJobseekerServer) JobSeekerLogin(context.Context, *JobSeekerLoginRequest) (*JobSeekerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerLogin not implemented")
}
func (UnimplementedJobseekerServer) AddProfile(context.Context, *AddProfileRequest) (*AddProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProfile not implemented")
}
func (UnimplementedJobseekerServer) JobSeekerGetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerGetProfile not implemented")
}
func (UnimplementedJobseekerServer) JobSeekerEditProfile(context.Context, *JobSeekerEditProfileRequest) (*JobSeekerEditProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerEditProfile not implemented")
}
func (UnimplementedJobseekerServer) JobSeekerOTPLogin(context.Context, *JobSeekerOTPLoginRequest) (*JobSeekerOTPLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerOTPLogin not implemented")
}
func (UnimplementedJobseekerServer) OtpVerification(context.Context, *OtpVerificationRequest) (*OtpVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OtpVerification not implemented")
}
func (UnimplementedJobseekerServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedJobseekerServer) mustEmbedUnimplementedJobseekerServer() {}

// UnsafeJobseekerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobseekerServer will
// result in compilation errors.
type UnsafeJobseekerServer interface {
	mustEmbedUnimplementedJobseekerServer()
}

func RegisterJobseekerServer(s grpc.ServiceRegistrar, srv JobseekerServer) {
	s.RegisterService(&Jobseeker_ServiceDesc, srv)
}

func _Jobseeker_JobSeekerSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).JobSeekerSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_JobSeekerSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).JobSeekerSignup(ctx, req.(*JobSeekerSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_JobSeekerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).JobSeekerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_JobSeekerLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).JobSeekerLogin(ctx, req.(*JobSeekerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_AddProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).AddProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_AddProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).AddProfile(ctx, req.(*AddProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_JobSeekerGetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).JobSeekerGetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_JobSeekerGetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).JobSeekerGetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_JobSeekerEditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerEditProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).JobSeekerEditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_JobSeekerEditProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).JobSeekerEditProfile(ctx, req.(*JobSeekerEditProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_JobSeekerOTPLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerOTPLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).JobSeekerOTPLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_JobSeekerOTPLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).JobSeekerOTPLogin(ctx, req.(*JobSeekerOTPLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_OtpVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OtpVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).OtpVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_OtpVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).OtpVerification(ctx, req.(*OtpVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobseeker_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobseeker_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jobseeker_ServiceDesc is the grpc.ServiceDesc for Jobseeker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jobseeker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jobseekerauth.Jobseeker",
	HandlerType: (*JobseekerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JobSeekerSignup",
			Handler:    _Jobseeker_JobSeekerSignup_Handler,
		},
		{
			MethodName: "JobSeekerLogin",
			Handler:    _Jobseeker_JobSeekerLogin_Handler,
		},
		{
			MethodName: "AddProfile",
			Handler:    _Jobseeker_AddProfile_Handler,
		},
		{
			MethodName: "JobSeekerGetProfile",
			Handler:    _Jobseeker_JobSeekerGetProfile_Handler,
		},
		{
			MethodName: "JobSeekerEditProfile",
			Handler:    _Jobseeker_JobSeekerEditProfile_Handler,
		},
		{
			MethodName: "JobSeekerOTPLogin",
			Handler:    _Jobseeker_JobSeekerOTPLogin_Handler,
		},
		{
			MethodName: "OtpVerification",
			Handler:    _Jobseeker_OtpVerification_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Jobseeker_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth/jobseeker/jobseeker.proto",
}
