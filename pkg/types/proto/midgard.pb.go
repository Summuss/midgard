// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed by
// a GNU GPL-3.0 license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: midgard.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingInput) Reset() {
	*x = PingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingInput) ProtoMessage() {}

func (x *PingInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingInput.ProtoReflect.Descriptor instead.
func (*PingInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{0}
}

type PingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	GoVersion string `protobuf:"bytes,2,opt,name=GoVersion,proto3" json:"GoVersion,omitempty"`
	BuildTime string `protobuf:"bytes,3,opt,name=BuildTime,proto3" json:"BuildTime,omitempty"`
}

func (x *PingOutput) Reset() {
	*x = PingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingOutput) ProtoMessage() {}

func (x *PingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingOutput.ProtoReflect.Descriptor instead.
func (*PingOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{1}
}

func (x *PingOutput) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PingOutput) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *PingOutput) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

type AllocateURLInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DesiredPath string `protobuf:"bytes,1,opt,name=DesiredPath,proto3" json:"DesiredPath,omitempty"`
	SourcePath  string `protobuf:"bytes,2,opt,name=SourcePath,proto3" json:"SourcePath,omitempty"`
}

func (x *AllocateURLInput) Reset() {
	*x = AllocateURLInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateURLInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateURLInput) ProtoMessage() {}

func (x *AllocateURLInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateURLInput.ProtoReflect.Descriptor instead.
func (*AllocateURLInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{2}
}

func (x *AllocateURLInput) GetDesiredPath() string {
	if x != nil {
		return x.DesiredPath
	}
	return ""
}

func (x *AllocateURLInput) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

type AllocateURLOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL     string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AllocateURLOutput) Reset() {
	*x = AllocateURLOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateURLOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateURLOutput) ProtoMessage() {}

func (x *AllocateURLOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateURLOutput.ProtoReflect.Descriptor instead.
func (*AllocateURLOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{3}
}

func (x *AllocateURLOutput) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *AllocateURLOutput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_midgard_proto protoreflect.FileDescriptor

var file_midgard_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x69, 0x64, 0x67, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x62, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x47,
	0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x47, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x3f, 0x0a,
	0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x7c,
	0x0a, 0x07, 0x4d, 0x69, 0x64, 0x67, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x55, 0x52, 0x4c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_midgard_proto_rawDescOnce sync.Once
	file_midgard_proto_rawDescData = file_midgard_proto_rawDesc
)

func file_midgard_proto_rawDescGZIP() []byte {
	file_midgard_proto_rawDescOnce.Do(func() {
		file_midgard_proto_rawDescData = protoimpl.X.CompressGZIP(file_midgard_proto_rawDescData)
	})
	return file_midgard_proto_rawDescData
}

var file_midgard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_midgard_proto_goTypes = []interface{}{
	(*PingInput)(nil),         // 0: proto.PingInput
	(*PingOutput)(nil),        // 1: proto.PingOutput
	(*AllocateURLInput)(nil),  // 2: proto.AllocateURLInput
	(*AllocateURLOutput)(nil), // 3: proto.AllocateURLOutput
}
var file_midgard_proto_depIdxs = []int32{
	0, // 0: proto.Midgard.Ping:input_type -> proto.PingInput
	2, // 1: proto.Midgard.AllocateURL:input_type -> proto.AllocateURLInput
	1, // 2: proto.Midgard.Ping:output_type -> proto.PingOutput
	3, // 3: proto.Midgard.AllocateURL:output_type -> proto.AllocateURLOutput
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_midgard_proto_init() }
func file_midgard_proto_init() {
	if File_midgard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_midgard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingInput); i {
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
		file_midgard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingOutput); i {
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
		file_midgard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateURLInput); i {
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
		file_midgard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateURLOutput); i {
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
			RawDescriptor: file_midgard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_midgard_proto_goTypes,
		DependencyIndexes: file_midgard_proto_depIdxs,
		MessageInfos:      file_midgard_proto_msgTypes,
	}.Build()
	File_midgard_proto = out.File
	file_midgard_proto_rawDesc = nil
	file_midgard_proto_goTypes = nil
	file_midgard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MidgardClient is the client API for Midgard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MidgardClient interface {
	Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error)
	AllocateURL(ctx context.Context, in *AllocateURLInput, opts ...grpc.CallOption) (*AllocateURLOutput, error)
}

type midgardClient struct {
	cc grpc.ClientConnInterface
}

func NewMidgardClient(cc grpc.ClientConnInterface) MidgardClient {
	return &midgardClient{cc}
}

func (c *midgardClient) Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error) {
	out := new(PingOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) AllocateURL(ctx context.Context, in *AllocateURLInput, opts ...grpc.CallOption) (*AllocateURLOutput, error) {
	out := new(AllocateURLOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/AllocateURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MidgardServer is the server API for Midgard service.
type MidgardServer interface {
	Ping(context.Context, *PingInput) (*PingOutput, error)
	AllocateURL(context.Context, *AllocateURLInput) (*AllocateURLOutput, error)
}

// UnimplementedMidgardServer can be embedded to have forward compatible implementations.
type UnimplementedMidgardServer struct {
}

func (*UnimplementedMidgardServer) Ping(context.Context, *PingInput) (*PingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMidgardServer) AllocateURL(context.Context, *AllocateURLInput) (*AllocateURLOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateURL not implemented")
}

func RegisterMidgardServer(s *grpc.Server, srv MidgardServer) {
	s.RegisterService(&_Midgard_serviceDesc, srv)
}

func _Midgard_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).Ping(ctx, req.(*PingInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_AllocateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateURLInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).AllocateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/AllocateURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).AllocateURL(ctx, req.(*AllocateURLInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Midgard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Midgard",
	HandlerType: (*MidgardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Midgard_Ping_Handler,
		},
		{
			MethodName: "AllocateURL",
			Handler:    _Midgard_AllocateURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "midgard.proto",
}
