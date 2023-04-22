// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: protos/common/common.proto

package gopb

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

// Describes the provider that produced a particular source of data
type Provider int32

const (
	Provider_None Provider = 0 // None, implying that the provider was not included, not that we're querying data that
	// was generated by no provider
	Provider_Polygon Provider = 1 // Data generated by Polygon
)

// Enum value maps for Provider.
var (
	Provider_name = map[int32]string{
		0: "None",
		1: "Polygon",
	}
	Provider_value = map[string]int32{
		"None":    0,
		"Polygon": 1,
	}
)

func (x Provider) Enum() *Provider {
	p := new(Provider)
	*p = x
	return p
}

func (x Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_common_common_proto_enumTypes[0].Descriptor()
}

func (Provider) Type() protoreflect.EnumType {
	return &file_protos_common_common_proto_enumTypes[0]
}

func (x Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Provider.Descriptor instead.
func (Provider) EnumDescriptor() ([]byte, []int) {
	return file_protos_common_common_proto_rawDescGZIP(), []int{0}
}

var File_protos_common_common_proto protoreflect.FileDescriptor

var file_protos_common_common_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x21, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x10, 0x01, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x65, 0x66,
	0x69, 0x6e, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_common_proto_rawDescOnce sync.Once
	file_protos_common_common_proto_rawDescData = file_protos_common_common_proto_rawDesc
)

func file_protos_common_common_proto_rawDescGZIP() []byte {
	file_protos_common_common_proto_rawDescOnce.Do(func() {
		file_protos_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_common_proto_rawDescData)
	})
	return file_protos_common_common_proto_rawDescData
}

var file_protos_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_common_common_proto_goTypes = []interface{}{
	(Provider)(0), // 0: protos.common.Provider
}
var file_protos_common_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_common_common_proto_init() }
func file_protos_common_common_proto_init() {
	if File_protos_common_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_common_proto_goTypes,
		DependencyIndexes: file_protos_common_common_proto_depIdxs,
		EnumInfos:         file_protos_common_common_proto_enumTypes,
	}.Build()
	File_protos_common_common_proto = out.File
	file_protos_common_common_proto_rawDesc = nil
	file_protos_common_common_proto_goTypes = nil
	file_protos_common_common_proto_depIdxs = nil
}
