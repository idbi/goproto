// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: integrations/integrations/v1/messages.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/idbi/goproto/integrations/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetIntegrationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *GetIntegrationsRequest) Reset() {
	*x = GetIntegrationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIntegrationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIntegrationsRequest) ProtoMessage() {}

func (x *GetIntegrationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIntegrationsRequest.ProtoReflect.Descriptor instead.
func (*GetIntegrationsRequest) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetIntegrationsRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetIntegrationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       *v1.Result        `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Integrations []*v1.Integration `protobuf:"bytes,2,rep,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *GetIntegrationsResponse) Reset() {
	*x = GetIntegrationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIntegrationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIntegrationsResponse) ProtoMessage() {}

func (x *GetIntegrationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIntegrationsResponse.ProtoReflect.Descriptor instead.
func (*GetIntegrationsResponse) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetIntegrationsResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetIntegrationsResponse) GetIntegrations() []*v1.Integration {
	if x != nil {
		return x.Integrations
	}
	return nil
}

type RegisterIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	TypeId  string `protobuf:"bytes,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
}

func (x *RegisterIntegrationRequest) Reset() {
	*x = RegisterIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIntegrationRequest) ProtoMessage() {}

func (x *RegisterIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIntegrationRequest.ProtoReflect.Descriptor instead.
func (*RegisterIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterIntegrationRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *RegisterIntegrationRequest) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

type RegisterIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result      *v1.Result      `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Integration *v1.Integration `protobuf:"bytes,2,opt,name=integration,proto3" json:"integration,omitempty"`
}

func (x *RegisterIntegrationResponse) Reset() {
	*x = RegisterIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIntegrationResponse) ProtoMessage() {}

func (x *RegisterIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIntegrationResponse.ProtoReflect.Descriptor instead.
func (*RegisterIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterIntegrationResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *RegisterIntegrationResponse) GetIntegration() *v1.Integration {
	if x != nil {
		return x.Integration
	}
	return nil
}

type RegisterIntegrationPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string                    `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	Properties    []*v1.IntegrationProperty `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *RegisterIntegrationPropertiesRequest) Reset() {
	*x = RegisterIntegrationPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIntegrationPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIntegrationPropertiesRequest) ProtoMessage() {}

func (x *RegisterIntegrationPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIntegrationPropertiesRequest.ProtoReflect.Descriptor instead.
func (*RegisterIntegrationPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterIntegrationPropertiesRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *RegisterIntegrationPropertiesRequest) GetProperties() []*v1.IntegrationProperty {
	if x != nil {
		return x.Properties
	}
	return nil
}

type RegisterIntegrationPropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RegisterIntegrationPropertiesResponse) Reset() {
	*x = RegisterIntegrationPropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterIntegrationPropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterIntegrationPropertiesResponse) ProtoMessage() {}

func (x *RegisterIntegrationPropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterIntegrationPropertiesResponse.ProtoReflect.Descriptor instead.
func (*RegisterIntegrationPropertiesResponse) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterIntegrationPropertiesResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetIntegrationPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
}

func (x *GetIntegrationPropertiesRequest) Reset() {
	*x = GetIntegrationPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIntegrationPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIntegrationPropertiesRequest) ProtoMessage() {}

