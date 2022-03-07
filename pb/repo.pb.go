// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/repo.proto

package pb

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

// request
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type ReqBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserOrOrg   string `protobuf:"bytes,1,opt,name=UserOrOrg,proto3" json:"UserOrOrg,omitempty"`
	ProjectName string `protobuf:"bytes,2,opt,name=ProjectName,proto3" json:"ProjectName,omitempty"`
}

func (x *ReqBase) Reset() {
	*x = ReqBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqBase) ProtoMessage() {}

func (x *ReqBase) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqBase.ProtoReflect.Descriptor instead.
func (*ReqBase) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{1}
}

func (x *ReqBase) GetUserOrOrg() string {
	if x != nil {
		return x.UserOrOrg
	}
	return ""
}

func (x *ReqBase) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type ReqList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserOrOrg   string `protobuf:"bytes,1,opt,name=UserOrOrg,proto3" json:"UserOrOrg,omitempty"`
	ProjectName string `protobuf:"bytes,2,opt,name=ProjectName,proto3" json:"ProjectName,omitempty"`
	TreePath    string `protobuf:"bytes,3,opt,name=TreePath,proto3" json:"TreePath,omitempty"`
}

func (x *ReqList) Reset() {
	*x = ReqList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqList) ProtoMessage() {}

func (x *ReqList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqList.ProtoReflect.Descriptor instead.
func (*ReqList) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{2}
}

func (x *ReqList) GetUserOrOrg() string {
	if x != nil {
		return x.UserOrOrg
	}
	return ""
}

func (x *ReqList) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ReqList) GetTreePath() string {
	if x != nil {
		return x.TreePath
	}
	return ""
}

type RespBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrueOrFalse bool `protobuf:"varint,1,opt,name=TrueOrFalse,proto3" json:"TrueOrFalse,omitempty"`
}

func (x *RespBool) Reset() {
	*x = RespBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespBool) ProtoMessage() {}

func (x *RespBool) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespBool.ProtoReflect.Descriptor instead.
func (*RespBool) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{3}
}

func (x *RespBool) GetTrueOrFalse() bool {
	if x != nil {
		return x.TrueOrFalse
	}
	return false
}

var File_pb_repo_proto protoreflect.FileDescriptor

var file_pb_repo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x49, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x4f, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x4f, 0x72, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x72, 0x4f, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x72, 0x4f, 0x72, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x72, 0x65, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x75, 0x65, 0x4f, 0x72, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x54, 0x72, 0x75, 0x65, 0x4f, 0x72, 0x46, 0x61, 0x6c,
	0x73, 0x65, 0x32, 0x79, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x01, 0x53, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x71, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_repo_proto_rawDescOnce sync.Once
	file_pb_repo_proto_rawDescData = file_pb_repo_proto_rawDesc
)

func file_pb_repo_proto_rawDescGZIP() []byte {
	file_pb_repo_proto_rawDescOnce.Do(func() {
		file_pb_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_repo_proto_rawDescData)
	})
	return file_pb_repo_proto_rawDescData
}

var file_pb_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_repo_proto_goTypes = []interface{}{
	(*Message)(nil),  // 0: pb.Message
	(*ReqBase)(nil),  // 1: pb.ReqBase
	(*ReqList)(nil),  // 2: pb.ReqList
	(*RespBool)(nil), // 3: pb.RespBool
}
var file_pb_repo_proto_depIdxs = []int32{
	0, // 0: pb.RepoService.S:input_type -> pb.Message
	1, // 1: pb.RepoService.Create:input_type -> pb.ReqBase
	2, // 2: pb.RepoService.List:input_type -> pb.ReqList
	0, // 3: pb.RepoService.S:output_type -> pb.Message
	3, // 4: pb.RepoService.Create:output_type -> pb.RespBool
	0, // 5: pb.RepoService.List:output_type -> pb.Message
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_repo_proto_init() }
func file_pb_repo_proto_init() {
	if File_pb_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pb_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqBase); i {
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
		file_pb_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqList); i {
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
		file_pb_repo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespBool); i {
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
			RawDescriptor: file_pb_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_repo_proto_goTypes,
		DependencyIndexes: file_pb_repo_proto_depIdxs,
		MessageInfos:      file_pb_repo_proto_msgTypes,
	}.Build()
	File_pb_repo_proto = out.File
	file_pb_repo_proto_rawDesc = nil
	file_pb_repo_proto_goTypes = nil
	file_pb_repo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RepoServiceClient is the client API for RepoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepoServiceClient interface {
	S(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Create(ctx context.Context, in *ReqBase, opts ...grpc.CallOption) (*RespBool, error)
	List(ctx context.Context, in *ReqList, opts ...grpc.CallOption) (*Message, error)
}

type repoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoServiceClient(cc grpc.ClientConnInterface) RepoServiceClient {
	return &repoServiceClient{cc}
}

func (c *repoServiceClient) S(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.RepoService/S", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) Create(ctx context.Context, in *ReqBase, opts ...grpc.CallOption) (*RespBool, error) {
	out := new(RespBool)
	err := c.cc.Invoke(ctx, "/pb.RepoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) List(ctx context.Context, in *ReqList, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.RepoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServiceServer is the server API for RepoService service.
type RepoServiceServer interface {
	S(context.Context, *Message) (*Message, error)
	Create(context.Context, *ReqBase) (*RespBool, error)
	List(context.Context, *ReqList) (*Message, error)
}

// UnimplementedRepoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepoServiceServer struct {
}

func (*UnimplementedRepoServiceServer) S(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S not implemented")
}
func (*UnimplementedRepoServiceServer) Create(context.Context, *ReqBase) (*RespBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRepoServiceServer) List(context.Context, *ReqList) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRepoServiceServer(s *grpc.Server, srv RepoServiceServer) {
	s.RegisterService(&_RepoService_serviceDesc, srv)
}

func _RepoService_S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RepoService/S",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).S(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RepoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).Create(ctx, req.(*ReqBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RepoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).List(ctx, req.(*ReqList))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RepoService",
	HandlerType: (*RepoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "S",
			Handler:    _RepoService_S_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RepoService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RepoService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/repo.proto",
}
