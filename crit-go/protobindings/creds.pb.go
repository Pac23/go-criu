// Code generated by protoc-gen-go. DO NOT EDIT.
// source: creds.proto

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

type CredsEntry struct {
	Uid                  *uint32  `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Gid                  *uint32  `protobuf:"varint,2,req,name=gid" json:"gid,omitempty"`
	Euid                 *uint32  `protobuf:"varint,3,req,name=euid" json:"euid,omitempty"`
	Egid                 *uint32  `protobuf:"varint,4,req,name=egid" json:"egid,omitempty"`
	Suid                 *uint32  `protobuf:"varint,5,req,name=suid" json:"suid,omitempty"`
	Sgid                 *uint32  `protobuf:"varint,6,req,name=sgid" json:"sgid,omitempty"`
	Fsuid                *uint32  `protobuf:"varint,7,req,name=fsuid" json:"fsuid,omitempty"`
	Fsgid                *uint32  `protobuf:"varint,8,req,name=fsgid" json:"fsgid,omitempty"`
	CapInh               []uint32 `protobuf:"varint,9,rep,name=cap_inh,json=capInh" json:"cap_inh,omitempty"`
	CapPrm               []uint32 `protobuf:"varint,10,rep,name=cap_prm,json=capPrm" json:"cap_prm,omitempty"`
	CapEff               []uint32 `protobuf:"varint,11,rep,name=cap_eff,json=capEff" json:"cap_eff,omitempty"`
	CapBnd               []uint32 `protobuf:"varint,12,rep,name=cap_bnd,json=capBnd" json:"cap_bnd,omitempty"`
	Secbits              *uint32  `protobuf:"varint,13,req,name=secbits" json:"secbits,omitempty"`
	Groups               []uint32 `protobuf:"varint,14,rep,name=groups" json:"groups,omitempty"`
	LsmProfile           *string  `protobuf:"bytes,15,opt,name=lsm_profile,json=lsmProfile" json:"lsm_profile,omitempty"`
	LsmSockcreate        *string  `protobuf:"bytes,16,opt,name=lsm_sockcreate,json=lsmSockcreate" json:"lsm_sockcreate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredsEntry) Reset()         { *m = CredsEntry{} }
func (m *CredsEntry) String() string { return proto.CompactTextString(m) }
func (*CredsEntry) ProtoMessage()    {}
func (*CredsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc577c816a5ed283, []int{0}
}

func (m *CredsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredsEntry.Unmarshal(m, b)
}
func (m *CredsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredsEntry.Marshal(b, m, deterministic)
}
func (m *CredsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredsEntry.Merge(m, src)
}
func (m *CredsEntry) XXX_Size() int {
	return xxx_messageInfo_CredsEntry.Size(m)
}
func (m *CredsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CredsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CredsEntry proto.InternalMessageInfo

func (m *CredsEntry) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *CredsEntry) GetGid() uint32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *CredsEntry) GetEuid() uint32 {
	if m != nil && m.Euid != nil {
		return *m.Euid
	}
	return 0
}

func (m *CredsEntry) GetEgid() uint32 {
	if m != nil && m.Egid != nil {
		return *m.Egid
	}
	return 0
}

func (m *CredsEntry) GetSuid() uint32 {
	if m != nil && m.Suid != nil {
		return *m.Suid
	}
	return 0
}

func (m *CredsEntry) GetSgid() uint32 {
	if m != nil && m.Sgid != nil {
		return *m.Sgid
	}
	return 0
}

func (m *CredsEntry) GetFsuid() uint32 {
	if m != nil && m.Fsuid != nil {
		return *m.Fsuid
	}
	return 0
}

func (m *CredsEntry) GetFsgid() uint32 {
	if m != nil && m.Fsgid != nil {
		return *m.Fsgid
	}
	return 0
}

func (m *CredsEntry) GetCapInh() []uint32 {
	if m != nil {
		return m.CapInh
	}
	return nil
}

func (m *CredsEntry) GetCapPrm() []uint32 {
	if m != nil {
		return m.CapPrm
	}
	return nil
}

func (m *CredsEntry) GetCapEff() []uint32 {
	if m != nil {
		return m.CapEff
	}
	return nil
}

func (m *CredsEntry) GetCapBnd() []uint32 {
	if m != nil {
		return m.CapBnd
	}
	return nil
}

