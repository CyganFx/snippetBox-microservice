// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/grpc/protos/news.protos

package protobuffs

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy protos package is being used.
const _ = proto.ProtoPackageIsVersion4

type NewsSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *NewsSendRequest) Reset() {
	*x = NewsSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsSendRequest) ProtoMessage() {}

func (x *NewsSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsSendRequest.ProtoReflect.Descriptor instead.
func (*NewsSendRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_news_proto_rawDescGZIP(), []int{0}
}

func (x *NewsSendRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewsSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title   string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string                 `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	Expires *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Expires,proto3" json:"Expires,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=Created,proto3" json:"Created,omitempty"`
}

func (x *NewsSendResponse) Reset() {
	*x = NewsSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsSendResponse) ProtoMessage() {}

func (x *NewsSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsSendResponse.ProtoReflect.Descriptor instead.
func (*NewsSendResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_news_proto_rawDescGZIP(), []int{1}
}

func (x *NewsSendResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewsSendResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsSendResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewsSendResponse) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *NewsSendResponse) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type NewsCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string                 `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string                 `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Expires *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Expires,proto3" json:"Expires,omitempty"`
}

func (x *NewsCreateRequest) Reset() {
	*x = NewsCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCreateRequest) ProtoMessage() {}

func (x *NewsCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCreateRequest.ProtoReflect.Descriptor instead.
func (*NewsCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_news_proto_rawDescGZIP(), []int{2}
}

func (x *NewsCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsCreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewsCreateRequest) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

type NewsCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *NewsCreateResponse) Reset() {
	*x = NewsCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsCreateResponse) ProtoMessage() {}

func (x *NewsCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_news_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsCreateResponse.ProtoReflect.Descriptor instead.
func (*NewsCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_news_proto_rawDescGZIP(), []int{3}
}

func (x *NewsCreateResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_grpc_protos_news_proto protoreflect.FileDescriptor

var file_api_grpc_protos_news_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x4e, 0x65, 0x77,
	0x73, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x79, 0x0a, 0x11, 0x4e, 0x65, 0x77,
	0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34,
	0x0a, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x4e,
	0x65, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65,
	0x77, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x5c, 0x43, 0x79, 0x67, 0x61, 0x6e, 0x46, 0x78, 0x5c, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x42, 0x6f, 0x78, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5c, 0x6e, 0x65, 0x77, 0x73, 0x5c, 0x61, 0x70, 0x69, 0x5c, 0x67, 0x72, 0x70, 0x63, 0x5c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_protos_news_proto_rawDescOnce sync.Once
	file_api_grpc_protos_news_proto_rawDescData = file_api_grpc_protos_news_proto_rawDesc
)

func file_api_grpc_protos_news_proto_rawDescGZIP() []byte {
	file_api_grpc_protos_news_proto_rawDescOnce.Do(func() {
		file_api_grpc_protos_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_protos_news_proto_rawDescData)
	})
	return file_api_grpc_protos_news_proto_rawDescData
}

var file_api_grpc_protos_news_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_grpc_protos_news_proto_goTypes = []interface{}{
	(*NewsSendRequest)(nil),       // 0: protos.NewsSendRequest
	(*NewsSendResponse)(nil),      // 1: protos.NewsSendResponse
	(*NewsCreateRequest)(nil),     // 2: protos.NewsCreateRequest
	(*NewsCreateResponse)(nil),    // 3: protos.NewsCreateResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_api_grpc_protos_news_proto_depIdxs = []int32{
	4, // 0: protos.NewsSendResponse.Expires:type_name -> google.protobuf.Timestamp
	4, // 1: protos.NewsSendResponse.Created:type_name -> google.protobuf.Timestamp
	4, // 2: protos.NewsCreateRequest.Expires:type_name -> google.protobuf.Timestamp
	0, // 3: protos.NewsService.SendNews:input_type -> protos.NewsSendRequest
	2, // 4: protos.NewsService.CreateNews:input_type -> protos.NewsCreateRequest
	1, // 5: protos.NewsService.SendNews:output_type -> protos.NewsSendResponse
	3, // 6: protos.NewsService.CreateNews:output_type -> protos.NewsCreateResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_grpc_protos_news_proto_init() }
func file_api_grpc_protos_news_proto_init() {
	if File_api_grpc_protos_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_protos_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsSendRequest); i {
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
		file_api_grpc_protos_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsSendResponse); i {
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
		file_api_grpc_protos_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCreateRequest); i {
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
		file_api_grpc_protos_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsCreateResponse); i {
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
			RawDescriptor: file_api_grpc_protos_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_protos_news_proto_goTypes,
		DependencyIndexes: file_api_grpc_protos_news_proto_depIdxs,
		MessageInfos:      file_api_grpc_protos_news_proto_msgTypes,
	}.Build()
	File_api_grpc_protos_news_proto = out.File
	file_api_grpc_protos_news_proto_rawDesc = nil
	file_api_grpc_protos_news_proto_goTypes = nil
	file_api_grpc_protos_news_proto_depIdxs = nil
}
