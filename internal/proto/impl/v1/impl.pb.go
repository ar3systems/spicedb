// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: impl/v1/impl.proto

package v1

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

type RelationMetadata_RelationKind int32

const (
	RelationMetadata_UNKNOWN_KIND RelationMetadata_RelationKind = 0
	RelationMetadata_RELATION     RelationMetadata_RelationKind = 1
	RelationMetadata_PERMISSION   RelationMetadata_RelationKind = 2
)

// Enum value maps for RelationMetadata_RelationKind.
var (
	RelationMetadata_RelationKind_name = map[int32]string{
		0: "UNKNOWN_KIND",
		1: "RELATION",
		2: "PERMISSION",
	}
	RelationMetadata_RelationKind_value = map[string]int32{
		"UNKNOWN_KIND": 0,
		"RELATION":     1,
		"PERMISSION":   2,
	}
)

func (x RelationMetadata_RelationKind) Enum() *RelationMetadata_RelationKind {
	p := new(RelationMetadata_RelationKind)
	*p = x
	return p
}

func (x RelationMetadata_RelationKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationMetadata_RelationKind) Descriptor() protoreflect.EnumDescriptor {
	return file_impl_v1_impl_proto_enumTypes[0].Descriptor()
}

func (RelationMetadata_RelationKind) Type() protoreflect.EnumType {
	return &file_impl_v1_impl_proto_enumTypes[0]
}

func (x RelationMetadata_RelationKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationMetadata_RelationKind.Descriptor instead.
func (RelationMetadata_RelationKind) EnumDescriptor() ([]byte, []int) {
	return file_impl_v1_impl_proto_rawDescGZIP(), []int{2, 0}
}

type DecodedZookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Types that are assignable to VersionOneof:
	//	*DecodedZookie_V1
	//	*DecodedZookie_V2
	VersionOneof isDecodedZookie_VersionOneof `protobuf_oneof:"version_oneof"`
}

func (x *DecodedZookie) Reset() {
	*x = DecodedZookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_impl_v1_impl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedZookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedZookie) ProtoMessage() {}

func (x *DecodedZookie) ProtoReflect() protoreflect.Message {
	mi := &file_impl_v1_impl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedZookie.ProtoReflect.Descriptor instead.
func (*DecodedZookie) Descriptor() ([]byte, []int) {
	return file_impl_v1_impl_proto_rawDescGZIP(), []int{0}
}

func (x *DecodedZookie) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (m *DecodedZookie) GetVersionOneof() isDecodedZookie_VersionOneof {
	if m != nil {
		return m.VersionOneof
	}
	return nil
}

func (x *DecodedZookie) GetV1() *DecodedZookie_V1Zookie {
	if x, ok := x.GetVersionOneof().(*DecodedZookie_V1); ok {
		return x.V1
	}
	return nil
}

func (x *DecodedZookie) GetV2() *DecodedZookie_V2Zookie {
	if x, ok := x.GetVersionOneof().(*DecodedZookie_V2); ok {
		return x.V2
	}
	return nil
}

type isDecodedZookie_VersionOneof interface {
	isDecodedZookie_VersionOneof()
}

type DecodedZookie_V1 struct {
	V1 *DecodedZookie_V1Zookie `protobuf:"bytes,2,opt,name=v1,proto3,oneof"`
}

type DecodedZookie_V2 struct {
	V2 *DecodedZookie_V2Zookie `protobuf:"bytes,3,opt,name=v2,proto3,oneof"`
}

func (*DecodedZookie_V1) isDecodedZookie_VersionOneof() {}

func (*DecodedZookie_V2) isDecodedZookie_VersionOneof() {}

type DocComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment string `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *DocComment) Reset() {
	*x = DocComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_impl_v1_impl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocComment) ProtoMessage() {}

func (x *DocComment) ProtoReflect() protoreflect.Message {
	mi := &file_impl_v1_impl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocComment.ProtoReflect.Descriptor instead.
func (*DocComment) Descriptor() ([]byte, []int) {
	return file_impl_v1_impl_proto_rawDescGZIP(), []int{1}
}

func (x *DocComment) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type RelationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind RelationMetadata_RelationKind `protobuf:"varint,1,opt,name=kind,proto3,enum=impl.v1.RelationMetadata_RelationKind" json:"kind,omitempty"`
}

func (x *RelationMetadata) Reset() {
	*x = RelationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_impl_v1_impl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationMetadata) ProtoMessage() {}

func (x *RelationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_impl_v1_impl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationMetadata.ProtoReflect.Descriptor instead.
func (*RelationMetadata) Descriptor() ([]byte, []int) {
	return file_impl_v1_impl_proto_rawDescGZIP(), []int{2}
}

func (x *RelationMetadata) GetKind() RelationMetadata_RelationKind {
	if x != nil {
		return x.Kind
	}
	return RelationMetadata_UNKNOWN_KIND
}

type DecodedZookie_V1Zookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision uint64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *DecodedZookie_V1Zookie) Reset() {
	*x = DecodedZookie_V1Zookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_impl_v1_impl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedZookie_V1Zookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedZookie_V1Zookie) ProtoMessage() {}

func (x *DecodedZookie_V1Zookie) ProtoReflect() protoreflect.Message {
	mi := &file_impl_v1_impl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedZookie_V1Zookie.ProtoReflect.Descriptor instead.
func (*DecodedZookie_V1Zookie) Descriptor() ([]byte, []int) {
	return file_impl_v1_impl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DecodedZookie_V1Zookie) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

type DecodedZookie_V2Zookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision string `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *DecodedZookie_V2Zookie) Reset() {
	*x = DecodedZookie_V2Zookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_impl_v1_impl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedZookie_V2Zookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedZookie_V2Zookie) ProtoMessage() {}

