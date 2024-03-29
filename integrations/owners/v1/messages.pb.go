// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: integrations/owners/v1/messages.proto

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

type CreateOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateOwnerRequest) Reset() {
	*x = CreateOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_owners_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerRequest) ProtoMessage() {}

func (x *CreateOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_owners_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerRequest.ProtoReflect.Descriptor instead.
func (*CreateOwnerRequest) Descriptor() ([]byte, []int) {
	return file_integrations_owners_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOwnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Owner  *v1.Owner  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *CreateOwnerResponse) Reset() {
	*x = CreateOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_owners_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerResponse) ProtoMessage() {}

func (x *CreateOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_owners_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerResponse.ProtoReflect.Descriptor instead.
func (*CreateOwnerResponse) Descriptor() ([]byte, []int) {
	return file_integrations_owners_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOwnerResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CreateOwnerResponse) GetOwner() *v1.Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type GetOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOwnerRequest) Reset() {
	*x = GetOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_owners_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerRequest) ProtoMessage() {}

func (x *GetOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_owners_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerRequest.ProtoReflect.Descriptor instead.
func (*GetOwnerRequest) Descriptor() ([]byte, []int) {
	return file_integrations_owners_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Owner  *v1.Owner  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *GetOwnerResponse) Reset() {
	*x = GetOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_owners_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerResponse) ProtoMessage() {}

func (x *GetOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_owners_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerResponse.ProtoReflect.Descriptor instead.
func (*GetOwnerResponse) Descriptor() ([]byte, []int) {
	return file_integrations_owners_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetOwnerResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *GetOwnerResponse) GetOwner() *v1.Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

var File_integrations_owners_v1_messages_proto protoreflect.FileDescriptor

var file_integrations_owners_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x01, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x8a, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64, 0x62,
	0x69, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0xca, 0x02, 0x24, 0x49, 0x44, 0x42, 0x49, 0x5c, 0x50, 0x48, 0x50, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5c, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x5c, 0x76, 0x31, 0xe2, 0x02, 0x2d, 0x49, 0x44, 0x42,
	0x49, 0x5c, 0x50, 0x48, 0x50, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x5c, 0x76,
	0x31, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_integrations_owners_v1_messages_proto_rawDescOnce sync.Once
	file_integrations_owners_v1_messages_proto_rawDescData = file_integrations_owners_v1_messages_proto_rawDesc
)

func file_integrations_owners_v1_messages_proto_rawDescGZIP() []byte {
	file_integrations_owners_v1_messages_proto_rawDescOnce.Do(func() {
		file_integrations_owners_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_owners_v1_messages_proto_rawDescData)
	})
	return file_integrations_owners_v1_messages_proto_rawDescData
}

var file_integrations_owners_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_integrations_owners_v1_messages_proto_goTypes = []interface{}{
	(*CreateOwnerRequest)(nil),  // 0: integrations.owners.v1.CreateOwnerRequest
	(*CreateOwnerResponse)(nil), // 1: integrations.owners.v1.CreateOwnerResponse
	(*GetOwnerRequest)(nil),     // 2: integrations.owners.v1.GetOwnerRequest
	(*GetOwnerResponse)(nil),    // 3: integrations.owners.v1.GetOwnerResponse
	(*v1.Result)(nil),           // 4: integrations.entities.v1.Result
	(*v1.Owner)(nil),            // 5: integrations.entities.v1.Owner
}
var file_integrations_owners_v1_messages_proto_depIdxs = []int32{
	4, // 0: integrations.owners.v1.CreateOwnerResponse.result:type_name -> integrations.entities.v1.Result
	5, // 1: integrations.owners.v1.CreateOwnerResponse.owner:type_name -> integrations.entities.v1.Owner
	4, // 2: integrations.owners.v1.GetOwnerResponse.result:type_name -> integrations.entities.v1.Result
	5, // 3: integrations.owners.v1.GetOwnerResponse.owner:type_name -> integrations.entities.v1.Owner
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_integrations_owners_v1_messages_proto_init() }
func file_integrations_owners_v1_messages_proto_init() {
	if File_integrations_owners_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integrations_owners_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOwnerRequest); i {
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
		file_integrations_owners_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOwnerResponse); i {
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
		file_integrations_owners_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerRequest); i {
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
		file_integrations_owners_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerResponse); i {
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
			RawDescriptor: file_integrations_owners_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integrations_owners_v1_messages_proto_goTypes,
		DependencyIndexes: file_integrations_owners_v1_messages_proto_depIdxs,
		MessageInfos:      file_integrations_owners_v1_messages_proto_msgTypes,
	}.Build()
	File_integrations_owners_v1_messages_proto = out.File
	file_integrations_owners_v1_messages_proto_rawDesc = nil
	file_integrations_owners_v1_messages_proto_goTypes = nil
	file_integrations_owners_v1_messages_proto_depIdxs = nil
}
