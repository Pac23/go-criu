// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pstree.proto

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

type PstreeEntry struct {
	Pid                  *uint32  `protobuf:"varint,1,req,name=pid" json:"pid,omitempty"`
	Ppid                 *uint32  `protobuf:"varint,2,req,name=ppid" json:"ppid,omitempty"`
	Pgid                 *uint32  `protobuf:"varint,3,req,name=pgid" json:"pgid,omitempty"`
	Sid                  *uint32  `protobuf:"varint,4,req,name=sid" json:"sid,omitempty"`
	Threads              []uint32 `protobuf:"varint,5,rep,name=threads" json:"threads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PstreeEntry) Reset()         { *m = PstreeEntry{} }
func (m *PstreeEntry) String() string { return proto.CompactTextString(m) }
func (*PstreeEntry) ProtoMessage()    {}
func (*PstreeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9197f030fa60d0c2, []int{0}
}

func (m *PstreeEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PstreeEntry.Unmarshal(m, b)
}
func (m *PstreeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PstreeEntry.Marshal(b, m, deterministic)
}
func (m *PstreeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PstreeEntry.Merge(m, src)
}
func (m *PstreeEntry) XXX_Size() int {
	return xxx_messageInfo_PstreeEntry.Size(m)
}
func (m *PstreeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PstreeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PstreeEntry proto.InternalMessageInfo

func (m *PstreeEntry) GetPid() uint32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *PstreeEntry) GetPpid() uint32 {
	if m != nil && m.Ppid != nil {
		return *m.Ppid
	}
	return 0
}

func (m *PstreeEntry) GetPgid() uint32 {
	if m != nil && m.Pgid != nil {
		return *m.Pgid
	}
	return 0
}

func (m *PstreeEntry) GetSid() uint32 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *PstreeEntry) GetThreads() []uint32 {
	if m != nil {
		return m.Threads
	}
	return nil
}

func init() {
	proto.RegisterType((*PstreeEntry)(nil), "pstree_entry")
}

func init() {
	proto.RegisterFile("pstree.proto", fileDescriptor_9197f030fa60d0c2)
}

var fileDescriptor_9197f030fa60d0c2 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0x2e, 0x29,
	0x4a, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x2a, 0x81, 0xf1, 0xe3, 0x53, 0xf3, 0x4a,
	0x8a, 0x2a, 0x85, 0x04, 0xb8, 0x98, 0x0b, 0x32, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34, 0x78, 0x83,
	0x40, 0x4c, 0x21, 0x21, 0x2e, 0x96, 0x02, 0x90, 0x10, 0x13, 0x58, 0x08, 0xcc, 0x06, 0x8b, 0xa5,
	0x03, 0xc5, 0x98, 0xa1, 0x62, 0x40, 0x36, 0x48, 0x67, 0x31, 0x50, 0x88, 0x05, 0xa2, 0x13, 0xc8,
	0x14, 0x92, 0xe0, 0x62, 0x2f, 0xc9, 0x28, 0x4a, 0x4d, 0x4c, 0x29, 0x96, 0x60, 0x55, 0x60, 0x06,
	0x8a, 0xc2, 0xb8, 0x4e, 0x02, 0x51, 0x7c, 0xfa, 0x60, 0xfb, 0x93, 0x32, 0xf3, 0x52, 0x32, 0xf3,
	0xd2, 0x8b, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xe6, 0x9a, 0xf1, 0x96, 0x00, 0x00, 0x00,
}