// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: resolver/v1/annotations.proto

package resolverv1

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var file_resolver_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*SchemaMetadata)(nil),
		Field:         59901,
		Name:          "clutch.resolver.v1.schema",
		Tag:           "bytes,59901,opt,name=schema",
		Filename:      "resolver/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldMetadata)(nil),
		Field:         59901,
		Name:          "clutch.resolver.v1.schema_field",
		Tag:           "bytes,59901,opt,name=schema_field",
		Filename:      "resolver/v1/annotations.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// Use a random high number that won't conflict with annotations from other
	// libraries.
	//
	// optional clutch.resolver.v1.SchemaMetadata schema = 59901;
	E_Schema = &file_resolver_v1_annotations_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// Use a random high number that won't conflict with annotations from other
	// libraries.
	//
	// optional clutch.resolver.v1.FieldMetadata schema_field = 59901;
	E_SchemaField = &file_resolver_v1_annotations_proto_extTypes[1]
)

var File_resolver_v1_annotations_proto protoreflect.FileDescriptor

var file_resolver_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x5d, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0xd3, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x65,
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0xd3,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_resolver_v1_annotations_proto_goTypes = []interface{}{
	(*descriptor.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
	(*SchemaMetadata)(nil),            // 2: clutch.resolver.v1.SchemaMetadata
	(*FieldMetadata)(nil),             // 3: clutch.resolver.v1.FieldMetadata
}
var file_resolver_v1_annotations_proto_depIdxs = []int32{
	0, // 0: clutch.resolver.v1.schema:extendee -> google.protobuf.MessageOptions
	1, // 1: clutch.resolver.v1.schema_field:extendee -> google.protobuf.FieldOptions
	2, // 2: clutch.resolver.v1.schema:type_name -> clutch.resolver.v1.SchemaMetadata
	3, // 3: clutch.resolver.v1.schema_field:type_name -> clutch.resolver.v1.FieldMetadata
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resolver_v1_annotations_proto_init() }
func file_resolver_v1_annotations_proto_init() {
	if File_resolver_v1_annotations_proto != nil {
		return
	}
	file_resolver_v1_schema_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resolver_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_resolver_v1_annotations_proto_goTypes,
		DependencyIndexes: file_resolver_v1_annotations_proto_depIdxs,
		ExtensionInfos:    file_resolver_v1_annotations_proto_extTypes,
	}.Build()
	File_resolver_v1_annotations_proto = out.File
	file_resolver_v1_annotations_proto_rawDesc = nil
	file_resolver_v1_annotations_proto_goTypes = nil
	file_resolver_v1_annotations_proto_depIdxs = nil
}
