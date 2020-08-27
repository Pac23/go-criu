// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rlimit.proto

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

type RlimitEntry struct {
	Cur                  *uint64  `protobuf:"varint,1,req,name=cur" json:"cur,omitempty"`
	Max                  *uint64  `protobuf:"varint,2,req,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RlimitEntry) Reset()         { *m = RlimitEntry{} }
func (m *RlimitEntry) String() string { return proto.CompactTextString(m) }
func (*RlimitEntry) ProtoMessage()    {}
func (*RlimitEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_68ead0126a3792e7, []int{0}
}

func (m *RlimitEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RlimitEntry.Unmarshal(m, b)
}
func (m *RlimitEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RlimitEntry.Marshal(b, m, deterministic)
}
func (m *RlimitEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RlimitEntry.Merge(m, src)
}
func (m *RlimitEntry) XXX_Size() int {
	return xxx_messageInfo_RlimitEntry.Size(m)
}
func (m *RlimitEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RlimitEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RlimitEntry proto.InternalMessageInfo

func (m *RlimitEntry) GetCur() uint64 {
	if m != nil && m.Cur != nil {
		return *m.Cur
	}
	return 0
}

func (m *RlimitEntry) GetMax() uint64 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

func init() {
	proto.RegisterType((*RlimitEntry)(nil), "rlimit_entry")
}

func init() {
	proto.RegisterFile("rlimit.proto", fileDescriptor_68ead0126a3792e7)
}

var fileDescriptor_68ead0126a3792e7 = []byte{
	// 89 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xca, 0xc9, 0xcc,
	0xcd, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0x82, 0xf1, 0xe3, 0x53, 0xf3, 0x4a,
	0x8a, 0x2a, 0x85, 0x04, 0xb8, 0x98, 0x93, 0x4b, 0x8b, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x82,
	0x40, 0x4c, 0x90, 0x48, 0x6e, 0x62, 0x85, 0x04, 0x13, 0x44, 0x04, 0xc8, 0x74, 0x12, 0x88, 0xe2,
	0xd3, 0x07, 0xeb, 0x4e, 0xca, 0xcc, 0x4b, 0xc9, 0xcc, 0x4b, 0x2f, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x25, 0xba, 0xbf, 0x34, 0x54, 0x00, 0x00, 0x00,
}
