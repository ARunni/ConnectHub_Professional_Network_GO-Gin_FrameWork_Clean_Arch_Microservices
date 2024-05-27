// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/pb/job/recruiter/job.proto

package recruiter

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

const (
	RecruiterJob_PostJob_FullMethodName    = "/job.RecruiterJob/PostJob"
	RecruiterJob_GetAllJobs_FullMethodName = "/job.RecruiterJob/GetAllJobs"
	RecruiterJob_GetOneJob_FullMethodName  = "/job.RecruiterJob/GetOneJob"
	RecruiterJob_DeleteAJob_FullMethodName = "/job.RecruiterJob/DeleteAJob"
	RecruiterJob_UpdateAJob_FullMethodName = "/job.RecruiterJob/UpdateAJob"
)

// RecruiterJobClient is the client API for RecruiterJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecruiterJobClient interface {
	PostJob(ctx context.Context, in *JobOpeningRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error)
	GetAllJobs(ctx context.Context, in *GetAllJobsRequest, opts ...grpc.CallOption) (*GetAllJobsResponse, error)
	GetOneJob(ctx context.Context, in *GetAJobRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error)
	DeleteAJob(ctx context.Context, in *DeleteAJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateAJob(ctx context.Context, in *UpdateAJobRequest, opts ...grpc.CallOption) (*UpdateAJobResponse, error)
}

type recruiterJobClient struct {
	cc grpc.ClientConnInterface
}

func NewRecruiterJobClient(cc grpc.ClientConnInterface) RecruiterJobClient {
	return &recruiterJobClient{cc}
}

func (c *recruiterJobClient) PostJob(ctx context.Context, in *JobOpeningRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error) {
	out := new(JobOpeningResponse)
	err := c.cc.Invoke(ctx, RecruiterJob_PostJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterJobClient) GetAllJobs(ctx context.Context, in *GetAllJobsRequest, opts ...grpc.CallOption) (*GetAllJobsResponse, error) {
	out := new(GetAllJobsResponse)
	err := c.cc.Invoke(ctx, RecruiterJob_GetAllJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterJobClient) GetOneJob(ctx context.Context, in *GetAJobRequest, opts ...grpc.CallOption) (*JobOpeningResponse, error) {
	out := new(JobOpeningResponse)
	err := c.cc.Invoke(ctx, RecruiterJob_GetOneJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterJobClient) DeleteAJob(ctx context.Context, in *DeleteAJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RecruiterJob_DeleteAJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruiterJobClient) UpdateAJob(ctx context.Context, in *UpdateAJobRequest, opts ...grpc.CallOption) (*UpdateAJobResponse, error) {
	out := new(UpdateAJobResponse)
	err := c.cc.Invoke(ctx, RecruiterJob_UpdateAJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecruiterJobServer is the server API for RecruiterJob service.
// All implementations must embed UnimplementedRecruiterJobServer
// for forward compatibility
type RecruiterJobServer interface {
	PostJob(context.Context, *JobOpeningRequest) (*JobOpeningResponse, error)
	GetAllJobs(context.Context, *GetAllJobsRequest) (*GetAllJobsResponse, error)
	GetOneJob(context.Context, *GetAJobRequest) (*JobOpeningResponse, error)
	DeleteAJob(context.Context, *DeleteAJobRequest) (*emptypb.Empty, error)
	UpdateAJob(context.Context, *UpdateAJobRequest) (*UpdateAJobResponse, error)
	mustEmbedUnimplementedRecruiterJobServer()
}

// UnimplementedRecruiterJobServer must be embedded to have forward compatible implementations.
type UnimplementedRecruiterJobServer struct {
}

func (UnimplementedRecruiterJobServer) PostJob(context.Context, *JobOpeningRequest) (*JobOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostJob not implemented")
}
func (UnimplementedRecruiterJobServer) GetAllJobs(context.Context, *GetAllJobsRequest) (*GetAllJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllJobs not implemented")
}
func (UnimplementedRecruiterJobServer) GetOneJob(context.Context, *GetAJobRequest) (*JobOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneJob not implemented")
}
func (UnimplementedRecruiterJobServer) DeleteAJob(context.Context, *DeleteAJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAJob not implemented")
}
func (UnimplementedRecruiterJobServer) UpdateAJob(context.Context, *UpdateAJobRequest) (*UpdateAJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAJob not implemented")
}
func (UnimplementedRecruiterJobServer) mustEmbedUnimplementedRecruiterJobServer() {}

// UnsafeRecruiterJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecruiterJobServer will
// result in compilation errors.
type UnsafeRecruiterJobServer interface {
	mustEmbedUnimplementedRecruiterJobServer()
}

func RegisterRecruiterJobServer(s grpc.ServiceRegistrar, srv RecruiterJobServer) {
	s.RegisterService(&RecruiterJob_ServiceDesc, srv)
}

func _RecruiterJob_PostJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterJobServer).PostJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecruiterJob_PostJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterJobServer).PostJob(ctx, req.(*JobOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruiterJob_GetAllJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterJobServer).GetAllJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecruiterJob_GetAllJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterJobServer).GetAllJobs(ctx, req.(*GetAllJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruiterJob_GetOneJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterJobServer).GetOneJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecruiterJob_GetOneJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterJobServer).GetOneJob(ctx, req.(*GetAJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruiterJob_DeleteAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterJobServer).DeleteAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecruiterJob_DeleteAJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterJobServer).DeleteAJob(ctx, req.(*DeleteAJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruiterJob_UpdateAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruiterJobServer).UpdateAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecruiterJob_UpdateAJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruiterJobServer).UpdateAJob(ctx, req.(*UpdateAJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecruiterJob_ServiceDesc is the grpc.ServiceDesc for RecruiterJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecruiterJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.RecruiterJob",
	HandlerType: (*RecruiterJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostJob",
			Handler:    _RecruiterJob_PostJob_Handler,
		},
		{
			MethodName: "GetAllJobs",
			Handler:    _RecruiterJob_GetAllJobs_Handler,
		},
		{
			MethodName: "GetOneJob",
			Handler:    _RecruiterJob_GetOneJob_Handler,
		},
		{
			MethodName: "DeleteAJob",
			Handler:    _RecruiterJob_DeleteAJob_Handler,
		},
		{
			MethodName: "UpdateAJob",
			Handler:    _RecruiterJob_UpdateAJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/job/recruiter/job.proto",
}