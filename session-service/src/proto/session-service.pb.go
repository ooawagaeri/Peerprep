// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: session-service.proto

package proto

import (
	_ "cs3219-project-ay2223s1-g33/gateway"
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

type ValidateTokenErrorCode int32

const (
	ValidateTokenErrorCode_VALIDATE_TOKEN_NO_ERROR       ValidateTokenErrorCode = 0
	ValidateTokenErrorCode_VALIDATE_TOKEN_ERROR_INVALID  ValidateTokenErrorCode = 1
	ValidateTokenErrorCode_VALIDATE_TOKEN_ERROR_EXPIRED  ValidateTokenErrorCode = 2
	ValidateTokenErrorCode_VALIDATE_TOKEN_ERROR_INTERNAL ValidateTokenErrorCode = 100
)

// Enum value maps for ValidateTokenErrorCode.
var (
	ValidateTokenErrorCode_name = map[int32]string{
		0:   "VALIDATE_TOKEN_NO_ERROR",
		1:   "VALIDATE_TOKEN_ERROR_INVALID",
		2:   "VALIDATE_TOKEN_ERROR_EXPIRED",
		100: "VALIDATE_TOKEN_ERROR_INTERNAL",
	}
	ValidateTokenErrorCode_value = map[string]int32{
		"VALIDATE_TOKEN_NO_ERROR":       0,
		"VALIDATE_TOKEN_ERROR_INVALID":  1,
		"VALIDATE_TOKEN_ERROR_EXPIRED":  2,
		"VALIDATE_TOKEN_ERROR_INTERNAL": 100,
	}
)

func (x ValidateTokenErrorCode) Enum() *ValidateTokenErrorCode {
	p := new(ValidateTokenErrorCode)
	*p = x
	return p
}

func (x ValidateTokenErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidateTokenErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_session_service_proto_enumTypes[0].Descriptor()
}

func (ValidateTokenErrorCode) Type() protoreflect.EnumType {
	return &file_session_service_proto_enumTypes[0]
}

func (x ValidateTokenErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidateTokenErrorCode.Descriptor instead.
func (ValidateTokenErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_session_service_proto_rawDescGZIP(), []int{0}
}

type AddBlacklistErrorCode int32

const (
	AddBlacklistErrorCode_ADD_BLACKLIST_NO_ERROR       AddBlacklistErrorCode = 0
	AddBlacklistErrorCode_ADD_BLACKLIST_ERROR_INTERNAL AddBlacklistErrorCode = 100
)

// Enum value maps for AddBlacklistErrorCode.
var (
	AddBlacklistErrorCode_name = map[int32]string{
		0:   "ADD_BLACKLIST_NO_ERROR",
		100: "ADD_BLACKLIST_ERROR_INTERNAL",
	}
	AddBlacklistErrorCode_value = map[string]int32{
		"ADD_BLACKLIST_NO_ERROR":       0,
		"ADD_BLACKLIST_ERROR_INTERNAL": 100,
	}
)

func (x AddBlacklistErrorCode) Enum() *AddBlacklistErrorCode {
	p := new(AddBlacklistErrorCode)
	*p = x
	return p
}

func (x AddBlacklistErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddBlacklistErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_session_service_proto_enumTypes[1].Descriptor()
}

func (AddBlacklistErrorCode) Type() protoreflect.EnumType {
	return &file_session_service_proto_enumTypes[1]
}

func (x AddBlacklistErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddBlacklistErrorCode.Descriptor instead.
func (AddBlacklistErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_session_service_proto_rawDescGZIP(), []int{1}
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_session_service_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ErrorCode ValidateTokenErrorCode `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=session_service.ValidateTokenErrorCode" json:"error_code,omitempty"`
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_session_service_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateTokenResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ValidateTokenResponse) GetErrorCode() ValidateTokenErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ValidateTokenErrorCode_VALIDATE_TOKEN_NO_ERROR
}

type AddBlacklistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AddBlacklistRequest) Reset() {
	*x = AddBlacklistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlacklistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlacklistRequest) ProtoMessage() {}

func (x *AddBlacklistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlacklistRequest.ProtoReflect.Descriptor instead.
func (*AddBlacklistRequest) Descriptor() ([]byte, []int) {
	return file_session_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddBlacklistRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddBlacklistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode AddBlacklistErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=session_service.AddBlacklistErrorCode" json:"error_code,omitempty"`
}

func (x *AddBlacklistResponse) Reset() {
	*x = AddBlacklistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlacklistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlacklistResponse) ProtoMessage() {}

func (x *AddBlacklistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlacklistResponse.ProtoReflect.Descriptor instead.
func (*AddBlacklistResponse) Descriptor() ([]byte, []int) {
	return file_session_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddBlacklistResponse) GetErrorCode() AddBlacklistErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return AddBlacklistErrorCode_ADD_BLACKLIST_NO_ERROR
}

var File_session_service_proto protoreflect.FileDescriptor

var file_session_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x75, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x46, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x64,
	0x64, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x9c, 0x01, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x20,
	0x0a, 0x1c, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x01,
	0x12, 0x20, 0x0a, 0x1c, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x10, 0x64, 0x2a, 0x55, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x16, 0x41, 0x44, 0x44, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x4c, 0x49, 0x53, 0x54, 0x5f,
	0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44,
	0x44, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x64, 0x32, 0xcd, 0x01, 0x0a,
	0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5e, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x25, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5b, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21,
	0x63, 0x73, 0x33, 0x32, 0x31, 0x39, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x61,
	0x79, 0x32, 0x32, 0x32, 0x33, 0x73, 0x31, 0x2d, 0x67, 0x33, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_service_proto_rawDescOnce sync.Once
	file_session_service_proto_rawDescData = file_session_service_proto_rawDesc
)

func file_session_service_proto_rawDescGZIP() []byte {
	file_session_service_proto_rawDescOnce.Do(func() {
		file_session_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_service_proto_rawDescData)
	})
	return file_session_service_proto_rawDescData
}

