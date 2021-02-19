// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/apierror/all.proto

package apierror

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ErrorType defines the type of the error.
type ErrorType int32

const (
	ErrorType_UnknownErrorType ErrorType = 0
	// ErrorType_AccountAlreadyExists is returned when trying to create a duplicate account.
	ErrorType_AccountAlreadyExists ErrorType = 1
)

var ErrorType_name = map[int32]string{
	0: "UnknownErrorType",
	1: "AccountAlreadyExists",
}

var ErrorType_value = map[string]int32{
	"UnknownErrorType":     0,
	"AccountAlreadyExists": 1,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}

func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47a3f2c2a170d0d1, []int{0}
}

// APIError defines the error model returned by an API in case something went wrong
type APIError struct {
	// Code of the error returned, corresponding to gRPC or HTTP code, depending on your client.
	Code int32 `protobuf:"varint,1,opt,name=Code,json=code,proto3" json:"code,omitempty"`
	// Type is an enum to help handling errors programatically.
	Type ErrorType `protobuf:"varint,2,opt,name=Type,json=type,proto3,enum=apierror.ErrorType" json:"type,omitempty"`
	// Message is a human readable message providing more details about the error.
	Message string `protobuf:"bytes,3,opt,name=Message,json=message,proto3" json:"message,omitempty"`
	// DocURL is a direct link to the specific error type documentation, when applicable.
	DocURL               string   `protobuf:"bytes,4,opt,name=DocURL,json=doc_url,proto3" json:"doc_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIError) Reset()         { *m = APIError{} }
func (m *APIError) String() string { return proto.CompactTextString(m) }
func (*APIError) ProtoMessage()    {}
func (*APIError) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a3f2c2a170d0d1, []int{0}
}

func (m *APIError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIError.Unmarshal(m, b)
}

func (m *APIError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIError.Marshal(b, m, deterministic)
}

func (m *APIError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIError.Merge(m, src)
}

func (m *APIError) XXX_Size() int {
	return xxx_messageInfo_APIError.Size(m)
}

func (m *APIError) XXX_DiscardUnknown() {
	xxx_messageInfo_APIError.DiscardUnknown(m)
}

var xxx_messageInfo_APIError proto.InternalMessageInfo

func (m *APIError) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *APIError) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_UnknownErrorType
}

func (m *APIError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *APIError) GetDocURL() string {
	if m != nil {
		return m.DocURL
	}
	return ""
}

func init() {
	proto.RegisterEnum("apierror.ErrorType", ErrorType_name, ErrorType_value)
	proto.RegisterType((*APIError)(nil), "apierror.APIError")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/apierror/all.proto", fileDescriptor_47a3f2c2a170d0d1)
}

var fileDescriptor_47a3f2c2a170d0d1 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x3f, 0x4b, 0xc3, 0x40,
	0x14, 0x7f, 0x2f, 0xc6, 0x36, 0xbd, 0x41, 0xc2, 0xd5, 0xe1, 0x28, 0x22, 0x45, 0x1d, 0x82, 0x43,
	0xaa, 0xed, 0x22, 0x6e, 0xad, 0x16, 0x15, 0x14, 0x4a, 0xb1, 0x8b, 0x8b, 0x24, 0x97, 0xa3, 0x96,
	0xa6, 0xf7, 0xc2, 0xe5, 0xaa, 0x76, 0x73, 0x74, 0xf4, 0xb3, 0x38, 0xf9, 0x31, 0x1c, 0xfd, 0x28,
	0x8e, 0x92, 0x4a, 0x23, 0xc5, 0xed, 0xbd, 0xdf, 0x9f, 0xf7, 0xf8, 0xfd, 0x58, 0x7b, 0x3c, 0xb1,
	0x0f, 0xf3, 0x38, 0x94, 0x34, 0x6b, 0x51, 0xa6, 0x74, 0x1c, 0xe9, 0xe9, 0xdf, 0xf0, 0x78, 0xdc,
	0x8a, 0xb2, 0x89, 0x32, 0x86, 0x4c, 0x2b, 0x4a, 0xd3, 0x30, 0x33, 0x64, 0x89, 0x7b, 0x2b, 0x6c,
	0xef, 0x1d, 0x99, 0xd7, 0x1d, 0x5c, 0xf5, 0x8b, 0x85, 0xef, 0x32, 0xf7, 0x8c, 0x12, 0x25, 0xb0,
	0x89, 0xc1, 0x66, 0x8f, 0x79, 0x20, 0x20, 0x80, 0x23, 0x18, 0xc0, 0xd0, 0x95, 0x94, 0x28, 0xde,
	0x61, 0xee, 0xed, 0x22, 0x53, 0xc2, 0x69, 0x62, 0xb0, 0xd5, 0xae, 0x87, 0xab, 0x2b, 0xe1, 0xd2,
	0x5e, 0x50, 0xeb, 0x26, 0xbb, 0xc8, 0x14, 0x3f, 0x60, 0xd5, 0x1b, 0x95, 0xe7, 0xd1, 0x58, 0x89,
	0x8d, 0x26, 0x06, 0xb5, 0x35, 0x49, 0x75, 0xf6, 0x4b, 0xf1, 0x7d, 0x56, 0x39, 0x27, 0x39, 0x1a,
	0x5e, 0x0b, 0xf7, 0xbf, 0x28, 0x21, 0x79, 0x3f, 0x37, 0xe9, 0x69, 0xc5, 0x03, 0x1f, 0x04, 0x1c,
	0x5e, 0xb0, 0x5a, 0xf9, 0x91, 0x0b, 0xe6, 0x8f, 0xf4, 0x54, 0xd3, 0x93, 0x2e, 0x31, 0x1f, 0x1a,
	0x8e, 0x07, 0x7c, 0x87, 0x6d, 0x77, 0xa5, 0xa4, 0xb9, 0xb6, 0xdd, 0xd4, 0xa8, 0x28, 0x59, 0xf4,
	0x9f, 0x27, 0xb9, 0xcd, 0x7d, 0x2c, 0xd8, 0x86, 0x23, 0xa0, 0x77, 0xc2, 0xea, 0x96, 0xc2, 0x58,
	0x4f, 0x8b, 0x28, 0x65, 0x9c, 0x4b, 0x1c, 0xe0, 0x5d, 0x59, 0xd1, 0x0b, 0xc2, 0x2b, 0xc2, 0x1b,
	0xc2, 0x07, 0xc2, 0x17, 0xc2, 0x37, 0xe2, 0xa7, 0x03, 0x71, 0x65, 0x59, 0x64, 0xe7, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0x80, 0xe6, 0xe1, 0x7e, 0x01, 0x00, 0x00,
}
