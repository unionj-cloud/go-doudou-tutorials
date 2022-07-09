// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: helloworld_openapi3.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GreetingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetingReq) Reset() {
	*x = GreetingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingReq) ProtoMessage() {}

func (x *GreetingReq) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingReq.ProtoReflect.Descriptor instead.
func (*GreetingReq) Descriptor() ([]byte, []int) {
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingReq) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type GreetingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GreetingResp) Reset() {
	*x = GreetingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResp) ProtoMessage() {}

func (x *GreetingResp) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResp.ProtoReflect.Descriptor instead.
func (*GreetingResp) Descriptor() ([]byte, []int) {
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
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
		mi := &file_helloworld_openapi3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[2]
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
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{2}
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

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	PageNo int32    `protobuf:"varint,2,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	Size   int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[3]
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
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{3}
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

	Dept int32  `protobuf:"varint,1,opt,name=dept,proto3" json:"dept,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PageFilter) Reset() {
	*x = PageFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageFilter) ProtoMessage() {}

func (x *PageFilter) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[4]
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
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{4}
}

func (x *PageFilter) GetDept() int32 {
	if x != nil {
		return x.Dept
	}
	return 0
}

func (x *PageFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

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
		mi := &file_helloworld_openapi3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageQuery) ProtoMessage() {}

func (x *PageQuery) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[5]
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
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{5}
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

type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{6}
}

type PageRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasNext  bool   `protobuf:"varint,1,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
	Items    *Items `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
	PageNo   int32  `protobuf:"varint,3,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	PageSize int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total    int32  `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PageRet) Reset() {
	*x = PageRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRet) ProtoMessage() {}

func (x *PageRet) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[7]
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
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{7}
}

func (x *PageRet) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

func (x *PageRet) GetItems() *Items {
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

//*
//这当然也是注释
type UserVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dept string `protobuf:"bytes,1,opt,name=dept,proto3" json:"dept,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// 这是一个注释
	// 这也是一个注释
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // 这是一个注释
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserVo) Reset() {
	*x = UserVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVo) ProtoMessage() {}

func (x *UserVo) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[8]
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
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{8}
}

func (x *UserVo) GetDept() string {
	if x != nil {
		return x.Dept
	}
	return ""
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

//POST_greetingParameters holds parameters to POST_greeting
type POSTGreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationXWwwFormUrlencoded *GreetingReq `protobuf:"bytes,1,opt,name=application_x_www_form_urlencoded,json=applicationXWwwFormUrlencoded,proto3" json:"application_x_www_form_urlencoded,omitempty"`
}

func (x *POSTGreetingRequest) Reset() {
	*x = POSTGreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_openapi3_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *POSTGreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*POSTGreetingRequest) ProtoMessage() {}

func (x *POSTGreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_openapi3_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use POSTGreetingRequest.ProtoReflect.Descriptor instead.
func (*POSTGreetingRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_openapi3_proto_rawDescGZIP(), []int{9}
}

func (x *POSTGreetingRequest) GetApplicationXWwwFormUrlencoded() *GreetingReq {
	if x != nil {
		return x.ApplicationXWwwFormUrlencoded
	}
	return nil
}

var File_helloworld_openapi3_proto protoreflect.FileDescriptor

var file_helloworld_openapi3_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x67, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x32,
	0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x33, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x34, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x65, 0x70,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x56, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x22, 0x81, 0x01, 0x0a, 0x13, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6a, 0x0a, 0x21, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x78, 0x5f, 0x77, 0x77, 0x77, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x52, 0x1d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x58, 0x57, 0x77, 0x77, 0x46, 0x6f, 0x72, 0x6d, 0x55, 0x72, 0x6c, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x64, 0x32, 0x72, 0x0a, 0x13, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33, 0x12, 0x5b, 0x0a, 0x0c, 0x50,
	0x4f, 0x53, 0x54, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x33, 0x2e, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x33, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x6a, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_openapi3_proto_rawDescOnce sync.Once
	file_helloworld_openapi3_proto_rawDescData = file_helloworld_openapi3_proto_rawDesc
)

func file_helloworld_openapi3_proto_rawDescGZIP() []byte {
	file_helloworld_openapi3_proto_rawDescOnce.Do(func() {
		file_helloworld_openapi3_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_openapi3_proto_rawDescData)
	})
	return file_helloworld_openapi3_proto_rawDescData
}

var file_helloworld_openapi3_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_helloworld_openapi3_proto_goTypes = []interface{}{
	(*GreetingReq)(nil),         // 0: helloworld_openapi3.GreetingReq
	(*GreetingResp)(nil),        // 1: helloworld_openapi3.GreetingResp
	(*Order)(nil),               // 2: helloworld_openapi3.Order
	(*Page)(nil),                // 3: helloworld_openapi3.Page
	(*PageFilter)(nil),          // 4: helloworld_openapi3.PageFilter
	(*PageQuery)(nil),           // 5: helloworld_openapi3.PageQuery
	(*Items)(nil),               // 6: helloworld_openapi3.Items
	(*PageRet)(nil),             // 7: helloworld_openapi3.PageRet
	(*UserVo)(nil),              // 8: helloworld_openapi3.UserVo
	(*POSTGreetingRequest)(nil), // 9: helloworld_openapi3.POSTGreetingRequest
}
var file_helloworld_openapi3_proto_depIdxs = []int32{
	2, // 0: helloworld_openapi3.Page.orders:type_name -> helloworld_openapi3.Order
	4, // 1: helloworld_openapi3.PageQuery.filter:type_name -> helloworld_openapi3.PageFilter
	3, // 2: helloworld_openapi3.PageQuery.page:type_name -> helloworld_openapi3.Page
	6, // 3: helloworld_openapi3.PageRet.items:type_name -> helloworld_openapi3.Items
	0, // 4: helloworld_openapi3.POSTGreetingRequest.application_x_www_form_urlencoded:type_name -> helloworld_openapi3.GreetingReq
	9, // 5: helloworld_openapi3.Helloworld_openapi3.POSTGreeting:input_type -> helloworld_openapi3.POSTGreetingRequest
	1, // 6: helloworld_openapi3.Helloworld_openapi3.POSTGreeting:output_type -> helloworld_openapi3.GreetingResp
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_helloworld_openapi3_proto_init() }
func file_helloworld_openapi3_proto_init() {
	if File_helloworld_openapi3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_openapi3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingReq); i {
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
		file_helloworld_openapi3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingResp); i {
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
		file_helloworld_openapi3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_openapi3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_openapi3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_openapi3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_openapi3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
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
		file_helloworld_openapi3_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_openapi3_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_openapi3_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*POSTGreetingRequest); i {
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
			RawDescriptor: file_helloworld_openapi3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_openapi3_proto_goTypes,
		DependencyIndexes: file_helloworld_openapi3_proto_depIdxs,
		MessageInfos:      file_helloworld_openapi3_proto_msgTypes,
	}.Build()
	File_helloworld_openapi3_proto = out.File
	file_helloworld_openapi3_proto_rawDesc = nil
	file_helloworld_openapi3_proto_goTypes = nil
	file_helloworld_openapi3_proto_depIdxs = nil
}
