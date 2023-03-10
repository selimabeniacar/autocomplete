// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: trie.proto

package triepb

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

type Trie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fanout []*Trie `protobuf:"bytes,1,rep,name=fanout,proto3" json:"fanout,omitempty"`
	Freqs  int32   `protobuf:"varint,2,opt,name=freqs,proto3" json:"freqs,omitempty"`
	End    bool    `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Trie) Reset() {
	*x = Trie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trie) ProtoMessage() {}

func (x *Trie) ProtoReflect() protoreflect.Message {
	mi := &file_trie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trie.ProtoReflect.Descriptor instead.
func (*Trie) Descriptor() ([]byte, []int) {
	return file_trie_proto_rawDescGZIP(), []int{0}
}

func (x *Trie) GetFanout() []*Trie {
	if x != nil {
		return x.Fanout
	}
	return nil
}

func (x *Trie) GetFreqs() int32 {
	if x != nil {
		return x.Freqs
	}
	return 0
}

func (x *Trie) GetEnd() bool {
	if x != nil {
		return x.End
	}
	return false
}

var File_trie_proto protoreflect.FileDescriptor

var file_trie_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x72,
	0x69, 0x65, 0x22, 0x52, 0x0a, 0x04, 0x54, 0x72, 0x69, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x61,
	0x6e, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x72, 0x69,
	0x65, 0x2e, 0x54, 0x72, 0x69, 0x65, 0x52, 0x06, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x72, 0x65, 0x71, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66,
	0x72, 0x65, 0x71, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x74, 0x72, 0x69,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trie_proto_rawDescOnce sync.Once
	file_trie_proto_rawDescData = file_trie_proto_rawDesc
)

func file_trie_proto_rawDescGZIP() []byte {
	file_trie_proto_rawDescOnce.Do(func() {
		file_trie_proto_rawDescData = protoimpl.X.CompressGZIP(file_trie_proto_rawDescData)
	})
	return file_trie_proto_rawDescData
}

var file_trie_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trie_proto_goTypes = []interface{}{
	(*Trie)(nil), // 0: trie.Trie
}
var file_trie_proto_depIdxs = []int32{
	0, // 0: trie.Trie.fanout:type_name -> trie.Trie
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trie_proto_init() }
func file_trie_proto_init() {
	if File_trie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trie); i {
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
			RawDescriptor: file_trie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trie_proto_goTypes,
		DependencyIndexes: file_trie_proto_depIdxs,
		MessageInfos:      file_trie_proto_msgTypes,
	}.Build()
	File_trie_proto = out.File
	file_trie_proto_rawDesc = nil
	file_trie_proto_goTypes = nil
	file_trie_proto_depIdxs = nil
}
