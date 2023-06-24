// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
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

	Name string           `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Meta *structpb.Struct `protobuf:"bytes,2,opt,name=Meta,json=meta,proto3" json:"Meta,omitempty"`
	Data *SecretValue     `protobuf:"bytes,3,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
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

func (x *Secret) GetData() *SecretValue {
	if x != nil {
		return x.Data
	}
	return nil
}

type SecretValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*SecretValue_AuthData
	//	*SecretValue_Text
	//	*SecretValue_Any
	//	*SecretValue_CreditCardData
	Value isSecretValue_Value `protobuf_oneof:"Value"`
}

func (x *SecretValue) Reset() {
	*x = SecretValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretValue) ProtoMessage() {}

func (x *SecretValue) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SecretValue.ProtoReflect.Descriptor instead.
func (*SecretValue) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{1}
}

func (m *SecretValue) GetValue() isSecretValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SecretValue) GetAuthData() *AuthenticationData {
	if x, ok := x.GetValue().(*SecretValue_AuthData); ok {
		return x.AuthData
	}
	return nil
}

func (x *SecretValue) GetText() string {
	if x, ok := x.GetValue().(*SecretValue_Text); ok {
		return x.Text
	}
	return ""
}

func (x *SecretValue) GetAny() *anypb.Any {
	if x, ok := x.GetValue().(*SecretValue_Any); ok {
		return x.Any
	}
	return nil
}

func (x *SecretValue) GetCreditCardData() *CreditCardData {
	if x, ok := x.GetValue().(*SecretValue_CreditCardData); ok {
		return x.CreditCardData
	}
	return nil
}

type isSecretValue_Value interface {
	isSecretValue_Value()
}

type SecretValue_AuthData struct {
	AuthData *AuthenticationData `protobuf:"bytes,4,opt,name=AuthData,json=auth_data,proto3,oneof"`
}

type SecretValue_Text struct {
	Text string `protobuf:"bytes,5,opt,name=Text,json=text,proto3,oneof"`
}

type SecretValue_Any struct {
	Any *anypb.Any `protobuf:"bytes,6,opt,name=Any,json=any,proto3,oneof"`
}

type SecretValue_CreditCardData struct {
	CreditCardData *CreditCardData `protobuf:"bytes,7,opt,name=CreditCardData,json=credit_card_data,proto3,oneof"`
}

func (*SecretValue_AuthData) isSecretValue_Value() {}

func (*SecretValue_Text) isSecretValue_Value() {}

func (*SecretValue_Any) isSecretValue_Value() {}

func (*SecretValue_CreditCardData) isSecretValue_Value() {}

type AuthenticationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login       string `protobuf:"bytes,1,opt,name=Login,json=login,proto3" json:"Login,omitempty"`
	RawPassword string `protobuf:"bytes,2,opt,name=RawPassword,json=raw_password,proto3" json:"RawPassword,omitempty"`
}

func (x *AuthenticationData) Reset() {
	*x = AuthenticationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_data_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationData) ProtoMessage() {}

func (x *AuthenticationData) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AuthenticationData.ProtoReflect.Descriptor instead.
func (*AuthenticationData) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationData) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthenticationData) GetRawPassword() string {
	if x != nil {
		return x.RawPassword
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
		mi := &file_proto_data_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditCardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditCardData) ProtoMessage() {}

func (x *CreditCardData) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreditCardData.ProtoReflect.Descriptor instead.
func (*CreditCardData) Descriptor() ([]byte, []int) {
	return file_proto_data_secret_proto_rawDescGZIP(), []int{3}
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

var File_proto_data_secret_proto protoreflect.FileDescriptor

var file_proto_data_secret_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x06, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd5,
	0x01, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39,
	0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x28, 0x0a, 0x03, 0x41, 0x6e, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x48, 0x00, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x10, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x07, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x52, 0x61, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x06, 0x43, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x52, 0x03, 0x43, 0x76, 0x76,
	0x52, 0x03, 0x43, 0x56, 0x56, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x73, 0x75, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x70,
	0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*SecretValue)(nil),        // 1: secret.SecretValue
	(*AuthenticationData)(nil), // 2: secret.AuthenticationData
	(*CreditCardData)(nil),     // 3: secret.CreditCardData
	(*structpb.Struct)(nil),    // 4: google.protobuf.Struct
	(*anypb.Any)(nil),          // 5: google.protobuf.Any
}
var file_proto_data_secret_proto_depIdxs = []int32{
	4, // 0: secret.Secret.Meta:type_name -> google.protobuf.Struct
	1, // 1: secret.Secret.Data:type_name -> secret.SecretValue
	2, // 2: secret.SecretValue.AuthData:type_name -> secret.AuthenticationData
	5, // 3: secret.SecretValue.Any:type_name -> google.protobuf.Any
	3, // 4: secret.SecretValue.CreditCardData:type_name -> secret.CreditCardData
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
			switch v := v.(*SecretValue); i {
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
		file_proto_data_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_proto_data_secret_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SecretValue_AuthData)(nil),
		(*SecretValue_Text)(nil),
		(*SecretValue_Any)(nil),
		(*SecretValue_CreditCardData)(nil),
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
