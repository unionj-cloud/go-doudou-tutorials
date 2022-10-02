//*
// Generated by go-doudou v1.2.4.
// Don't edit!
//
// Version No.: v20221002

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: transport/grpc/helloworld.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyboardLayout int32

const (
	KeyboardLayout_UNKNOWN KeyboardLayout = 0
	KeyboardLayout_QWERTZ  KeyboardLayout = 1
	KeyboardLayout_AZERTY  KeyboardLayout = 2
	KeyboardLayout_QWERTY  KeyboardLayout = 3
)

// Enum value maps for KeyboardLayout.
var (
	KeyboardLayout_name = map[int32]string{
		0: "UNKNOWN",
		1: "QWERTZ",
		2: "AZERTY",
		3: "QWERTY",
	}
	KeyboardLayout_value = map[string]int32{
		"UNKNOWN": 0,
		"QWERTZ":  1,
		"AZERTY":  2,
		"QWERTY":  3,
	}
)

func (x KeyboardLayout) Enum() *KeyboardLayout {
	p := new(KeyboardLayout)
	*p = x
	return p
}

func (x KeyboardLayout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyboardLayout) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_grpc_helloworld_proto_enumTypes[0].Descriptor()
}

func (KeyboardLayout) Type() protoreflect.EnumType {
	return &file_transport_grpc_helloworld_proto_enumTypes[0]
}

func (x KeyboardLayout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyboardLayout.Descriptor instead.
func (KeyboardLayout) EnumDescriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{0}
}

type ByeRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 入参
	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *ByeRpcRequest) Reset() {
	*x = ByeRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeRpcRequest) ProtoMessage() {}

func (x *ByeRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeRpcRequest.ProtoReflect.Descriptor instead.
func (*ByeRpcRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *ByeRpcRequest) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type ByeRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 出参
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ByeRpcResponse) Reset() {
	*x = ByeRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeRpcResponse) ProtoMessage() {}

func (x *ByeRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeRpcResponse.ProtoReflect.Descriptor instead.
func (*ByeRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *ByeRpcResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GreetingRpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 入参
	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetingRpcRequest) Reset() {
	*x = GreetingRpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingRpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRpcRequest) ProtoMessage() {}

func (x *GreetingRpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRpcRequest.ProtoReflect.Descriptor instead.
func (*GreetingRpcRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *GreetingRpcRequest) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type GreetingRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 出参
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GreetingRpcResponse) Reset() {
	*x = GreetingRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRpcResponse) ProtoMessage() {}

func (x *GreetingRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRpcResponse.ProtoReflect.Descriptor instead.
func (*GreetingRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *GreetingRpcResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Keyboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layout  KeyboardLayout `protobuf:"varint,1,opt,name=layout,proto3,enum=helloworld.KeyboardLayout" json:"layout,omitempty"`
	Backlit bool           `protobuf:"varint,2,opt,name=backlit,proto3" json:"backlit,omitempty"`
}

func (x *Keyboard) Reset() {
	*x = Keyboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyboard) ProtoMessage() {}

func (x *Keyboard) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyboard.ProtoReflect.Descriptor instead.
func (*Keyboard) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{4}
}

func (x *Keyboard) GetLayout() KeyboardLayout {
	if x != nil {
		return x.Layout
	}
	return KeyboardLayout_UNKNOWN
}

func (x *Keyboard) GetBacklit() bool {
	if x != nil {
		return x.Backlit
	}
	return false
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Col  string `protobuf:"bytes,1,opt,name=col,proto3" json:"col,omitempty"`
	Sort string `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{5}
}

func (x *Order) GetCol() string {
	if x != nil {
		return x.Col
	}
	return ""
}

func (x *Order) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 排序规则
	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	// 页码
	PageNo int32 `protobuf:"varint,2,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	// 每页行数
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{6}
}

func (x *Page) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *Page) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *Page) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PageFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 真实姓名，前缀匹配
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 所属部门ID
	Dept int32 `protobuf:"varint,2,opt,name=dept,proto3" json:"dept,omitempty"`
}

func (x *PageFilter) Reset() {
	*x = PageFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageFilter) ProtoMessage() {}

func (x *PageFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageFilter.ProtoReflect.Descriptor instead.
func (*PageFilter) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{7}
}

func (x *PageFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PageFilter) GetDept() int32 {
	if x != nil {
		return x.Dept
	}
	return 0
}

// 分页筛选条件
type PageQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *PageFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Page   *Page       `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageQuery) Reset() {
	*x = PageQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageQuery) ProtoMessage() {}

func (x *PageQuery) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageQuery.ProtoReflect.Descriptor instead.
func (*PageQuery) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{8}
}

func (x *PageQuery) GetFilter() *PageFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *PageQuery) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type PageRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items    *anypb.Any `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	PageNo   int32      `protobuf:"varint,2,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	PageSize int32      `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total    int32      `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	HasNext  bool       `protobuf:"varint,5,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
}

func (x *PageRet) Reset() {
	*x = PageRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRet) ProtoMessage() {}

func (x *PageRet) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRet.ProtoReflect.Descriptor instead.
func (*PageRet) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{9}
}

func (x *PageRet) GetItems() *anypb.Any {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PageRet) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *PageRet) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageRet) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageRet) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

type UserVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Dept  string `protobuf:"bytes,4,opt,name=dept,proto3" json:"dept,omitempty"`
}

func (x *UserVo) Reset() {
	*x = UserVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_helloworld_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVo) ProtoMessage() {}

func (x *UserVo) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_helloworld_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVo.ProtoReflect.Descriptor instead.
func (*UserVo) Descriptor() ([]byte, []int) {
	return file_transport_grpc_helloworld_proto_rawDescGZIP(), []int{10}
}