func (x *GetIntegrationPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIntegrationPropertiesRequest.ProtoReflect.Descriptor instead.
func (*GetIntegrationPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetIntegrationPropertiesRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

type GetIntegrationPropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     *v1.Result                `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Properties []*v1.IntegrationProperty `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *GetIntegrationPropertiesResponse) Reset() {
	*x = GetIntegrationPropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIntegrationPropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIntegrationPropertiesResponse) ProtoMessage() {}

func (x *GetIntegrationPropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIntegrationPropertiesResponse.ProtoReflect.Descriptor instead.
func (*GetIntegrationPropertiesResponse) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *GetIntegrationPropertiesResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetIntegrationPropertiesResponse) GetProperties() []*v1.IntegrationProperty {
	if x != nil {
		return x.Properties
	}
	return nil
}

type RemoveIntegrationPropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string   `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	PropertyIds   []string `protobuf:"bytes,2,rep,name=property_ids,json=propertyIds,proto3" json:"property_ids,omitempty"`
}

func (x *RemoveIntegrationPropertiesRequest) Reset() {
	*x = RemoveIntegrationPropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveIntegrationPropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveIntegrationPropertiesRequest) ProtoMessage() {}

func (x *RemoveIntegrationPropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveIntegrationPropertiesRequest.ProtoReflect.Descriptor instead.
func (*RemoveIntegrationPropertiesRequest) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveIntegrationPropertiesRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *RemoveIntegrationPropertiesRequest) GetPropertyIds() []string {
	if x != nil {
		return x.PropertyIds
	}
	return nil
}

type RemoveIntegrationPropertiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RemoveIntegrationPropertiesResponse) Reset() {
	*x = RemoveIntegrationPropertiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_integrations_v1_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveIntegrationPropertiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveIntegrationPropertiesResponse) ProtoMessage() {}

func (x *RemoveIntegrationPropertiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_integrations_v1_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveIntegrationPropertiesResponse.ProtoReflect.Descriptor instead.
func (*RemoveIntegrationPropertiesResponse) Descriptor() ([]byte, []int) {
	return file_integrations_integrations_v1_messages_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveIntegrationPropertiesResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_integrations_integrations_v1_messages_proto protoreflect.FileDescriptor

var file_integrations_integrations_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xff, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x49, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x64, 0x0a, 0x1a,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x0b,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa6, 0x01, 0x0a, 0x24, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x4d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x61,
	0x0a, 0x25, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x52, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x22, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x16, 0xfa, 0x42, 0x13, 0x92, 0x01, 0x10, 0x08, 0x01, 0x10, 0x64, 0x18, 0x01, 0x22,
	0x08, 0x72, 0x06, 0xd0, 0x01, 0x01, 0xb0, 0x01, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x49, 0x64, 0x73, 0x22, 0x5f, 0x0a, 0x23, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x9c, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64, 0x62, 0x69, 0x2f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0xca, 0x02, 0x2a, 0x49, 0x44, 0x42, 0x49, 0x5c, 0x50, 0x48, 0x50, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x76, 0x31,
	0xe2, 0x02, 0x33, 0x49, 0x44, 0x42, 0x49, 0x5c, 0x50, 0x48, 0x50, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x76, 0x31, 0x5c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integrations_integrations_v1_messages_proto_rawDescOnce sync.Once
	file_integrations_integrations_v1_messages_proto_rawDescData = file_integrations_integrations_v1_messages_proto_rawDesc
)

func file_integrations_integrations_v1_messages_proto_rawDescGZIP() []byte {
	file_integrations_integrations_v1_messages_proto_rawDescOnce.Do(func() {
		file_integrations_integrations_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_integrations_v1_messages_proto_rawDescData)
	})
	return file_integrations_integrations_v1_messages_proto_rawDescData
}

var file_integrations_integrations_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_integrations_integrations_v1_messages_proto_goTypes = []interface{}{
	(*GetIntegrationsRequest)(nil),                // 0: integrations.integrations.v1.GetIntegrationsRequest
	(*GetIntegrationsResponse)(nil),               // 1: integrations.integrations.v1.GetIntegrationsResponse
	(*RegisterIntegrationRequest)(nil),            // 2: integrations.integrations.v1.RegisterIntegrationRequest
	(*RegisterIntegrationResponse)(nil),           // 3: integrations.integrations.v1.RegisterIntegrationResponse
	(*RegisterIntegrationPropertiesRequest)(nil),  // 4: integrations.integrations.v1.RegisterIntegrationPropertiesRequest
	(*RegisterIntegrationPropertiesResponse)(nil), // 5: integrations.integrations.v1.RegisterIntegrationPropertiesResponse
	(*GetIntegrationPropertiesRequest)(nil),       // 6: integrations.integrations.v1.GetIntegrationPropertiesRequest
	(*GetIntegrationPropertiesResponse)(nil),      // 7: integrations.integrations.v1.GetIntegrationPropertiesResponse
	(*RemoveIntegrationPropertiesRequest)(nil),    // 8: integrations.integrations.v1.RemoveIntegrationPropertiesRequest
	(*RemoveIntegrationPropertiesResponse)(nil),   // 9: integrations.integrations.v1.RemoveIntegrationPropertiesResponse
	(*v1.Result)(nil),                             // 10: integrations.entities.v1.Result
	(*v1.Integration)(nil),                        // 11: integrations.entities.v1.Integration
	(*v1.IntegrationProperty)(nil),                // 12: integrations.entities.v1.IntegrationProperty
}
var file_integrations_integrations_v1_messages_proto_depIdxs = []int32{
	10, // 0: integrations.integrations.v1.GetIntegrationsResponse.result:type_name -> integrations.entities.v1.Result
	11, // 1: integrations.integrations.v1.GetIntegrationsResponse.integrations:type_name -> integrations.entities.v1.Integration
	10, // 2: integrations.integrations.v1.RegisterIntegrationResponse.result:type_name -> integrations.entities.v1.Result
	11, // 3: integrations.integrations.v1.RegisterIntegrationResponse.integration:type_name -> integrations.entities.v1.Integration
	12, // 4: integrations.integrations.v1.RegisterIntegrationPropertiesRequest.properties:type_name -> integrations.entities.v1.IntegrationProperty
	10, // 5: integrations.integrations.v1.RegisterIntegrationPropertiesResponse.result:type_name -> integrations.entities.v1.Result
	10, // 6: integrations.integrations.v1.GetIntegrationPropertiesResponse.result:type_name -> integrations.entities.v1.Result
	12, // 7: integrations.integrations.v1.GetIntegrationPropertiesResponse.properties:type_name -> integrations.entities.v1.IntegrationProperty
	10, // 8: integrations.integrations.v1.RemoveIntegrationPropertiesResponse.result:type_name -> integrations.entities.v1.Result
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_integrations_integrations_v1_messages_proto_init() }
func file_integrations_integrations_v1_messages_proto_init() {
	if File_integrations_integrations_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integrations_integrations_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIntegrationsRequest); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIntegrationsResponse); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIntegrationRequest); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIntegrationResponse); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIntegrationPropertiesRequest); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterIntegrationPropertiesResponse); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIntegrationPropertiesRequest); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIntegrationPropertiesResponse); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveIntegrationPropertiesRequest); i {
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
		file_integrations_integrations_v1_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveIntegrationPropertiesResponse); i {
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
			RawDescriptor: file_integrations_integrations_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integrations_integrations_v1_messages_proto_goTypes,
		DependencyIndexes: file_integrations_integrations_v1_messages_proto_depIdxs,
		MessageInfos:      file_integrations_integrations_v1_messages_proto_msgTypes,
	}.Build()
	File_integrations_integrations_v1_messages_proto = out.File
	file_integrations_integrations_v1_messages_proto_rawDesc = nil
	file_integrations_integrations_v1_messages_proto_goTypes = nil
	file_integrations_integrations_v1_messages_proto_depIdxs = nil
}
