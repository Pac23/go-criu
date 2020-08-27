// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core-x86.proto

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

type UserX86RegsMode int32

const (
	UserX86RegsMode_NATIVE UserX86RegsMode = 1
	UserX86RegsMode_COMPAT UserX86RegsMode = 2
)

var UserX86RegsMode_name = map[int32]string{
	1: "NATIVE",
	2: "COMPAT",
}

var UserX86RegsMode_value = map[string]int32{
	"NATIVE": 1,
	"COMPAT": 2,
}

func (x UserX86RegsMode) Enum() *UserX86RegsMode {
	p := new(UserX86RegsMode)
	*p = x
	return p
}

func (x UserX86RegsMode) String() string {
	return proto.EnumName(UserX86RegsMode_name, int32(x))
}

func (x *UserX86RegsMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserX86RegsMode_value, data, "UserX86RegsMode")
	if err != nil {
		return err
	}
	*x = UserX86RegsMode(value)
	return nil
}

func (UserX86RegsMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b0422cff64b59e29, []int{0}
}

// Reusing entry for both 64 and 32 bits register sets
type UserX86RegsEntry struct {
	R15                  *uint64          `protobuf:"varint,1,req,name=r15" json:"r15,omitempty"`
	R14                  *uint64          `protobuf:"varint,2,req,name=r14" json:"r14,omitempty"`
	R13                  *uint64          `protobuf:"varint,3,req,name=r13" json:"r13,omitempty"`
	R12                  *uint64          `protobuf:"varint,4,req,name=r12" json:"r12,omitempty"`
	Bp                   *uint64          `protobuf:"varint,5,req,name=bp" json:"bp,omitempty"`
	Bx                   *uint64          `protobuf:"varint,6,req,name=bx" json:"bx,omitempty"`
	R11                  *uint64          `protobuf:"varint,7,req,name=r11" json:"r11,omitempty"`
	R10                  *uint64          `protobuf:"varint,8,req,name=r10" json:"r10,omitempty"`
	R9                   *uint64          `protobuf:"varint,9,req,name=r9" json:"r9,omitempty"`
	R8                   *uint64          `protobuf:"varint,10,req,name=r8" json:"r8,omitempty"`
	Ax                   *uint64          `protobuf:"varint,11,req,name=ax" json:"ax,omitempty"`
	Cx                   *uint64          `protobuf:"varint,12,req,name=cx" json:"cx,omitempty"`
	Dx                   *uint64          `protobuf:"varint,13,req,name=dx" json:"dx,omitempty"`
	Si                   *uint64          `protobuf:"varint,14,req,name=si" json:"si,omitempty"`
	Di                   *uint64          `protobuf:"varint,15,req,name=di" json:"di,omitempty"`
	OrigAx               *uint64          `protobuf:"varint,16,req,name=orig_ax,json=origAx" json:"orig_ax,omitempty"`
	Ip                   *uint64          `protobuf:"varint,17,req,name=ip" json:"ip,omitempty"`
	Cs                   *uint64          `protobuf:"varint,18,req,name=cs" json:"cs,omitempty"`
	Flags                *uint64          `protobuf:"varint,19,req,name=flags" json:"flags,omitempty"`
	Sp                   *uint64          `protobuf:"varint,20,req,name=sp" json:"sp,omitempty"`
	Ss                   *uint64          `protobuf:"varint,21,req,name=ss" json:"ss,omitempty"`
	FsBase               *uint64          `protobuf:"varint,22,req,name=fs_base,json=fsBase" json:"fs_base,omitempty"`
	GsBase               *uint64          `protobuf:"varint,23,req,name=gs_base,json=gsBase" json:"gs_base,omitempty"`
	Ds                   *uint64          `protobuf:"varint,24,req,name=ds" json:"ds,omitempty"`
	Es                   *uint64          `protobuf:"varint,25,req,name=es" json:"es,omitempty"`
	Fs                   *uint64          `protobuf:"varint,26,req,name=fs" json:"fs,omitempty"`
	Gs                   *uint64          `protobuf:"varint,27,req,name=gs" json:"gs,omitempty"`
	Mode                 *UserX86RegsMode `protobuf:"varint,28,opt,name=mode,enum=UserX86RegsMode,def=1" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserX86RegsEntry) Reset()         { *m = UserX86RegsEntry{} }
func (m *UserX86RegsEntry) String() string { return proto.CompactTextString(m) }
func (*UserX86RegsEntry) ProtoMessage()    {}
func (*UserX86RegsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0422cff64b59e29, []int{0}
}

func (m *UserX86RegsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserX86RegsEntry.Unmarshal(m, b)
}
func (m *UserX86RegsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserX86RegsEntry.Marshal(b, m, deterministic)
}
func (m *UserX86RegsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserX86RegsEntry.Merge(m, src)
}
func (m *UserX86RegsEntry) XXX_Size() int {
	return xxx_messageInfo_UserX86RegsEntry.Size(m)
}
func (m *UserX86RegsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserX86RegsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserX86RegsEntry proto.InternalMessageInfo

const Default_UserX86RegsEntry_Mode UserX86RegsMode = UserX86RegsMode_NATIVE

func (m *UserX86RegsEntry) GetR15() uint64 {
	if m != nil && m.R15 != nil {
		return *m.R15
	}
	return 0
}

func (m *UserX86RegsEntry) GetR14() uint64 {
	if m != nil && m.R14 != nil {
		return *m.R14
	}
	return 0
}

func (m *UserX86RegsEntry) GetR13() uint64 {
	if m != nil && m.R13 != nil {
		return *m.R13
	}
	return 0
}

func (m *UserX86RegsEntry) GetR12() uint64 {
	if m != nil && m.R12 != nil {
		return *m.R12
	}
	return 0
}

func (m *UserX86RegsEntry) GetBp() uint64 {
	if m != nil && m.Bp != nil {
		return *m.Bp
	}
	return 0
}

func (m *UserX86RegsEntry) GetBx() uint64 {
	if m != nil && m.Bx != nil {
		return *m.Bx
	}
	return 0
}

func (m *UserX86RegsEntry) GetR11() uint64 {
	if m != nil && m.R11 != nil {
		return *m.R11
	}
	return 0
}

func (m *UserX86RegsEntry) GetR10() uint64 {
	if m != nil && m.R10 != nil {
		return *m.R10
	}
	return 0
}

func (m *UserX86RegsEntry) GetR9() uint64 {
	if m != nil && m.R9 != nil {
		return *m.R9
	}
	return 0
}

func (m *UserX86RegsEntry) GetR8() uint64 {
	if m != nil && m.R8 != nil {
		return *m.R8
	}
	return 0
}

func (m *UserX86RegsEntry) GetAx() uint64 {
	if m != nil && m.Ax != nil {
		return *m.Ax
	}
	return 0
}

func (m *UserX86RegsEntry) GetCx() uint64 {
	if m != nil && m.Cx != nil {
		return *m.Cx
	}
	return 0
}

func (m *UserX86RegsEntry) GetDx() uint64 {
	if m != nil && m.Dx != nil {
		return *m.Dx
	}
	return 0
}

func (m *UserX86RegsEntry) GetSi() uint64 {
	if m != nil && m.Si != nil {
		return *m.Si
	}
	return 0
}

func (m *UserX86RegsEntry) GetDi() uint64 {
	if m != nil && m.Di != nil {
		return *m.Di
	}
	return 0
}

func (m *UserX86RegsEntry) GetOrigAx() uint64 {
	if m != nil && m.OrigAx != nil {
		return *m.OrigAx
	}
	return 0
}

func (m *UserX86RegsEntry) GetIp() uint64 {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return 0
}

func (m *UserX86RegsEntry) GetCs() uint64 {
	if m != nil && m.Cs != nil {
		return *m.Cs
	}
	return 0
}

func (m *UserX86RegsEntry) GetFlags() uint64 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *UserX86RegsEntry) GetSp() uint64 {
	if m != nil && m.Sp != nil {
		return *m.Sp
	}
	return 0
}

func (m *UserX86RegsEntry) GetSs() uint64 {
	if m != nil && m.Ss != nil {
		return *m.Ss
	}
	return 0
}

func (m *UserX86RegsEntry) GetFsBase() uint64 {
	if m != nil && m.FsBase != nil {
		return *m.FsBase
	}
	return 0
}

func (m *UserX86RegsEntry) GetGsBase() uint64 {
	if m != nil && m.GsBase != nil {
		return *m.GsBase
	}
	return 0
}

func (m *UserX86RegsEntry) GetDs() uint64 {
	if m != nil && m.Ds != nil {
		return *m.Ds
	}
	return 0
}

func (m *UserX86RegsEntry) GetEs() uint64 {
	if m != nil && m.Es != nil {
		return *m.Es
	}
	return 0
}

func (m *UserX86RegsEntry) GetFs() uint64 {
	if m != nil && m.Fs != nil {
		return *m.Fs
	}
	return 0
}

func (m *UserX86RegsEntry) GetGs() uint64 {
	if m != nil && m.Gs != nil {
		return *m.Gs
	}
	return 0
}

func (m *UserX86RegsEntry) GetMode() UserX86RegsMode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return Default_UserX86RegsEntry_Mode
}

type UserX86XsaveEntry struct {
	// standart xsave features
	XstateBv *uint64 `protobuf:"varint,1,req,name=xstate_bv,json=xstateBv" json:"xstate_bv,omitempty"`
	// AVX components: 16x 256-bit ymm registers, hi 128 bits
	YmmhSpace []uint32 `protobuf:"varint,2,rep,name=ymmh_space,json=ymmhSpace" json:"ymmh_space,omitempty"`
	// MPX components
	BndregState []uint64 `protobuf:"varint,3,rep,name=bndreg_state,json=bndregState" json:"bndreg_state,omitempty"`
	BndcsrState []uint64 `protobuf:"varint,4,rep,name=bndcsr_state,json=bndcsrState" json:"bndcsr_state,omitempty"`
	// AVX512 components: k0-k7, ZMM_Hi256, Hi16_ZMM
	OpmaskReg []uint64 `protobuf:"varint,5,rep,name=opmask_reg,json=opmaskReg" json:"opmask_reg,omitempty"`
	ZmmUpper  []uint64 `protobuf:"varint,6,rep,name=zmm_upper,json=zmmUpper" json:"zmm_upper,omitempty"`
	Hi16Zmm   []uint64 `protobuf:"varint,7,rep,name=hi16_zmm,json=hi16Zmm" json:"hi16_zmm,omitempty"`
	// Protected keys
	Pkru                 []uint32 `protobuf:"varint,8,rep,name=pkru" json:"pkru,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserX86XsaveEntry) Reset()         { *m = UserX86XsaveEntry{} }
func (m *UserX86XsaveEntry) String() string { return proto.CompactTextString(m) }
func (*UserX86XsaveEntry) ProtoMessage()    {}
func (*UserX86XsaveEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0422cff64b59e29, []int{1}
}

func (m *UserX86XsaveEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserX86XsaveEntry.Unmarshal(m, b)
}
func (m *UserX86XsaveEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserX86XsaveEntry.Marshal(b, m, deterministic)
}
func (m *UserX86XsaveEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserX86XsaveEntry.Merge(m, src)
}
func (m *UserX86XsaveEntry) XXX_Size() int {
	return xxx_messageInfo_UserX86XsaveEntry.Size(m)
}
func (m *UserX86XsaveEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserX86XsaveEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserX86XsaveEntry proto.InternalMessageInfo

func (m *UserX86XsaveEntry) GetXstateBv() uint64 {
	if m != nil && m.XstateBv != nil {
		return *m.XstateBv
	}
	return 0
}

func (m *UserX86XsaveEntry) GetYmmhSpace() []uint32 {
	if m != nil {
		return m.YmmhSpace
	}
	return nil
}

func (m *UserX86XsaveEntry) GetBndregState() []uint64 {
	if m != nil {
		return m.BndregState
	}
	return nil
}

func (m *UserX86XsaveEntry) GetBndcsrState() []uint64 {
	if m != nil {
		return m.BndcsrState
	}
	return nil
}

func (m *UserX86XsaveEntry) GetOpmaskReg() []uint64 {
	if m != nil {
		return m.OpmaskReg
	}
	return nil
}

func (m *UserX86XsaveEntry) GetZmmUpper() []uint64 {
	if m != nil {
		return m.ZmmUpper
	}
	return nil
}

func (m *UserX86XsaveEntry) GetHi16Zmm() []uint64 {
	if m != nil {
		return m.Hi16Zmm
	}
	return nil
}

func (m *UserX86XsaveEntry) GetPkru() []uint32 {
	if m != nil {
		return m.Pkru
	}
	return nil
}

type UserX86FpregsEntry struct {
	// fxsave data
	Cwd       *uint32  `protobuf:"varint,1,req,name=cwd" json:"cwd,omitempty"`
	Swd       *uint32  `protobuf:"varint,2,req,name=swd" json:"swd,omitempty"`
	Twd       *uint32  `protobuf:"varint,3,req,name=twd" json:"twd,omitempty"`
	Fop       *uint32  `protobuf:"varint,4,req,name=fop" json:"fop,omitempty"`
	Rip       *uint64  `protobuf:"varint,5,req,name=rip" json:"rip,omitempty"`
	Rdp       *uint64  `protobuf:"varint,6,req,name=rdp" json:"rdp,omitempty"`
	Mxcsr     *uint32  `protobuf:"varint,7,req,name=mxcsr" json:"mxcsr,omitempty"`
	MxcsrMask *uint32  `protobuf:"varint,8,req,name=mxcsr_mask,json=mxcsrMask" json:"mxcsr_mask,omitempty"`
	StSpace   []uint32 `protobuf:"varint,9,rep,name=st_space,json=stSpace" json:"st_space,omitempty"`
	XmmSpace  []uint32 `protobuf:"varint,10,rep,name=xmm_space,json=xmmSpace" json:"xmm_space,omitempty"`
	// Unused, but present for backward compatibility
	Padding []uint32 `protobuf:"varint,11,rep,name=padding" json:"padding,omitempty"`
	// xsave extension
	Xsave                *UserX86XsaveEntry `protobuf:"bytes,13,opt,name=xsave" json:"xsave,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserX86FpregsEntry) Reset()         { *m = UserX86FpregsEntry{} }
func (m *UserX86FpregsEntry) String() string { return proto.CompactTextString(m) }
func (*UserX86FpregsEntry) ProtoMessage()    {}
func (*UserX86FpregsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0422cff64b59e29, []int{2}
}

func (m *UserX86FpregsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserX86FpregsEntry.Unmarshal(m, b)
}
func (m *UserX86FpregsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserX86FpregsEntry.Marshal(b, m, deterministic)
}
func (m *UserX86FpregsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserX86FpregsEntry.Merge(m, src)
}
func (m *UserX86FpregsEntry) XXX_Size() int {
	return xxx_messageInfo_UserX86FpregsEntry.Size(m)
}
func (m *UserX86FpregsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserX86FpregsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserX86FpregsEntry proto.InternalMessageInfo

func (m *UserX86FpregsEntry) GetCwd() uint32 {
	if m != nil && m.Cwd != nil {
		return *m.Cwd
	}
	return 0
}

func (m *UserX86FpregsEntry) GetSwd() uint32 {
	if m != nil && m.Swd != nil {
		return *m.Swd
	}
	return 0
}

func (m *UserX86FpregsEntry) GetTwd() uint32 {
	if m != nil && m.Twd != nil {
		return *m.Twd
	}
	return 0
}

func (m *UserX86FpregsEntry) GetFop() uint32 {
	if m != nil && m.Fop != nil {
		return *m.Fop
	}
	return 0
}

func (m *UserX86FpregsEntry) GetRip() uint64 {
	if m != nil && m.Rip != nil {
		return *m.Rip
	}
	return 0
}

func (m *UserX86FpregsEntry) GetRdp() uint64 {
	if m != nil && m.Rdp != nil {
		return *m.Rdp
	}
	return 0
}

func (m *UserX86FpregsEntry) GetMxcsr() uint32 {
	if m != nil && m.Mxcsr != nil {
		return *m.Mxcsr
	}
	return 0
}

func (m *UserX86FpregsEntry) GetMxcsrMask() uint32 {
	if m != nil && m.MxcsrMask != nil {
		return *m.MxcsrMask
	}
	return 0
}

func (m *UserX86FpregsEntry) GetStSpace() []uint32 {
	if m != nil {
		return m.StSpace
	}
	return nil
}

func (m *UserX86FpregsEntry) GetXmmSpace() []uint32 {
	if m != nil {
		return m.XmmSpace
	}
	return nil
}

func (m *UserX86FpregsEntry) GetPadding() []uint32 {
	if m != nil {
		return m.Padding
	}
	return nil
}

func (m *UserX86FpregsEntry) GetXsave() *UserX86XsaveEntry {
	if m != nil {
		return m.Xsave
	}
	return nil
}

type UserDescT struct {
	EntryNumber *uint32 `protobuf:"varint,1,req,name=entry_number,json=entryNumber" json:"entry_number,omitempty"`
	// this is for GDT, not for MSRs - 32-bit base
	BaseAddr             *uint32  `protobuf:"varint,2,req,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
	Limit                *uint32  `protobuf:"varint,3,req,name=limit" json:"limit,omitempty"`
	Seg_32Bit            *bool    `protobuf:"varint,4,req,name=seg_32bit,json=seg32bit" json:"seg_32bit,omitempty"`
	ContentsH            *bool    `protobuf:"varint,5,req,name=contents_h,json=contentsH" json:"contents_h,omitempty"`
	ContentsL            *bool    `protobuf:"varint,6,req,name=contents_l,json=contentsL" json:"contents_l,omitempty"`
	ReadExecOnly         *bool    `protobuf:"varint,7,req,name=read_exec_only,json=readExecOnly,def=1" json:"read_exec_only,omitempty"`
	LimitInPages         *bool    `protobuf:"varint,8,req,name=limit_in_pages,json=limitInPages" json:"limit_in_pages,omitempty"`
	SegNotPresent        *bool    `protobuf:"varint,9,req,name=seg_not_present,json=segNotPresent,def=1" json:"seg_not_present,omitempty"`
	Useable              *bool    `protobuf:"varint,10,req,name=useable" json:"useable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDescT) Reset()         { *m = UserDescT{} }
func (m *UserDescT) String() string { return proto.CompactTextString(m) }
func (*UserDescT) ProtoMessage()    {}
func (*UserDescT) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0422cff64b59e29, []int{3}
}

func (m *UserDescT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDescT.Unmarshal(m, b)
}
func (m *UserDescT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDescT.Marshal(b, m, deterministic)
}
func (m *UserDescT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDescT.Merge(m, src)
}
func (m *UserDescT) XXX_Size() int {
	return xxx_messageInfo_UserDescT.Size(m)
}
func (m *UserDescT) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDescT.DiscardUnknown(m)
}

var xxx_messageInfo_UserDescT proto.InternalMessageInfo

const Default_UserDescT_ReadExecOnly bool = true
const Default_UserDescT_SegNotPresent bool = true

func (m *UserDescT) GetEntryNumber() uint32 {
	if m != nil && m.EntryNumber != nil {
		return *m.EntryNumber
	}
	return 0
}

func (m *UserDescT) GetBaseAddr() uint32 {
	if m != nil && m.BaseAddr != nil {
		return *m.BaseAddr
	}
	return 0
}

func (m *UserDescT) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *UserDescT) GetSeg_32Bit() bool {
	if m != nil && m.Seg_32Bit != nil {
		return *m.Seg_32Bit
	}
	return false
}

func (m *UserDescT) GetContentsH() bool {
	if m != nil && m.ContentsH != nil {
		return *m.ContentsH
	}
	return false
}

func (m *UserDescT) GetContentsL() bool {
	if m != nil && m.ContentsL != nil {
		return *m.ContentsL
	}
	return false
}

func (m *UserDescT) GetReadExecOnly() bool {
	if m != nil && m.ReadExecOnly != nil {
		return *m.ReadExecOnly
	}
	return Default_UserDescT_ReadExecOnly
}

func (m *UserDescT) GetLimitInPages() bool {
	if m != nil && m.LimitInPages != nil {
		return *m.LimitInPages
	}
	return false
}

func (m *UserDescT) GetSegNotPresent() bool {
	if m != nil && m.SegNotPresent != nil {
		return *m.SegNotPresent
	}
	return Default_UserDescT_SegNotPresent
}

func (m *UserDescT) GetUseable() bool {
	if m != nil && m.Useable != nil {
		return *m.Useable
	}
	return false
}

type ThreadInfoX86 struct {
	ClearTidAddr         *uint64             `protobuf:"varint,1,req,name=clear_tid_addr,json=clearTidAddr" json:"clear_tid_addr,omitempty"`
	Gpregs               *UserX86RegsEntry   `protobuf:"bytes,2,req,name=gpregs" json:"gpregs,omitempty"`
	Fpregs               *UserX86FpregsEntry `protobuf:"bytes,3,req,name=fpregs" json:"fpregs,omitempty"`
	Tls                  []*UserDescT        `protobuf:"bytes,4,rep,name=tls" json:"tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ThreadInfoX86) Reset()         { *m = ThreadInfoX86{} }
func (m *ThreadInfoX86) String() string { return proto.CompactTextString(m) }
func (*ThreadInfoX86) ProtoMessage()    {}
func (*ThreadInfoX86) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0422cff64b59e29, []int{4}
}

func (m *ThreadInfoX86) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadInfoX86.Unmarshal(m, b)
}
func (m *ThreadInfoX86) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadInfoX86.Marshal(b, m, deterministic)
}
func (m *ThreadInfoX86) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadInfoX86.Merge(m, src)
}
func (m *ThreadInfoX86) XXX_Size() int {
	return xxx_messageInfo_ThreadInfoX86.Size(m)
}
func (m *ThreadInfoX86) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadInfoX86.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadInfoX86 proto.InternalMessageInfo

