// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.11.2
// source: payment.proto

package payment

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

type CreditCardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCardNumber          string `protobuf:"bytes,1,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	CreditCardCvv             int32  `protobuf:"varint,2,opt,name=credit_card_cvv,json=creditCardCvv,proto3" json:"credit_card_cvv,omitempty"`
	CreditCardExpirationYear  int32  `protobuf:"varint,3,opt,name=credit_card_expiration_year,json=creditCardExpirationYear,proto3" json:"credit_card_expiration_year,omitempty"`
	CreditCardExpirationMonth int32  `protobuf:"varint,4,opt,name=credit_card_expiration_month,json=creditCardExpirationMonth,proto3" json:"credit_card_expiration_month,omitempty"`
}

func (x *CreditCardInfo) Reset() {
	*x = CreditCardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditCardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCardInfo) ProtoMessage() {}

func (x *CreditCardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditCardInfo.ProtoReflect.Descriptor instead.
func (*CreditCardInfo) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreditCardInfo) GetCreditCardNumber() string {
	if x != nil {
		return x.CreditCardNumber
	}
	return ""
}

func (x *CreditCardInfo) GetCreditCardCvv() int32 {
	if x != nil {
		return x.CreditCardCvv
	}
	return 0
}

func (x *CreditCardInfo) GetCreditCardExpirationYear() int32 {
	if x != nil {
		return x.CreditCardExpirationYear
	}
	return 0
}

func (x *CreditCardInfo) GetCreditCardExpirationMonth() int32 {
	if x != nil {
		return x.CreditCardExpirationMonth
	}
	return 0
}

type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Units        int64  `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	Nanos        int32  `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type ChargeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount     *Money          `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CreditCard *CreditCardInfo `protobuf:"bytes,2,opt,name=credit_card,json=creditCard,proto3" json:"credit_card,omitempty"`
}

func (x *ChargeRequest) Reset() {
	*x = ChargeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeRequest) ProtoMessage() {}

func (x *ChargeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeRequest.ProtoReflect.Descriptor instead.
func (*ChargeRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *ChargeRequest) GetAmount() *Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *ChargeRequest) GetCreditCard() *CreditCardInfo {
	if x != nil {
		return x.CreditCard
	}
	return nil
}

type ChargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *ChargeResponse) Reset() {
	*x = ChargeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeResponse) ProtoMessage() {}

func (x *ChargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeResponse.ProtoReflect.Descriptor instead.
func (*ChargeResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *ChargeResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x76, 0x76, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x43, 0x76,
	0x76, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72,
	0x12, 0x3f, 0x0a, 0x1c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x71, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x22, 0x37,
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0x4d, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_payment_proto_goTypes = []interface{}{
	(*CreditCardInfo)(nil), // 0: payment.CreditCardInfo
	(*Money)(nil),          // 1: payment.Money
	(*ChargeRequest)(nil),  // 2: payment.ChargeRequest
	(*ChargeResponse)(nil), // 3: payment.ChargeResponse
}
var file_payment_proto_depIdxs = []int32{
	1, // 0: payment.ChargeRequest.amount:type_name -> payment.Money
	0, // 1: payment.ChargeRequest.credit_card:type_name -> payment.CreditCardInfo
	2, // 2: payment.PaymentService.Charge:input_type -> payment.ChargeRequest
	3, // 3: payment.PaymentService.Charge:output_type -> payment.ChargeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditCardInfo); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
		file_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeRequest); i {
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
		file_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeResponse); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}
