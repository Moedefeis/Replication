// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: grpc/auction.proto

package grpc

import (
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

type OperationId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OperationId) Reset() {
	*x = OperationId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationId) ProtoMessage() {}

func (x *OperationId) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationId.ProtoReflect.Descriptor instead.
func (*OperationId) Descriptor() ([]byte, []int) {
	return file_grpc_auction_proto_rawDescGZIP(), []int{0}
}

func (x *OperationId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32        `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Opid   *OperationId `protobuf:"bytes,2,opt,name=opid,proto3" json:"opid,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_grpc_auction_proto_rawDescGZIP(), []int{1}
}

func (x *Amount) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Amount) GetOpid() *OperationId {
	if x != nil {
		return x.Opid
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_auction_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ServerId) Reset() {
	*x = ServerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerId) ProtoMessage() {}

func (x *ServerId) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerId.ProtoReflect.Descriptor instead.
func (*ServerId) Descriptor() ([]byte, []int) {
	return file_grpc_auction_proto_rawDescGZIP(), []int{3}
}

func (x *ServerId) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_grpc_auction_proto_rawDescGZIP(), []int{4}
}

var File_grpc_auction_proto protoreflect.FileDescriptor

var file_grpc_auction_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04,
	0x6f, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x04, 0x6f,
	0x70, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x1e, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0xc0, 0x01, 0x0a, 0x07, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x0c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x72, 0x61,
	0x73, 0x68, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x0a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x32, 0x4d, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x61, 0x64,
	0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x6f, 0x65, 0x64, 0x65, 0x66, 0x65, 0x69, 0x73, 0x2f, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_auction_proto_rawDescOnce sync.Once
	file_grpc_auction_proto_rawDescData = file_grpc_auction_proto_rawDesc
)

func file_grpc_auction_proto_rawDescGZIP() []byte {
	file_grpc_auction_proto_rawDescOnce.Do(func() {
		file_grpc_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_auction_proto_rawDescData)
	})
	return file_grpc_auction_proto_rawDescData
}

var file_grpc_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_auction_proto_goTypes = []interface{}{
	(*OperationId)(nil),           // 0: grpc.OperationId
	(*Amount)(nil),                // 1: grpc.Amount
	(*Response)(nil),              // 2: grpc.Response
	(*ServerId)(nil),              // 3: grpc.ServerId
	(*Void)(nil),                  // 4: grpc.Void
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_grpc_auction_proto_depIdxs = []int32{
	0, // 0: grpc.Amount.opid:type_name -> grpc.OperationId
	1, // 1: grpc.auction.bid:input_type -> grpc.Amount
	0, // 2: grpc.auction.result:input_type -> grpc.OperationId
	3, // 3: grpc.auction.crashed:input_type -> grpc.ServerId
	5, // 4: grpc.auction.startAuction:input_type -> google.protobuf.Timestamp
	0, // 5: grpc.serverNode.addOperationToExecutionOrder:input_type -> grpc.OperationId
	2, // 6: grpc.auction.bid:output_type -> grpc.Response
	1, // 7: grpc.auction.result:output_type -> grpc.Amount
	4, // 8: grpc.auction.crashed:output_type -> grpc.Void
	4, // 9: grpc.auction.startAuction:output_type -> grpc.Void
	4, // 10: grpc.serverNode.addOperationToExecutionOrder:output_type -> grpc.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_auction_proto_init() }
func file_grpc_auction_proto_init() {
	if File_grpc_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationId); i {
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
		file_grpc_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_grpc_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grpc_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerId); i {
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
		file_grpc_auction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_grpc_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_auction_proto_goTypes,
		DependencyIndexes: file_grpc_auction_proto_depIdxs,
		MessageInfos:      file_grpc_auction_proto_msgTypes,
	}.Build()
	File_grpc_auction_proto = out.File
	file_grpc_auction_proto_rawDesc = nil
	file_grpc_auction_proto_goTypes = nil
	file_grpc_auction_proto_depIdxs = nil
}
