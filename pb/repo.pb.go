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

type ReqUpdateOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastCommitID string `protobuf:"bytes,1,opt,name=LastCommitID,proto3" json:"LastCommitID,omitempty"`
	OldBranch    string `protobuf:"bytes,2,opt,name=OldBranch,proto3" json:"OldBranch,omitempty"`
	NewBranch    string `protobuf:"bytes,3,opt,name=NewBranch,proto3" json:"NewBranch,omitempty"`
	OldTreeName  string `protobuf:"bytes,4,opt,name=OldTreeName,proto3" json:"OldTreeName,omitempty"`
	NewTreeName  string `protobuf:"bytes,5,opt,name=NewTreeName,proto3" json:"NewTreeName,omitempty"`
	Message      string `protobuf:"bytes,6,opt,name=Message,proto3" json:"Message,omitempty"`
	Content      string `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`
	IsNewFile    bool   `protobuf:"varint,8,opt,name=IsNewFile,proto3" json:"IsNewFile,omitempty"`
	UserOrOrg    string `protobuf:"bytes,9,opt,name=UserOrOrg,proto3" json:"UserOrOrg,omitempty"`
	ProjectName  string `protobuf:"bytes,10,opt,name=ProjectName,proto3" json:"ProjectName,omitempty"`
}

func (x *ReqUpdateOptions) Reset() {
	*x = ReqUpdateOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateOptions) ProtoMessage() {}

func (x *ReqUpdateOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReqUpdateOptions.ProtoReflect.Descriptor instead.
func (*ReqUpdateOptions) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{3}
}

func (x *ReqUpdateOptions) GetLastCommitID() string {
	if x != nil {
		return x.LastCommitID
	}
	return ""
}

func (x *ReqUpdateOptions) GetOldBranch() string {
	if x != nil {
		return x.OldBranch
	}
	return ""
}

func (x *ReqUpdateOptions) GetNewBranch() string {
	if x != nil {
		return x.NewBranch
	}
	return ""
}

func (x *ReqUpdateOptions) GetOldTreeName() string {
	if x != nil {
		return x.OldTreeName
	}
	return ""
}

func (x *ReqUpdateOptions) GetNewTreeName() string {
	if x != nil {
		return x.NewTreeName
	}
	return ""
}

func (x *ReqUpdateOptions) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ReqUpdateOptions) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReqUpdateOptions) GetIsNewFile() bool {
	if x != nil {
		return x.IsNewFile
	}
	return false
}

func (x *ReqUpdateOptions) GetUserOrOrg() string {
	if x != nil {
		return x.UserOrOrg
	}
	return ""
}

func (x *ReqUpdateOptions) GetProjectName() string {
	if x != nil {
		return x.ProjectName
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
		mi := &file_pb_repo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespBool) ProtoMessage() {}

func (x *RespBool) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[4]
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
	return file_pb_repo_proto_rawDescGZIP(), []int{4}
}

func (x *RespBool) GetTrueOrFalse() bool {
	if x != nil {
		return x.TrueOrFalse
	}
	return false
}

type RespStructList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	CommitId string `protobuf:"bytes,2,opt,name=CommitId,proto3" json:"CommitId,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	When     string `protobuf:"bytes,4,opt,name=When,proto3" json:"When,omitempty"`
}

func (x *RespStructList) Reset() {
	*x = RespStructList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespStructList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespStructList) ProtoMessage() {}

func (x *RespStructList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespStructList.ProtoReflect.Descriptor instead.
func (*RespStructList) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{5}
}

func (x *RespStructList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RespStructList) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *RespStructList) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RespStructList) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

type RespStructNewest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	AuthorName string `protobuf:"bytes,3,opt,name=AuthorName,proto3" json:"AuthorName,omitempty"`
	When       string `protobuf:"bytes,4,opt,name=When,proto3" json:"When,omitempty"`
}

func (x *RespStructNewest) Reset() {
	*x = RespStructNewest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespStructNewest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespStructNewest) ProtoMessage() {}

func (x *RespStructNewest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespStructNewest.ProtoReflect.Descriptor instead.
func (*RespStructNewest) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{6}
}

func (x *RespStructNewest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RespStructNewest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RespStructNewest) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *RespStructNewest) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

type RespList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Newest *RespStructNewest `protobuf:"bytes,1,opt,name=Newest,proto3" json:"Newest,omitempty"`
	List   []*RespStructList `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *RespList) Reset() {
	*x = RespList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_repo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespList) ProtoMessage() {}

