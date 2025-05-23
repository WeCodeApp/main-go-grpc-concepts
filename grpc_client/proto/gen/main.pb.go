// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: main.proto

package mainpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             int32                  `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B             int32                  `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_main_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sum           int32                  `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	mi := &file_main_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_main_proto protoreflect.FileDescriptor

const file_main_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"main.proto\x12\n" +
	"calculator\x1a\rgreeter.proto\"(\n" +
	"\n" +
	"AddRequest\x12\f\n" +
	"\x01a\x18\x01 \x01(\x05R\x01a\x12\f\n" +
	"\x01b\x18\x02 \x01(\x05R\x01b\"\x1f\n" +
	"\vAddResponse\x12\x10\n" +
	"\x03sum\x18\x01 \x01(\x05R\x03sum2C\n" +
	"\tCalculate\x126\n" +
	"\x03Add\x12\x16.calculator.AddRequest\x1a\x17.calculator.AddResponse2G\n" +
	"\aGreeter\x12<\n" +
	"\x05Greet\x12\x18.calculator.HelloRequest\x1a\x19.calculator.HelloResponseB\x13Z\x11/proto/gen;mainpbb\x06proto3"

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData []byte
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_main_proto_rawDesc), len(file_main_proto_rawDesc)))
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_main_proto_goTypes = []any{
	(*AddRequest)(nil),    // 0: calculator.AddRequest
	(*AddResponse)(nil),   // 1: calculator.AddResponse
	(*HelloRequest)(nil),  // 2: calculator.HelloRequest
	(*HelloResponse)(nil), // 3: calculator.HelloResponse
}
var file_main_proto_depIdxs = []int32{
	0, // 0: calculator.Calculate.Add:input_type -> calculator.AddRequest
	2, // 1: calculator.Greeter.Greet:input_type -> calculator.HelloRequest
	1, // 2: calculator.Calculate.Add:output_type -> calculator.AddResponse
	3, // 3: calculator.Greeter.Greet:output_type -> calculator.HelloResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	file_greeter_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_main_proto_rawDesc), len(file_main_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}
