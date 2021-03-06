// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/response/v1/helloresponse.proto

package helloresponse

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A response message containing a greeting
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_91a3254c15d42d03, []int{0}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReply)(nil), "protos.response.v1.HelloReply")
}

func init() {
	proto.RegisterFile("protos/response/v1/helloresponse.proto", fileDescriptor_91a3254c15d42d03)
}

var fileDescriptor_91a3254c15d42d03 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2f, 0x33, 0xd4, 0xcf,
	0x48, 0xcd, 0xc9, 0xc9, 0x87, 0x09, 0xe8, 0x81, 0x15, 0x08, 0x09, 0x41, 0xd4, 0xe9, 0xc1, 0x85,
	0xcb, 0x0c, 0x95, 0xd4, 0xb8, 0xb8, 0x3c, 0x40, 0x4a, 0x83, 0x52, 0x0b, 0x72, 0x2a, 0x85, 0x24,
	0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x60, 0x5c, 0x27, 0xe7, 0x28, 0xc7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xec, 0xc4, 0xcc, 0x92, 0x7c, 0x23, 0xfd, 0xe2, 0xc4, 0xdc, 0x82, 0x9c, 0x54, 0xb0, 0xa9,
	0xfa, 0xe9, 0xf9, 0xba, 0x98, 0xce, 0xb0, 0x46, 0x71, 0x46, 0x12, 0x1b, 0x58, 0x85, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x32, 0x22, 0x2b, 0x41, 0xb1, 0x00, 0x00, 0x00,
}
