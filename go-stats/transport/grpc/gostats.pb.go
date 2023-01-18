//*
// Generated by go-doudou v2.0.3.
// Don't edit!
//
// Version No.: v20221225

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: transport/grpc/gostats.proto

package grpc

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

type GetShelvesShelfBooksBookRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shelf int32  `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	Book  string `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetShelvesShelfBooksBookRpcRequest) Reset() {
	*x = GetShelvesShelfBooksBookRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_gostats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelvesShelfBooksBookRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelvesShelfBooksBookRpcRequest) ProtoMessage() {}

func (x *GetShelvesShelfBooksBookRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_gostats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelvesShelfBooksBookRpcRequest.ProtoReflect.Descriptor instead.
func (*GetShelvesShelfBooksBookRpcRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_gostats_proto_rawDescGZIP(), []int{0}
}

func (x *GetShelvesShelfBooksBookRpcRequest) GetShelf() int32 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *GetShelvesShelfBooksBookRpcRequest) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

type GetShelvesShelfBooksBookRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetShelvesShelfBooksBookRpcResponse) Reset() {
	*x = GetShelvesShelfBooksBookRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_gostats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelvesShelfBooksBookRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelvesShelfBooksBookRpcResponse) ProtoMessage() {}

func (x *GetShelvesShelfBooksBookRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_gostats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelvesShelfBooksBookRpcResponse.ProtoReflect.Descriptor instead.
func (*GetShelvesShelfBooksBookRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_gostats_proto_rawDescGZIP(), []int{1}
}

func (x *GetShelvesShelfBooksBookRpcResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type LargestRemainderRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PercentageRespVo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *LargestRemainderRpcResponse) Reset() {
	*x = LargestRemainderRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_gostats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LargestRemainderRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LargestRemainderRpcResponse) ProtoMessage() {}

func (x *LargestRemainderRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_gostats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LargestRemainderRpcResponse.ProtoReflect.Descriptor instead.
func (*LargestRemainderRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_gostats_proto_rawDescGZIP(), []int{2}
}

func (x *LargestRemainderRpcResponse) GetData() []*PercentageRespVo {
	if x != nil {
		return x.Data
	}
	return nil
}

// request vo
type PercentageReqVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key value pairs
	Data []*PercentageVo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// digit number after dot
	Places int32 `protobuf:"varint,2,opt,name=places,proto3" json:"places,omitempty"`
}

func (x *PercentageReqVo) Reset() {
	*x = PercentageReqVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_gostats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PercentageReqVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PercentageReqVo) ProtoMessage() {}

func (x *PercentageReqVo) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_gostats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PercentageReqVo.ProtoReflect.Descriptor instead.
func (*PercentageReqVo) Descriptor() ([]byte, []int) {
	return file_transport_grpc_gostats_proto_rawDescGZIP(), []int{3}
}

func (x *PercentageReqVo) GetData() []*PercentageVo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PercentageReqVo) GetPlaces() int32 {
	if x != nil {
		return x.Places
	}
	return 0
}

