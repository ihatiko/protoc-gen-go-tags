// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: protoc/tags/tags.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Custom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Custom) Reset() {
	*x = Custom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_tags_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Custom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Custom) ProtoMessage() {}

func (x *Custom) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_tags_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Custom.ProtoReflect.Descriptor instead.
func (*Custom) Descriptor() ([]byte, []int) {
	return file_protoc_tags_tags_proto_rawDescGZIP(), []int{0}
}

func (x *Custom) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Custom) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var file_protoc_tags_tags_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847930,
		Name:          "tags.json",
		Tag:           "bytes,847930,opt,name=json",
		Filename:      "protoc/tags/tags.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847931,
		Name:          "tags.validate",
		Tag:           "bytes,847931,opt,name=validate",
		Filename:      "protoc/tags/tags.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847932,
		Name:          "tags.query",
		Tag:           "bytes,847932,opt,name=query",
		Filename:      "protoc/tags/tags.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847933,
		Name:          "tags.param",
		Tag:           "bytes,847933,opt,name=param",
		Filename:      "protoc/tags/tags.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: ([]*Custom)(nil),
		Field:         847934,
		Name:          "tags.custom",
		Tag:           "bytes,847934,rep,name=custom",
		Filename:      "protoc/tags/tags.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847935,
		Name:          "tags.db",
		Tag:           "bytes,847935,opt,name=db",
		Filename:      "protoc/tags/tags.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847936,
		Name:          "tags.header",
		Tag:           "bytes,847936,opt,name=header",
		Filename:      "protoc/tags/tags.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string json = 847930;
	E_Json = &file_protoc_tags_tags_proto_extTypes[0]
	// optional string validate = 847931;
	E_Validate = &file_protoc_tags_tags_proto_extTypes[1]
	// optional string query = 847932;
	E_Query = &file_protoc_tags_tags_proto_extTypes[2]
	// optional string param = 847933;
	E_Param = &file_protoc_tags_tags_proto_extTypes[3]
	// repeated tags.Custom custom = 847934;
	E_Custom = &file_protoc_tags_tags_proto_extTypes[4]
	// optional string db = 847935;
	E_Db = &file_protoc_tags_tags_proto_extTypes[5]
	// optional string header = 847936;
	E_Header = &file_protoc_tags_tags_proto_extTypes[6]
)

var File_protoc_tags_tags_proto protoreflect.FileDescriptor

var file_protoc_tags_tags_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x74, 0x61,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x30, 0x0a, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x33, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xba, 0xe0, 0x33, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x3b, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xbb, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x35, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbc, 0xe0, 0x33,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x35, 0x0a, 0x05, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x3a, 0x45, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0xe0, 0x33, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x3a, 0x2f, 0x0a, 0x02, 0x64, 0x62, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf,
	0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x3a, 0x37, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xc0, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x68, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x3b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_tags_tags_proto_rawDescOnce sync.Once
	file_protoc_tags_tags_proto_rawDescData = file_protoc_tags_tags_proto_rawDesc
)

func file_protoc_tags_tags_proto_rawDescGZIP() []byte {
	file_protoc_tags_tags_proto_rawDescOnce.Do(func() {
		file_protoc_tags_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_tags_tags_proto_rawDescData)
	})
	return file_protoc_tags_tags_proto_rawDescData
}

var file_protoc_tags_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protoc_tags_tags_proto_goTypes = []interface{}{
	(*Custom)(nil),                    // 0: tags.Custom
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_protoc_tags_tags_proto_depIdxs = []int32{
	1, // 0: tags.json:extendee -> google.protobuf.FieldOptions
	1, // 1: tags.validate:extendee -> google.protobuf.FieldOptions
	1, // 2: tags.query:extendee -> google.protobuf.FieldOptions
	1, // 3: tags.param:extendee -> google.protobuf.FieldOptions
	1, // 4: tags.custom:extendee -> google.protobuf.FieldOptions
	1, // 5: tags.db:extendee -> google.protobuf.FieldOptions
	1, // 6: tags.header:extendee -> google.protobuf.FieldOptions
	0, // 7: tags.custom:type_name -> tags.Custom
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	7, // [7:8] is the sub-list for extension type_name
	0, // [0:7] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_tags_tags_proto_init() }
func file_protoc_tags_tags_proto_init() {
	if File_protoc_tags_tags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_tags_tags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Custom); i {
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
			RawDescriptor: file_protoc_tags_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 7,
			NumServices:   0,
		},
		GoTypes:           file_protoc_tags_tags_proto_goTypes,
		DependencyIndexes: file_protoc_tags_tags_proto_depIdxs,
		MessageInfos:      file_protoc_tags_tags_proto_msgTypes,
		ExtensionInfos:    file_protoc_tags_tags_proto_extTypes,
	}.Build()
	File_protoc_tags_tags_proto = out.File
	file_protoc_tags_tags_proto_rawDesc = nil
	file_protoc_tags_tags_proto_goTypes = nil
	file_protoc_tags_tags_proto_depIdxs = nil
}
