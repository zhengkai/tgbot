// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: reply.proto

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

type ReplyParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId                uint64           `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	ChatId                   string           `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool             `protobuf:"varint,3,opt,name=allow_sending_without_reply,json=allowSendingWithoutReply,proto3" json:"allow_sending_without_reply,omitempty"`
	Quote                    string           `protobuf:"bytes,4,opt,name=quote,proto3" json:"quote,omitempty"`
	QuoteParseMode           string           `protobuf:"bytes,5,opt,name=quote_parse_mode,json=quoteParseMode,proto3" json:"quote_parse_mode,omitempty"`
	QuoteEntities            []*MessageEntity `protobuf:"bytes,6,rep,name=quote_entities,json=quoteEntities,proto3" json:"quote_entities,omitempty"`
	QuotePosition            uint64           `protobuf:"varint,7,opt,name=quote_position,json=quotePosition,proto3" json:"quote_position,omitempty"`
}

func (x *ReplyParameters) Reset() {
	*x = ReplyParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyParameters) ProtoMessage() {}

func (x *ReplyParameters) ProtoReflect() protoreflect.Message {
	mi := &file_reply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyParameters.ProtoReflect.Descriptor instead.
func (*ReplyParameters) Descriptor() ([]byte, []int) {
	return file_reply_proto_rawDescGZIP(), []int{0}
}

func (x *ReplyParameters) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReplyParameters) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ReplyParameters) GetAllowSendingWithoutReply() bool {
	if x != nil {
		return x.AllowSendingWithoutReply
	}
	return false
}

func (x *ReplyParameters) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *ReplyParameters) GetQuoteParseMode() string {
	if x != nil {
		return x.QuoteParseMode
	}
	return ""
}

func (x *ReplyParameters) GetQuoteEntities() []*MessageEntity {
	if x != nil {
		return x.QuoteEntities
	}
	return nil
}

func (x *ReplyParameters) GetQuotePosition() uint64 {
	if x != nil {
		return x.QuotePosition
	}
	return 0
}

var File_reply_proto protoreflect.FileDescriptor

var file_reply_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x0a, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02,
	0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reply_proto_rawDescOnce sync.Once
	file_reply_proto_rawDescData = file_reply_proto_rawDesc
)

func file_reply_proto_rawDescGZIP() []byte {
	file_reply_proto_rawDescOnce.Do(func() {
		file_reply_proto_rawDescData = protoimpl.X.CompressGZIP(file_reply_proto_rawDescData)
	})
	return file_reply_proto_rawDescData
}

var file_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reply_proto_goTypes = []any{
	(*ReplyParameters)(nil), // 0: pb.ReplyParameters
	(*MessageEntity)(nil),   // 1: pb.MessageEntity
}
var file_reply_proto_depIdxs = []int32{
	1, // 0: pb.ReplyParameters.quote_entities:type_name -> pb.MessageEntity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reply_proto_init() }
func file_reply_proto_init() {
	if File_reply_proto != nil {
		return
	}
	file_misc_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reply_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReplyParameters); i {
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
			RawDescriptor: file_reply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reply_proto_goTypes,
		DependencyIndexes: file_reply_proto_depIdxs,
		MessageInfos:      file_reply_proto_msgTypes,
	}.Build()
	File_reply_proto = out.File
	file_reply_proto_rawDesc = nil
	file_reply_proto_goTypes = nil
	file_reply_proto_depIdxs = nil
}
