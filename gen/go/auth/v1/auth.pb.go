// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: auth/v1/auth.proto

package auth

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

var File_auth_v1_auth_proto protoreflect.FileDescriptor

var file_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x70, 0x92,
	0x41, 0x52, 0x12, 0x3c, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x12, 0x1f, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x47, 0x6f, 0x53, 0x68, 0x6f, 0x70, 0x27, 0x73, 0x20, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x30,
	0x1a, 0x0f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x67, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x64, 0x65,
	0x76, 0x2a, 0x01, 0x02, 0x5a, 0x19, 0x67, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_v1_auth_proto_goTypes = []interface{}{}
var file_auth_v1_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_proto_init() }
func file_auth_v1_auth_proto_init() {
	if File_auth_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_proto_depIdxs,
	}.Build()
	File_auth_v1_auth_proto = out.File
	file_auth_v1_auth_proto_rawDesc = nil
	file_auth_v1_auth_proto_goTypes = nil
	file_auth_v1_auth_proto_depIdxs = nil
}