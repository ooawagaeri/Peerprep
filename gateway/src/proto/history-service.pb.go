// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: history-service.proto

package proto

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

type GetAttemptHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset     int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	QuestionId uint64 `protobuf:"varint,3,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
}

func (x *GetAttemptHistoryRequest) Reset() {
	*x = GetAttemptHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttemptHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttemptHistoryRequest) ProtoMessage() {}

func (x *GetAttemptHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttemptHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetAttemptHistoryRequest) Descriptor() ([]byte, []int) {
	return file_history_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAttemptHistoryRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAttemptHistoryRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAttemptHistoryRequest) GetQuestionId() uint64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

type GetAttemptHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attempts     []*HistoryAttempt `protobuf:"bytes,1,rep,name=attempts,proto3" json:"attempts,omitempty"`
	TotalCount   int32             `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	ErrorMessage string            `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *GetAttemptHistoryResponse) Reset() {
	*x = GetAttemptHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttemptHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttemptHistoryResponse) ProtoMessage() {}

func (x *GetAttemptHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttemptHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetAttemptHistoryResponse) Descriptor() ([]byte, []int) {
	return file_history_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAttemptHistoryResponse) GetAttempts() []*HistoryAttempt {
	if x != nil {
		return x.Attempts
	}
	return nil
}

func (x *GetAttemptHistoryResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetAttemptHistoryResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetAttemptSubmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttemptId uint64 `protobuf:"varint,1,opt,name=attempt_id,json=attemptId,proto3" json:"attempt_id,omitempty"`
}

func (x *GetAttemptSubmissionRequest) Reset() {
	*x = GetAttemptSubmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttemptSubmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttemptSubmissionRequest) ProtoMessage() {}

func (x *GetAttemptSubmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttemptSubmissionRequest.ProtoReflect.Descriptor instead.
func (*GetAttemptSubmissionRequest) Descriptor() ([]byte, []int) {
	return file_history_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAttemptSubmissionRequest) GetAttemptId() uint64 {
	if x != nil {
		return x.AttemptId
	}
	return 0
}

type GetAttemptSubmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attempt      *HistoryAttempt `protobuf:"bytes,1,opt,name=attempt,proto3" json:"attempt,omitempty"`
	ErrorMessage string          `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *GetAttemptSubmissionResponse) Reset() {
	*x = GetAttemptSubmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttemptSubmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttemptSubmissionResponse) ProtoMessage() {}

