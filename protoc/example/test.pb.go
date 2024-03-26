// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: protoc/example/test.proto

package example

import (
	_ "github.com/ihatiko/protoc-gen-go-tags/protoc/tags"
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

type ExampleCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FineType:
	//
	//	*ExampleCase_FieldString
	//	*ExampleCase_FieldInt32
	FineType isExampleCase_FineType `protobuf_oneof:"fine_type"`
	Field_1  string                 `protobuf:"bytes,3,opt,name=field_1,json=field1,proto3" json:"field_1,omitempty"`
}

func (x *ExampleCase) Reset() {
	*x = ExampleCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_example_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleCase) ProtoMessage() {}

func (x *ExampleCase) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_example_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleCase.ProtoReflect.Descriptor instead.
func (*ExampleCase) Descriptor() ([]byte, []int) {
	return file_protoc_example_test_proto_rawDescGZIP(), []int{0}
}

func (m *ExampleCase) GetFineType() isExampleCase_FineType {
	if m != nil {
		return m.FineType
	}
	return nil
}

func (x *ExampleCase) GetFieldString() string {
	if x, ok := x.GetFineType().(*ExampleCase_FieldString); ok {
		return x.FieldString
	}
	return ""
}

func (x *ExampleCase) GetFieldInt32() int32 {
	if x, ok := x.GetFineType().(*ExampleCase_FieldInt32); ok {
		return x.FieldInt32
	}
	return 0
}

func (x *ExampleCase) GetField_1() string {
	if x != nil {
		return x.Field_1
	}
	return ""
}

type isExampleCase_FineType interface {
	isExampleCase_FineType()
}

type ExampleCase_FieldString struct {
	FieldString string `protobuf:"bytes,1,opt,name=field_string,json=fieldString,proto3,oneof" validate:"uuid"`
}

type ExampleCase_FieldInt32 struct {
	FieldInt32 int32 `protobuf:"varint,2,opt,name=field_int32,json=fieldInt32,proto3,oneof" validate:"uuid"`
}

func (*ExampleCase_FieldString) isExampleCase_FineType() {}

func (*ExampleCase_FieldInt32) isExampleCase_FineType() {}

var File_protoc_example_test_proto protoreflect.FileDescriptor

var file_protoc_example_test_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x74, 0x61, 0x67,
	0x73, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0c,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xda, 0x83, 0x9e, 0x03, 0x04, 0x75, 0x75, 0x69, 0x64, 0x48, 0x00, 0x52,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x0b,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x09, 0xda, 0x83, 0x9e, 0x03, 0x04, 0x75, 0x75, 0x69, 0x64, 0x48, 0x00, 0x52, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x42, 0x0b, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_example_test_proto_rawDescOnce sync.Once
	file_protoc_example_test_proto_rawDescData = file_protoc_example_test_proto_rawDesc
)

func file_protoc_example_test_proto_rawDescGZIP() []byte {
	file_protoc_example_test_proto_rawDescOnce.Do(func() {
		file_protoc_example_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_example_test_proto_rawDescData)
	})
	return file_protoc_example_test_proto_rawDescData
}

var file_protoc_example_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protoc_example_test_proto_goTypes = []interface{}{
	(*ExampleCase)(nil), // 0: example.ExampleCase
}
var file_protoc_example_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_example_test_proto_init() }
func file_protoc_example_test_proto_init() {
	if File_protoc_example_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_example_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleCase); i {
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
	file_protoc_example_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExampleCase_FieldString)(nil),
		(*ExampleCase_FieldInt32)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoc_example_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoc_example_test_proto_goTypes,
		DependencyIndexes: file_protoc_example_test_proto_depIdxs,
		MessageInfos:      file_protoc_example_test_proto_msgTypes,
	}.Build()
	File_protoc_example_test_proto = out.File
	file_protoc_example_test_proto_rawDesc = nil
	file_protoc_example_test_proto_goTypes = nil
	file_protoc_example_test_proto_depIdxs = nil
}