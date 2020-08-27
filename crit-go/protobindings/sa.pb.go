// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sa.proto

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

type SaEntry struct {
	Sigaction            *uint64  `protobuf:"varint,1,req,name=sigaction" json:"sigaction,omitempty"`
	Flags                *uint64  `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Restorer             *uint64  `protobuf:"varint,3,req,name=restorer" json:"restorer,omitempty"`
	Mask                 *uint64  `protobuf:"varint,4,req,name=mask" json:"mask,omitempty"`
	CompatSigaction      *bool    `protobuf:"varint,5,opt,name=compat_sigaction,json=compatSigaction" json:"compat_sigaction,omitempty"`
	MaskExtended         *uint64  `protobuf:"varint,6,opt,name=mask_extended,json=maskExtended" json:"mask_extended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaEntry) Reset()         { *m = SaEntry{} }
func (m *SaEntry) String() string { return proto.CompactTextString(m) }
func (*SaEntry) ProtoMessage()    {}
func (*SaEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_8291c9e9087af49a, []int{0}
}

func (m *SaEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaEntry.Unmarshal(m, b)
}
func (m *SaEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaEntry.Marshal(b, m, deterministic)
}
func (m *SaEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaEntry.Merge(m, src)
}
func (m *SaEntry) XXX_Size() int {
	return xxx_messageInfo_SaEntry.Size(m)
}
func (m *SaEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SaEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SaEntry proto.InternalMessageInfo

func (m *SaEntry) GetSigaction() uint64 {
	if m != nil && m.Sigaction != nil {
		return *m.Sigaction
	}
	return 0
}

func (m *SaEntry) GetFlags() uint64 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *SaEntry) GetRestorer() uint64 {
	if m != nil && m.Restorer != nil {
		return *m.Restorer
	}
	return 0
}

func (m *SaEntry) GetMask() uint64 {
	if m != nil && m.Mask != nil {
		return *m.Mask
	}
	return 0
}

func (m *SaEntry) GetCompatSigaction() bool {
	if m != nil && m.CompatSigaction != nil {
		return *m.CompatSigaction
	}
	return false
}

func (m *SaEntry) GetMaskExtended() uint64 {
	if m != nil && m.MaskExtended != nil {
		return *m.MaskExtended
	}
	return 0
}

func init() {
	proto.RegisterType((*SaEntry)(nil), "sa_entry")
}

func init() {
	proto.RegisterFile("sa.proto", fileDescriptor_8291c9e9087af49a)
}

var fileDescriptor_8291c9e9087af49a = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x4e, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xca, 0x2f, 0x28, 0x29, 0x86, 0xb0, 0x95, 0x1e, 0x32, 0x82,
	0x24, 0xe2, 0x53, 0xf3, 0x4a, 0x8a, 0x2a, 0x85, 0x94, 0xb9, 0x38, 0x8b, 0x33, 0xd3, 0x13, 0x93,
	0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x9c, 0x58, 0x2f, 0xd9, 0x33, 0x71,
	0x30, 0x06, 0x21, 0xc4, 0x85, 0xa4, 0xb9, 0x58, 0xd3, 0x72, 0x12, 0xd3, 0x8b, 0x25, 0x98, 0x90,
	0x15, 0x40, 0xc4, 0x84, 0x14, 0xb9, 0x38, 0x8a, 0x52, 0x8b, 0x4b, 0xf2, 0x8b, 0x52, 0x8b, 0x24,
	0x98, 0x91, 0xe5, 0xe1, 0xc2, 0x42, 0x92, 0x5c, 0x2c, 0xb9, 0x89, 0xc5, 0xd9, 0x12, 0x2c, 0xc8,
	0xd2, 0x60, 0x21, 0x21, 0x4d, 0x2e, 0x81, 0xe4, 0xfc, 0xdc, 0x82, 0xc4, 0x92, 0x78, 0x84, 0x33,
	0x58, 0x15, 0x18, 0x35, 0x38, 0x82, 0xf8, 0x21, 0xe2, 0xc1, 0x70, 0x57, 0x68, 0x71, 0xf1, 0x82,
	0xb4, 0xc4, 0xa7, 0x56, 0x94, 0xa4, 0xe6, 0xa5, 0xa4, 0xa6, 0x48, 0xb0, 0x01, 0xd5, 0xc1, 0x8d,
	0xe3, 0x01, 0xc9, 0xb9, 0x42, 0xa5, 0x9c, 0x04, 0xa2, 0xf8, 0xf4, 0xc1, 0xbe, 0x4d, 0xca, 0xcc,
	0x4b, 0xc9, 0xcc, 0x4b, 0x2f, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x8a, 0x38, 0xbc, 0x0c,
	0x01, 0x00, 0x00,
}
