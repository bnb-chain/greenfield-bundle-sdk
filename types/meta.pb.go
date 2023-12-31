// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/meta.proto

package types

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

// HashAlgo specifies the hash algorithm used to verify the integrity of the objects in the bundle.
type HashAlgo int32

const (
	HashAlgo_Unknown HashAlgo = 0
	HashAlgo_SHA256  HashAlgo = 1
)

// Enum value maps for HashAlgo.
var (
	HashAlgo_name = map[int32]string{
		0: "Unknown",
		1: "SHA256",
	}
	HashAlgo_value = map[string]int32{
		"Unknown": 0,
		"SHA256":  1,
	}
)

func (x HashAlgo) Enum() *HashAlgo {
	p := new(HashAlgo)
	*p = x
	return p
}

func (x HashAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_meta_proto_enumTypes[0].Descriptor()
}

func (HashAlgo) Type() protoreflect.EnumType {
	return &file_proto_meta_proto_enumTypes[0]
}

func (x HashAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgo.Descriptor instead.
func (HashAlgo) EnumDescriptor() ([]byte, []int) {
	return file_proto_meta_proto_rawDescGZIP(), []int{0}
}

// BundleVersion specifies the bundle version used for the current bundle.
type BundleVersion int32

const (
	BundleVersion_V1 BundleVersion = 0
)

// Enum value maps for BundleVersion.
var (
	BundleVersion_name = map[int32]string{
		0: "V1",
	}
	BundleVersion_value = map[string]int32{
		"V1": 0,
	}
)

func (x BundleVersion) Enum() *BundleVersion {
	p := new(BundleVersion)
	*p = x
	return p
}

func (x BundleVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BundleVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_meta_proto_enumTypes[1].Descriptor()
}

func (BundleVersion) Type() protoreflect.EnumType {
	return &file_proto_meta_proto_enumTypes[1]
}

func (x BundleVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BundleVersion.Descriptor instead.
func (BundleVersion) EnumDescriptor() ([]byte, []int) {
	return file_proto_meta_proto_rawDescGZIP(), []int{1}
}

// ObjectMeta defines metadata for each object in the bundle.
type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the object.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// offset indicates the object's starting position in the bundle.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// size is the size of the object.
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// hash_algo indicates the hash algorithm used for computing the hash.
	HashAlgo HashAlgo `protobuf:"varint,4,opt,name=hash_algo,json=hashAlgo,proto3,enum=HashAlgo" json:"hash_algo,omitempty"`
	// hash is the hash result of the object.
	Hash []byte `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	// content_type indicates the content type of the object.
	ContentType string `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// tags indicates the tags of the object, each object can be tagged with several attributes.
	Tags map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_proto_meta_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectMeta) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ObjectMeta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ObjectMeta) GetHashAlgo() HashAlgo {
	if x != nil {
		return x.HashAlgo
	}
	return HashAlgo_Unknown
}

func (x *ObjectMeta) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ObjectMeta) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ObjectMeta) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// BundleMeta stores the metadata of all the objects in the bundle.
type BundleMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// meta indicates the object meta arrays of the bundle.
	Meta []*ObjectMeta `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty"`
}

func (x *BundleMeta) Reset() {
	*x = BundleMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleMeta) ProtoMessage() {}

func (x *BundleMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleMeta.ProtoReflect.Descriptor instead.
func (*BundleMeta) Descriptor() ([]byte, []int) {
	return file_proto_meta_proto_rawDescGZIP(), []int{1}
}

func (x *BundleMeta) GetMeta() []*ObjectMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_proto_meta_proto protoreflect.FileDescriptor

var file_proto_meta_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x26, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x52,
	0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x2d, 0x0a, 0x0a, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x2a, 0x23, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x2a, 0x17, 0x0a, 0x0d, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x31, 0x10,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_meta_proto_rawDescOnce sync.Once
	file_proto_meta_proto_rawDescData = file_proto_meta_proto_rawDesc
)

func file_proto_meta_proto_rawDescGZIP() []byte {
	file_proto_meta_proto_rawDescOnce.Do(func() {
		file_proto_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_meta_proto_rawDescData)
	})
	return file_proto_meta_proto_rawDescData
}

var file_proto_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_meta_proto_goTypes = []interface{}{
	(HashAlgo)(0),      // 0: HashAlgo
	(BundleVersion)(0), // 1: BundleVersion
	(*ObjectMeta)(nil), // 2: ObjectMeta
	(*BundleMeta)(nil), // 3: BundleMeta
	nil,                // 4: ObjectMeta.TagsEntry
}
var file_proto_meta_proto_depIdxs = []int32{
	0, // 0: ObjectMeta.hash_algo:type_name -> HashAlgo
	4, // 1: ObjectMeta.tags:type_name -> ObjectMeta.TagsEntry
	2, // 2: BundleMeta.meta:type_name -> ObjectMeta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_meta_proto_init() }
func file_proto_meta_proto_init() {
	if File_proto_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
		file_proto_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleMeta); i {
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
			RawDescriptor: file_proto_meta_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_meta_proto_goTypes,
		DependencyIndexes: file_proto_meta_proto_depIdxs,
		EnumInfos:         file_proto_meta_proto_enumTypes,
		MessageInfos:      file_proto_meta_proto_msgTypes,
	}.Build()
	File_proto_meta_proto = out.File
	file_proto_meta_proto_rawDesc = nil
	file_proto_meta_proto_goTypes = nil
	file_proto_meta_proto_depIdxs = nil
}
