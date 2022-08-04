// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: keyboard_message.proto

package pb

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

type Keyboard_Layout int32

const (
	Keyboard_UNKNOWN Keyboard_Layout = 0
	Keyboard_QWERTY  Keyboard_Layout = 1
	Keyboard_QWERTZ  Keyboard_Layout = 2 // Germany
	Keyboard_AZERTY  Keyboard_Layout = 3 // French
)

// Enum value maps for Keyboard_Layout.
var (
	Keyboard_Layout_name = map[int32]string{
		0: "UNKNOWN",
		1: "QWERTY",
		2: "QWERTZ",
		3: "AZERTY",
	}
	Keyboard_Layout_value = map[string]int32{
		"UNKNOWN": 0,
		"QWERTY":  1,
		"QWERTZ":  2,
		"AZERTY":  3,
	}
)

func (x Keyboard_Layout) Enum() *Keyboard_Layout {
	p := new(Keyboard_Layout)
	*p = x
	return p
}

func (x Keyboard_Layout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Keyboard_Layout) Descriptor() protoreflect.EnumDescriptor {
	return file_keyboard_message_proto_enumTypes[0].Descriptor()
}

func (Keyboard_Layout) Type() protoreflect.EnumType {
	return &file_keyboard_message_proto_enumTypes[0]
}

func (x Keyboard_Layout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Keyboard_Layout.Descriptor instead.
func (Keyboard_Layout) EnumDescriptor() ([]byte, []int) {
	return file_keyboard_message_proto_rawDescGZIP(), []int{0, 0}
}

type Keyboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layout  Keyboard_Layout `protobuf:"varint,1,opt,name=layout,proto3,enum=pb.Keyboard_Layout" json:"layout,omitempty"`
	Backlit bool            `protobuf:"varint,2,opt,name=backlit,proto3" json:"backlit,omitempty"`
}

func (x *Keyboard) Reset() {
	*x = Keyboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyboard) ProtoMessage() {}

func (x *Keyboard) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyboard.ProtoReflect.Descriptor instead.
func (*Keyboard) Descriptor() ([]byte, []int) {
	return file_keyboard_message_proto_rawDescGZIP(), []int{0}
}

func (x *Keyboard) GetLayout() Keyboard_Layout {
	if x != nil {
		return x.Layout
	}
	return Keyboard_UNKNOWN
}

func (x *Keyboard) GetBacklit() bool {
	if x != nil {
		return x.Backlit
	}
	return false
}

var File_keyboard_message_proto protoreflect.FileDescriptor

var file_keyboard_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x8c, 0x01, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4b,
	0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x06,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x74,
	0x22, 0x39, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x57, 0x45, 0x52, 0x54,
	0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x57, 0x45, 0x52, 0x54, 0x5a, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x5a, 0x45, 0x52, 0x54, 0x59, 0x10, 0x03, 0x42, 0x05, 0x5a, 0x03, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keyboard_message_proto_rawDescOnce sync.Once
	file_keyboard_message_proto_rawDescData = file_keyboard_message_proto_rawDesc
)

func file_keyboard_message_proto_rawDescGZIP() []byte {
	file_keyboard_message_proto_rawDescOnce.Do(func() {
		file_keyboard_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_keyboard_message_proto_rawDescData)
	})
	return file_keyboard_message_proto_rawDescData
}

var file_keyboard_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_keyboard_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_keyboard_message_proto_goTypes = []interface{}{
	(Keyboard_Layout)(0), // 0: pb.Keyboard.Layout
	(*Keyboard)(nil),     // 1: pb.Keyboard
}
var file_keyboard_message_proto_depIdxs = []int32{
	0, // 0: pb.Keyboard.layout:type_name -> pb.Keyboard.Layout
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_keyboard_message_proto_init() }
func file_keyboard_message_proto_init() {
	if File_keyboard_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keyboard_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyboard); i {
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
			RawDescriptor: file_keyboard_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keyboard_message_proto_goTypes,
		DependencyIndexes: file_keyboard_message_proto_depIdxs,
		EnumInfos:         file_keyboard_message_proto_enumTypes,
		MessageInfos:      file_keyboard_message_proto_msgTypes,
	}.Build()
	File_keyboard_message_proto = out.File
	file_keyboard_message_proto_rawDesc = nil
	file_keyboard_message_proto_goTypes = nil
	file_keyboard_message_proto_depIdxs = nil
}