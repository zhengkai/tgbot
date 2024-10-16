// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: callback.proto

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

// https://core.telegram.org/bots/api#callbackquery
//
// This object represents an incoming callback query from a callback
// button in an inline keyboard. If the button that originated the query was
// attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From            *User    `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Message         *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	InlineMessageId string   `protobuf:"bytes,4,opt,name=inline_message_id,json=inlineMessageId,proto3" json:"inline_message_id,omitempty"`
	ChatInstance    string   `protobuf:"bytes,5,opt,name=chat_instance,json=chatInstance,proto3" json:"chat_instance,omitempty"`
	Data            string   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	GameShortName   string   `protobuf:"bytes,7,opt,name=game_short_name,json=gameShortName,proto3" json:"game_short_name,omitempty"`
}

func (x *CallbackQuery) Reset() {
	*x = CallbackQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_callback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackQuery) ProtoMessage() {}

func (x *CallbackQuery) ProtoReflect() protoreflect.Message {
	mi := &file_callback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackQuery.ProtoReflect.Descriptor instead.
func (*CallbackQuery) Descriptor() ([]byte, []int) {
	return file_callback_proto_rawDescGZIP(), []int{0}
}

func (x *CallbackQuery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CallbackQuery) GetFrom() *User {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *CallbackQuery) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *CallbackQuery) GetInlineMessageId() string {
	if x != nil {
		return x.InlineMessageId
	}
	return ""
}

func (x *CallbackQuery) GetChatInstance() string {
	if x != nil {
		return x.ChatInstance
	}
	return ""
}

func (x *CallbackQuery) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CallbackQuery) GetGameShortName() string {
	if x != nil {
		return x.GameShortName
	}
	return ""
}

var File_callback_proto protoreflect.FileDescriptor

var file_callback_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf1, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_callback_proto_rawDescOnce sync.Once
	file_callback_proto_rawDescData = file_callback_proto_rawDesc
)

func file_callback_proto_rawDescGZIP() []byte {
	file_callback_proto_rawDescOnce.Do(func() {
		file_callback_proto_rawDescData = protoimpl.X.CompressGZIP(file_callback_proto_rawDescData)
	})
	return file_callback_proto_rawDescData
}

var file_callback_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_callback_proto_goTypes = []any{
	(*CallbackQuery)(nil), // 0: pb.CallbackQuery
	(*User)(nil),          // 1: pb.User
	(*Message)(nil),       // 2: pb.Message
}
var file_callback_proto_depIdxs = []int32{
	1, // 0: pb.CallbackQuery.from:type_name -> pb.User
	2, // 1: pb.CallbackQuery.message:type_name -> pb.Message
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_callback_proto_init() }
func file_callback_proto_init() {
	if File_callback_proto != nil {
		return
	}
	file_user_proto_init()
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_callback_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CallbackQuery); i {
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
			RawDescriptor: file_callback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_callback_proto_goTypes,
		DependencyIndexes: file_callback_proto_depIdxs,
		MessageInfos:      file_callback_proto_msgTypes,
	}.Build()
	File_callback_proto = out.File
	file_callback_proto_rawDesc = nil
	file_callback_proto_goTypes = nil
	file_callback_proto_depIdxs = nil
}
