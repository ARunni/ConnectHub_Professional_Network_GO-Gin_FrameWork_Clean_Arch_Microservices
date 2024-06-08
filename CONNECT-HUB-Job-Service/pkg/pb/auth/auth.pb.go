// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: pkg/pb/auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDetailsByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *GetDetailsByIdRequest) Reset() {
	*x = GetDetailsByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailsByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailsByIdRequest) ProtoMessage() {}

func (x *GetDetailsByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailsByIdRequest.ProtoReflect.Descriptor instead.
func (*GetDetailsByIdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetDetailsByIdRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type GetDetailsByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetDetailsByIdResponse) Reset() {
	*x = GetDetailsByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailsByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailsByIdResponse) ProtoMessage() {}

func (x *GetDetailsByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailsByIdResponse.ProtoReflect.Descriptor instead.
func (*GetDetailsByIdResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GetDetailsByIdResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetDetailsByIdResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_pkg_pb_auth_auth_proto protoreflect.FileDescriptor

var file_pkg_pb_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32,
	0x60, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x41, 0x75, 0x74, 0x68, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_auth_auth_proto_rawDescOnce sync.Once
	file_pkg_pb_auth_auth_proto_rawDescData = file_pkg_pb_auth_auth_proto_rawDesc
)

func file_pkg_pb_auth_auth_proto_rawDescGZIP() []byte {
	file_pkg_pb_auth_auth_proto_rawDescOnce.Do(func() {
		file_pkg_pb_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_auth_auth_proto_rawDescData)
	})
	return file_pkg_pb_auth_auth_proto_rawDescData
}

var file_pkg_pb_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_auth_auth_proto_goTypes = []interface{}{
	(*GetDetailsByIdRequest)(nil),  // 0: job_auth.GetDetailsByIdRequest
	(*GetDetailsByIdResponse)(nil), // 1: job_auth.GetDetailsByIdResponse
}
var file_pkg_pb_auth_auth_proto_depIdxs = []int32{
	0, // 0: job_auth.jobAuth.GetDetailsById:input_type -> job_auth.GetDetailsByIdRequest
	1, // 1: job_auth.jobAuth.GetDetailsById:output_type -> job_auth.GetDetailsByIdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_auth_auth_proto_init() }
func file_pkg_pb_auth_auth_proto_init() {
	if File_pkg_pb_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailsByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailsByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_auth_auth_proto_goTypes,
		DependencyIndexes: file_pkg_pb_auth_auth_proto_depIdxs,
		MessageInfos:      file_pkg_pb_auth_auth_proto_msgTypes,
	}.Build()
	File_pkg_pb_auth_auth_proto = out.File
	file_pkg_pb_auth_auth_proto_rawDesc = nil
	file_pkg_pb_auth_auth_proto_goTypes = nil
	file_pkg_pb_auth_auth_proto_depIdxs = nil
}