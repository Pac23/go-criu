// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cpuinfo.proto

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

type CpuinfoX86EntryVendor int32

const (
	CpuinfoX86Entry_UNKNOWN CpuinfoX86EntryVendor = 0
	CpuinfoX86Entry_INTEL   CpuinfoX86EntryVendor = 1
	CpuinfoX86Entry_AMD     CpuinfoX86EntryVendor = 2
)

var CpuinfoX86EntryVendor_name = map[int32]string{
	0: "UNKNOWN",
	1: "INTEL",
	2: "AMD",
}

var CpuinfoX86EntryVendor_value = map[string]int32{
	"UNKNOWN": 0,
	"INTEL":   1,
	"AMD":     2,
}

func (x CpuinfoX86EntryVendor) Enum() *CpuinfoX86EntryVendor {
	p := new(CpuinfoX86EntryVendor)
	*p = x
	return p
}

func (x CpuinfoX86EntryVendor) String() string {
	return proto.EnumName(CpuinfoX86EntryVendor_name, int32(x))
}

func (x *CpuinfoX86EntryVendor) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CpuinfoX86EntryVendor_value, data, "CpuinfoX86EntryVendor")
	if err != nil {
		return err
	}
	*x = CpuinfoX86EntryVendor(value)
	return nil
}

func (CpuinfoX86EntryVendor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a713993171b850, []int{0, 0}
}

type CpuinfoPpc64EntryEndianness int32

const (
	CpuinfoPpc64Entry_BIGENDIAN    CpuinfoPpc64EntryEndianness = 0
	CpuinfoPpc64Entry_LITTLEENDIAN CpuinfoPpc64EntryEndianness = 1
)

var CpuinfoPpc64EntryEndianness_name = map[int32]string{
	0: "BIGENDIAN",
	1: "LITTLEENDIAN",
}

var CpuinfoPpc64EntryEndianness_value = map[string]int32{
	"BIGENDIAN":    0,
	"LITTLEENDIAN": 1,
}

func (x CpuinfoPpc64EntryEndianness) Enum() *CpuinfoPpc64EntryEndianness {
	p := new(CpuinfoPpc64EntryEndianness)
	*p = x
	return p
}

func (x CpuinfoPpc64EntryEndianness) String() string {
	return proto.EnumName(CpuinfoPpc64EntryEndianness_name, int32(x))
}

func (x *CpuinfoPpc64EntryEndianness) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CpuinfoPpc64EntryEndianness_value, data, "CpuinfoPpc64EntryEndianness")
	if err != nil {
		return err
	}
	*x = CpuinfoPpc64EntryEndianness(value)
	return nil
}

func (CpuinfoPpc64EntryEndianness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a713993171b850, []int{1, 0}
}

