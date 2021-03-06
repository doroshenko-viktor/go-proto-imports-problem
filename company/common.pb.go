// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: company/common.proto

package company

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CompanyEnum int32

const (
	CompanyEnum_VAR1 CompanyEnum = 0
	CompanyEnum_VAR2 CompanyEnum = 1
)

// Enum value maps for CompanyEnum.
var (
	CompanyEnum_name = map[int32]string{
		0: "VAR1",
		1: "VAR2",
	}
	CompanyEnum_value = map[string]int32{
		"VAR1": 0,
		"VAR2": 1,
	}
)

func (x CompanyEnum) Enum() *CompanyEnum {
	p := new(CompanyEnum)
	*p = x
	return p
}

func (x CompanyEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompanyEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_company_common_proto_enumTypes[0].Descriptor()
}

func (CompanyEnum) Type() protoreflect.EnumType {
	return &file_company_common_proto_enumTypes[0]
}

func (x CompanyEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompanyEnum.Descriptor instead.
func (CompanyEnum) EnumDescriptor() ([]byte, []int) {
	return file_company_common_proto_rawDescGZIP(), []int{0}
}

var File_company_common_proto protoreflect.FileDescriptor

var file_company_common_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2a,
	0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08,
	0x0a, 0x04, 0x56, 0x41, 0x52, 0x31, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x41, 0x52, 0x32,
	0x10, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_common_proto_rawDescOnce sync.Once
	file_company_common_proto_rawDescData = file_company_common_proto_rawDesc
)

func file_company_common_proto_rawDescGZIP() []byte {
	file_company_common_proto_rawDescOnce.Do(func() {
		file_company_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_common_proto_rawDescData)
	})
	return file_company_common_proto_rawDescData
}

var file_company_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_company_common_proto_goTypes = []interface{}{
	(CompanyEnum)(0), // 0: company.CompanyEnum
}
var file_company_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_company_common_proto_init() }
func file_company_common_proto_init() {
	if File_company_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_company_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_company_common_proto_goTypes,
		DependencyIndexes: file_company_common_proto_depIdxs,
		EnumInfos:         file_company_common_proto_enumTypes,
	}.Build()
	File_company_common_proto = out.File
	file_company_common_proto_rawDesc = nil
	file_company_common_proto_goTypes = nil
	file_company_common_proto_depIdxs = nil
}
