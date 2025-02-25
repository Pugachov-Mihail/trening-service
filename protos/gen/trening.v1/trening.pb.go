// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: protos/trening.proto

package trening_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetTreningList
type GetTreningListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTreningListRequest) Reset() {
	*x = GetTreningListRequest{}
	mi := &file_protos_trening_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTreningListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreningListRequest) ProtoMessage() {}

func (x *GetTreningListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreningListRequest.ProtoReflect.Descriptor instead.
func (*GetTreningListRequest) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{0}
}

func (x *GetTreningListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetTreningListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetTreningListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TreningList   []*GetTreningList      `protobuf:"bytes,1,rep,name=treningList,proto3" json:"treningList,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTreningListResponse) Reset() {
	*x = GetTreningListResponse{}
	mi := &file_protos_trening_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTreningListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreningListResponse) ProtoMessage() {}

func (x *GetTreningListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreningListResponse.ProtoReflect.Descriptor instead.
func (*GetTreningListResponse) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{1}
}

func (x *GetTreningListResponse) GetTreningList() []*GetTreningList {
	if x != nil {
		return x.TreningList
	}
	return nil
}

type GetTreningList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Raiting       float32                `protobuf:"fixed32,5,opt,name=raiting,proto3" json:"raiting,omitempty"`
	TrenerInfo    *Trener                `protobuf:"bytes,6,opt,name=trenerInfo,proto3" json:"trenerInfo,omitempty"`
	TrenningInfo  []*Exercise            `protobuf:"bytes,7,rep,name=trenningInfo,proto3" json:"trenningInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTreningList) Reset() {
	*x = GetTreningList{}
	mi := &file_protos_trening_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTreningList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreningList) ProtoMessage() {}

func (x *GetTreningList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreningList.ProtoReflect.Descriptor instead.
func (*GetTreningList) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{2}
}

func (x *GetTreningList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetTreningList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetTreningList) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetTreningList) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GetTreningList) GetRaiting() float32 {
	if x != nil {
		return x.Raiting
	}
	return 0
}

func (x *GetTreningList) GetTrenerInfo() *Trener {
	if x != nil {
		return x.TrenerInfo
	}
	return nil
}

func (x *GetTreningList) GetTrenningInfo() []*Exercise {
	if x != nil {
		return x.TrenningInfo
	}
	return nil
}

type GetCurrentTreningResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Raiting       float32                `protobuf:"fixed32,5,opt,name=raiting,proto3" json:"raiting,omitempty"`
	TrenerInfo    *Trener                `protobuf:"bytes,6,opt,name=trenerInfo,proto3" json:"trenerInfo,omitempty"`
	TrenningInfo  []*Exercise            `protobuf:"bytes,7,rep,name=trenningInfo,proto3" json:"trenningInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCurrentTreningResponse) Reset() {
	*x = GetCurrentTreningResponse{}
	mi := &file_protos_trening_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentTreningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentTreningResponse) ProtoMessage() {}

func (x *GetCurrentTreningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentTreningResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentTreningResponse) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrentTreningResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCurrentTreningResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetCurrentTreningResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetCurrentTreningResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GetCurrentTreningResponse) GetRaiting() float32 {
	if x != nil {
		return x.Raiting
	}
	return 0
}

func (x *GetCurrentTreningResponse) GetTrenerInfo() *Trener {
	if x != nil {
		return x.TrenerInfo
	}
	return nil
}

func (x *GetCurrentTreningResponse) GetTrenningInfo() []*Exercise {
	if x != nil {
		return x.TrenningInfo
	}
	return nil
}

type Trener struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trener) Reset() {
	*x = Trener{}
	mi := &file_protos_trening_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trener) ProtoMessage() {}

func (x *Trener) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trener.ProtoReflect.Descriptor instead.
func (*Trener) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{4}
}

func (x *Trener) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trener) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Exercise struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Avatar        string                 `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Video         string                 `protobuf:"bytes,3,opt,name=video,proto3" json:"video,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Approach      []*Approach            `protobuf:"bytes,5,rep,name=approach,proto3" json:"approach,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	mi := &file_protos_trening_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exercise.ProtoReflect.Descriptor instead.
func (*Exercise) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{5}
}

func (x *Exercise) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Exercise) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Exercise) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

func (x *Exercise) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Exercise) GetApproach() []*Approach {
	if x != nil {
		return x.Approach
	}
	return nil
}

