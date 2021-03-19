// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: github.com/openbank/openbank/v1/atm/all.proto

package atm

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ATM holds all details about an ATM.
type ATM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique identifier of an ATM.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
	// BankID is the identifier of the bank associated with the ATM.
	BankID string `protobuf:"bytes,2,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	// Name is the name of ATM.
	Name string `protobuf:"bytes,3,opt,name=Name,json=name,proto3" json:"name,omitempty"`
	// Address is the ATM's address.
	Address *types.Address `protobuf:"bytes,4,opt,name=Address,json=address,proto3" json:"address,omitempty"`
	// Location is the ATM longitude and latitude.
	Location *types.Location `protobuf:"bytes,5,opt,name=Location,json=location,proto3" json:"location,omitempty"`
	// Description is the ATM's description.
	Description string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"description,omitempty"`
	// Metadata is the ATM's metadata.
	Metadata string `protobuf:"bytes,7,opt,name=Metadata,json=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ATM) Reset() {
	*x = ATM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ATM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ATM) ProtoMessage() {}

func (x *ATM) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ATM.ProtoReflect.Descriptor instead.
func (*ATM) Descriptor() ([]byte, []int) {
	return file_github_com_openbank_openbank_v1_atm_all_proto_rawDescGZIP(), []int{0}
}

func (x *ATM) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ATM) GetBankID() string {
	if x != nil {
		return x.BankID
	}
	return ""
}

func (x *ATM) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ATM) GetAddress() *types.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ATM) GetLocation() *types.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ATM) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ATM) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

// CreateATMRequest is a request to create an ATM.
type CreateATMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BankID is the bank identifier that owned the ATM
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	// BankID is the identifier of the bank associated with the ATM.
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"name,omitempty"`
	// Address is the ATM's address.
	Address *types.Address `protobuf:"bytes,3,opt,name=Address,json=3,proto3" json:"3,omitempty"`
	// Location is the ATM's longitude and latitude
	Location *types.Location `protobuf:"bytes,4,opt,name=Location,json=location,proto3" json:"location,omitempty"`
	// Description is the ATM's description.
	Description string `protobuf:"bytes,5,opt,name=Description,json=description,proto3" json:"description,omitempty"`
	// Metadata is the ATM's metadata.
	Metadata string `protobuf:"bytes,6,opt,name=Metadata,json=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CreateATMRequest) Reset() {
	*x = CreateATMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateATMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateATMRequest) ProtoMessage() {}

func (x *CreateATMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateATMRequest.ProtoReflect.Descriptor instead.
func (*CreateATMRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openbank_openbank_v1_atm_all_proto_rawDescGZIP(), []int{1}
}

func (x *CreateATMRequest) GetBankID() string {
	if x != nil {
		return x.BankID
	}
	return ""
}

func (x *CreateATMRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateATMRequest) GetAddress() *types.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CreateATMRequest) GetLocation() *types.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *CreateATMRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateATMRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

// CreateATMResponse is the response envelope for creating an ATM.
type CreateATMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ATM_ID is the unique identifier of the ATM.
	ATM_ID string `protobuf:"bytes,1,opt,name=ATM_ID,json=atm_id,proto3" json:"atm_id,omitempty"`
}

func (x *CreateATMResponse) Reset() {
	*x = CreateATMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateATMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateATMResponse) ProtoMessage() {}

func (x *CreateATMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateATMResponse.ProtoReflect.Descriptor instead.
func (*CreateATMResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openbank_openbank_v1_atm_all_proto_rawDescGZIP(), []int{2}
}

func (x *CreateATMResponse) GetATM_ID() string {
	if x != nil {
		return x.ATM_ID
	}
	return ""
}

// GetATMRequest is the request envelope for retrieving a specific ATM's information.
type GetATMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ATM_ID is a unique identifier of the ATM.
	ATM_ID string `protobuf:"bytes,1,opt,name=ATM_ID,json=atm_id,proto3" json:"atm_id,omitempty"`
}

func (x *GetATMRequest) Reset() {
	*x = GetATMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetATMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetATMRequest) ProtoMessage() {}

func (x *GetATMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetATMRequest.ProtoReflect.Descriptor instead.
func (*GetATMRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openbank_openbank_v1_atm_all_proto_rawDescGZIP(), []int{3}
}

func (x *GetATMRequest) GetATM_ID() string {
	if x != nil {
		return x.ATM_ID
	}
	return ""
}

// GetATMsResponse is the response envelope for retrieving ATM information.
type GetATMsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result is the list of the ATMs.
	Result []*ATM `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"result,omitempty"`
}

func (x *GetATMsResponse) Reset() {
	*x = GetATMsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetATMsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetATMsResponse) ProtoMessage() {}

func (x *GetATMsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetATMsResponse.ProtoReflect.Descriptor instead.
func (*GetATMsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openbank_openbank_v1_atm_all_proto_rawDescGZIP(), []int{4}
}

func (x *GetATMsResponse) GetResult() []*ATM {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_github_com_openbank_openbank_v1_atm_all_proto protoreflect.FileDescriptor

var file_github_com_openbank_openbank_v1_atm_all_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x74, 0x6d, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x74, 0x6d, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb3, 0x02, 0x0a, 0x03, 0x41, 0x54, 0x4d, 0x12, 0x1a, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x52, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x37, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x9e, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x54, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x42,
	0x61, 0x6e, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x01, 0x33,
	0x12, 0x37, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a,
	0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x54, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06,
	0x41, 0x54, 0x4d, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x06, 0x61, 0x74, 0x6d, 0x5f, 0x69, 0x64,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x54, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x41, 0x54, 0x4d,
	0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x06, 0x61, 0x74, 0x6d, 0x5f, 0x69, 0x64, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x54, 0x4d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x74, 0x6d, 0x2e, 0x41,
	0x54, 0x4d, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32, 0xe9,
	0x07, 0x0a, 0x0a, 0x41, 0x54, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd6, 0x02,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x54, 0x4d, 0x12, 0x12, 0x2e, 0x61, 0x74, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x54, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x61,
	0x74, 0x6d, 0x2e, 0x41, 0x54, 0x4d, 0x22, 0xa9, 0x02, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92,
	0x41, 0x86, 0x02, 0x0a, 0x03, 0x41, 0x54, 0x4d, 0x12, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x54, 0x4d, 0x1a, 0x4c, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x41, 0x54, 0x4d, 0x2c, 0x20, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x64, 0x20, 0x49, 0x44, 0x2e, 0x4a, 0x41, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x3a,
	0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e,
	0x12, 0x18, 0x0a, 0x16, 0x1a, 0x14, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x74, 0x6d, 0x41, 0x54, 0x4d, 0x4a, 0x31, 0x0a, 0x03, 0x34, 0x30,
	0x34, 0x12, 0x2a, 0x0a, 0x28, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20,
	0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x62, 0x2a, 0x0a,
	0x28, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x1e, 0x0a, 0x1c, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x62, 0x6e, 0x6b, 0x2e, 0x74, 0x6f,
	0x2f, 0x61, 0x74, 0x6d, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12,
	0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x6d, 0x73, 0x2f, 0x7b, 0x41, 0x54, 0x4d, 0x5f, 0x49,
	0x44, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xdf, 0x02, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x54,
	0x4d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x74, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x54, 0x4d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xa3, 0x02, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92, 0x41, 0x89, 0x02, 0x0a, 0x03, 0x41,
	0x54, 0x4d, 0x12, 0x1b, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x6c, 0x6c,
	0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x41, 0x54, 0x4d, 0x73, 0x1a,
	0x37, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x20, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x41, 0x54, 0x4d, 0x73, 0x2e, 0x4a, 0x4d, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12,
	0x46, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79,
	0x2e, 0x12, 0x24, 0x0a, 0x22, 0x1a, 0x20, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x74, 0x6d, 0x47, 0x65, 0x74, 0x41, 0x54, 0x4d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x31, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x2a,
	0x0a, 0x28, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x62, 0x2a, 0x0a, 0x28, 0x0a, 0x06,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x1e, 0x0a, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x62, 0x6e, 0x6b, 0x2e, 0x74, 0x6f, 0x2f, 0x61, 0x74,
	0x6d, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x74, 0x6d, 0x73, 0x30, 0x00, 0x12, 0x9a, 0x02, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x54, 0x4d, 0x12, 0x15, 0x2e, 0x61, 0x74, 0x6d, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x54, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x61, 0x74, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x54, 0x4d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd9, 0x01, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92, 0x41,
	0xbc, 0x01, 0x0a, 0x03, 0x41, 0x54, 0x4d, 0x12, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x6e, 0x20, 0x41, 0x54, 0x4d, 0x1a, 0x25, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20,
	0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x41, 0x54, 0x4d, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x69, 0x74, 0x73, 0x20, 0x69, 0x64, 0x2e, 0x4a, 0x52, 0x0a,
	0x03, 0x32, 0x30, 0x31, 0x12, 0x4b, 0x0a, 0x21, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x12, 0x26, 0x0a, 0x24, 0x1a, 0x22, 0x23,
	0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x74, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x54, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x62, 0x2b, 0x0a, 0x29, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x1f, 0x0a,
	0x1d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x62, 0x6e,
	0x6b, 0x2e, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x6d, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x6d,
	0x73, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0xd6, 0x06, 0x0a, 0x0e, 0x74,
	0x6f, 0x2e, 0x62, 0x6e, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x74, 0x6d, 0x48, 0x01, 0x50,
	0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x74, 0x6d, 0x3b, 0x61, 0x74, 0x6d, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00,
	0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x92,
	0x41, 0x80, 0x06, 0x0a, 0x03, 0x32, 0x2e, 0x30, 0x12, 0x4a, 0x0a, 0x07, 0x41, 0x54, 0x4d, 0x20,
	0x41, 0x50, 0x49, 0x12, 0x38, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x61, 0x64, 0x20, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x41, 0x54, 0x4d, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x32, 0x05, 0x31,
	0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x5f, 0x0a, 0x03, 0x34,
	0x30, 0x30, 0x12, 0x58, 0x0a, 0x56, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20,
	0x62, 0x6f, 0x64, 0x79, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x74, 0x65, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x4c, 0x0a, 0x03,
	0x34, 0x30, 0x31, 0x12, 0x45, 0x0a, 0x43, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30,
	0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x31, 0x0a, 0x03,
	0x34, 0x30, 0x34, 0x12, 0x2a, 0x0a, 0x28, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x52,
	0x46, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x3f, 0x0a, 0x3d, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x6e, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x6f, 0x63, 0x63, 0x75,
	0x72, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x5a, 0x89, 0x02, 0x0a, 0x86, 0x02, 0x0a, 0x06, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0xfb, 0x01, 0x08, 0x03, 0x12, 0x7d, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x41, 0x54, 0x4d, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x43,
	0x6f, 0x64, 0x65, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x28, 0x04, 0x32, 0x08, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x61, 0x0a, 0x2d, 0x0a, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x62, 0x6e, 0x6b, 0x2e, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x6d, 0x2e, 0x72, 0x65,
	0x61, 0x64, 0x12, 0x0d, 0x56, 0x69, 0x65, 0x77, 0x20, 0x61, 0x74, 0x6d, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x0a, 0x30, 0x0a, 0x1d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x62, 0x6e, 0x6b, 0x2e, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x6d, 0x2e, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x0f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x20, 0x61, 0x74, 0x6d, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_openbank_openbank_v1_atm_all_proto_rawDescOnce sync.Once
	file_github_com_openbank_openbank_v1_atm_all_proto_rawDescData = file_github_com_openbank_openbank_v1_atm_all_proto_rawDesc
)

func file_github_com_openbank_openbank_v1_atm_all_proto_rawDescGZIP() []byte {
	file_github_com_openbank_openbank_v1_atm_all_proto_rawDescOnce.Do(func() {
		file_github_com_openbank_openbank_v1_atm_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_openbank_openbank_v1_atm_all_proto_rawDescData)
	})
	return file_github_com_openbank_openbank_v1_atm_all_proto_rawDescData
}

var (
	file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
	file_github_com_openbank_openbank_v1_atm_all_proto_goTypes  = []interface{}{
		(*ATM)(nil),               // 0: atm.ATM
		(*CreateATMRequest)(nil),  // 1: atm.CreateATMRequest
		(*CreateATMResponse)(nil), // 2: atm.CreateATMResponse
		(*GetATMRequest)(nil),     // 3: atm.GetATMRequest
		(*GetATMsResponse)(nil),   // 4: atm.GetATMsResponse
		(*types.Address)(nil),     // 5: types.Address
		(*types.Location)(nil),    // 6: types.Location
		(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
	}
)

var file_github_com_openbank_openbank_v1_atm_all_proto_depIdxs = []int32{
	5, // 0: atm.ATM.Address:type_name -> types.Address
	6, // 1: atm.ATM.Location:type_name -> types.Location
	5, // 2: atm.CreateATMRequest.Address:type_name -> types.Address
	6, // 3: atm.CreateATMRequest.Location:type_name -> types.Location
	0, // 4: atm.GetATMsResponse.Result:type_name -> atm.ATM
	3, // 5: atm.ATMService.GetATM:input_type -> atm.GetATMRequest
	7, // 6: atm.ATMService.GetATMs:input_type -> google.protobuf.Empty
	1, // 7: atm.ATMService.CreateATM:input_type -> atm.CreateATMRequest
	0, // 8: atm.ATMService.GetATM:output_type -> atm.ATM
	4, // 9: atm.ATMService.GetATMs:output_type -> atm.GetATMsResponse
	2, // 10: atm.ATMService.CreateATM:output_type -> atm.CreateATMResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_openbank_openbank_v1_atm_all_proto_init() }
func file_github_com_openbank_openbank_v1_atm_all_proto_init() {
	if File_github_com_openbank_openbank_v1_atm_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ATM); i {
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
		file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateATMRequest); i {
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
		file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateATMResponse); i {
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
		file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetATMRequest); i {
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
		file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetATMsResponse); i {
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
			RawDescriptor: file_github_com_openbank_openbank_v1_atm_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_openbank_openbank_v1_atm_all_proto_goTypes,
		DependencyIndexes: file_github_com_openbank_openbank_v1_atm_all_proto_depIdxs,
		MessageInfos:      file_github_com_openbank_openbank_v1_atm_all_proto_msgTypes,
	}.Build()
	File_github_com_openbank_openbank_v1_atm_all_proto = out.File
	file_github_com_openbank_openbank_v1_atm_all_proto_rawDesc = nil
	file_github_com_openbank_openbank_v1_atm_all_proto_goTypes = nil
	file_github_com_openbank_openbank_v1_atm_all_proto_depIdxs = nil
}