func (m *ThreadInfoX86) GetClearTidAddr() uint64 {
	if m != nil && m.ClearTidAddr != nil {
		return *m.ClearTidAddr
	}
	return 0
}

func (m *ThreadInfoX86) GetGpregs() *UserX86RegsEntry {
	if m != nil {
		return m.Gpregs
	}
	return nil
}

func (m *ThreadInfoX86) GetFpregs() *UserX86FpregsEntry {
	if m != nil {
		return m.Fpregs
	}
	return nil
}

func (m *ThreadInfoX86) GetTls() []*UserDescT {
	if m != nil {
		return m.Tls
	}
	return nil
}

func init() {
	proto.RegisterEnum("UserX86RegsMode", UserX86RegsMode_name, UserX86RegsMode_value)
	proto.RegisterType((*UserX86RegsEntry)(nil), "user_x86_regs_entry")
	proto.RegisterType((*UserX86XsaveEntry)(nil), "user_x86_xsave_entry")
	proto.RegisterType((*UserX86FpregsEntry)(nil), "user_x86_fpregs_entry")
	proto.RegisterType((*UserDescT)(nil), "user_desc_t")
	proto.RegisterType((*ThreadInfoX86)(nil), "thread_info_x86")
}

func init() {
	proto.RegisterFile("core-x86.proto", fileDescriptor_b0422cff64b59e29)
}

