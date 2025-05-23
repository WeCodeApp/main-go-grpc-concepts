// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/main.proto

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

type ChatMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_proto_main_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NumberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumberRequest) Reset() {
	*x = NumberRequest{}
	mi := &file_proto_main_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberRequest) ProtoMessage() {}

func (x *NumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberRequest.ProtoReflect.Descriptor instead.
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

func (x *NumberRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type NumberResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Sum           int32                  `protobuf:"varint,2,opt,name=sum,proto3" json:"sum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumberResponse) Reset() {
	*x = NumberResponse{}
	mi := &file_proto_main_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberResponse) ProtoMessage() {}

func (x *NumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberResponse.ProtoReflect.Descriptor instead.
func (*NumberResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

func (x *NumberResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *NumberResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             int32                  `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B             int32                  `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_proto_main_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
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
	return file_proto_main_proto_rawDescGZIP(), []int{3}
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
	mi := &file_proto_main_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
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
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *AddResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type FibonacciRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	N             int32                  `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	mi := &file_proto_main_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{5}
}

func (x *FibonacciRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type FibonacciResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	mi := &file_proto_main_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{6}
}

func (x *FibonacciResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_proto_main_proto protoreflect.FileDescriptor

const file_proto_main_proto_rawDesc = "" +
	"\n" +
	"\x10proto/main.proto\x12\n" +
	"calculator\"'\n" +
	"\vChatMessage\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"'\n" +
	"\rNumberRequest\x12\x16\n" +
	"\x06number\x18\x01 \x01(\x05R\x06number\":\n" +
	"\x0eNumberResponse\x12\x16\n" +
	"\x06number\x18\x01 \x01(\x05R\x06number\x12\x10\n" +
	"\x03sum\x18\x02 \x01(\x05R\x03sum\"(\n" +
	"\n" +
	"AddRequest\x12\f\n" +
	"\x01a\x18\x01 \x01(\x05R\x01a\x12\f\n" +
	"\x01b\x18\x02 \x01(\x05R\x01b\"\x1f\n" +
	"\vAddResponse\x12\x10\n" +
	"\x03sum\x18\x01 \x01(\x05R\x03sum\" \n" +
	"\x10FibonacciRequest\x12\f\n" +
	"\x01n\x18\x01 \x01(\x05R\x01n\"+\n" +
	"\x11FibonacciResponse\x12\x16\n" +
	"\x06number\x18\x01 \x01(\x05R\x06number2\x9d\x02\n" +
	"\tCalculate\x126\n" +
	"\x03Add\x12\x16.calculator.AddRequest\x1a\x17.calculator.AddResponse\x12R\n" +
	"\x11GenerateFibonacci\x12\x1c.calculator.FibonacciRequest\x1a\x1d.calculator.FibonacciResponse0\x01\x12F\n" +
	"\vSendNumbers\x12\x19.calculator.NumberRequest\x1a\x1a.calculator.NumberResponse(\x01\x12<\n" +
	"\x04Chat\x12\x17.calculator.ChatMessage\x1a\x17.calculator.ChatMessage(\x010\x01B\x13Z\x11/proto/gen;mainpbb\x06proto3"

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData []byte
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_main_proto_rawDesc), len(file_proto_main_proto_rawDesc)))
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_main_proto_goTypes = []any{
	(*ChatMessage)(nil),       // 0: calculator.ChatMessage
	(*NumberRequest)(nil),     // 1: calculator.NumberRequest
	(*NumberResponse)(nil),    // 2: calculator.NumberResponse
	(*AddRequest)(nil),        // 3: calculator.AddRequest
	(*AddResponse)(nil),       // 4: calculator.AddResponse
	(*FibonacciRequest)(nil),  // 5: calculator.FibonacciRequest
	(*FibonacciResponse)(nil), // 6: calculator.FibonacciResponse
}
var file_proto_main_proto_depIdxs = []int32{
	3, // 0: calculator.Calculate.Add:input_type -> calculator.AddRequest
	5, // 1: calculator.Calculate.GenerateFibonacci:input_type -> calculator.FibonacciRequest
	1, // 2: calculator.Calculate.SendNumbers:input_type -> calculator.NumberRequest
	0, // 3: calculator.Calculate.Chat:input_type -> calculator.ChatMessage
	4, // 4: calculator.Calculate.Add:output_type -> calculator.AddResponse
	6, // 5: calculator.Calculate.GenerateFibonacci:output_type -> calculator.FibonacciResponse
	2, // 6: calculator.Calculate.SendNumbers:output_type -> calculator.NumberResponse
	0, // 7: calculator.Calculate.Chat:output_type -> calculator.ChatMessage
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_main_proto_rawDesc), len(file_proto_main_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}
