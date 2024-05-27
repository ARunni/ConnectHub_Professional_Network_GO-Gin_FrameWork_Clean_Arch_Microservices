// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/pb/auth/admin/admin.proto

package admin

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
	Admin_AdminLogin_FullMethodName          = "/adminauth.Admin/AdminLogin"
	Admin_GetJobseekers_FullMethodName       = "/adminauth.Admin/GetJobseekers"
	Admin_BlockJobseeker_FullMethodName      = "/adminauth.Admin/BlockJobseeker"
	Admin_UnBlockJobseeker_FullMethodName    = "/adminauth.Admin/UnBlockJobseeker"
	Admin_GetRecruiters_FullMethodName       = "/adminauth.Admin/GetRecruiters"
	Admin_BlockRecruiter_FullMethodName      = "/adminauth.Admin/BlockRecruiter"
	Admin_UnBlockRecruiter_FullMethodName    = "/adminauth.Admin/UnBlockRecruiter"
	Admin_GetJobseekerDetails_FullMethodName = "/adminauth.Admin/GetJobseekerDetails"
	Admin_GetRecruiterDetails_FullMethodName = "/adminauth.Admin/GetRecruiterDetails"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	AdminLogin(ctx context.Context, in *AdminLoginInRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	GetJobseekers(ctx context.Context, in *GetJobseekerRequest, opts ...grpc.CallOption) (*GetJobseekerResponse, error)
	BlockJobseeker(ctx context.Context, in *BlockJobseekerRequest, opts ...grpc.CallOption) (*BlockJobseekerResponse, error)
	UnBlockJobseeker(ctx context.Context, in *UnBlockJobseekerRequest, opts ...grpc.CallOption) (*UnBlockJobseekerResponse, error)
	GetRecruiters(ctx context.Context, in *GetRecruiterRequest, opts ...grpc.CallOption) (*GetRecruitersResponse, error)
	BlockRecruiter(ctx context.Context, in *BlockRecruiterRequest, opts ...grpc.CallOption) (*BlockRecruiterResponse, error)
	UnBlockRecruiter(ctx context.Context, in *UnBlockRecruiterRequest, opts ...grpc.CallOption) (*UnBlockRecruiterResponse, error)
	GetJobseekerDetails(ctx context.Context, in *GetJobseekerDetailsRequest, opts ...grpc.CallOption) (*GetJobseekerDetailsResponse, error)
	GetRecruiterDetails(ctx context.Context, in *GetRecruiterDetailsRequest, opts ...grpc.CallOption) (*GetRecruiterDetailsResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AdminLogin(ctx context.Context, in *AdminLoginInRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, Admin_AdminLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetJobseekers(ctx context.Context, in *GetJobseekerRequest, opts ...grpc.CallOption) (*GetJobseekerResponse, error) {
	out := new(GetJobseekerResponse)
	err := c.cc.Invoke(ctx, Admin_GetJobseekers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) BlockJobseeker(ctx context.Context, in *BlockJobseekerRequest, opts ...grpc.CallOption) (*BlockJobseekerResponse, error) {
	out := new(BlockJobseekerResponse)
	err := c.cc.Invoke(ctx, Admin_BlockJobseeker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UnBlockJobseeker(ctx context.Context, in *UnBlockJobseekerRequest, opts ...grpc.CallOption) (*UnBlockJobseekerResponse, error) {
	out := new(UnBlockJobseekerResponse)
	err := c.cc.Invoke(ctx, Admin_UnBlockJobseeker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetRecruiters(ctx context.Context, in *GetRecruiterRequest, opts ...grpc.CallOption) (*GetRecruitersResponse, error) {
	out := new(GetRecruitersResponse)
	err := c.cc.Invoke(ctx, Admin_GetRecruiters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) BlockRecruiter(ctx context.Context, in *BlockRecruiterRequest, opts ...grpc.CallOption) (*BlockRecruiterResponse, error) {
	out := new(BlockRecruiterResponse)
	err := c.cc.Invoke(ctx, Admin_BlockRecruiter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UnBlockRecruiter(ctx context.Context, in *UnBlockRecruiterRequest, opts ...grpc.CallOption) (*UnBlockRecruiterResponse, error) {
	out := new(UnBlockRecruiterResponse)
	err := c.cc.Invoke(ctx, Admin_UnBlockRecruiter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetJobseekerDetails(ctx context.Context, in *GetJobseekerDetailsRequest, opts ...grpc.CallOption) (*GetJobseekerDetailsResponse, error) {
	out := new(GetJobseekerDetailsResponse)
	err := c.cc.Invoke(ctx, Admin_GetJobseekerDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetRecruiterDetails(ctx context.Context, in *GetRecruiterDetailsRequest, opts ...grpc.CallOption) (*GetRecruiterDetailsResponse, error) {
	out := new(GetRecruiterDetailsResponse)
	err := c.cc.Invoke(ctx, Admin_GetRecruiterDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	AdminLogin(context.Context, *AdminLoginInRequest) (*AdminLoginResponse, error)
	GetJobseekers(context.Context, *GetJobseekerRequest) (*GetJobseekerResponse, error)
	BlockJobseeker(context.Context, *BlockJobseekerRequest) (*BlockJobseekerResponse, error)
	UnBlockJobseeker(context.Context, *UnBlockJobseekerRequest) (*UnBlockJobseekerResponse, error)
	GetRecruiters(context.Context, *GetRecruiterRequest) (*GetRecruitersResponse, error)
	BlockRecruiter(context.Context, *BlockRecruiterRequest) (*BlockRecruiterResponse, error)
	UnBlockRecruiter(context.Context, *UnBlockRecruiterRequest) (*UnBlockRecruiterResponse, error)
	GetJobseekerDetails(context.Context, *GetJobseekerDetailsRequest) (*GetJobseekerDetailsResponse, error)
	GetRecruiterDetails(context.Context, *GetRecruiterDetailsRequest) (*GetRecruiterDetailsResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AdminLogin(context.Context, *AdminLoginInRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServer) GetJobseekers(context.Context, *GetJobseekerRequest) (*GetJobseekerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobseekers not implemented")
}
func (UnimplementedAdminServer) BlockJobseeker(context.Context, *BlockJobseekerRequest) (*BlockJobseekerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockJobseeker not implemented")
}
func (UnimplementedAdminServer) UnBlockJobseeker(context.Context, *UnBlockJobseekerRequest) (*UnBlockJobseekerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockJobseeker not implemented")
}
func (UnimplementedAdminServer) GetRecruiters(context.Context, *GetRecruiterRequest) (*GetRecruitersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecruiters not implemented")
}
func (UnimplementedAdminServer) BlockRecruiter(context.Context, *BlockRecruiterRequest) (*BlockRecruiterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockRecruiter not implemented")
}
func (UnimplementedAdminServer) UnBlockRecruiter(context.Context, *UnBlockRecruiterRequest) (*UnBlockRecruiterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockRecruiter not implemented")
}
func (UnimplementedAdminServer) GetJobseekerDetails(context.Context, *GetJobseekerDetailsRequest) (*GetJobseekerDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobseekerDetails not implemented")
}
func (UnimplementedAdminServer) GetRecruiterDetails(context.Context, *GetRecruiterDetailsRequest) (*GetRecruiterDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecruiterDetails not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLogin(ctx, req.(*AdminLoginInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetJobseekers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobseekerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetJobseekers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetJobseekers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetJobseekers(ctx, req.(*GetJobseekerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_BlockJobseeker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockJobseekerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).BlockJobseeker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_BlockJobseeker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).BlockJobseeker(ctx, req.(*BlockJobseekerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UnBlockJobseeker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnBlockJobseekerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UnBlockJobseeker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_UnBlockJobseeker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UnBlockJobseeker(ctx, req.(*UnBlockJobseekerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetRecruiters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecruiterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetRecruiters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetRecruiters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetRecruiters(ctx, req.(*GetRecruiterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_BlockRecruiter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRecruiterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).BlockRecruiter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_BlockRecruiter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).BlockRecruiter(ctx, req.(*BlockRecruiterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UnBlockRecruiter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnBlockRecruiterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UnBlockRecruiter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_UnBlockRecruiter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UnBlockRecruiter(ctx, req.(*UnBlockRecruiterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetJobseekerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobseekerDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetJobseekerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetJobseekerDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetJobseekerDetails(ctx, req.(*GetJobseekerDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetRecruiterDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecruiterDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetRecruiterDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetRecruiterDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetRecruiterDetails(ctx, req.(*GetRecruiterDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adminauth.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _Admin_AdminLogin_Handler,
		},
		{
			MethodName: "GetJobseekers",
			Handler:    _Admin_GetJobseekers_Handler,
		},
		{
			MethodName: "BlockJobseeker",
			Handler:    _Admin_BlockJobseeker_Handler,
		},
		{
			MethodName: "UnBlockJobseeker",
			Handler:    _Admin_UnBlockJobseeker_Handler,
		},
		{
			MethodName: "GetRecruiters",
			Handler:    _Admin_GetRecruiters_Handler,
		},
		{
			MethodName: "BlockRecruiter",
			Handler:    _Admin_BlockRecruiter_Handler,
		},
		{
			MethodName: "UnBlockRecruiter",
			Handler:    _Admin_UnBlockRecruiter_Handler,
		},
		{
			MethodName: "GetJobseekerDetails",
			Handler:    _Admin_GetJobseekerDetails_Handler,
		},
		{
			MethodName: "GetRecruiterDetails",
			Handler:    _Admin_GetRecruiterDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth/admin/admin.proto",
}
