//*
// Generated by go-doudou v2.0.1.
// Don't edit!
//
// Version No.: v20221121

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: transport/grpc/wordcloudseg.proto

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

type NestedAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedAny []*anypb.Any `protobuf:"bytes,1,rep,name=nested_any,json=nestedAny,proto3" json:"nested_any,omitempty"`
}

func (x *NestedAny) Reset() {
	*x = NestedAny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_wordcloudseg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedAny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedAny) ProtoMessage() {}

func (x *NestedAny) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_wordcloudseg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedAny.ProtoReflect.Descriptor instead.
func (*NestedAny) Descriptor() ([]byte, []int) {
	return file_transport_grpc_wordcloudseg_proto_rawDescGZIP(), []int{0}
}

func (x *NestedAny) GetNestedAny() []*anypb.Any {
	if x != nil {
		return x.NestedAny
	}
	return nil
}

type SegPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文本
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// 文本语言
	// 仅支持中文和英文
	// text language
	// only support zh and en
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *SegPayload) Reset() {
	*x = SegPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_wordcloudseg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegPayload) ProtoMessage() {}

func (x *SegPayload) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_wordcloudseg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegPayload.ProtoReflect.Descriptor instead.
func (*SegPayload) Descriptor() ([]byte, []int) {
	return file_transport_grpc_wordcloudseg_proto_rawDescGZIP(), []int{1}
}

func (x *SegPayload) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SegPayload) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type SegResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WordFreq []*NestedAny `protobuf:"bytes,1,rep,name=word_freq,json=WordFreq,proto3" json:"word_freq,omitempty"`
}

func (x *SegResult) Reset() {
	*x = SegResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_wordcloudseg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegResult) ProtoMessage() {}

func (x *SegResult) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_wordcloudseg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegResult.ProtoReflect.Descriptor instead.
func (*SegResult) Descriptor() ([]byte, []int) {
	return file_transport_grpc_wordcloudseg_proto_rawDescGZIP(), []int{2}
}

func (x *SegResult) GetWordFreq() []*NestedAny {
	if x != nil {
		return x.WordFreq
	}
	return nil
}

var File_transport_grpc_wordcloudseg_proto protoreflect.FileDescriptor

var file_transport_grpc_wordcloudseg_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x77, 0x6f, 0x72, 0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x65, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x6f, 0x72, 0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73,
	0x65, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a,
	0x09, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x79, 0x12, 0x33, 0x0a, 0x0a, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x79, 0x22,
	0x34, 0x0a, 0x0a, 0x53, 0x65, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x42, 0x0a, 0x09, 0x53, 0x65, 0x67, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x73, 0x65, 0x67, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x79, 0x52,
	0x08, 0x57, 0x6f, 0x72, 0x64, 0x46, 0x72, 0x65, 0x71, 0x32, 0x54, 0x0a, 0x13, 0x57, 0x6f, 0x72,
	0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x67, 0x52, 0x70, 0x63, 0x12, 0x19, 0x2e, 0x77, 0x6f, 0x72,
	0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x67, 0x2e, 0x53, 0x65, 0x67, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x73, 0x65, 0x67, 0x2e, 0x53, 0x65, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e,
	0x69, 0x6f, 0x6e, 0x6a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x6f,
	0x75, 0x64, 0x6f, 0x75, 0x2d, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x77,
	0x6f, 0x72, 0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x73, 0x65, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_grpc_wordcloudseg_proto_rawDescOnce sync.Once
	file_transport_grpc_wordcloudseg_proto_rawDescData = file_transport_grpc_wordcloudseg_proto_rawDesc
)

func file_transport_grpc_wordcloudseg_proto_rawDescGZIP() []byte {
	file_transport_grpc_wordcloudseg_proto_rawDescOnce.Do(func() {
		file_transport_grpc_wordcloudseg_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_grpc_wordcloudseg_proto_rawDescData)
	})
	return file_transport_grpc_wordcloudseg_proto_rawDescData
}

var file_transport_grpc_wordcloudseg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transport_grpc_wordcloudseg_proto_goTypes = []interface{}{
	(*NestedAny)(nil),  // 0: wordcloud_seg.NestedAny
	(*SegPayload)(nil), // 1: wordcloud_seg.SegPayload
	(*SegResult)(nil),  // 2: wordcloud_seg.SegResult
	(*anypb.Any)(nil),  // 3: google.protobuf.Any
}
var file_transport_grpc_wordcloudseg_proto_depIdxs = []int32{
	3, // 0: wordcloud_seg.NestedAny.nested_any:type_name -> google.protobuf.Any
	0, // 1: wordcloud_seg.SegResult.word_freq:type_name -> wordcloud_seg.NestedAny
	1, // 2: wordcloud_seg.WordcloudSegService.SegRpc:input_type -> wordcloud_seg.SegPayload
	2, // 3: wordcloud_seg.WordcloudSegService.SegRpc:output_type -> wordcloud_seg.SegResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transport_grpc_wordcloudseg_proto_init() }
func file_transport_grpc_wordcloudseg_proto_init() {
	if File_transport_grpc_wordcloudseg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_grpc_wordcloudseg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedAny); i {
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
		file_transport_grpc_wordcloudseg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegPayload); i {
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
		file_transport_grpc_wordcloudseg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegResult); i {
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
			RawDescriptor: file_transport_grpc_wordcloudseg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_grpc_wordcloudseg_proto_goTypes,
		DependencyIndexes: file_transport_grpc_wordcloudseg_proto_depIdxs,
		MessageInfos:      file_transport_grpc_wordcloudseg_proto_msgTypes,
	}.Build()
	File_transport_grpc_wordcloudseg_proto = out.File
	file_transport_grpc_wordcloudseg_proto_rawDesc = nil
	file_transport_grpc_wordcloudseg_proto_goTypes = nil
	file_transport_grpc_wordcloudseg_proto_depIdxs = nil
}