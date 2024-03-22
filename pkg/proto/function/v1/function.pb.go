// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: apollo/proto/function/v1/function.proto

package function

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

type FunctionStatus int32

const (
	FunctionStatus_FUNCTION_STATUS_UNSPECIFIED FunctionStatus = 0
	FunctionStatus_CREATED                     FunctionStatus = 1
	FunctionStatus_INITIALIZED                 FunctionStatus = 2
)

// Enum value maps for FunctionStatus.
var (
	FunctionStatus_name = map[int32]string{
		0: "FUNCTION_STATUS_UNSPECIFIED",
		1: "CREATED",
		2: "INITIALIZED",
	}
	FunctionStatus_value = map[string]int32{
		"FUNCTION_STATUS_UNSPECIFIED": 0,
		"CREATED":                     1,
		"INITIALIZED":                 2,
	}
)

func (x FunctionStatus) Enum() *FunctionStatus {
	p := new(FunctionStatus)
	*p = x
	return p
}

func (x FunctionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FunctionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_apollo_proto_function_v1_function_proto_enumTypes[0].Descriptor()
}

func (FunctionStatus) Type() protoreflect.EnumType {
	return &file_apollo_proto_function_v1_function_proto_enumTypes[0]
}

func (x FunctionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FunctionStatus.Descriptor instead.
func (FunctionStatus) EnumDescriptor() ([]byte, []int) {
	return file_apollo_proto_function_v1_function_proto_rawDescGZIP(), []int{0}
}

var File_apollo_proto_function_v1_function_proto protoreflect.FileDescriptor

var file_apollo_proto_function_v1_function_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x6f, 0x6c, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2a, 0x4f, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a,
	0x45, 0x44, 0x10, 0x02, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6e, 0x6e, 0x69, 0x73, 0x68, 0x69, 0x6c, 0x67, 0x65, 0x72, 0x74,
	0x2f, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apollo_proto_function_v1_function_proto_rawDescOnce sync.Once
	file_apollo_proto_function_v1_function_proto_rawDescData = file_apollo_proto_function_v1_function_proto_rawDesc
)

func file_apollo_proto_function_v1_function_proto_rawDescGZIP() []byte {
	file_apollo_proto_function_v1_function_proto_rawDescOnce.Do(func() {
		file_apollo_proto_function_v1_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_apollo_proto_function_v1_function_proto_rawDescData)
	})
	return file_apollo_proto_function_v1_function_proto_rawDescData
}

var file_apollo_proto_function_v1_function_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apollo_proto_function_v1_function_proto_goTypes = []interface{}{
	(FunctionStatus)(0), // 0: apollo.proto.function.v1.FunctionStatus
}
var file_apollo_proto_function_v1_function_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apollo_proto_function_v1_function_proto_init() }
func file_apollo_proto_function_v1_function_proto_init() {
	if File_apollo_proto_function_v1_function_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apollo_proto_function_v1_function_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apollo_proto_function_v1_function_proto_goTypes,
		DependencyIndexes: file_apollo_proto_function_v1_function_proto_depIdxs,
		EnumInfos:         file_apollo_proto_function_v1_function_proto_enumTypes,
	}.Build()
	File_apollo_proto_function_v1_function_proto = out.File
	file_apollo_proto_function_v1_function_proto_rawDesc = nil
	file_apollo_proto_function_v1_function_proto_goTypes = nil
	file_apollo_proto_function_v1_function_proto_depIdxs = nil
}
