// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

package protobindings

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

type Timeval struct {
	TvSec                *uint64  `protobuf:"varint,1,req,name=tv_sec,json=tvSec" json:"tv_sec,omitempty"`
	TvUsec               *uint64  `protobuf:"varint,2,req,name=tv_usec,json=tvUsec" json:"tv_usec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timeval) Reset()         { *m = Timeval{} }
func (m *Timeval) String() string { return proto.CompactTextString(m) }
func (*Timeval) ProtoMessage()    {}
func (*Timeval) Descriptor() ([]byte, []int) {
	return fileDescriptor_49a92d779a28c7fd, []int{0}
}

func (m *Timeval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timeval.Unmarshal(m, b)
}
func (m *Timeval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timeval.Marshal(b, m, deterministic)
}
func (m *Timeval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timeval.Merge(m, src)
}
func (m *Timeval) XXX_Size() int {
	return xxx_messageInfo_Timeval.Size(m)
}
func (m *Timeval) XXX_DiscardUnknown() {
	xxx_messageInfo_Timeval.DiscardUnknown(m)
}

var xxx_messageInfo_Timeval proto.InternalMessageInfo

func (m *Timeval) GetTvSec() uint64 {
	if m != nil && m.TvSec != nil {
		return *m.TvSec
	}
	return 0
}

func (m *Timeval) GetTvUsec() uint64 {
	if m != nil && m.TvUsec != nil {
		return *m.TvUsec
	}
	return 0
}

func init() {
	proto.RegisterType((*Timeval)(nil), "timeval")
}

func init() {
	proto.RegisterFile("time.proto", fileDescriptor_49a92d779a28c7fd)
}

var fileDescriptor_49a92d779a28c7fd = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe4, 0x62, 0x07, 0xf1, 0xca, 0x12, 0x73, 0x84,
	0x44, 0xb9, 0xd8, 0x4a, 0xca, 0xe2, 0x8b, 0x53, 0x93, 0x25, 0x18, 0x15, 0x98, 0x34, 0x58, 0x82,
	0x58, 0x4b, 0xca, 0x82, 0x53, 0x93, 0x85, 0xc4, 0x81, 0x2a, 0xca, 0xe2, 0x4b, 0x41, 0xe2, 0x4c,
	0x60, 0x71, 0xa0, 0xaa, 0x50, 0x20, 0xcf, 0x49, 0x20, 0x8a, 0x4f, 0x1f, 0x6c, 0x48, 0x52, 0x66,
	0x5e, 0x4a, 0x66, 0x5e, 0x7a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x34, 0xa2, 0xf5, 0x7b, 0x59,
	0x00, 0x00, 0x00,
}
