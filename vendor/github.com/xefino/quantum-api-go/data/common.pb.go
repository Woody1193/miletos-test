// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: protos/frontend/data/common.proto

package data

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

// HTTP method to use when describing a link associated with a DTO
type Method int32

const (
	Method_Post   Method = 0
	Method_Get    Method = 1
	Method_Put    Method = 2
	Method_Delete Method = 3
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "Post",
		1: "Get",
		2: "Put",
		3: "Delete",
	}
	Method_value = map[string]int32{
		"Post":   0,
		"Get":    1,
		"Put":    2,
		"Delete": 3,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_frontend_data_common_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_protos_frontend_data_common_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_protos_frontend_data_common_proto_rawDescGZIP(), []int{0}
}

// Link to return with a DTO to allow the front-end to retrieve additional
// information about the DTO
type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A human-readable name for the link
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The HTTP method to use when requesting data from the link
	Method Method `protobuf:"varint,2,opt,name=method,proto3,enum=protos.frontend.data.Method" json:"method"` 
	// The URL to call for this link
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_data_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_data_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_protos_frontend_data_common_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_Post
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_protos_frontend_data_common_proto protoreflect.FileDescriptor

var file_protos_frontend_data_common_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x04, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x2a, 0x30, 0x0a,
	0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x75,
	0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x03, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65,
	0x66, 0x69, 0x6e, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_frontend_data_common_proto_rawDescOnce sync.Once
	file_protos_frontend_data_common_proto_rawDescData = file_protos_frontend_data_common_proto_rawDesc
)

func file_protos_frontend_data_common_proto_rawDescGZIP() []byte {
	file_protos_frontend_data_common_proto_rawDescOnce.Do(func() {
		file_protos_frontend_data_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_frontend_data_common_proto_rawDescData)
	})
	return file_protos_frontend_data_common_proto_rawDescData
}

var file_protos_frontend_data_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_frontend_data_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_frontend_data_common_proto_goTypes = []interface{}{
	(Method)(0),  // 0: protos.frontend.data.Method
	(*Link)(nil), // 1: protos.frontend.data.Link
}
var file_protos_frontend_data_common_proto_depIdxs = []int32{
	0, // 0: protos.frontend.data.Link.method:type_name -> protos.frontend.data.Method
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_frontend_data_common_proto_init() }
func file_protos_frontend_data_common_proto_init() {
	if File_protos_frontend_data_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_frontend_data_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_protos_frontend_data_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_frontend_data_common_proto_goTypes,
		DependencyIndexes: file_protos_frontend_data_common_proto_depIdxs,
		EnumInfos:         file_protos_frontend_data_common_proto_enumTypes,
		MessageInfos:      file_protos_frontend_data_common_proto_msgTypes,
	}.Build()
	File_protos_frontend_data_common_proto = out.File
	file_protos_frontend_data_common_proto_rawDesc = nil
	file_protos_frontend_data_common_proto_goTypes = nil
	file_protos_frontend_data_common_proto_depIdxs = nil
}