func (x *UserVo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserVo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserVo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserVo) GetDept() string {
	if x != nil {
		return x.Dept
	}
	return ""
}

var File_transport_grpc_helloworld_proto protoreflect.FileDescriptor

var file_transport_grpc_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x42, 0x79, 0x65, 0x52,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x24, 0x0a, 0x0e, 0x42, 0x79, 0x65, 0x52, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x12, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x29, 0x0a,
	0x13, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x52, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x6c,
	0x69, 0x74, 0x22, 0x2d, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x22, 0x5e, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x22, 0x61, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x22, 0x56, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x56, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x70,
	0x74, 0x2a, 0x41, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x51, 0x57, 0x45, 0x52, 0x54, 0x5a, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x5a, 0x45, 0x52, 0x54, 0x59, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x57, 0x45, 0x52,
	0x54, 0x59, 0x10, 0x03, 0x32, 0xcc, 0x02, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x70, 0x63, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x70, 0x63, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x42, 0x79, 0x65, 0x52, 0x70, 0x63,
	0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x79,
	0x65, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x79, 0x65, 0x52, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x42, 0x69, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x70, 0x63, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x38, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x70, 0x63, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x28, 0x01, 0x12, 0x38, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x70, 0x63, 0x12, 0x11, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x10, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x6a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_grpc_helloworld_proto_rawDescOnce sync.Once
	file_transport_grpc_helloworld_proto_rawDescData = file_transport_grpc_helloworld_proto_rawDesc
)

func file_transport_grpc_helloworld_proto_rawDescGZIP() []byte {
	file_transport_grpc_helloworld_proto_rawDescOnce.Do(func() {
		file_transport_grpc_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_grpc_helloworld_proto_rawDescData)
	})
	return file_transport_grpc_helloworld_proto_rawDescData
}

var file_transport_grpc_helloworld_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transport_grpc_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_transport_grpc_helloworld_proto_goTypes = []interface{}{
	(KeyboardLayout)(0),         // 0: helloworld.KeyboardLayout
	(*ByeRpcRequest)(nil),       // 1: helloworld.ByeRpcRequest
	(*ByeRpcResponse)(nil),      // 2: helloworld.ByeRpcResponse
	(*GreetingRpcRequest)(nil),  // 3: helloworld.GreetingRpcRequest
	(*GreetingRpcResponse)(nil), // 4: helloworld.GreetingRpcResponse
	(*Keyboard)(nil),            // 5: helloworld.Keyboard
	(*Order)(nil),               // 6: helloworld.Order
	(*Page)(nil),                // 7: helloworld.Page
	(*PageFilter)(nil),          // 8: helloworld.PageFilter
	(*PageQuery)(nil),           // 9: helloworld.PageQuery
	(*PageRet)(nil),             // 10: helloworld.PageRet
	(*UserVo)(nil),              // 11: helloworld.UserVo
	(*anypb.Any)(nil),           // 12: google.protobuf.Any
}
var file_transport_grpc_helloworld_proto_depIdxs = []int32{
	0,  // 0: helloworld.Keyboard.layout:type_name -> helloworld.KeyboardLayout
	6,  // 1: helloworld.Page.orders:type_name -> helloworld.Order
	8,  // 2: helloworld.PageQuery.filter:type_name -> helloworld.PageFilter
	7,  // 3: helloworld.PageQuery.page:type_name -> helloworld.Page
	12, // 4: helloworld.PageRet.items:type_name -> google.protobuf.Any
	3,  // 5: helloworld.HelloworldRpc.GreetingRpc:input_type -> helloworld.GreetingRpcRequest
	1,  // 6: helloworld.HelloworldRpc.ByeRpc:input_type -> helloworld.ByeRpcRequest
	6,  // 7: helloworld.HelloworldRpc.BiStreamRpc:input_type -> helloworld.Order
	6,  // 8: helloworld.HelloworldRpc.ClientStreamRpc:input_type -> helloworld.Order
	6,  // 9: helloworld.HelloworldRpc.ServerStreamRpc:input_type -> helloworld.Order
	4,  // 10: helloworld.HelloworldRpc.GreetingRpc:output_type -> helloworld.GreetingRpcResponse
	2,  // 11: helloworld.HelloworldRpc.ByeRpc:output_type -> helloworld.ByeRpcResponse
	7,  // 12: helloworld.HelloworldRpc.BiStreamRpc:output_type -> helloworld.Page
	7,  // 13: helloworld.HelloworldRpc.ClientStreamRpc:output_type -> helloworld.Page
	7,  // 14: helloworld.HelloworldRpc.ServerStreamRpc:output_type -> helloworld.Page
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_transport_grpc_helloworld_proto_init() }
func file_transport_grpc_helloworld_proto_init() {
	if File_transport_grpc_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_grpc_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeRpcRequest); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeRpcResponse); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingRpcRequest); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingRpcResponse); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyboard); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageFilter); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageQuery); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRet); i {
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
		file_transport_grpc_helloworld_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVo); i {
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
			RawDescriptor: file_transport_grpc_helloworld_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_grpc_helloworld_proto_goTypes,
		DependencyIndexes: file_transport_grpc_helloworld_proto_depIdxs,
		EnumInfos:         file_transport_grpc_helloworld_proto_enumTypes,
		MessageInfos:      file_transport_grpc_helloworld_proto_msgTypes,
	}.Build()
	File_transport_grpc_helloworld_proto = out.File
	file_transport_grpc_helloworld_proto_rawDesc = nil
	file_transport_grpc_helloworld_proto_goTypes = nil
	file_transport_grpc_helloworld_proto_depIdxs = nil
}
