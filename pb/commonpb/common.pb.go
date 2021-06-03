// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: pb/commonpb/common.proto

package commonpb

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

type ResolveParamType int32

const (
	ResolveParamType_RESOLVE_PARAM_TYPE_INT    ResolveParamType = 0
	ResolveParamType_RESOLVE_PARAM_TYPE_STRING ResolveParamType = 1
	ResolveParamType_RESOLVE_PARAM_TYPE_FLOAT  ResolveParamType = 2
)

// Enum value maps for ResolveParamType.
var (
	ResolveParamType_name = map[int32]string{
		0: "RESOLVE_PARAM_TYPE_INT",
		1: "RESOLVE_PARAM_TYPE_STRING",
		2: "RESOLVE_PARAM_TYPE_FLOAT",
	}
	ResolveParamType_value = map[string]int32{
		"RESOLVE_PARAM_TYPE_INT":    0,
		"RESOLVE_PARAM_TYPE_STRING": 1,
		"RESOLVE_PARAM_TYPE_FLOAT":  2,
	}
)

func (x ResolveParamType) Enum() *ResolveParamType {
	p := new(ResolveParamType)
	*p = x
	return p
}

func (x ResolveParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResolveParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_commonpb_common_proto_enumTypes[0].Descriptor()
}

func (ResolveParamType) Type() protoreflect.EnumType {
	return &file_pb_commonpb_common_proto_enumTypes[0]
}

func (x ResolveParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResolveParamType.Descriptor instead.
func (ResolveParamType) EnumDescriptor() ([]byte, []int) {
	return file_pb_commonpb_common_proto_rawDescGZIP(), []int{0}
}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseCode int32     `protobuf:"varint,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	Headers      []*Header `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_commonpb_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_commonpb_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_pb_commonpb_common_proto_rawDescGZIP(), []int{0}
}

func (x *ClientResponse) GetResponseCode() int32 {
	if x != nil {
		return x.ResponseCode
	}
	return 0
}

func (x *ClientResponse) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_commonpb_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_pb_commonpb_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_pb_commonpb_common_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ResolveParamValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int   int64   `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
	Text  string  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Float float64 `protobuf:"fixed64,3,opt,name=float,proto3" json:"float,omitempty"`
}

func (x *ResolveParamValue) Reset() {
	*x = ResolveParamValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_commonpb_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveParamValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveParamValue) ProtoMessage() {}

func (x *ResolveParamValue) ProtoReflect() protoreflect.Message {
	mi := &file_pb_commonpb_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveParamValue.ProtoReflect.Descriptor instead.
func (*ResolveParamValue) Descriptor() ([]byte, []int) {
	return file_pb_commonpb_common_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveParamValue) GetInt() int64 {
	if x != nil {
		return x.Int
	}
	return 0
}

func (x *ResolveParamValue) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ResolveParamValue) GetFloat() float64 {
	if x != nil {
		return x.Float
	}
	return 0
}

type ResolveParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ResolveParamType  ResolveParamType   `protobuf:"varint,2,opt,name=resolve_param_type,json=resolveParamType,proto3,enum=common.ResolveParamType" json:"resolve_param_type,omitempty"`
	ResolveParamValue *ResolveParamValue `protobuf:"bytes,3,opt,name=resolve_param_value,json=resolveParamValue,proto3" json:"resolve_param_value,omitempty"`
}

func (x *ResolveParam) Reset() {
	*x = ResolveParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_commonpb_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveParam) ProtoMessage() {}

func (x *ResolveParam) ProtoReflect() protoreflect.Message {
	mi := &file_pb_commonpb_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveParam.ProtoReflect.Descriptor instead.
func (*ResolveParam) Descriptor() ([]byte, []int) {
	return file_pb_commonpb_common_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResolveParam) GetResolveParamType() ResolveParamType {
	if x != nil {
		return x.ResolveParamType
	}
	return ResolveParamType_RESOLVE_PARAM_TYPE_INT
}

func (x *ResolveParam) GetResolveParamValue() *ResolveParamValue {
	if x != nil {
		return x.ResolveParamValue
	}
	return nil
}

var File_pb_commonpb_common_proto protoreflect.FileDescriptor

var file_pb_commonpb_common_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0x5f, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4f, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x6b,
	0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1c, 0x0a,
	0x18, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x69, 0x74, 0x61,
	0x35, 0x36, 0x33, 0x37, 0x2f, 0x71, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_commonpb_common_proto_rawDescOnce sync.Once
	file_pb_commonpb_common_proto_rawDescData = file_pb_commonpb_common_proto_rawDesc
)

func file_pb_commonpb_common_proto_rawDescGZIP() []byte {
	file_pb_commonpb_common_proto_rawDescOnce.Do(func() {
		file_pb_commonpb_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_commonpb_common_proto_rawDescData)
	})
	return file_pb_commonpb_common_proto_rawDescData
}

var file_pb_commonpb_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_commonpb_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_commonpb_common_proto_goTypes = []interface{}{
	(ResolveParamType)(0),     // 0: common.ResolveParamType
	(*ClientResponse)(nil),    // 1: common.ClientResponse
	(*Header)(nil),            // 2: common.Header
	(*ResolveParamValue)(nil), // 3: common.ResolveParamValue
	(*ResolveParam)(nil),      // 4: common.ResolveParam
}
var file_pb_commonpb_common_proto_depIdxs = []int32{
	2, // 0: common.ClientResponse.headers:type_name -> common.Header
	0, // 1: common.ResolveParam.resolve_param_type:type_name -> common.ResolveParamType
	3, // 2: common.ResolveParam.resolve_param_value:type_name -> common.ResolveParamValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_commonpb_common_proto_init() }
func file_pb_commonpb_common_proto_init() {
	if File_pb_commonpb_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_commonpb_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_pb_commonpb_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_pb_commonpb_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveParamValue); i {
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
		file_pb_commonpb_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveParam); i {
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
			RawDescriptor: file_pb_commonpb_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_commonpb_common_proto_goTypes,
		DependencyIndexes: file_pb_commonpb_common_proto_depIdxs,
		EnumInfos:         file_pb_commonpb_common_proto_enumTypes,
		MessageInfos:      file_pb_commonpb_common_proto_msgTypes,
	}.Build()
	File_pb_commonpb_common_proto = out.File
	file_pb_commonpb_common_proto_rawDesc = nil
	file_pb_commonpb_common_proto_goTypes = nil
	file_pb_commonpb_common_proto_depIdxs = nil
}