type CpuinfoX86Entry struct {
	VendorId             *CpuinfoX86EntryVendor `protobuf:"varint,1,req,name=vendor_id,json=vendorId,enum=CpuinfoX86EntryVendor" json:"vendor_id,omitempty"`
	CpuFamily            *uint32                `protobuf:"varint,2,req,name=cpu_family,json=cpuFamily" json:"cpu_family,omitempty"`
	Model                *uint32                `protobuf:"varint,3,req,name=model" json:"model,omitempty"`
	Stepping             *uint32                `protobuf:"varint,4,req,name=stepping" json:"stepping,omitempty"`
	CapabilityVer        *uint32                `protobuf:"varint,5,req,name=capability_ver,json=capabilityVer" json:"capability_ver,omitempty"`
	Capability           []uint32               `protobuf:"varint,6,rep,name=capability" json:"capability,omitempty"`
	ModelId              *string                `protobuf:"bytes,7,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	XfeaturesMask        *uint64                `protobuf:"varint,8,opt,name=xfeatures_mask,json=xfeaturesMask" json:"xfeatures_mask,omitempty"`
	XsaveSize            *uint32                `protobuf:"varint,9,opt,name=xsave_size,json=xsaveSize" json:"xsave_size,omitempty"`
	XsaveSizeMax         *uint32                `protobuf:"varint,10,opt,name=xsave_size_max,json=xsaveSizeMax" json:"xsave_size_max,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CpuinfoX86Entry) Reset()         { *m = CpuinfoX86Entry{} }
func (m *CpuinfoX86Entry) String() string { return proto.CompactTextString(m) }
func (*CpuinfoX86Entry) ProtoMessage()    {}
func (*CpuinfoX86Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a713993171b850, []int{0}
}

func (m *CpuinfoX86Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuinfoX86Entry.Unmarshal(m, b)
}
func (m *CpuinfoX86Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuinfoX86Entry.Marshal(b, m, deterministic)
}
func (m *CpuinfoX86Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuinfoX86Entry.Merge(m, src)
}
func (m *CpuinfoX86Entry) XXX_Size() int {
	return xxx_messageInfo_CpuinfoX86Entry.Size(m)
}
func (m *CpuinfoX86Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuinfoX86Entry.DiscardUnknown(m)
}

var xxx_messageInfo_CpuinfoX86Entry proto.InternalMessageInfo

func (m *CpuinfoX86Entry) GetVendorId() CpuinfoX86EntryVendor {
	if m != nil && m.VendorId != nil {
		return *m.VendorId
	}
	return CpuinfoX86Entry_UNKNOWN
}

func (m *CpuinfoX86Entry) GetCpuFamily() uint32 {
	if m != nil && m.CpuFamily != nil {
		return *m.CpuFamily
	}
	return 0
}

func (m *CpuinfoX86Entry) GetModel() uint32 {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return 0
}

func (m *CpuinfoX86Entry) GetStepping() uint32 {
	if m != nil && m.Stepping != nil {
		return *m.Stepping
	}
	return 0
}

func (m *CpuinfoX86Entry) GetCapabilityVer() uint32 {
	if m != nil && m.CapabilityVer != nil {
		return *m.CapabilityVer
	}
	return 0
}

func (m *CpuinfoX86Entry) GetCapability() []uint32 {
	if m != nil {
		return m.Capability
	}
	return nil
}

func (m *CpuinfoX86Entry) GetModelId() string {
	if m != nil && m.ModelId != nil {
		return *m.ModelId
	}
	return ""
}

func (m *CpuinfoX86Entry) GetXfeaturesMask() uint64 {
	if m != nil && m.XfeaturesMask != nil {
		return *m.XfeaturesMask
	}
	return 0
}

func (m *CpuinfoX86Entry) GetXsaveSize() uint32 {
	if m != nil && m.XsaveSize != nil {
		return *m.XsaveSize
	}
	return 0
}

func (m *CpuinfoX86Entry) GetXsaveSizeMax() uint32 {
	if m != nil && m.XsaveSizeMax != nil {
		return *m.XsaveSizeMax
	}
	return 0
}

type CpuinfoPpc64Entry struct {
	Endian               *CpuinfoPpc64EntryEndianness `protobuf:"varint,1,req,name=endian,enum=CpuinfoPpc64EntryEndianness" json:"endian,omitempty"`
	Hwcap                []uint64                     `protobuf:"varint,2,rep,name=hwcap" json:"hwcap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CpuinfoPpc64Entry) Reset()         { *m = CpuinfoPpc64Entry{} }
func (m *CpuinfoPpc64Entry) String() string { return proto.CompactTextString(m) }
func (*CpuinfoPpc64Entry) ProtoMessage()    {}
func (*CpuinfoPpc64Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a713993171b850, []int{1}
}

func (m *CpuinfoPpc64Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuinfoPpc64Entry.Unmarshal(m, b)
}
func (m *CpuinfoPpc64Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuinfoPpc64Entry.Marshal(b, m, deterministic)
}
func (m *CpuinfoPpc64Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuinfoPpc64Entry.Merge(m, src)
}
func (m *CpuinfoPpc64Entry) XXX_Size() int {
	return xxx_messageInfo_CpuinfoPpc64Entry.Size(m)
}
func (m *CpuinfoPpc64Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuinfoPpc64Entry.DiscardUnknown(m)
}

var xxx_messageInfo_CpuinfoPpc64Entry proto.InternalMessageInfo

func (m *CpuinfoPpc64Entry) GetEndian() CpuinfoPpc64EntryEndianness {
	if m != nil && m.Endian != nil {
		return *m.Endian
	}
	return CpuinfoPpc64Entry_BIGENDIAN
}

func (m *CpuinfoPpc64Entry) GetHwcap() []uint64 {
	if m != nil {
		return m.Hwcap
	}
	return nil
}

type CpuinfoS390Entry struct {
	Hwcap                []uint64 `protobuf:"varint,2,rep,name=hwcap" json:"hwcap,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CpuinfoS390Entry) Reset()         { *m = CpuinfoS390Entry{} }
func (m *CpuinfoS390Entry) String() string { return proto.CompactTextString(m) }
func (*CpuinfoS390Entry) ProtoMessage()    {}
func (*CpuinfoS390Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a713993171b850, []int{2}
}

func (m *CpuinfoS390Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuinfoS390Entry.Unmarshal(m, b)
}
func (m *CpuinfoS390Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuinfoS390Entry.Marshal(b, m, deterministic)
}
func (m *CpuinfoS390Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuinfoS390Entry.Merge(m, src)
}
func (m *CpuinfoS390Entry) XXX_Size() int {
	return xxx_messageInfo_CpuinfoS390Entry.Size(m)
}
func (m *CpuinfoS390Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuinfoS390Entry.DiscardUnknown(m)
}

var xxx_messageInfo_CpuinfoS390Entry proto.InternalMessageInfo

func (m *CpuinfoS390Entry) GetHwcap() []uint64 {
	if m != nil {
		return m.Hwcap
	}
	return nil
}

type CpuinfoEntry struct {
	//
	// Usually on SMP system there should be same CPUs
	// installed, but it might happen that system carries
	// various CPUs so @repeated used.
	X86Entry             []*CpuinfoX86Entry   `protobuf:"bytes,1,rep,name=x86_entry,json=x86Entry" json:"x86_entry,omitempty"`
	Ppc64Entry           []*CpuinfoPpc64Entry `protobuf:"bytes,2,rep,name=ppc64_entry,json=ppc64Entry" json:"ppc64_entry,omitempty"`
	S390Entry            []*CpuinfoS390Entry  `protobuf:"bytes,3,rep,name=s390_entry,json=s390Entry" json:"s390_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CpuinfoEntry) Reset()         { *m = CpuinfoEntry{} }
func (m *CpuinfoEntry) String() string { return proto.CompactTextString(m) }
func (*CpuinfoEntry) ProtoMessage()    {}
func (*CpuinfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a713993171b850, []int{3}
}

func (m *CpuinfoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpuinfoEntry.Unmarshal(m, b)
}
func (m *CpuinfoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpuinfoEntry.Marshal(b, m, deterministic)
}
func (m *CpuinfoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpuinfoEntry.Merge(m, src)
}
func (m *CpuinfoEntry) XXX_Size() int {
	return xxx_messageInfo_CpuinfoEntry.Size(m)
}
func (m *CpuinfoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CpuinfoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CpuinfoEntry proto.InternalMessageInfo

func (m *CpuinfoEntry) GetX86Entry() []*CpuinfoX86Entry {
	if m != nil {
		return m.X86Entry
	}
	return nil
}

func (m *CpuinfoEntry) GetPpc64Entry() []*CpuinfoPpc64Entry {
	if m != nil {
		return m.Ppc64Entry
	}
	return nil
}

func (m *CpuinfoEntry) GetS390Entry() []*CpuinfoS390Entry {
	if m != nil {
		return m.S390Entry
	}
	return nil
}

func init() {
	proto.RegisterEnum("CpuinfoX86EntryVendor", CpuinfoX86EntryVendor_name, CpuinfoX86EntryVendor_value)
	proto.RegisterEnum("CpuinfoPpc64EntryEndianness", CpuinfoPpc64EntryEndianness_name, CpuinfoPpc64EntryEndianness_value)
	proto.RegisterType((*CpuinfoX86Entry)(nil), "cpuinfo_x86_entry")
	proto.RegisterType((*CpuinfoPpc64Entry)(nil), "cpuinfo_ppc64_entry")
	proto.RegisterType((*CpuinfoS390Entry)(nil), "cpuinfo_s390_entry")
	proto.RegisterType((*CpuinfoEntry)(nil), "cpuinfo_entry")
}

func init() {
	proto.RegisterFile("cpuinfo.proto", fileDescriptor_77a713993171b850)
}

var fileDescriptor_77a713993171b850 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xb3, 0xb6, 0xc9, 0xe9, 0x52, 0x05, 0x6f, 0x17, 0x1e, 0x12, 0x50, 0x45, 0x20,
	0x15, 0x24, 0x32, 0x54, 0xa0, 0xc0, 0xe5, 0xa6, 0x15, 0x14, 0xd1, 0x06, 0x29, 0x14, 0x90, 0xb8,
	0x89, 0xbc, 0xd6, 0x1d, 0x16, 0x6d, 0x12, 0xc5, 0x59, 0x49, 0x79, 0x0d, 0xc4, 0x53, 0xf0, 0x92,
	0x3b, 0x71, 0xb2, 0x74, 0x52, 0x7b, 0x15, 0x9f, 0xbf, 0x7f, 0xe7, 0xd3, 0x27, 0x60, 0xcd, 0x92,
	0x6b, 0x11, 0x2d, 0x62, 0x37, 0x49, 0xe3, 0x2c, 0x76, 0xfe, 0xe9, 0x70, 0xbf, 0x52, 0xc2, 0xfc,
	0xdd, 0x30, 0xe4, 0x51, 0x96, 0x6e, 0xc8, 0x10, 0xcc, 0x35, 0x8f, 0xe6, 0x71, 0x1a, 0x8a, 0x39,
	0xd5, 0x7a, 0x8d, 0x7e, 0x77, 0x70, 0xe2, 0xee, 0x60, 0x6e, 0xc9, 0x04, 0x46, 0xf9, 0xf5, 0xe6,
	0xe4, 0x21, 0x00, 0x52, 0xe1, 0x82, 0xad, 0xc4, 0x72, 0x43, 0x1b, 0xe8, 0x68, 0x05, 0x26, 0x2a,
	0x1f, 0x94, 0x40, 0x8e, 0xa1, 0xb9, 0x8a, 0xe7, 0x7c, 0x49, 0x75, 0x75, 0x53, 0x1a, 0xe4, 0x01,
	0x18, 0x32, 0xe3, 0x49, 0x22, 0xa2, 0x2b, 0x7a, 0xa0, 0x2e, 0x6a, 0x9b, 0x3c, 0x85, 0xee, 0x8c,
	0x25, 0xec, 0x52, 0x2c, 0x45, 0xb6, 0x09, 0xd7, 0x3c, 0xa5, 0x4d, 0x45, 0x58, 0x5b, 0xf5, 0x1b,
	0x4f, 0xc9, 0x23, 0xcc, 0x5b, 0x0b, 0xb4, 0xd5, 0xd3, 0x11, 0xb9, 0xa3, 0x90, 0x13, 0x30, 0x54,
	0xae, 0xa2, 0x9d, 0x76, 0x4f, 0xeb, 0x9b, 0x41, 0x5b, 0xd9, 0x58, 0x32, 0x66, 0xc8, 0x17, 0x9c,
	0x65, 0xd7, 0x29, 0x97, 0xe1, 0x8a, 0xc9, 0x5f, 0xd4, 0x40, 0xe0, 0x20, 0xb0, 0x6a, 0x75, 0x82,
	0x62, 0xd1, 0x59, 0x2e, 0xd9, 0x9a, 0x87, 0x52, 0xfc, 0xe1, 0xd4, 0x44, 0x04, 0x3b, 0x53, 0xca,
	0x17, 0x14, 0xc8, 0x13, 0x8c, 0x52, 0x5f, 0x63, 0x98, 0x9c, 0x82, 0x42, 0x0e, 0x6b, 0x64, 0xc2,
	0x72, 0xe7, 0x19, 0xb4, 0xca, 0x51, 0x91, 0x0e, 0xb4, 0xbf, 0xfa, 0x9f, 0xfc, 0xcf, 0xdf, 0x7d,
	0xfb, 0x1e, 0x31, 0xa1, 0xe9, 0xf9, 0xd3, 0xd1, 0xd8, 0xd6, 0x48, 0x1b, 0xf4, 0xb3, 0xc9, 0x85,
	0xdd, 0x70, 0xfe, 0x6a, 0x70, 0x74, 0x3b, 0xf0, 0x24, 0x99, 0x0d, 0x5f, 0x57, 0x2f, 0xf3, 0x16,
	0x5a, 0x18, 0x41, 0xb0, 0xa8, 0x7a, 0x96, 0xc7, 0xee, 0x1e, 0xca, 0x2d, 0x91, 0x88, 0x4b, 0x19,
	0x54, 0x78, 0x31, 0xfb, 0x9f, 0xbf, 0x71, 0x24, 0xf8, 0x2a, 0x3a, 0xb6, 0x57, 0x1a, 0xce, 0x0b,
	0x80, 0x2d, 0x4b, 0x2c, 0x30, 0xcf, 0xbd, 0x8f, 0x23, 0xff, 0xc2, 0x3b, 0x2b, 0xea, 0xb2, 0xe1,
	0x70, 0xec, 0x4d, 0xa7, 0xe3, 0x51, 0xa5, 0x68, 0xce, 0x73, 0x20, 0xb7, 0xe9, 0xe4, 0xab, 0xf7,
	0x2f, 0xab, 0x9a, 0xf6, 0x87, 0xfe, 0xaf, 0xd5, 0xbb, 0x56, 0x71, 0xa7, 0x60, 0xd6, 0xbb, 0x83,
	0xe5, 0xeb, 0xfd, 0xce, 0x80, 0xec, 0x6e, 0x55, 0x60, 0xe0, 0x71, 0xa4, 0x1c, 0xde, 0x40, 0xe7,
	0x4e, 0x57, 0x2a, 0x7c, 0x67, 0x70, 0xbc, 0xaf, 0xe3, 0x00, 0x94, 0x51, 0xba, 0x0d, 0x00, 0xb6,
	0xd5, 0xe1, 0xae, 0x15, 0x5e, 0x47, 0xee, 0x6e, 0xe1, 0x81, 0x59, 0x9c, 0x95, 0xcf, 0xb9, 0xfd,
	0xa3, 0x7b, 0xaa, 0xfe, 0x88, 0x4b, 0x81, 0xf3, 0x88, 0xae, 0xe4, 0x4d, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0xd9, 0xce, 0x3f, 0x29, 0x03, 0x00, 0x00,
}