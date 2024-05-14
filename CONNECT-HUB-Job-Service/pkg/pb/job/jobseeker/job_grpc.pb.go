// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/job/jobseeker/job.proto

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
	JobseekerJob_JobSeekerGetAllJobs_FullMethodName = "/job.JobseekerJob/JobSeekerGetAllJobs"
)

// JobseekerJobClient is the client API for JobseekerJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobseekerJobClient interface {
	JobSeekerGetAllJobs(ctx context.Context, in *JobSeekerGetAllJobsRequest, opts ...grpc.CallOption) (*JobSeekerGetAllJobsResponse, error)
}

type jobseekerJobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobseekerJobClient(cc grpc.ClientConnInterface) JobseekerJobClient {
	return &jobseekerJobClient{cc}
}

func (c *jobseekerJobClient) JobSeekerGetAllJobs(ctx context.Context, in *JobSeekerGetAllJobsRequest, opts ...grpc.CallOption) (*JobSeekerGetAllJobsResponse, error) {
	out := new(JobSeekerGetAllJobsResponse)
	err := c.cc.Invoke(ctx, JobseekerJob_JobSeekerGetAllJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobseekerJobServer is the server API for JobseekerJob service.
// All implementations must embed UnimplementedJobseekerJobServer
// for forward compatibility
type JobseekerJobServer interface {
	JobSeekerGetAllJobs(context.Context, *JobSeekerGetAllJobsRequest) (*JobSeekerGetAllJobsResponse, error)
	mustEmbedUnimplementedJobseekerJobServer()
}

// UnimplementedJobseekerJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobseekerJobServer struct {
}

func (UnimplementedJobseekerJobServer) JobSeekerGetAllJobs(context.Context, *JobSeekerGetAllJobsRequest) (*JobSeekerGetAllJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerGetAllJobs not implemented")
}
func (UnimplementedJobseekerJobServer) mustEmbedUnimplementedJobseekerJobServer() {}

// UnsafeJobseekerJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobseekerJobServer will
// result in compilation errors.
type UnsafeJobseekerJobServer interface {
	mustEmbedUnimplementedJobseekerJobServer()
}

func RegisterJobseekerJobServer(s grpc.ServiceRegistrar, srv JobseekerJobServer) {
	s.RegisterService(&JobseekerJob_ServiceDesc, srv)
}

func _JobseekerJob_JobSeekerGetAllJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerGetAllJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobseekerJobServer).JobSeekerGetAllJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobseekerJob_JobSeekerGetAllJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobseekerJobServer).JobSeekerGetAllJobs(ctx, req.(*JobSeekerGetAllJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobseekerJob_ServiceDesc is the grpc.ServiceDesc for JobseekerJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobseekerJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.JobseekerJob",
	HandlerType: (*JobseekerJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JobSeekerGetAllJobs",
			Handler:    _JobseekerJob_JobSeekerGetAllJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/job/jobseeker/job.proto",
}