var fileDescriptor_b0422cff64b59e29 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x55, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x85, 0x1e, 0x96, 0xa8, 0xd1, 0xc3, 0xea, 0xd8, 0x4e, 0x26, 0x49, 0x5b, 0xa4, 0x42, 0x17,
	0x46, 0x9a, 0x2a, 0xb1, 0x9c, 0x06, 0x4e, 0x36, 0x85, 0x5d, 0x04, 0x68, 0x80, 0xc6, 0x31, 0x18,
	0xb7, 0x8b, 0x6c, 0x06, 0x14, 0x39, 0xa2, 0x89, 0x88, 0x0f, 0x70, 0x28, 0x97, 0xee, 0xb6, 0x9f,
	0xd2, 0x8f, 0xe8, 0x07, 0xf4, 0x0f, 0xba, 0xea, 0xe7, 0xf4, 0x3e, 0x86, 0x86, 0x8d, 0x78, 0x37,
	0xe7, 0x9c, 0x7b, 0x75, 0xef, 0xdc, 0x39, 0x97, 0x12, 0x93, 0x30, 0x2f, 0xcd, 0xf7, 0xf5, 0xd1,
	0xcb, 0x79, 0x51, 0xe6, 0x55, 0xfe, 0x50, 0xe4, 0x45, 0x65, 0xf9, 0x3c, 0xfb, 0xab, 0x2b, 0x76,
	0x36, 0xd6, 0x94, 0x1a, 0x64, 0x5d, 0x9a, 0xd8, 0x6a, 0x93, 0x55, 0xe5, 0x95, 0x9c, 0x8a, 0x4e,
	0x79, 0xf0, 0x83, 0x6a, 0x3d, 0x6e, 0xef, 0x77, 0x7d, 0x3c, 0x32, 0xf3, 0x42, 0xb5, 0x1b, 0xe6,
	0x05, 0x33, 0x87, 0xaa, 0xd3, 0x30, 0x87, 0xcc, 0x2c, 0x54, 0xb7, 0x61, 0x16, 0x72, 0x22, 0xda,
	0xcb, 0x42, 0x6d, 0x11, 0x01, 0x27, 0xc2, 0xb5, 0xea, 0x39, 0x5c, 0x73, 0xc6, 0x81, 0xea, 0x37,
	0x19, 0x07, 0xcc, 0x3c, 0x57, 0x5e, 0xc3, 0x3c, 0xc7, 0x9c, 0xf2, 0x95, 0x1a, 0x70, 0x4e, 0xf9,
	0x8a, 0xf0, 0x91, 0x12, 0x0e, 0x1f, 0x21, 0x0e, 0x6a, 0x35, 0x64, 0x1c, 0xd4, 0x88, 0xc3, 0x5a,
	0x8d, 0x18, 0x87, 0x84, 0xa3, 0x5a, 0x8d, 0x19, 0x47, 0x84, 0x6d, 0xa2, 0x26, 0x8c, 0x6d, 0x42,
	0x7a, 0xa2, 0xb6, 0x9d, 0x9e, 0xc8, 0xfb, 0xa2, 0x9f, 0x97, 0x49, 0xac, 0xe1, 0x47, 0xa7, 0x44,
	0xf6, 0x10, 0x1e, 0x53, 0x62, 0x52, 0xa8, 0x2f, 0x38, 0x30, 0xa1, 0xcb, 0x84, 0x56, 0x49, 0x57,
	0xc8, 0xca, 0x5d, 0xb1, 0xb5, 0x5a, 0x07, 0xb1, 0x55, 0x3b, 0x44, 0x31, 0xa0, 0x72, 0x85, 0xda,
	0x75, 0xe5, 0x28, 0xcb, 0x5a, 0xb5, 0xe7, 0xb0, 0xc5, 0x72, 0x2b, 0xab, 0x97, 0x81, 0x35, 0xea,
	0x1e, 0x97, 0x5b, 0xd9, 0x13, 0x40, 0x28, 0xc4, 0x4e, 0xb8, 0xcf, 0x42, 0xcc, 0x02, 0x36, 0x6c,
	0x95, 0x72, 0x0d, 0x53, 0x05, 0x63, 0xd5, 0x03, 0xc6, 0x86, 0xf0, 0xca, 0xaa, 0x87, 0x8c, 0x57,
	0x84, 0xa1, 0xa9, 0x47, 0x8c, 0xa1, 0xa3, 0x67, 0xa2, 0x9b, 0xe6, 0x91, 0x51, 0x5f, 0x3e, 0x6e,
	0xed, 0x4f, 0x16, 0x3b, 0xf3, 0xdb, 0x06, 0x40, 0xe9, 0x75, 0xef, 0xf4, 0xf8, 0xfc, 0xed, 0x6f,
	0x6f, 0x7c, 0x0a, 0x9c, 0xfd, 0xd9, 0x16, 0xbb, 0xd7, 0x41, 0xb5, 0x0d, 0x2e, 0x8d, 0xb3, 0xc9,
	0x23, 0x31, 0xa8, 0x6d, 0x15, 0x54, 0x46, 0x2f, 0x2f, 0x9d, 0x59, 0x3c, 0x26, 0x4e, 0x2e, 0xe5,
	0x57, 0x42, 0x5c, 0xa5, 0xe9, 0x85, 0xb6, 0x45, 0x10, 0x1a, 0x30, 0x4e, 0x67, 0x7f, 0xec, 0x0f,
	0x90, 0xf9, 0x80, 0x84, 0xfc, 0x46, 0x8c, 0x96, 0x59, 0x04, 0x25, 0x35, 0x25, 0x80, 0x8f, 0x3a,
	0x90, 0x3e, 0x64, 0xee, 0x03, 0x52, 0x2e, 0x24, 0xb4, 0xa5, 0x0b, 0xe9, 0x5e, 0x87, 0x00, 0xc7,
	0x21, 0x50, 0x24, 0x2f, 0xd2, 0xc0, 0x7e, 0xc2, 0xe6, 0xc1, 0x68, 0x18, 0x30, 0x60, 0xc6, 0x37,
	0x31, 0x36, 0xf8, 0x47, 0x9a, 0xea, 0x4d, 0x51, 0x98, 0x12, 0x6c, 0x87, 0xaa, 0x07, 0xc4, 0xaf,
	0x88, 0xe5, 0x03, 0xe1, 0x5d, 0x24, 0x07, 0x2f, 0x35, 0x10, 0xe0, 0x40, 0xd4, 0xfa, 0x88, 0x3f,
	0xa6, 0xa9, 0x94, 0xa2, 0x5b, 0x7c, 0x2a, 0x37, 0x60, 0x43, 0xec, 0x9a, 0xce, 0xb3, 0xbf, 0xdb,
	0x62, 0xef, 0x7a, 0x0a, 0xab, 0xe2, 0xf6, 0xb6, 0x84, 0xbf, 0x47, 0x34, 0x80, 0xb1, 0x8f, 0x47,
	0x64, 0x2c, 0x30, 0x6d, 0x66, 0x2c, 0x33, 0x15, 0x30, 0x1d, 0x66, 0x2a, 0x66, 0x56, 0x79, 0x41,
	0xdb, 0x02, 0x0c, 0x1c, 0xc9, 0xfb, 0x49, 0xb3, 0x2e, 0x78, 0x24, 0x26, 0x2a, 0xdc, 0xc2, 0xe0,
	0x11, 0x4d, 0x96, 0xd6, 0x70, 0x7d, 0xda, 0x99, 0xb1, 0xcf, 0x00, 0xc7, 0x40, 0x07, 0x8d, 0x17,
	0xa7, 0xe5, 0x81, 0x59, 0x13, 0xf3, 0x0e, 0x08, 0xbc, 0xa9, 0xad, 0xdc, 0x43, 0x0c, 0xe8, 0x4a,
	0x7d, 0x5b, 0xf1, 0x33, 0xe0, 0x13, 0xc2, 0x84, 0x58, 0x13, 0xa4, 0x79, 0x40, 0xb0, 0xa8, 0x44,
	0xbf, 0x08, 0xa2, 0x28, 0xc9, 0x62, 0xd8, 0x2f, 0x4a, 0x73, 0x50, 0x7e, 0x27, 0xb6, 0xc8, 0x08,
	0xb0, 0x57, 0xad, 0xfd, 0xe1, 0x62, 0x6f, 0x7e, 0x97, 0x3f, 0x7c, 0x8e, 0x99, 0xfd, 0xd7, 0x16,
	0x43, 0xd2, 0x23, 0x63, 0x43, 0x5d, 0xe1, 0xbb, 0x92, 0xae, 0xb3, 0x4d, 0xba, 0x84, 0x87, 0xe1,
	0xc1, 0x0d, 0x89, 0x3b, 0x25, 0x0a, 0xdb, 0x42, 0xe7, 0x6b, 0x28, 0x57, 0xba, 0x31, 0x7a, 0x48,
	0x1c, 0x03, 0xc6, 0x19, 0xac, 0x93, 0x34, 0xa9, 0xdc, 0x34, 0x19, 0x60, 0x8a, 0x05, 0x37, 0x1d,
	0x2e, 0x96, 0xa0, 0xe0, 0x54, 0x3d, 0xdf, 0x03, 0x82, 0x30, 0x0e, 0x28, 0xcc, 0xb3, 0x0a, 0x4a,
	0x58, 0x7d, 0x41, 0x13, 0xf6, 0xfc, 0x41, 0xc3, 0xfc, 0x7c, 0x4b, 0x5e, 0xd3, 0xb8, 0x6f, 0xc8,
	0xbf, 0xc8, 0x27, 0x62, 0x52, 0x9a, 0x20, 0xd2, 0xa6, 0x36, 0xa1, 0xce, 0xb3, 0xf5, 0x15, 0x4d,
	0xdf, 0x7b, 0xdd, 0xad, 0xca, 0x8d, 0xf1, 0x47, 0xa8, 0xbd, 0x01, 0xe9, 0x3d, 0x28, 0xf2, 0x5b,
	0x31, 0xa1, 0x7e, 0x74, 0x92, 0xe9, 0x22, 0x88, 0x61, 0x33, 0x3d, 0xfa, 0xb9, 0x11, 0xb1, 0x6f,
	0xb3, 0x33, 0xe4, 0xe4, 0x53, 0xb1, 0x8d, 0xcd, 0x66, 0x79, 0xa5, 0xc1, 0x49, 0x16, 0xca, 0xd0,
	0x17, 0xae, 0xf9, 0xc9, 0x31, 0x88, 0xa7, 0x79, 0x75, 0xc6, 0x12, 0xbe, 0x03, 0xcc, 0x2f, 0x58,
	0xae, 0x0d, 0x7d, 0xf7, 0x3c, 0xbf, 0x81, 0xb3, 0x7f, 0x5a, 0x62, 0xbb, 0xba, 0xa0, 0xe6, 0x92,
	0x6c, 0x95, 0xe3, 0x0b, 0xc0, 0xdb, 0x4c, 0xc2, 0xb5, 0x09, 0x4a, 0x5d, 0x25, 0x11, 0x0f, 0x90,
	0x56, 0xf3, 0x64, 0xeb, 0xdf, 0x1f, 0xdb, 0x5e, 0xcb, 0x1f, 0x91, 0x78, 0x9e, 0x44, 0x34, 0xcb,
	0x43, 0xd1, 0x8b, 0xc9, 0xcb, 0x34, 0xe5, 0xe1, 0x62, 0x77, 0x7e, 0xc7, 0xff, 0x41, 0x93, 0xea,
	0x42, 0xe5, 0x5c, 0xf4, 0x78, 0x01, 0xe8, 0x05, 0x86, 0x8b, 0x7b, 0xf3, 0x3b, 0x17, 0xc3, 0x77,
	0x51, 0xf2, 0x6b, 0x30, 0xff, 0xda, 0xd2, 0xfe, 0x0e, 0x17, 0xa3, 0xf9, 0x0d, 0x2f, 0xf8, 0x28,
	0x3c, 0x79, 0x2a, 0xe4, 0xe7, 0x1f, 0x21, 0x29, 0x84, 0xfb, 0x0c, 0x4d, 0x5b, 0x78, 0xfe, 0xe9,
	0xfd, 0xbb, 0xb3, 0xe3, 0xf3, 0x69, 0xfb, 0x64, 0xfa, 0x71, 0xf2, 0x8c, 0xfe, 0xbe, 0x96, 0x49,
	0x86, 0x66, 0xb4, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x72, 0xdf, 0x26, 0xe3, 0x06, 0x00,
	0x00,
}
