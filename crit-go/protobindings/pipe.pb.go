// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipe.proto

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

type PipeEntry struct {
	Id                   *uint32    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	PipeId               *uint32    `protobuf:"varint,2,req,name=pipe_id,json=pipeId" json:"pipe_id,omitempty"`
	Flags                *uint32    `protobuf:"varint,3,req,name=flags" json:"flags,omitempty"`
	Fown                 *FownEntry `protobuf:"bytes,4,req,name=fown" json:"fown,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PipeEntry) Reset()         { *m = PipeEntry{} }
func (m *PipeEntry) String() string { return proto.CompactTextString(m) }
func (*PipeEntry) ProtoMessage()    {}
func (*PipeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{0}
}

func (m *PipeEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipeEntry.Unmarshal(m, b)
}
func (m *PipeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipeEntry.Marshal(b, m, deterministic)
}
func (m *PipeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipeEntry.Merge(m, src)
}
func (m *PipeEntry) XXX_Size() int {
	return xxx_messageInfo_PipeEntry.Size(m)
}
func (m *PipeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PipeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PipeEntry proto.InternalMessageInfo

func (m *PipeEntry) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PipeEntry) GetPipeId() uint32 {
	if m != nil && m.PipeId != nil {
		return *m.PipeId
	}
	return 0
}

func (m *PipeEntry) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *PipeEntry) GetFown() *FownEntry {
	if m != nil {
		return m.Fown
	}
	return nil
}

func init() {
	proto.RegisterType((*PipeEntry)(nil), "pipe_entry")
}

func init() {
	proto.RegisterFile("pipe.proto", fileDescriptor_d1979e1a5fdc07ed)
}

var fileDescriptor_d1979e1a5fdc07ed = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0x2c, 0x48,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xca, 0x2f, 0x28, 0x29, 0x86, 0xb1, 0xd3, 0xf2,
	0xcb, 0xf3, 0x20, 0x6c, 0xa5, 0x62, 0x88, 0xaa, 0xf8, 0xd4, 0xbc, 0x92, 0xa2, 0x4a, 0x21, 0x3e,
	0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x26, 0x0d, 0xde, 0x20, 0x20, 0x4b, 0x48, 0x9c, 0x8b,
	0x1d, 0x2c, 0x0b, 0x14, 0x64, 0x02, 0x0b, 0xb2, 0x81, 0xb8, 0x9e, 0x29, 0x42, 0xd2, 0x5c, 0xac,
	0x69, 0x39, 0x89, 0xe9, 0xc5, 0x12, 0xcc, 0x20, 0x61, 0x27, 0xd6, 0x4b, 0xf6, 0x4c, 0x1c, 0x8c,
	0x41, 0x10, 0x31, 0x21, 0x79, 0x2e, 0x16, 0x90, 0x0d, 0x12, 0x2c, 0x40, 0x39, 0x6e, 0x23, 0x6e,
	0x3d, 0x10, 0x07, 0x62, 0x41, 0x10, 0x58, 0xc2, 0x49, 0x20, 0x8a, 0x4f, 0x1f, 0x6c, 0x7d, 0x52,
	0x66, 0x5e, 0x4a, 0x66, 0x5e, 0x7a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x43, 0x4c, 0xd5,
	0xab, 0x00, 0x00, 0x00,
}