// result vo
type PercentageRespVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value   int32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Key     string  `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Percent float64 `protobuf:"fixed64,3,opt,name=percent,proto3" json:"percent,omitempty"`
	// formatted percentage
	PercentFormatted string `protobuf:"bytes,4,opt,name=percentFormatted,proto3" json:"percentFormatted,omitempty"`
}

func (x *PercentageRespVo) Reset() {
	*x = PercentageRespVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_gostats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PercentageRespVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PercentageRespVo) ProtoMessage() {}

func (x *PercentageRespVo) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_gostats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PercentageRespVo.ProtoReflect.Descriptor instead.
func (*PercentageRespVo) Descriptor() ([]byte, []int) {
	return file_transport_grpc_gostats_proto_rawDescGZIP(), []int{4}
}

func (x *PercentageRespVo) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PercentageRespVo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PercentageRespVo) GetPercent() float64 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *PercentageRespVo) GetPercentFormatted() string {
	if x != nil {
		return x.PercentFormatted
	}
	return ""
}

// key value pair
type PercentageVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number for something
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// unique key for distinguishing something
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PercentageVo) Reset() {
	*x = PercentageVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_gostats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PercentageVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PercentageVo) ProtoMessage() {}

func (x *PercentageVo) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_gostats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PercentageVo.ProtoReflect.Descriptor instead.
func (*PercentageVo) Descriptor() ([]byte, []int) {
	return file_transport_grpc_gostats_proto_rawDescGZIP(), []int{5}
}

func (x *PercentageVo) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PercentageVo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_transport_grpc_gostats_proto protoreflect.FileDescriptor

var file_transport_grpc_gostats_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x6f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x4e, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x39, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x1b, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x56, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x55, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x56, 0x6f, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x56, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x0c,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x32, 0xe5, 0x01, 0x0a, 0x0e, 0x47, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x13, 0x4c, 0x61, 0x72, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x70, 0x63, 0x12, 0x19,
	0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x56, 0x6f, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7a, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x70, 0x63, 0x12,
	0x2c, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x76, 0x65, 0x73, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c,
	0x76, 0x65, 0x73, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x6f, 0x2d, 0x64, 0x6f, 0x75, 0x64, 0x6f, 0x75, 0x2d, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_transport_grpc_gostats_proto_rawDescOnce sync.Once
	file_transport_grpc_gostats_proto_rawDescData = file_transport_grpc_gostats_proto_rawDesc
)

func file_transport_grpc_gostats_proto_rawDescGZIP() []byte {
	file_transport_grpc_gostats_proto_rawDescOnce.Do(func() {
		file_transport_grpc_gostats_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_grpc_gostats_proto_rawDescData)
	})
	return file_transport_grpc_gostats_proto_rawDescData
}

var file_transport_grpc_gostats_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transport_grpc_gostats_proto_goTypes = []interface{}{
	(*GetShelvesShelfBooksBookRpcRequest)(nil),  // 0: go_stats.GetShelvesShelfBooksBookRpcRequest
	(*GetShelvesShelfBooksBookRpcResponse)(nil), // 1: go_stats.GetShelvesShelfBooksBookRpcResponse
	(*LargestRemainderRpcResponse)(nil),         // 2: go_stats.LargestRemainderRpcResponse
	(*PercentageReqVo)(nil),                     // 3: go_stats.PercentageReqVo
	(*PercentageRespVo)(nil),                    // 4: go_stats.PercentageRespVo
	(*PercentageVo)(nil),                        // 5: go_stats.PercentageVo
}
var file_transport_grpc_gostats_proto_depIdxs = []int32{
	4, // 0: go_stats.LargestRemainderRpcResponse.data:type_name -> go_stats.PercentageRespVo
	5, // 1: go_stats.PercentageReqVo.data:type_name -> go_stats.PercentageVo
	3, // 2: go_stats.GoStatsService.LargestRemainderRpc:input_type -> go_stats.PercentageReqVo
	0, // 3: go_stats.GoStatsService.GetShelvesShelfBooksBookRpc:input_type -> go_stats.GetShelvesShelfBooksBookRpcRequest
	2, // 4: go_stats.GoStatsService.LargestRemainderRpc:output_type -> go_stats.LargestRemainderRpcResponse
	1, // 5: go_stats.GoStatsService.GetShelvesShelfBooksBookRpc:output_type -> go_stats.GetShelvesShelfBooksBookRpcResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transport_grpc_gostats_proto_init() }
func file_transport_grpc_gostats_proto_init() {
	if File_transport_grpc_gostats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_grpc_gostats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelvesShelfBooksBookRpcRequest); i {
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
		file_transport_grpc_gostats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelvesShelfBooksBookRpcResponse); i {
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
		file_transport_grpc_gostats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LargestRemainderRpcResponse); i {
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
		file_transport_grpc_gostats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PercentageReqVo); i {
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
		file_transport_grpc_gostats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PercentageRespVo); i {
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
		file_transport_grpc_gostats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PercentageVo); i {
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
			RawDescriptor: file_transport_grpc_gostats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_grpc_gostats_proto_goTypes,
		DependencyIndexes: file_transport_grpc_gostats_proto_depIdxs,
		MessageInfos:      file_transport_grpc_gostats_proto_msgTypes,
	}.Build()
	File_transport_grpc_gostats_proto = out.File
	file_transport_grpc_gostats_proto_rawDesc = nil
	file_transport_grpc_gostats_proto_goTypes = nil
	file_transport_grpc_gostats_proto_depIdxs = nil
}