func (m *CredsEntry) GetSecbits() uint32 {
	if m != nil && m.Secbits != nil {
		return *m.Secbits
	}
	return 0
}

func (m *CredsEntry) GetGroups() []uint32 {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *CredsEntry) GetLsmProfile() string {
	if m != nil && m.LsmProfile != nil {
		return *m.LsmProfile
	}
	return ""
}

func (m *CredsEntry) GetLsmSockcreate() string {
	if m != nil && m.LsmSockcreate != nil {
		return *m.LsmSockcreate
	}
	return ""
}

func init() {
	proto.RegisterType((*CredsEntry)(nil), "creds_entry")
}

func init() {
	proto.RegisterFile("creds.proto", fileDescriptor_bc577c816a5ed283)
}

var fileDescriptor_bc577c816a5ed283 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x91, 0xb1, 0x6e, 0x83, 0x30,
	0x14, 0x45, 0x95, 0x90, 0x40, 0xf3, 0x28, 0x14, 0x59, 0x55, 0xeb, 0xad, 0xa8, 0x52, 0xa5, 0x4c,
	0xed, 0x3f, 0x44, 0xea, 0xd0, 0x2d, 0x4a, 0xb7, 0x2e, 0x51, 0x00, 0x9b, 0x58, 0x0d, 0x06, 0xf9,
	0x91, 0xa1, 0xdf, 0xd0, 0x9f, 0x8e, 0x9f, 0x8d, 0xd9, 0xee, 0x3d, 0xf7, 0xa0, 0x27, 0x61, 0x48,
	0x6b, 0x23, 0x1a, 0x7c, 0x1f, 0x4c, 0x3f, 0xf6, 0xaf, 0xff, 0xd1, 0xd4, 0x8f, 0x42, 0x8f, 0xe6,
	0x8f, 0x15, 0x10, 0x5d, 0x55, 0xc3, 0x17, 0xe5, 0x72, 0x9b, 0x1d, 0x28, 0x12, 0x69, 0x2d, 0x59,
	0x7a, 0x62, 0x23, 0x63, 0xb0, 0x12, 0x24, 0x45, 0x0e, 0xb9, 0xec, 0x18, 0x69, 0xab, 0x89, 0x4d,
	0x1e, 0x92, 0xb7, 0xf6, 0x0c, 0x27, 0x0f, 0xc9, 0x8b, 0x27, 0x46, 0xde, 0x23, 0xac, 0xa5, 0x13,
	0x13, 0x07, 0x7d, 0xf1, 0x94, 0xd4, 0xbb, 0x40, 0xc9, 0x7d, 0x86, 0xa4, 0x3e, 0x0d, 0x47, 0xa5,
	0xcf, 0x7c, 0x53, 0x46, 0x96, 0xc7, 0xb6, 0x7e, 0xe9, 0x73, 0x18, 0x06, 0xd3, 0x71, 0x98, 0x87,
	0xbd, 0xe9, 0xc2, 0x20, 0xa4, 0xe4, 0xe9, 0x3c, 0x7c, 0x4a, 0x19, 0x86, 0x4a, 0x37, 0xfc, 0x7e,
	0x1e, 0x76, 0xba, 0x61, 0x1c, 0x12, 0x14, 0x75, 0xa5, 0x46, 0xe4, 0x99, 0xbb, 0x1d, 0x2a, 0x7b,
	0x82, 0xb8, 0x35, 0xfd, 0x75, 0x40, 0x9e, 0xfb, 0x2f, 0x7c, 0x63, 0x2f, 0x90, 0x5e, 0xb0, 0xb3,
	0xc7, 0x7b, 0xa9, 0x2e, 0x82, 0x3f, 0x94, 0x8b, 0xed, 0xe6, 0x00, 0x16, 0xed, 0x3d, 0x61, 0x6f,
	0x90, 0x93, 0x80, 0x7d, 0xfd, 0x6b, 0xff, 0xf6, 0x69, 0x14, 0xbc, 0x70, 0x4e, 0x66, 0xe9, 0xf7,
	0x0c, 0x77, 0xc5, 0x4f, 0xfe, 0xe1, 0xde, 0xa5, 0x52, 0xba, 0x51, 0xba, 0xc5, 0x5b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x30, 0x95, 0x04, 0xe2, 0xad, 0x01, 0x00, 0x00,
}