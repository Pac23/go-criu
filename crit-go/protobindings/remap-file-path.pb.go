// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remap-file-path.proto

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

type RemapType int32

const (
	RemapType_LINKED RemapType = 0
	RemapType_GHOST  RemapType = 1
	RemapType_PROCFS RemapType = 2
)

var RemapType_name = map[int32]string{
	0: "LINKED",
	1: "GHOST",
	2: "PROCFS",
}

var RemapType_value = map[string]int32{
	"LINKED": 0,
	"GHOST":  1,
	"PROCFS": 2,
}

func (x RemapType) Enum() *RemapType {
	p := new(RemapType)
	*p = x
	return p
}

func (x RemapType) String() string {
	return proto.EnumName(RemapType_name, int32(x))
}

func (x *RemapType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RemapType_value, data, "RemapType")
	if err != nil {
		return err
	}
	*x = RemapType(value)
	return nil
}

func (RemapType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b5dc4ad613cf4f5, []int{0}
}

type RemapFilePathEntry struct {
	OrigId               *uint32    `protobuf:"varint,1,req,name=orig_id,json=origId" json:"orig_id,omitempty"`
	RemapId              *uint32    `protobuf:"varint,2,req,name=remap_id,json=remapId" json:"remap_id,omitempty"`
	RemapType            *RemapType `protobuf:"varint,3,opt,name=remap_type,json=remapType,enum=RemapType" json:"remap_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RemapFilePathEntry) Reset()         { *m = RemapFilePathEntry{} }
func (m *RemapFilePathEntry) String() string { return proto.CompactTextString(m) }
func (*RemapFilePathEntry) ProtoMessage()    {}
func (*RemapFilePathEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5dc4ad613cf4f5, []int{0}
}

func (m *RemapFilePathEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemapFilePathEntry.Unmarshal(m, b)
}
func (m *RemapFilePathEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemapFilePathEntry.Marshal(b, m, deterministic)
}
func (m *RemapFilePathEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemapFilePathEntry.Merge(m, src)
}
func (m *RemapFilePathEntry) XXX_Size() int {
	return xxx_messageInfo_RemapFilePathEntry.Size(m)
}
func (m *RemapFilePathEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RemapFilePathEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RemapFilePathEntry proto.InternalMessageInfo

func (m *RemapFilePathEntry) GetOrigId() uint32 {
	if m != nil && m.OrigId != nil {
		return *m.OrigId
	}
	return 0
}

func (m *RemapFilePathEntry) GetRemapId() uint32 {
	if m != nil && m.RemapId != nil {
		return *m.RemapId
	}
	return 0
}

func (m *RemapFilePathEntry) GetRemapType() RemapType {
	if m != nil && m.RemapType != nil {
		return *m.RemapType
	}
	return RemapType_LINKED
}

func init() {
	proto.RegisterEnum("RemapType", RemapType_name, RemapType_value)
	proto.RegisterType((*RemapFilePathEntry)(nil), "remap_file_path_entry")
}

func init() {
	proto.RegisterFile("remap-file-path.proto", fileDescriptor_5b5dc4ad613cf4f5)
}

var fileDescriptor_5b5dc4ad613cf4f5 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0xcd, 0x4d,
	0x2c, 0xd0, 0x4d, 0xcb, 0xcc, 0x49, 0xd5, 0x2d, 0x48, 0x2c, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0x2a, 0x87, 0x4a, 0xc4, 0x83, 0x24, 0xe2, 0x41, 0x12, 0xf1, 0xa9, 0x79, 0x25, 0x45,
	0x95, 0x42, 0xe2, 0x5c, 0xec, 0xf9, 0x45, 0x99, 0xe9, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x4c,
	0x1a, 0xbc, 0x41, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x90, 0x24, 0x17, 0x07, 0x44, 0x07, 0x50, 0x86,
	0x09, 0x2c, 0xc3, 0x0e, 0xe6, 0x03, 0xa5, 0xb4, 0xb8, 0xb8, 0x20, 0x52, 0x25, 0x95, 0x05, 0xa9,
	0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xdc, 0x7a, 0x08, 0xa1, 0x20, 0x4e, 0x30, 0x3b, 0x04,
	0xc8, 0xd4, 0xd2, 0x47, 0x56, 0x2b, 0xc4, 0xc5, 0xc5, 0xe6, 0xe3, 0xe9, 0xe7, 0xed, 0xea, 0x22,
	0xc0, 0x20, 0xc4, 0xc9, 0xc5, 0xea, 0xee, 0xe1, 0x1f, 0x1c, 0x22, 0xc0, 0x08, 0x12, 0x0e, 0x08,
	0xf2, 0x77, 0x76, 0x0b, 0x16, 0x60, 0x72, 0x12, 0x88, 0xe2, 0xd3, 0x07, 0xbb, 0x39, 0x29, 0x33,
	0x2f, 0x25, 0x33, 0x2f, 0xbd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xca, 0xd4, 0x4b, 0xd3,
	0x00, 0x00, 0x00,
}