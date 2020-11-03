// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/mastiff.proto

package mastiff

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 创建Id
type CreateIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
}

func (x *CreateIdReq) Reset() {
	*x = CreateIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastiff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdReq) ProtoMessage() {}

func (x *CreateIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastiff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdReq.ProtoReflect.Descriptor instead.
func (*CreateIdReq) Descriptor() ([]byte, []int) {
	return file_proto_mastiff_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIdReq) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type CreateIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CreateIdRes) Reset() {
	*x = CreateIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastiff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdRes) ProtoMessage() {}

func (x *CreateIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastiff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdRes.ProtoReflect.Descriptor instead.
func (*CreateIdRes) Descriptor() ([]byte, []int) {
	return file_proto_mastiff_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIdRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_mastiff_proto protoreflect.FileDescriptor

var file_proto_mastiff_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x69, 0x66, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x69, 0x66,
	0x66, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x25, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x60, 0x0a, 0x07, 0x4d, 0x61, 0x73,
	0x74, 0x69, 0x66, 0x66, 0x12, 0x55, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x69, 0x66, 0x66, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x6d,
	0x61, 0x73, 0x74, 0x69, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x0f, 0x5a, 0x0d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x61, 0x73, 0x74, 0x69, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mastiff_proto_rawDescOnce sync.Once
	file_proto_mastiff_proto_rawDescData = file_proto_mastiff_proto_rawDesc
)

func file_proto_mastiff_proto_rawDescGZIP() []byte {
	file_proto_mastiff_proto_rawDescOnce.Do(func() {
		file_proto_mastiff_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mastiff_proto_rawDescData)
	})
	return file_proto_mastiff_proto_rawDescData
}

var file_proto_mastiff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_mastiff_proto_goTypes = []interface{}{
	(*CreateIdReq)(nil), // 0: pb_mastiff.CreateIdReq
	(*CreateIdRes)(nil), // 1: pb_mastiff.CreateIdRes
}
var file_proto_mastiff_proto_depIdxs = []int32{
	0, // 0: pb_mastiff.Mastiff.createId:input_type -> pb_mastiff.CreateIdReq
	1, // 1: pb_mastiff.Mastiff.createId:output_type -> pb_mastiff.CreateIdRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mastiff_proto_init() }
func file_proto_mastiff_proto_init() {
	if File_proto_mastiff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mastiff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdReq); i {
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
		file_proto_mastiff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdRes); i {
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
			RawDescriptor: file_proto_mastiff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mastiff_proto_goTypes,
		DependencyIndexes: file_proto_mastiff_proto_depIdxs,
		MessageInfos:      file_proto_mastiff_proto_msgTypes,
	}.Build()
	File_proto_mastiff_proto = out.File
	file_proto_mastiff_proto_rawDesc = nil
	file_proto_mastiff_proto_goTypes = nil
	file_proto_mastiff_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MastiffClient is the client API for Mastiff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MastiffClient interface {
	CreateId(ctx context.Context, in *CreateIdReq, opts ...grpc.CallOption) (*CreateIdRes, error)
}

type mastiffClient struct {
	cc grpc.ClientConnInterface
}

func NewMastiffClient(cc grpc.ClientConnInterface) MastiffClient {
	return &mastiffClient{cc}
}

func (c *mastiffClient) CreateId(ctx context.Context, in *CreateIdReq, opts ...grpc.CallOption) (*CreateIdRes, error) {
	out := new(CreateIdRes)
	err := c.cc.Invoke(ctx, "/pb_mastiff.Mastiff/createId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MastiffServer is the server API for Mastiff service.
type MastiffServer interface {
	CreateId(context.Context, *CreateIdReq) (*CreateIdRes, error)
}

// UnimplementedMastiffServer can be embedded to have forward compatible implementations.
type UnimplementedMastiffServer struct {
}

func (*UnimplementedMastiffServer) CreateId(context.Context, *CreateIdReq) (*CreateIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateId not implemented")
}

func RegisterMastiffServer(s *grpc.Server, srv MastiffServer) {
	s.RegisterService(&_Mastiff_serviceDesc, srv)
}

func _Mastiff_CreateId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MastiffServer).CreateId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_mastiff.Mastiff/CreateId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MastiffServer).CreateId(ctx, req.(*CreateIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mastiff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_mastiff.Mastiff",
	HandlerType: (*MastiffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createId",
			Handler:    _Mastiff_CreateId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mastiff.proto",
}