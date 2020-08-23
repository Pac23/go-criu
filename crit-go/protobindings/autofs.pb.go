// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autofs.proto

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

type AutofsEntry struct {
	Fd                   *int32   `protobuf:"varint,1,req,name=fd" json:"fd,omitempty"`
	Pgrp                 *int32   `protobuf:"varint,2,req,name=pgrp" json:"pgrp,omitempty"`
	Timeout              *int32   `protobuf:"varint,3,req,name=timeout" json:"timeout,omitempty"`
	Minproto             *int32   `protobuf:"varint,4,req,name=minproto" json:"minproto,omitempty"`
	Maxproto             *int32   `protobuf:"varint,5,req,name=maxproto" json:"maxproto,omitempty"`
	Mode                 *int32   `protobuf:"varint,6,req,name=mode" json:"mode,omitempty"`
	Uid                  *int32   `protobuf:"varint,7,opt,name=uid" json:"uid,omitempty"`
	Gid                  *int32   `protobuf:"varint,8,opt,name=gid" json:"gid,omitempty"`
	ReadFd               *int32   `protobuf:"varint,9,opt,name=read_fd,json=readFd" json:"read_fd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutofsEntry) Reset()         { *m = AutofsEntry{} }
func (m *AutofsEntry) String() string { return proto.CompactTextString(m) }
func (*AutofsEntry) ProtoMessage()    {}
func (*AutofsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1a3f381d63bcd3, []int{0}
}

func (m *AutofsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofsEntry.Unmarshal(m, b)
}
func (m *AutofsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofsEntry.Marshal(b, m, deterministic)
}
func (m *AutofsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofsEntry.Merge(m, src)
}
func (m *AutofsEntry) XXX_Size() int {
	return xxx_messageInfo_AutofsEntry.Size(m)
}
func (m *AutofsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AutofsEntry proto.InternalMessageInfo

func (m *AutofsEntry) GetFd() int32 {
	if m != nil && m.Fd != nil {
		return *m.Fd
	}
	return 0
}

func (m *AutofsEntry) GetPgrp() int32 {
	if m != nil && m.Pgrp != nil {
		return *m.Pgrp
	}
	return 0
}

func (m *AutofsEntry) GetTimeout() int32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *AutofsEntry) GetMinproto() int32 {
	if m != nil && m.Minproto != nil {
		return *m.Minproto
	}
	return 0
}

func (m *AutofsEntry) GetMaxproto() int32 {
	if m != nil && m.Maxproto != nil {
		return *m.Maxproto
	}
	return 0
}

func (m *AutofsEntry) GetMode() int32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

func (m *AutofsEntry) GetUid() int32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *AutofsEntry) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *AutofsEntry) GetReadFd() int32 {
	if m != nil && m.ReadFd != nil {
		return *m.ReadFd
	}
	return 0
}

func init() {
	proto.RegisterType((*AutofsEntry)(nil), "autofs_entry")
}

func init() {
	proto.RegisterFile("autofs.proto", fileDescriptor_2f1a3f381d63bcd3)
}

var fileDescriptor_2f1a3f381d63bcd3 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8e, 0x3d, 0x8e, 0x83, 0x30,
	0x10, 0x85, 0xc5, 0x3f, 0x3b, 0x5a, 0x21, 0x34, 0xcd, 0x8e, 0xb6, 0x42, 0xa9, 0x52, 0x25, 0x77,
	0x48, 0x91, 0x03, 0xa4, 0x4c, 0x83, 0x88, 0x6c, 0x90, 0x0b, 0x6c, 0x64, 0x8c, 0x94, 0x1c, 0x34,
	0xf7, 0x89, 0x3d, 0x40, 0xba, 0xef, 0x7d, 0x4f, 0x9a, 0x37, 0xf0, 0xdb, 0x2d, 0xce, 0xf4, 0xf3,
	0x69, 0xb2, 0xc6, 0x99, 0xc3, 0x3b, 0xda, 0x45, 0x2b, 0xb5, 0xb3, 0x2f, 0xac, 0x20, 0xee, 0x05,
	0x45, 0x4d, 0x7c, 0xcc, 0x6e, 0x9e, 0x10, 0x21, 0x9d, 0x06, 0x3b, 0x51, 0xcc, 0x86, 0x19, 0x09,
	0x0a, 0xa7, 0x46, 0x69, 0x16, 0x47, 0x09, 0xeb, 0x3d, 0xe2, 0x3f, 0x94, 0xa3, 0xd2, 0x7c, 0x9a,
	0x52, 0xae, 0xbe, 0x99, 0xbb, 0xee, 0xb9, 0x76, 0xd9, 0xd6, 0x6d, 0x39, 0xac, 0x8c, 0x46, 0x48,
	0xca, 0xd7, 0x95, 0xc0, 0x58, 0x43, 0xb2, 0x28, 0x41, 0x45, 0x13, 0x79, 0x15, 0x30, 0x98, 0xc1,
	0x9b, 0x72, 0x35, 0x1e, 0xf1, 0x0f, 0x0a, 0x2b, 0x3b, 0xd1, 0xfa, 0x97, 0x7f, 0xd8, 0xe6, 0x21,
	0x5e, 0xc5, 0xa5, 0xbe, 0x57, 0x67, 0x3e, 0xfd, 0x50, 0x5a, 0x28, 0x3d, 0xcc, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x4e, 0x72, 0x16, 0xf8, 0x00, 0x00, 0x00,
}
