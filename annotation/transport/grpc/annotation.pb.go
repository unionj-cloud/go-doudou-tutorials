//*
// Generated by go-doudou v2.0.0.
// Don't edit!
//
// Version No.: v20221024

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: transport/grpc/annotation.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoleEnum int32

const (
	RoleEnum_GUEST RoleEnum = 0
	RoleEnum_USER  RoleEnum = 1
	RoleEnum_ADMIN RoleEnum = 2
)

// Enum value maps for RoleEnum.
var (
	RoleEnum_name = map[int32]string{
		0: "GUEST",
		1: "USER",
		2: "ADMIN",
	}
	RoleEnum_value = map[string]int32{
		"GUEST": 0,
		"USER":  1,
		"ADMIN": 2,
	}
)

func (x RoleEnum) Enum() *RoleEnum {
	p := new(RoleEnum)
	*p = x
	return p
}

func (x RoleEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_grpc_annotation_proto_enumTypes[0].Descriptor()
}

func (RoleEnum) Type() protoreflect.EnumType {
	return &file_transport_grpc_annotation_proto_enumTypes[0]
}

func (x RoleEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleEnum.Descriptor instead.
func (RoleEnum) EnumDescriptor() ([]byte, []int) {
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{0}
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,json=User,proto3" json:"user,omitempty"`
	Pass string `protobuf:"bytes,2,opt,name=pass,json=Pass,proto3" json:"pass,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Auth) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type GetAdminRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAdminRpcResponse) Reset() {
	*x = GetAdminRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRpcResponse) ProtoMessage() {}

func (x *GetAdminRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRpcResponse.ProtoReflect.Descriptor instead.
func (*GetAdminRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{1}
}

func (x *GetAdminRpcResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetGuestRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetGuestRpcResponse) Reset() {
	*x = GetGuestRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuestRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuestRpcResponse) ProtoMessage() {}

func (x *GetGuestRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuestRpcResponse.ProtoReflect.Descriptor instead.
func (*GetGuestRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{2}
}

func (x *GetGuestRpcResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetUserRpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserRpcResponse) Reset() {
	*x = GetUserRpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRpcResponse) ProtoMessage() {}

func (x *GetUserRpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRpcResponse.ProtoReflect.Descriptor instead.
func (*GetUserRpcResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRpcResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Col  string `protobuf:"bytes,1,opt,name=col,json=Col,proto3" json:"col,omitempty"`
	Sort string `protobuf:"bytes,2,opt,name=sort,json=Sort,proto3" json:"sort,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[4]
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
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{4}
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
	Orders []*Order `protobuf:"bytes,1,rep,name=orders,json=Orders,proto3" json:"orders,omitempty"`
	// 页码
	PageNo int32 `protobuf:"varint,2,opt,name=page_no,json=PageNo,proto3" json:"page_no,omitempty"`
	// 每页行数
	Size int32 `protobuf:"varint,3,opt,name=size,json=Size,proto3" json:"size,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[5]
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
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{5}
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
	Name string `protobuf:"bytes,1,opt,name=name,json=Name,proto3" json:"name,omitempty"`
	// 所属部门ID
	Dept int32 `protobuf:"varint,2,opt,name=dept,json=Dept,proto3" json:"dept,omitempty"`
}

func (x *PageFilter) Reset() {
	*x = PageFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageFilter) ProtoMessage() {}

func (x *PageFilter) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[6]
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
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{6}
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

	Filter *PageFilter `protobuf:"bytes,1,opt,name=filter,json=Filter,proto3" json:"filter,omitempty"`
	Page   *Page       `protobuf:"bytes,2,opt,name=page,json=Page,proto3" json:"page,omitempty"`
}

func (x *PageQuery) Reset() {
	*x = PageQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageQuery) ProtoMessage() {}

func (x *PageQuery) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[7]
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
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{7}
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

	Items    *anypb.Any `protobuf:"bytes,1,opt,name=items,json=Items,proto3" json:"items,omitempty"`
	PageNo   int32      `protobuf:"varint,2,opt,name=page_no,json=PageNo,proto3" json:"page_no,omitempty"`
	PageSize int32      `protobuf:"varint,3,opt,name=page_size,json=PageSize,proto3" json:"page_size,omitempty"`
	Total    int32      `protobuf:"varint,4,opt,name=total,json=Total,proto3" json:"total,omitempty"`
	HasNext  bool       `protobuf:"varint,5,opt,name=has_next,json=HasNext,proto3" json:"has_next,omitempty"`
}

func (x *PageRet) Reset() {
	*x = PageRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRet) ProtoMessage() {}

func (x *PageRet) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[8]
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
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{8}
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

	Id    int32  `protobuf:"varint,1,opt,name=id,json=Id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,json=Name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,json=Phone,proto3" json:"phone,omitempty"`
	Dept  string `protobuf:"bytes,4,opt,name=dept,json=Dept,proto3" json:"dept,omitempty"`
}

func (x *UserVo) Reset() {
	*x = UserVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_annotation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVo) ProtoMessage() {}

func (x *UserVo) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_annotation_proto_msgTypes[9]
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
	return file_transport_grpc_annotation_proto_rawDescGZIP(), []int{9}
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

var File_transport_grpc_annotation_proto protoreflect.FileDescriptor

var file_transport_grpc_annotation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x50, 0x61, 0x73, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x70, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x6f, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x53, 0x6f, 0x72, 0x74, 0x22, 0x5e, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x44, 0x65, 0x70, 0x74, 0x22, 0x61, 0x0a, 0x09, 0x50, 0x61,
	0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0x9c, 0x01,
	0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x48, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x22, 0x56, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x65, 0x70, 0x74, 0x2a, 0x2a, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55,
	0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02,
	0x32, 0xe9, 0x01, 0x0a, 0x11, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_transport_grpc_annotation_proto_rawDescOnce sync.Once
	file_transport_grpc_annotation_proto_rawDescData = file_transport_grpc_annotation_proto_rawDesc
)

func file_transport_grpc_annotation_proto_rawDescGZIP() []byte {
	file_transport_grpc_annotation_proto_rawDescOnce.Do(func() {
		file_transport_grpc_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_grpc_annotation_proto_rawDescData)
	})
	return file_transport_grpc_annotation_proto_rawDescData
}

var file_transport_grpc_annotation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transport_grpc_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_transport_grpc_annotation_proto_goTypes = []interface{}{
	(RoleEnum)(0),               // 0: annotation.RoleEnum
	(*Auth)(nil),                // 1: annotation.Auth
	(*GetAdminRpcResponse)(nil), // 2: annotation.GetAdminRpcResponse
	(*GetGuestRpcResponse)(nil), // 3: annotation.GetGuestRpcResponse
	(*GetUserRpcResponse)(nil),  // 4: annotation.GetUserRpcResponse
	(*Order)(nil),               // 5: annotation.Order
	(*Page)(nil),                // 6: annotation.Page
	(*PageFilter)(nil),          // 7: annotation.PageFilter
	(*PageQuery)(nil),           // 8: annotation.PageQuery
	(*PageRet)(nil),             // 9: annotation.PageRet
	(*UserVo)(nil),              // 10: annotation.UserVo
	(*anypb.Any)(nil),           // 11: google.protobuf.Any
	(*emptypb.Empty)(nil),       // 12: google.protobuf.Empty
}
var file_transport_grpc_annotation_proto_depIdxs = []int32{
	5,  // 0: annotation.Page.orders:type_name -> annotation.Order
	7,  // 1: annotation.PageQuery.filter:type_name -> annotation.PageFilter
	6,  // 2: annotation.PageQuery.page:type_name -> annotation.Page
	11, // 3: annotation.PageRet.items:type_name -> google.protobuf.Any
	12, // 4: annotation.AnnotationService.GetGuestRpc:input_type -> google.protobuf.Empty
	12, // 5: annotation.AnnotationService.GetUserRpc:input_type -> google.protobuf.Empty
	12, // 6: annotation.AnnotationService.GetAdminRpc:input_type -> google.protobuf.Empty
	3,  // 7: annotation.AnnotationService.GetGuestRpc:output_type -> annotation.GetGuestRpcResponse
	4,  // 8: annotation.AnnotationService.GetUserRpc:output_type -> annotation.GetUserRpcResponse
	2,  // 9: annotation.AnnotationService.GetAdminRpc:output_type -> annotation.GetAdminRpcResponse
	7,  // [7:10] is the sub-list for method output_type
	4,  // [4:7] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_transport_grpc_annotation_proto_init() }
func file_transport_grpc_annotation_proto_init() {
	if File_transport_grpc_annotation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_grpc_annotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_transport_grpc_annotation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminRpcResponse); i {
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
		file_transport_grpc_annotation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuestRpcResponse); i {
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
		file_transport_grpc_annotation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRpcResponse); i {
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
		file_transport_grpc_annotation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_grpc_annotation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_grpc_annotation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_grpc_annotation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_grpc_annotation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_grpc_annotation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_transport_grpc_annotation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_grpc_annotation_proto_goTypes,
		DependencyIndexes: file_transport_grpc_annotation_proto_depIdxs,
		EnumInfos:         file_transport_grpc_annotation_proto_enumTypes,
		MessageInfos:      file_transport_grpc_annotation_proto_msgTypes,
	}.Build()
	File_transport_grpc_annotation_proto = out.File
	file_transport_grpc_annotation_proto_rawDesc = nil
	file_transport_grpc_annotation_proto_goTypes = nil
	file_transport_grpc_annotation_proto_depIdxs = nil
}