var file_session_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_session_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_session_service_proto_goTypes = []interface{}{
	(ValidateTokenErrorCode)(0),   // 0: session_service.ValidateTokenErrorCode
	(AddBlacklistErrorCode)(0),    // 1: session_service.AddBlacklistErrorCode
	(*ValidateTokenRequest)(nil),  // 2: session_service.ValidateTokenRequest
	(*ValidateTokenResponse)(nil), // 3: session_service.ValidateTokenResponse
	(*AddBlacklistRequest)(nil),   // 4: session_service.AddBlacklistRequest
	(*AddBlacklistResponse)(nil),  // 5: session_service.AddBlacklistResponse
}
var file_session_service_proto_depIdxs = []int32{
	0, // 0: session_service.ValidateTokenResponse.error_code:type_name -> session_service.ValidateTokenErrorCode
	1, // 1: session_service.AddBlacklistResponse.error_code:type_name -> session_service.AddBlacklistErrorCode
	2, // 2: session_service.SessionService.ValidateToken:input_type -> session_service.ValidateTokenRequest
	4, // 3: session_service.SessionService.AddBlacklist:input_type -> session_service.AddBlacklistRequest
	3, // 4: session_service.SessionService.ValidateToken:output_type -> session_service.ValidateTokenResponse
	5, // 5: session_service.SessionService.AddBlacklist:output_type -> session_service.AddBlacklistResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_session_service_proto_init() }
func file_session_service_proto_init() {
	if File_session_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_session_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
		file_session_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenResponse); i {
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
		file_session_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlacklistRequest); i {
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
		file_session_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlacklistResponse); i {
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
			RawDescriptor: file_session_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_session_service_proto_goTypes,
		DependencyIndexes: file_session_service_proto_depIdxs,
		EnumInfos:         file_session_service_proto_enumTypes,
		MessageInfos:      file_session_service_proto_msgTypes,
	}.Build()
	File_session_service_proto = out.File
	file_session_service_proto_rawDesc = nil
	file_session_service_proto_goTypes = nil
	file_session_service_proto_depIdxs = nil
}
