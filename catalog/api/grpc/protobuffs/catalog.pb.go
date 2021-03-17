// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: api/grpc/protos/catalog.proto

package protobuffs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ProductSendRequest) Reset() {
	*x = ProductSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSendRequest) ProtoMessage() {}

func (x *ProductSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSendRequest.ProtoReflect.Descriptor instead.
func (*ProductSendRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *ProductSendRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Category    string  `protobuf:"bytes,3,opt,name=Category,proto3" json:"Category,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float32 `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ProductSendResponse) Reset() {
	*x = ProductSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSendResponse) ProtoMessage() {}

func (x *ProductSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSendResponse.ProtoReflect.Descriptor instead.
func (*ProductSendResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSendResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductSendResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductSendResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductSendResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductSendResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_api_grpc_protos_catalog_proto protoreflect.FileDescriptor

var file_api_grpc_protos_catalog_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x8f,
	0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x32, 0x5a, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e,
	0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x42, 0x6f, 0x78, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c,
	0x61, 0x70, 0x69, 0x5c, 0x67, 0x72, 0x70, 0x63, 0x5c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x66, 0x73, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_protos_catalog_proto_rawDescOnce sync.Once
	file_api_grpc_protos_catalog_proto_rawDescData = file_api_grpc_protos_catalog_proto_rawDesc
)

func file_api_grpc_protos_catalog_proto_rawDescGZIP() []byte {
	file_api_grpc_protos_catalog_proto_rawDescOnce.Do(func() {
		file_api_grpc_protos_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_protos_catalog_proto_rawDescData)
	})
	return file_api_grpc_protos_catalog_proto_rawDescData
}

var file_api_grpc_protos_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_protos_catalog_proto_goTypes = []interface{}{
	(*ProductSendRequest)(nil),  // 0: protos.ProductSendRequest
	(*ProductSendResponse)(nil), // 1: protos.ProductSendResponse
}
var file_api_grpc_protos_catalog_proto_depIdxs = []int32{
	0, // 0: protos.CatalogService.SendProduct:input_type -> protos.ProductSendRequest
	1, // 1: protos.CatalogService.SendProduct:output_type -> protos.ProductSendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_protos_catalog_proto_init() }
func file_api_grpc_protos_catalog_proto_init() {
	if File_api_grpc_protos_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_protos_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSendRequest); i {
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
		file_api_grpc_protos_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSendResponse); i {
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
			RawDescriptor: file_api_grpc_protos_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_protos_catalog_proto_goTypes,
		DependencyIndexes: file_api_grpc_protos_catalog_proto_depIdxs,
		MessageInfos:      file_api_grpc_protos_catalog_proto_msgTypes,
	}.Build()
	File_api_grpc_protos_catalog_proto = out.File
	file_api_grpc_protos_catalog_proto_rawDesc = nil
	file_api_grpc_protos_catalog_proto_goTypes = nil
	file_api_grpc_protos_catalog_proto_depIdxs = nil
}