type Approach struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Weigth        float32                `protobuf:"fixed32,1,opt,name=Weigth,proto3" json:"Weigth,omitempty"`
	Repeat        int32                  `protobuf:"varint,2,opt,name=Repeat,proto3" json:"Repeat,omitempty"`
	Rest          int64                  `protobuf:"varint,3,opt,name=Rest,proto3" json:"Rest,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Approach) Reset() {
	*x = Approach{}
	mi := &file_protos_trening_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Approach) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approach) ProtoMessage() {}

func (x *Approach) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Approach.ProtoReflect.Descriptor instead.
func (*Approach) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{6}
}

func (x *Approach) GetWeigth() float32 {
	if x != nil {
		return x.Weigth
	}
	return 0
}

func (x *Approach) GetRepeat() int32 {
	if x != nil {
		return x.Repeat
	}
	return 0
}

func (x *Approach) GetRest() int64 {
	if x != nil {
		return x.Rest
	}
	return 0
}

// GetCurrentTrening
type TreningIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TreningIdRequest) Reset() {
	*x = TreningIdRequest{}
	mi := &file_protos_trening_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TreningIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreningIdRequest) ProtoMessage() {}

func (x *TreningIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreningIdRequest.ProtoReflect.Descriptor instead.
func (*TreningIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{7}
}

func (x *TreningIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// AddTreningUser
type AddTreningRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTreningRequest) Reset() {
	*x = AddTreningRequest{}
	mi := &file_protos_trening_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTreningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTreningRequest) ProtoMessage() {}

func (x *AddTreningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_trening_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTreningRequest.ProtoReflect.Descriptor instead.
func (*AddTreningRequest) Descriptor() ([]byte, []int) {
	return file_protos_trening_proto_rawDescGZIP(), []int{8}
}

func (x *AddTreningRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddTreningRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_protos_trening_proto protoreflect.FileDescriptor

var file_protos_trening_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65,
	0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4b, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x74, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x74, 0x72, 0x65,
	0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x69,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x72, 0x61, 0x69, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0a, 0x74, 0x72, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x65, 0x72,
	0x52, 0x0a, 0x74, 0x72, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x0c,
	0x74, 0x72, 0x65, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x0c, 0x74,
	0x72, 0x65, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x72, 0x61, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x27, 0x0a, 0x0a, 0x74, 0x72, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x0a,
	0x74, 0x72, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x0c, 0x74, 0x72,
	0x65, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x0c, 0x74, 0x72, 0x65,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x38, 0x0a, 0x06, 0x54, 0x72, 0x65,
	0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x61,
	0x63, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x61, 0x63, 0x68, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x61, 0x63, 0x68, 0x22, 0x4e, 0x0a,
	0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x61, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69,
	0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a,
	0x10, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xd9,
	0x02, 0x0a, 0x0e, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x65,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x65, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x72,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x54, 0x72,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x74, 0x72,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protos_trening_proto_rawDescOnce sync.Once
	file_protos_trening_proto_rawDescData []byte
)

func file_protos_trening_proto_rawDescGZIP() []byte {
	file_protos_trening_proto_rawDescOnce.Do(func() {
		file_protos_trening_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_trening_proto_rawDesc), len(file_protos_trening_proto_rawDesc)))
	})
	return file_protos_trening_proto_rawDescData
}

var file_protos_trening_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_trening_proto_goTypes = []any{
	(*GetTreningListRequest)(nil),     // 0: GetTreningListRequest
	(*GetTreningListResponse)(nil),    // 1: GetTreningListResponse
	(*GetTreningList)(nil),            // 2: GetTreningList
	(*GetCurrentTreningResponse)(nil), // 3: GetCurrentTreningResponse
	(*Trener)(nil),                    // 4: Trener
	(*Exercise)(nil),                  // 5: Exercise
	(*Approach)(nil),                  // 6: Approach
	(*TreningIdRequest)(nil),          // 7: TreningIdRequest
	(*AddTreningRequest)(nil),         // 8: AddTreningRequest
}
var file_protos_trening_proto_depIdxs = []int32{
	2,  // 0: GetTreningListResponse.treningList:type_name -> GetTreningList
	4,  // 1: GetTreningList.trenerInfo:type_name -> Trener
	5,  // 2: GetTreningList.trenningInfo:type_name -> Exercise
	4,  // 3: GetCurrentTreningResponse.trenerInfo:type_name -> Trener
	5,  // 4: GetCurrentTreningResponse.trenningInfo:type_name -> Exercise
	6,  // 5: Exercise.approach:type_name -> Approach
	0,  // 6: TreningService.GetTreningList:input_type -> GetTreningListRequest
	8,  // 7: TreningService.GetUserTrenings:input_type -> AddTreningRequest
	8,  // 8: TreningService.AddTreningUser:input_type -> AddTreningRequest
	8,  // 9: TreningService.DeletedTreningUser:input_type -> AddTreningRequest
	7,  // 10: TreningService.GetCurrentTrening:input_type -> TreningIdRequest
	1,  // 11: TreningService.GetTreningList:output_type -> GetTreningListResponse
	1,  // 12: TreningService.GetUserTrenings:output_type -> GetTreningListResponse
	1,  // 13: TreningService.AddTreningUser:output_type -> GetTreningListResponse
	1,  // 14: TreningService.DeletedTreningUser:output_type -> GetTreningListResponse
	3,  // 15: TreningService.GetCurrentTrening:output_type -> GetCurrentTreningResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_protos_trening_proto_init() }
func file_protos_trening_proto_init() {
	if File_protos_trening_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_trening_proto_rawDesc), len(file_protos_trening_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_trening_proto_goTypes,
		DependencyIndexes: file_protos_trening_proto_depIdxs,
		MessageInfos:      file_protos_trening_proto_msgTypes,
	}.Build()
	File_protos_trening_proto = out.File
	file_protos_trening_proto_goTypes = nil
	file_protos_trening_proto_depIdxs = nil
}
