// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: messagesproto/v1/messenger.proto

package messengerproto

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

type ExchangeResponse_StatusType int32

const (
	ExchangeResponse_STATUS_TYPE_UNSPECIFIED ExchangeResponse_StatusType = 0
	ExchangeResponse_STATUS_TYPE_SUCCESS     ExchangeResponse_StatusType = 1
	ExchangeResponse_STATUS_TYPE_AUTHFAIL    ExchangeResponse_StatusType = -1
)

// Enum value maps for ExchangeResponse_StatusType.
var (
	ExchangeResponse_StatusType_name = map[int32]string{
		0:  "STATUS_TYPE_UNSPECIFIED",
		1:  "STATUS_TYPE_SUCCESS",
		-1: "STATUS_TYPE_AUTHFAIL",
	}
	ExchangeResponse_StatusType_value = map[string]int32{
		"STATUS_TYPE_UNSPECIFIED": 0,
		"STATUS_TYPE_SUCCESS":     1,
		"STATUS_TYPE_AUTHFAIL":    -1,
	}
)

func (x ExchangeResponse_StatusType) Enum() *ExchangeResponse_StatusType {
	p := new(ExchangeResponse_StatusType)
	*p = x
	return p
}

func (x ExchangeResponse_StatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangeResponse_StatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_messagesproto_v1_messenger_proto_enumTypes[0].Descriptor()
}

func (ExchangeResponse_StatusType) Type() protoreflect.EnumType {
	return &file_messagesproto_v1_messenger_proto_enumTypes[0]
}

func (x ExchangeResponse_StatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangeResponse_StatusType.Descriptor instead.
func (ExchangeResponse_StatusType) EnumDescriptor() ([]byte, []int) {
	return file_messagesproto_v1_messenger_proto_rawDescGZIP(), []int{2, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorName   string `protobuf:"bytes,1,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	AuthorHash   string `protobuf:"bytes,2,opt,name=author_hash,json=authorHash,proto3" json:"author_hash,omitempty"`
	ReceiverName string `protobuf:"bytes,3,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	ReceiverHash string `protobuf:"bytes,4,opt,name=receiver_hash,json=receiverHash,proto3" json:"receiver_hash,omitempty"`
	Subject      string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	CreatedAt    int32  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Body         string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagesproto_v1_messenger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_messagesproto_v1_messenger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_messagesproto_v1_messenger_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *Message) GetAuthorHash() string {
	if x != nil {
		return x.AuthorHash
	}
	return ""
}

func (x *Message) GetReceiverName() string {
	if x != nil {
		return x.ReceiverName
	}
	return ""
}

func (x *Message) GetReceiverHash() string {
	if x != nil {
		return x.ReceiverHash
	}
	return ""
}

func (x *Message) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Message) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type ExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthHash string     `protobuf:"bytes,1,opt,name=auth_hash,json=authHash,proto3" json:"auth_hash,omitempty"`
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ExchangeRequest) Reset() {
	*x = ExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagesproto_v1_messenger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRequest) ProtoMessage() {}

func (x *ExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messagesproto_v1_messenger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRequest) Descriptor() ([]byte, []int) {
	return file_messagesproto_v1_messenger_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRequest) GetAuthHash() string {
	if x != nil {
		return x.AuthHash
	}
	return ""
}

func (x *ExchangeRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        ExchangeResponse_StatusType `protobuf:"varint,1,opt,name=status,proto3,enum=messagesproto.v1.ExchangeResponse_StatusType" json:"status,omitempty"`
	StatusMessage string                      `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	Messages      []*Message                  `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ExchangeResponse) Reset() {
	*x = ExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagesproto_v1_messenger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeResponse) ProtoMessage() {}

func (x *ExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messagesproto_v1_messenger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeResponse) Descriptor() ([]byte, []int) {
	return file_messagesproto_v1_messenger_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeResponse) GetStatus() ExchangeResponse_StatusType {
	if x != nil {
		return x.Status
	}
	return ExchangeResponse_STATUS_TYPE_UNSPECIFIED
}

func (x *ExchangeResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *ExchangeResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_messagesproto_v1_messenger_proto protoreflect.FileDescriptor

var file_messagesproto_v1_messenger_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x22, 0xe2, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x65, 0x0a, 0x0f, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x9e, 0x02, 0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x21,
	0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x46, 0x41, 0x49, 0x4c, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x32, 0x67, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messagesproto_v1_messenger_proto_rawDescOnce sync.Once
	file_messagesproto_v1_messenger_proto_rawDescData = file_messagesproto_v1_messenger_proto_rawDesc
)

func file_messagesproto_v1_messenger_proto_rawDescGZIP() []byte {
	file_messagesproto_v1_messenger_proto_rawDescOnce.Do(func() {
		file_messagesproto_v1_messenger_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagesproto_v1_messenger_proto_rawDescData)
	})
	return file_messagesproto_v1_messenger_proto_rawDescData
}

var file_messagesproto_v1_messenger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messagesproto_v1_messenger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messagesproto_v1_messenger_proto_goTypes = []interface{}{
	(ExchangeResponse_StatusType)(0), // 0: messagesproto.v1.ExchangeResponse.StatusType
	(*Message)(nil),                  // 1: messagesproto.v1.Message
	(*ExchangeRequest)(nil),          // 2: messagesproto.v1.ExchangeRequest
	(*ExchangeResponse)(nil),         // 3: messagesproto.v1.ExchangeResponse
}
var file_messagesproto_v1_messenger_proto_depIdxs = []int32{
	1, // 0: messagesproto.v1.ExchangeRequest.messages:type_name -> messagesproto.v1.Message
	0, // 1: messagesproto.v1.ExchangeResponse.status:type_name -> messagesproto.v1.ExchangeResponse.StatusType
	1, // 2: messagesproto.v1.ExchangeResponse.messages:type_name -> messagesproto.v1.Message
	2, // 3: messagesproto.v1.MessengerService.Exchange:input_type -> messagesproto.v1.ExchangeRequest
	3, // 4: messagesproto.v1.MessengerService.Exchange:output_type -> messagesproto.v1.ExchangeResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_messagesproto_v1_messenger_proto_init() }
func file_messagesproto_v1_messenger_proto_init() {
	if File_messagesproto_v1_messenger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messagesproto_v1_messenger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_messagesproto_v1_messenger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRequest); i {
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
		file_messagesproto_v1_messenger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeResponse); i {
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
			RawDescriptor: file_messagesproto_v1_messenger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messagesproto_v1_messenger_proto_goTypes,
		DependencyIndexes: file_messagesproto_v1_messenger_proto_depIdxs,
		EnumInfos:         file_messagesproto_v1_messenger_proto_enumTypes,
		MessageInfos:      file_messagesproto_v1_messenger_proto_msgTypes,
	}.Build()
	File_messagesproto_v1_messenger_proto = out.File
	file_messagesproto_v1_messenger_proto_rawDesc = nil
	file_messagesproto_v1_messenger_proto_goTypes = nil
	file_messagesproto_v1_messenger_proto_depIdxs = nil
}
