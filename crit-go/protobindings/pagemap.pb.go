// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pagemap.proto

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

type PagemapHead struct {
	PagesId              *uint32  `protobuf:"varint,1,req,name=pages_id,json=pagesId" json:"pages_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagemapHead) Reset()         { *m = PagemapHead{} }
func (m *PagemapHead) String() string { return proto.CompactTextString(m) }
func (*PagemapHead) ProtoMessage()    {}
func (*PagemapHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_5521d8432b2813e0, []int{0}
}

func (m *PagemapHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagemapHead.Unmarshal(m, b)
}
func (m *PagemapHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagemapHead.Marshal(b, m, deterministic)
}
func (m *PagemapHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagemapHead.Merge(m, src)
}
func (m *PagemapHead) XXX_Size() int {
	return xxx_messageInfo_PagemapHead.Size(m)
}
func (m *PagemapHead) XXX_DiscardUnknown() {
	xxx_messageInfo_PagemapHead.DiscardUnknown(m)
}

var xxx_messageInfo_PagemapHead proto.InternalMessageInfo

func (m *PagemapHead) GetPagesId() uint32 {
	if m != nil && m.PagesId != nil {
		return *m.PagesId
	}
	return 0
}

type PagemapEntry struct {
	Vaddr                *uint64  `protobuf:"varint,1,req,name=vaddr" json:"vaddr,omitempty"`
	NrPages              *uint32  `protobuf:"varint,2,req,name=nr_pages,json=nrPages" json:"nr_pages,omitempty"`
	InParent             *bool    `protobuf:"varint,3,opt,name=in_parent,json=inParent" json:"in_parent,omitempty"`
	Flags                *uint32  `protobuf:"varint,4,opt,name=flags" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagemapEntry) Reset()         { *m = PagemapEntry{} }
func (m *PagemapEntry) String() string { return proto.CompactTextString(m) }
func (*PagemapEntry) ProtoMessage()    {}
func (*PagemapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5521d8432b2813e0, []int{1}
}

func (m *PagemapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagemapEntry.Unmarshal(m, b)
}
func (m *PagemapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagemapEntry.Marshal(b, m, deterministic)
}
func (m *PagemapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagemapEntry.Merge(m, src)
}
func (m *PagemapEntry) XXX_Size() int {
	return xxx_messageInfo_PagemapEntry.Size(m)
}
func (m *PagemapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PagemapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PagemapEntry proto.InternalMessageInfo

func (m *PagemapEntry) GetVaddr() uint64 {
	if m != nil && m.Vaddr != nil {
		return *m.Vaddr
	}
	return 0
}

func (m *PagemapEntry) GetNrPages() uint32 {
	if m != nil && m.NrPages != nil {
		return *m.NrPages
	}
	return 0
}

func (m *PagemapEntry) GetInParent() bool {
	if m != nil && m.InParent != nil {
		return *m.InParent
	}
	return false
}

func (m *PagemapEntry) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func init() {
	proto.RegisterType((*PagemapHead)(nil), "pagemap_head")
	proto.RegisterType((*PagemapEntry)(nil), "pagemap_entry")
}

func init() {
	proto.RegisterFile("pagemap.proto", fileDescriptor_5521d8432b2813e0)
}

var fileDescriptor_5521d8432b2813e0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0x4c, 0x4f,
	0xcd, 0x4d, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xca, 0x2f, 0x28, 0x29, 0x86,
	0xb0, 0x95, 0x34, 0xb9, 0x78, 0xa0, 0x92, 0xf1, 0x19, 0xa9, 0x89, 0x29, 0x42, 0x92, 0x5c, 0x1c,
	0x20, 0x7e, 0x71, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x6f, 0x10, 0x3b, 0x98, 0xef,
	0x99, 0xa2, 0xd4, 0xcd, 0x08, 0x37, 0x28, 0x3e, 0x35, 0xaf, 0xa4, 0xa8, 0x52, 0x48, 0x9a, 0x8b,
	0xb5, 0x2c, 0x31, 0x25, 0xa5, 0x08, 0xac, 0x92, 0xc5, 0x89, 0xf5, 0x92, 0x3d, 0x13, 0x07, 0x63,
	0x10, 0x44, 0x0c, 0x64, 0x52, 0x5e, 0x51, 0x3c, 0x58, 0xb3, 0x04, 0x13, 0xc4, 0xa4, 0xbc, 0xa2,
	0x00, 0x10, 0x17, 0xa8, 0x8f, 0x33, 0x33, 0x0f, 0x28, 0x55, 0x04, 0x34, 0x46, 0x82, 0x59, 0x81,
	0x51, 0x83, 0x23, 0x88, 0x23, 0x33, 0x2f, 0x00, 0xcc, 0x17, 0x52, 0xe5, 0x62, 0x4d, 0xcb, 0x49,
	0x4c, 0x2f, 0x96, 0x60, 0x01, 0x4a, 0xf0, 0x3a, 0xf1, 0x5f, 0xb2, 0xe7, 0x91, 0xe2, 0x2a, 0x00,
	0xb9, 0x1e, 0x2c, 0x1c, 0x04, 0x91, 0x75, 0x12, 0x88, 0xe2, 0xd3, 0x07, 0x7b, 0x21, 0x29, 0x33,
	0x2f, 0x25, 0x33, 0x2f, 0xbd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x18, 0x0e, 0x82, 0x46, 0xe6,
	0x00, 0x00, 0x00,
}
