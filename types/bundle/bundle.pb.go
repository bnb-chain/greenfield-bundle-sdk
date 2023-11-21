// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/bundle.proto

package bundle

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

type HashAlgo int32

const (
	HashAlgo_SHA256 HashAlgo = 0
)

// Enum value maps for HashAlgo.
var (
	HashAlgo_name = map[int32]string{
		0: "SHA256",
	}
	HashAlgo_value = map[string]int32{
		"SHA256": 0,
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
	return file_proto_bundle_proto_enumTypes[0].Descriptor()
}

func (HashAlgo) Type() protoreflect.EnumType {
	return &file_proto_bundle_proto_enumTypes[0]
}

func (x HashAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgo.Descriptor instead.
func (HashAlgo) EnumDescriptor() ([]byte, []int) {
	return file_proto_bundle_proto_rawDescGZIP(), []int{0}
}

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Offset      uint64            `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size        uint64            `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	HashAlgo    HashAlgo          `protobuf:"varint,4,opt,name=hash_algo,json=hashAlgo,proto3,enum=HashAlgo" json:"hash_algo,omitempty"`
	Hash        []byte            `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	ContentType string            `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,7,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bundle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bundle_proto_msgTypes[0]
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
	return file_proto_bundle_proto_rawDescGZIP(), []int{0}
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
	return HashAlgo_SHA256
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

func (x *ObjectMeta) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type BundleMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metas []*ObjectMeta `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *BundleMeta) Reset() {
	*x = BundleMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bundle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleMeta) ProtoMessage() {}

func (x *BundleMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bundle_proto_msgTypes[1]
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
	return file_proto_bundle_proto_rawDescGZIP(), []int{1}
}

func (x *BundleMeta) GetMetas() []*ObjectMeta {
	if x != nil {
		return x.Metas
	}
	return nil
}

type Bundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	MetaSize uint64 `protobuf:"varint,2,opt,name=meta_size,json=metaSize,proto3" json:"meta_size,omitempty"`
	Meta     []byte `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	Data     []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Bundle) Reset() {
	*x = Bundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bundle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bundle) ProtoMessage() {}

func (x *Bundle) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bundle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bundle.ProtoReflect.Descriptor instead.
func (*Bundle) Descriptor() ([]byte, []int) {
	return file_proto_bundle_proto_rawDescGZIP(), []int{2}
}

func (x *Bundle) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Bundle) GetMetaSize() uint64 {
	if x != nil {
		return x.MetaSize
	}
	return 0
}

func (x *Bundle) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Bundle) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_bundle_proto protoreflect.FileDescriptor

var file_proto_bundle_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67,
	0x6f, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a,
	0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f,
	0x0a, 0x0a, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x05,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x22,
	0x67, 0x0a, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x16, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68,
	0x41, 0x6c, 0x67, 0x6f, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bundle_proto_rawDescOnce sync.Once
	file_proto_bundle_proto_rawDescData = file_proto_bundle_proto_rawDesc
)

func file_proto_bundle_proto_rawDescGZIP() []byte {
	file_proto_bundle_proto_rawDescOnce.Do(func() {
		file_proto_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bundle_proto_rawDescData)
	})
	return file_proto_bundle_proto_rawDescData
}

var file_proto_bundle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_bundle_proto_goTypes = []interface{}{
	(HashAlgo)(0),      // 0: HashAlgo
	(*ObjectMeta)(nil), // 1: ObjectMeta
	(*BundleMeta)(nil), // 2: BundleMeta
	(*Bundle)(nil),     // 3: Bundle
	nil,                // 4: ObjectMeta.AttributesEntry
}
var file_proto_bundle_proto_depIdxs = []int32{
	0, // 0: ObjectMeta.hash_algo:type_name -> HashAlgo
	4, // 1: ObjectMeta.attributes:type_name -> ObjectMeta.AttributesEntry
	1, // 2: BundleMeta.metas:type_name -> ObjectMeta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_bundle_proto_init() }
func file_proto_bundle_proto_init() {
	if File_proto_bundle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bundle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_bundle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_bundle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bundle); i {
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
			RawDescriptor: file_proto_bundle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bundle_proto_goTypes,
		DependencyIndexes: file_proto_bundle_proto_depIdxs,
		EnumInfos:         file_proto_bundle_proto_enumTypes,
		MessageInfos:      file_proto_bundle_proto_msgTypes,
	}.Build()
	File_proto_bundle_proto = out.File
	file_proto_bundle_proto_rawDesc = nil
	file_proto_bundle_proto_goTypes = nil
	file_proto_bundle_proto_depIdxs = nil
}