func (x *GetAttemptSubmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttemptSubmissionResponse.ProtoReflect.Descriptor instead.
func (*GetAttemptSubmissionResponse) Descriptor() ([]byte, []int) {
	return file_history_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAttemptSubmissionResponse) GetAttempt() *HistoryAttempt {
	if x != nil {
		return x.Attempt
	}
	return nil
}

func (x *GetAttemptSubmissionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type CreateCompletionSubmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completed *HistoryCompletion `protobuf:"bytes,1,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *CreateCompletionSubmissionRequest) Reset() {
	*x = CreateCompletionSubmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompletionSubmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompletionSubmissionRequest) ProtoMessage() {}

func (x *CreateCompletionSubmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompletionSubmissionRequest.ProtoReflect.Descriptor instead.
func (*CreateCompletionSubmissionRequest) Descriptor() ([]byte, []int) {
	return file_history_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCompletionSubmissionRequest) GetCompleted() *HistoryCompletion {
	if x != nil {
		return x.Completed
	}
	return nil
}

type CreateCompletionSubmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completed    *HistoryCompletion `protobuf:"bytes,1,opt,name=completed,proto3" json:"completed,omitempty"`
	ErrorMessage string             `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CreateCompletionSubmissionResponse) Reset() {
	*x = CreateCompletionSubmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompletionSubmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompletionSubmissionResponse) ProtoMessage() {}

func (x *CreateCompletionSubmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompletionSubmissionResponse.ProtoReflect.Descriptor instead.
func (*CreateCompletionSubmissionResponse) Descriptor() ([]byte, []int) {
	return file_history_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCompletionSubmissionResponse) GetCompleted() *HistoryCompletion {
	if x != nil {
		return x.Completed
	}
	return nil
}

func (x *CreateCompletionSubmissionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_history_service_proto protoreflect.FileDescriptor

var file_history_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x95, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x52,
	0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a,
	0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x22,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xf9, 0x02, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x73, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21,
	0x63, 0x73, 0x33, 0x32, 0x31, 0x39, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x61,
	0x79, 0x32, 0x32, 0x32, 0x33, 0x73, 0x31, 0x2d, 0x67, 0x33, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_history_service_proto_rawDescOnce sync.Once
	file_history_service_proto_rawDescData = file_history_service_proto_rawDesc
)

func file_history_service_proto_rawDescGZIP() []byte {
	file_history_service_proto_rawDescOnce.Do(func() {
		file_history_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_history_service_proto_rawDescData)
	})
	return file_history_service_proto_rawDescData
}

var file_history_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_history_service_proto_goTypes = []interface{}{
	(*GetAttemptHistoryRequest)(nil),           // 0: history_service.GetAttemptHistoryRequest
	(*GetAttemptHistoryResponse)(nil),          // 1: history_service.GetAttemptHistoryResponse
	(*GetAttemptSubmissionRequest)(nil),        // 2: history_service.GetAttemptSubmissionRequest
	(*GetAttemptSubmissionResponse)(nil),       // 3: history_service.GetAttemptSubmissionResponse
	(*CreateCompletionSubmissionRequest)(nil),  // 4: history_service.CreateCompletionSubmissionRequest
	(*CreateCompletionSubmissionResponse)(nil), // 5: history_service.CreateCompletionSubmissionResponse
	(*HistoryAttempt)(nil),                     // 6: common.HistoryAttempt
	(*HistoryCompletion)(nil),                  // 7: common.HistoryCompletion
}
var file_history_service_proto_depIdxs = []int32{
	6, // 0: history_service.GetAttemptHistoryResponse.attempts:type_name -> common.HistoryAttempt
	6, // 1: history_service.GetAttemptSubmissionResponse.attempt:type_name -> common.HistoryAttempt
	7, // 2: history_service.CreateCompletionSubmissionRequest.completed:type_name -> common.HistoryCompletion
	7, // 3: history_service.CreateCompletionSubmissionResponse.completed:type_name -> common.HistoryCompletion
	0, // 4: history_service.HistoryService.GetAttemptHistory:input_type -> history_service.GetAttemptHistoryRequest
	2, // 5: history_service.HistoryService.GetAttemptSubmission:input_type -> history_service.GetAttemptSubmissionRequest
	4, // 6: history_service.HistoryService.CreateCompletionSubmission:input_type -> history_service.CreateCompletionSubmissionRequest
	1, // 7: history_service.HistoryService.GetAttemptHistory:output_type -> history_service.GetAttemptHistoryResponse
	3, // 8: history_service.HistoryService.GetAttemptSubmission:output_type -> history_service.GetAttemptSubmissionResponse
	5, // 9: history_service.HistoryService.CreateCompletionSubmission:output_type -> history_service.CreateCompletionSubmissionResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_history_service_proto_init() }
func file_history_service_proto_init() {
	if File_history_service_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_history_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttemptHistoryRequest); i {
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
		file_history_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttemptHistoryResponse); i {
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
		file_history_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttemptSubmissionRequest); i {
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
		file_history_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttemptSubmissionResponse); i {
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
		file_history_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompletionSubmissionRequest); i {
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
		file_history_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompletionSubmissionResponse); i {
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
			RawDescriptor: file_history_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_history_service_proto_goTypes,
		DependencyIndexes: file_history_service_proto_depIdxs,
		MessageInfos:      file_history_service_proto_msgTypes,
	}.Build()
	File_history_service_proto = out.File
	file_history_service_proto_rawDesc = nil
	file_history_service_proto_goTypes = nil
	file_history_service_proto_depIdxs = nil
}