func (x *DecodedZookie_V2Zookie) ProtoReflect() protoreflect.Message {
	mi := &file_impl_v1_impl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedZookie_V2Zookie.ProtoReflect.Descriptor instead.
func (*DecodedZookie_V2Zookie) Descriptor() ([]byte, []int) {
	return file_impl_v1_impl_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DecodedZookie_V2Zookie) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

var File_impl_v1_impl_proto protoreflect.FileDescriptor

var file_impl_v1_impl_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0xf0, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x56, 0x31,
	0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x48, 0x00, 0x52, 0x02, 0x76, 0x31, 0x12, 0x31, 0x0a, 0x02,
	0x76, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x2e, 0x56, 0x32, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x48, 0x00, 0x52, 0x02, 0x76, 0x32, 0x1a,
	0x26, 0x0a, 0x08, 0x56, 0x31, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x26, 0x0a, 0x08, 0x56, 0x32, 0x5a, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x0f, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x22, 0x26, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x69, 0x6d,
	0x70, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x61, 0x64, 0x61, 0x6e, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d,
	0x70, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_impl_v1_impl_proto_rawDescOnce sync.Once
	file_impl_v1_impl_proto_rawDescData = file_impl_v1_impl_proto_rawDesc
)

func file_impl_v1_impl_proto_rawDescGZIP() []byte {
	file_impl_v1_impl_proto_rawDescOnce.Do(func() {
		file_impl_v1_impl_proto_rawDescData = protoimpl.X.CompressGZIP(file_impl_v1_impl_proto_rawDescData)
	})
	return file_impl_v1_impl_proto_rawDescData
}

var file_impl_v1_impl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_impl_v1_impl_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_impl_v1_impl_proto_goTypes = []interface{}{
	(RelationMetadata_RelationKind)(0), // 0: impl.v1.RelationMetadata.RelationKind
	(*DecodedZookie)(nil),              // 1: impl.v1.DecodedZookie
	(*DocComment)(nil),                 // 2: impl.v1.DocComment
	(*RelationMetadata)(nil),           // 3: impl.v1.RelationMetadata
	(*DecodedZookie_V1Zookie)(nil),     // 4: impl.v1.DecodedZookie.V1Zookie
	(*DecodedZookie_V2Zookie)(nil),     // 5: impl.v1.DecodedZookie.V2Zookie
}
var file_impl_v1_impl_proto_depIdxs = []int32{
	4, // 0: impl.v1.DecodedZookie.v1:type_name -> impl.v1.DecodedZookie.V1Zookie
	5, // 1: impl.v1.DecodedZookie.v2:type_name -> impl.v1.DecodedZookie.V2Zookie
	0, // 2: impl.v1.RelationMetadata.kind:type_name -> impl.v1.RelationMetadata.RelationKind
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_impl_v1_impl_proto_init() }
func file_impl_v1_impl_proto_init() {
	if File_impl_v1_impl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_impl_v1_impl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedZookie); i {
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
		file_impl_v1_impl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocComment); i {
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
		file_impl_v1_impl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationMetadata); i {
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
		file_impl_v1_impl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedZookie_V1Zookie); i {
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
		file_impl_v1_impl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedZookie_V2Zookie); i {
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
	file_impl_v1_impl_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DecodedZookie_V1)(nil),
		(*DecodedZookie_V2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_impl_v1_impl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_impl_v1_impl_proto_goTypes,
		DependencyIndexes: file_impl_v1_impl_proto_depIdxs,
		EnumInfos:         file_impl_v1_impl_proto_enumTypes,
		MessageInfos:      file_impl_v1_impl_proto_msgTypes,
	}.Build()
	File_impl_v1_impl_proto = out.File
	file_impl_v1_impl_proto_rawDesc = nil
	file_impl_v1_impl_proto_goTypes = nil
	file_impl_v1_impl_proto_depIdxs = nil
}