func (x *RespList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_repo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespList.ProtoReflect.Descriptor instead.
func (*RespList) Descriptor() ([]byte, []int) {
	return file_pb_repo_proto_rawDescGZIP(), []int{7}
}

func (x *RespList) GetNewest() *RespStructNewest {
	if x != nil {
		return x.Newest
	}
	return nil
}

func (x *RespList) GetList() []*RespStructList {
	if x != nil {
		return x.List
	}
	return nil
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
	0x50, 0x61, 0x74, 0x68, 0x22, 0xc8, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x4f, 0x6c, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4f, 0x6c, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x65, 0x77, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x65, 0x77, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x6c, 0x64,
	0x54, 0x72, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4f, 0x6c, 0x64, 0x54, 0x72, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e,
	0x65, 0x77, 0x54, 0x72, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x4f, 0x72, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x4f, 0x72, 0x67, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x2c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x72, 0x75, 0x65, 0x4f, 0x72, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x54, 0x72, 0x75, 0x65, 0x4f, 0x72, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x22, 0x68, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x70, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x68, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x57, 0x68, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x68, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x57, 0x68, 0x65, 0x6e, 0x22, 0x60, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x52, 0x06, 0x4e, 0x65, 0x77,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xb0, 0x01, 0x0a, 0x0b,
	0x52, 0x65, 0x70, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x42, 0x61,
	0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x71, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x06, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pb_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_repo_proto_goTypes = []interface{}{
	(*Message)(nil),          // 0: pb.Message
	(*ReqBase)(nil),          // 1: pb.ReqBase
	(*ReqList)(nil),          // 2: pb.ReqList
	(*ReqUpdateOptions)(nil), // 3: pb.ReqUpdateOptions
	(*RespBool)(nil),         // 4: pb.RespBool
	(*RespStructList)(nil),   // 5: pb.RespStructList
	(*RespStructNewest)(nil), // 6: pb.RespStructNewest
	(*RespList)(nil),         // 7: pb.RespList
}
var file_pb_repo_proto_depIdxs = []int32{
	6, // 0: pb.RespList.Newest:type_name -> pb.RespStructNewest
	5, // 1: pb.RespList.List:type_name -> pb.RespStructList
	1, // 2: pb.RepoService.Create:input_type -> pb.ReqBase
	1, // 3: pb.RepoService.Delete:input_type -> pb.ReqBase
	2, // 4: pb.RepoService.List:input_type -> pb.ReqList
	3, // 5: pb.RepoService.Editor:input_type -> pb.ReqUpdateOptions
	4, // 6: pb.RepoService.Create:output_type -> pb.RespBool
	4, // 7: pb.RepoService.Delete:output_type -> pb.RespBool
	7, // 8: pb.RepoService.List:output_type -> pb.RespList
	4, // 9: pb.RepoService.Editor:output_type -> pb.RespBool
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			switch v := v.(*ReqUpdateOptions); i {
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
		file_pb_repo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_repo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespStructList); i {
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
		file_pb_repo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespStructNewest); i {
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
		file_pb_repo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespList); i {
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
			NumMessages:   8,
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
	Create(ctx context.Context, in *ReqBase, opts ...grpc.CallOption) (*RespBool, error)
	Delete(ctx context.Context, in *ReqBase, opts ...grpc.CallOption) (*RespBool, error)
	List(ctx context.Context, in *ReqList, opts ...grpc.CallOption) (*RespList, error)
	Editor(ctx context.Context, in *ReqUpdateOptions, opts ...grpc.CallOption) (*RespBool, error)
}

type repoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoServiceClient(cc grpc.ClientConnInterface) RepoServiceClient {
	return &repoServiceClient{cc}
}

func (c *repoServiceClient) Create(ctx context.Context, in *ReqBase, opts ...grpc.CallOption) (*RespBool, error) {
	out := new(RespBool)
	err := c.cc.Invoke(ctx, "/pb.RepoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) Delete(ctx context.Context, in *ReqBase, opts ...grpc.CallOption) (*RespBool, error) {
	out := new(RespBool)
	err := c.cc.Invoke(ctx, "/pb.RepoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) List(ctx context.Context, in *ReqList, opts ...grpc.CallOption) (*RespList, error) {
	out := new(RespList)
	err := c.cc.Invoke(ctx, "/pb.RepoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoServiceClient) Editor(ctx context.Context, in *ReqUpdateOptions, opts ...grpc.CallOption) (*RespBool, error) {
	out := new(RespBool)
	err := c.cc.Invoke(ctx, "/pb.RepoService/Editor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServiceServer is the server API for RepoService service.
type RepoServiceServer interface {
	Create(context.Context, *ReqBase) (*RespBool, error)
	Delete(context.Context, *ReqBase) (*RespBool, error)
	List(context.Context, *ReqList) (*RespList, error)
	Editor(context.Context, *ReqUpdateOptions) (*RespBool, error)
}

// UnimplementedRepoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepoServiceServer struct {
}

func (*UnimplementedRepoServiceServer) Create(context.Context, *ReqBase) (*RespBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRepoServiceServer) Delete(context.Context, *ReqBase) (*RespBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRepoServiceServer) List(context.Context, *ReqList) (*RespList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedRepoServiceServer) Editor(context.Context, *ReqUpdateOptions) (*RespBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Editor not implemented")
}

func RegisterRepoServiceServer(s *grpc.Server, srv RepoServiceServer) {
	s.RegisterService(&_RepoService_serviceDesc, srv)
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

func _RepoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RepoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).Delete(ctx, req.(*ReqBase))
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

func _RepoService_Editor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServiceServer).Editor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RepoService/Editor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServiceServer).Editor(ctx, req.(*ReqUpdateOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RepoService",
	HandlerType: (*RepoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RepoService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RepoService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RepoService_List_Handler,
		},
		{
			MethodName: "Editor",
			Handler:    _RepoService_Editor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/repo.proto",
}
