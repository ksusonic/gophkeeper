// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: proto/data/secret.proto

package data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string           `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Meta       *structpb.Struct `protobuf:"bytes,2,opt,name=Meta,json=meta,proto3" json:"Meta,omitempty"`
	SecretData *Secret_Data     `protobuf:"bytes,3,opt,name=SecretData,json=secret_value,proto3" json:"SecretData,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Secret) GetMeta() *structpb.Struct {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Secret) GetSecretData() *Secret_Data {
	if x != nil {
		return x.SecretData
	}
	return nil
}

type AuthenticationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=Login,json=login,proto3" json:"Login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,json=password,proto3" json:"Password,omitempty"`
}

func (x *AuthenticationData) Reset() {
	*x = AuthenticationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationData) ProtoMessage() {}

func (x *AuthenticationData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationData.ProtoReflect.Descriptor instead.
func (*AuthenticationData) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationData) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthenticationData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreditCardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pan            string `protobuf:"bytes,1,opt,name=Pan,json=pan,proto3" json:"Pan,omitempty"`
	ChName         string `protobuf:"bytes,2,opt,name=ChName,json=ch_name,proto3" json:"ChName,omitempty"`
	ExpirationDate string `protobuf:"bytes,3,opt,name=ExpirationDate,json=expiration_date,proto3" json:"ExpirationDate,omitempty"`
}

func (x *CreditCardData) Reset() {
	*x = CreditCardData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditCardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCardData) ProtoMessage() {}

func (x *CreditCardData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditCardData.ProtoReflect.Descriptor instead.
func (*CreditCardData) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{2}
}

func (x *CreditCardData) GetPan() string {
	if x != nil {
		return x.Pan
	}
	return ""
}

func (x *CreditCardData) GetChName() string {
	if x != nil {
		return x.ChName
	}
	return ""
}

func (x *CreditCardData) GetExpirationDate() string {
	if x != nil {
		return x.ExpirationDate
	}
	return ""
}

type Secret_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*Secret_Data_Authentication
	//	*Secret_Data_Text
	//	*Secret_Data_Any
	//	*Secret_Data_CreditCardData
	Variant isSecret_Data_Variant `protobuf_oneof:"Variant"`
}

func (x *Secret_Data) Reset() {
	*x = Secret_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret_Data) ProtoMessage() {}

func (x *Secret_Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret_Data.ProtoReflect.Descriptor instead.
func (*Secret_Data) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Secret_Data) GetVariant() isSecret_Data_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *Secret_Data) GetAuthentication() *AuthenticationData {
	if x, ok := x.GetVariant().(*Secret_Data_Authentication); ok {
		return x.Authentication
	}
	return nil
}

func (x *Secret_Data) GetText() string {
	if x, ok := x.GetVariant().(*Secret_Data_Text); ok {
		return x.Text
	}
	return ""
}

func (x *Secret_Data) GetAny() *anypb.Any {
	if x, ok := x.GetVariant().(*Secret_Data_Any); ok {
		return x.Any
	}
	return nil
}

func (x *Secret_Data) GetCreditCardData() *CreditCardData {
	if x, ok := x.GetVariant().(*Secret_Data_CreditCardData); ok {
		return x.CreditCardData
	}
	return nil
}

type isSecret_Data_Variant interface {
	isSecret_Data_Variant()
}

type Secret_Data_Authentication struct {
	Authentication *AuthenticationData `protobuf:"bytes,11,opt,name=Authentication,json=authentication,proto3,oneof"`
}

type Secret_Data_Text struct {
	Text string `protobuf:"bytes,12,opt,name=Text,json=text,proto3,oneof"`
}

type Secret_Data_Any struct {
	Any *anypb.Any `protobuf:"bytes,13,opt,name=Any,json=any,proto3,oneof"`
}

type Secret_Data_CreditCardData struct {
	CreditCardData *CreditCardData `protobuf:"bytes,14,opt,name=CreditCardData,json=credit_card,proto3,oneof"`
}

func (*Secret_Data_Authentication) isSecret_Data_Variant() {}

func (*Secret_Data_Text) isSecret_Data_Variant() {}

func (*Secret_Data_Any) isSecret_Data_Variant() {}

func (*Secret_Data_CreditCardData) isSecret_Data_Variant() {}

var File_proto_data_secret_proto protoreflect.FileDescriptor

var file_proto_data_secret_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x06, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x4d, 0x65, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xd6, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x41, 0x6e, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x3d, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6e,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x61, 0x6e, 0x12, 0x17, 0x0a, 0x06, 0x43, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x03, 0x43, 0x76, 0x76, 0x52, 0x03, 0x43, 0x56, 0x56, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x73, 0x75,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_data_secret_proto_rawDescOnce sync.Once
	file_proto_data_secret_proto_rawDescData = file_proto_data_secret_proto_rawDesc
)

func file_proto_data_secret_proto_rawDescGZIP() []byte {
	file_proto_data_secret_proto_rawDescOnce.Do(func() {
		file_proto_data_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_data_secret_proto_rawDescData)
	})
	return file_proto_data_secret_proto_rawDescData
}

var file_proto_data_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_data_secret_proto_goTypes = []interface{}{
	(*Secret)(nil),             // 0: secret.Secret
	(*AuthenticationData)(nil), // 1: secret.AuthenticationData
	(*CreditCardData)(nil),     // 2: secret.CreditCardData
	(*Secret_Data)(nil),        // 3: secret.Secret.Data
	(*structpb.Struct)(nil),    // 4: google.protobuf.Struct
	(*anypb.Any)(nil),          // 5: google.protobuf.Any
}
var file_proto_data_secret_proto_depIdxs = []int32{
	4, // 0: secret.Secret.Meta:type_name -> google.protobuf.Struct
	3, // 1: secret.Secret.SecretData:type_name -> secret.Secret.Data
	1, // 2: secret.Secret.Data.Authentication:type_name -> secret.AuthenticationData
	5, // 3: secret.Secret.Data.Any:type_name -> google.protobuf.Any
	2, // 4: secret.Secret.Data.CreditCardData:type_name -> secret.CreditCardData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_data_secret_proto_init() }
func file_proto_data_secret_proto_init() {
	if File_proto_data_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_data_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_proto_data_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationData); i {
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
		file_proto_data_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditCardData); i {
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
		file_proto_data_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret_Data); i {
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
	file_proto_data_secret_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Secret_Data_Authentication)(nil),
		(*Secret_Data_Text)(nil),
		(*Secret_Data_Any)(nil),
		(*Secret_Data_CreditCardData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_data_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_data_secret_proto_goTypes,
		DependencyIndexes: file_proto_data_secret_proto_depIdxs,
		MessageInfos:      file_proto_data_secret_proto_msgTypes,
	}.Build()
	File_proto_data_secret_proto = out.File
	file_proto_data_secret_proto_rawDesc = nil
	file_proto_data_secret_proto_goTypes = nil
	file_proto_data_secret_proto_depIdxs = nil
}
