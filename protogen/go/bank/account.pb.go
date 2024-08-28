// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/bank/type/account.proto

package bank

import (
	date "google.golang.org/genproto/googleapis/type/date"
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

type CurrentBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber string `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
}

func (x *CurrentBalanceRequest) Reset() {
	*x = CurrentBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_type_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBalanceRequest) ProtoMessage() {}

func (x *CurrentBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBalanceRequest.ProtoReflect.Descriptor instead.
func (*CurrentBalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_account_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentBalanceRequest) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

type CurrentBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      float64    `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CurrentDate *date.Date `protobuf:"bytes,2,opt,name=current_date,proto3" json:"current_date,omitempty"`
}

func (x *CurrentBalanceResponse) Reset() {
	*x = CurrentBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_type_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBalanceResponse) ProtoMessage() {}

func (x *CurrentBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBalanceResponse.ProtoReflect.Descriptor instead.
func (*CurrentBalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_account_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentBalanceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CurrentBalanceResponse) GetCurrentDate() *date.Date {
	if x != nil {
		return x.CurrentDate
	}
	return nil
}

var File_proto_bank_type_account_proto protoreflect.FileDescriptor

var file_proto_bank_type_account_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x67, 0x0a, 0x16, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x44, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6a, 0x61,
	0x72, 0x61, 0x6d, 0x61, 0x75, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x62, 0x61, 0x6e, 0x6b, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bank_type_account_proto_rawDescOnce sync.Once
	file_proto_bank_type_account_proto_rawDescData = file_proto_bank_type_account_proto_rawDesc
)

func file_proto_bank_type_account_proto_rawDescGZIP() []byte {
	file_proto_bank_type_account_proto_rawDescOnce.Do(func() {
		file_proto_bank_type_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bank_type_account_proto_rawDescData)
	})
	return file_proto_bank_type_account_proto_rawDescData
}

var file_proto_bank_type_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_bank_type_account_proto_goTypes = []any{
	(*CurrentBalanceRequest)(nil),  // 0: bank.CurrentBalanceRequest
	(*CurrentBalanceResponse)(nil), // 1: bank.CurrentBalanceResponse
	(*date.Date)(nil),              // 2: google.type.Date
}
var file_proto_bank_type_account_proto_depIdxs = []int32{
	2, // 0: bank.CurrentBalanceResponse.current_date:type_name -> google.type.Date
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_bank_type_account_proto_init() }
func file_proto_bank_type_account_proto_init() {
	if File_proto_bank_type_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bank_type_account_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentBalanceRequest); i {
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
		file_proto_bank_type_account_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentBalanceResponse); i {
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
			RawDescriptor: file_proto_bank_type_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bank_type_account_proto_goTypes,
		DependencyIndexes: file_proto_bank_type_account_proto_depIdxs,
		MessageInfos:      file_proto_bank_type_account_proto_msgTypes,
	}.Build()
	File_proto_bank_type_account_proto = out.File
	file_proto_bank_type_account_proto_rawDesc = nil
	file_proto_bank_type_account_proto_goTypes = nil
	file_proto_bank_type_account_proto_depIdxs = nil
}
