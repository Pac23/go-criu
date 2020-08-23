// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core-mips.proto

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

type UserMipsRegsEntry struct {
	R0                   *uint64  `protobuf:"varint,1,req,name=r0" json:"r0,omitempty"`
	R1                   *uint64  `protobuf:"varint,2,req,name=r1" json:"r1,omitempty"`
	R2                   *uint64  `protobuf:"varint,3,req,name=r2" json:"r2,omitempty"`
	R3                   *uint64  `protobuf:"varint,4,req,name=r3" json:"r3,omitempty"`
	R4                   *uint64  `protobuf:"varint,5,req,name=r4" json:"r4,omitempty"`
	R5                   *uint64  `protobuf:"varint,6,req,name=r5" json:"r5,omitempty"`
	R6                   *uint64  `protobuf:"varint,7,req,name=r6" json:"r6,omitempty"`
	R7                   *uint64  `protobuf:"varint,8,req,name=r7" json:"r7,omitempty"`
	R8                   *uint64  `protobuf:"varint,9,req,name=r8" json:"r8,omitempty"`
	R9                   *uint64  `protobuf:"varint,10,req,name=r9" json:"r9,omitempty"`
	R10                  *uint64  `protobuf:"varint,11,req,name=r10" json:"r10,omitempty"`
	R11                  *uint64  `protobuf:"varint,12,req,name=r11" json:"r11,omitempty"`
	R12                  *uint64  `protobuf:"varint,13,req,name=r12" json:"r12,omitempty"`
	R13                  *uint64  `protobuf:"varint,14,req,name=r13" json:"r13,omitempty"`
	R14                  *uint64  `protobuf:"varint,15,req,name=r14" json:"r14,omitempty"`
	R15                  *uint64  `protobuf:"varint,16,req,name=r15" json:"r15,omitempty"`
	R16                  *uint64  `protobuf:"varint,17,req,name=r16" json:"r16,omitempty"`
	R17                  *uint64  `protobuf:"varint,18,req,name=r17" json:"r17,omitempty"`
	R18                  *uint64  `protobuf:"varint,19,req,name=r18" json:"r18,omitempty"`
	R19                  *uint64  `protobuf:"varint,20,req,name=r19" json:"r19,omitempty"`
	R20                  *uint64  `protobuf:"varint,21,req,name=r20" json:"r20,omitempty"`
	R21                  *uint64  `protobuf:"varint,22,req,name=r21" json:"r21,omitempty"`
	R22                  *uint64  `protobuf:"varint,23,req,name=r22" json:"r22,omitempty"`
	R23                  *uint64  `protobuf:"varint,24,req,name=r23" json:"r23,omitempty"`
	R24                  *uint64  `protobuf:"varint,25,req,name=r24" json:"r24,omitempty"`
	R25                  *uint64  `protobuf:"varint,26,req,name=r25" json:"r25,omitempty"`
	R26                  *uint64  `protobuf:"varint,27,req,name=r26" json:"r26,omitempty"`
	R27                  *uint64  `protobuf:"varint,28,req,name=r27" json:"r27,omitempty"`
	R28                  *uint64  `protobuf:"varint,29,req,name=r28" json:"r28,omitempty"`
	R29                  *uint64  `protobuf:"varint,30,req,name=r29" json:"r29,omitempty"`
	R30                  *uint64  `protobuf:"varint,31,req,name=r30" json:"r30,omitempty"`
	R31                  *uint64  `protobuf:"varint,32,req,name=r31" json:"r31,omitempty"`
	Lo                   *uint64  `protobuf:"varint,33,req,name=lo" json:"lo,omitempty"`
	Hi                   *uint64  `protobuf:"varint,34,req,name=hi" json:"hi,omitempty"`
	Cp0Epc               *uint64  `protobuf:"varint,35,req,name=cp0_epc,json=cp0Epc" json:"cp0_epc,omitempty"`
	Cp0Badvaddr          *uint64  `protobuf:"varint,36,req,name=cp0_badvaddr,json=cp0Badvaddr" json:"cp0_badvaddr,omitempty"`
	Cp0Status            *uint64  `protobuf:"varint,37,req,name=cp0_status,json=cp0Status" json:"cp0_status,omitempty"`
	Cp0Cause             *uint64  `protobuf:"varint,38,req,name=cp0_cause,json=cp0Cause" json:"cp0_cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMipsRegsEntry) Reset()         { *m = UserMipsRegsEntry{} }
func (m *UserMipsRegsEntry) String() string { return proto.CompactTextString(m) }
func (*UserMipsRegsEntry) ProtoMessage()    {}
func (*UserMipsRegsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c3e8602be8a6a2b, []int{0}
}

func (m *UserMipsRegsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMipsRegsEntry.Unmarshal(m, b)
}
func (m *UserMipsRegsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMipsRegsEntry.Marshal(b, m, deterministic)
}
func (m *UserMipsRegsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMipsRegsEntry.Merge(m, src)
}
func (m *UserMipsRegsEntry) XXX_Size() int {
	return xxx_messageInfo_UserMipsRegsEntry.Size(m)
}
func (m *UserMipsRegsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMipsRegsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserMipsRegsEntry proto.InternalMessageInfo

func (m *UserMipsRegsEntry) GetR0() uint64 {
	if m != nil && m.R0 != nil {
		return *m.R0
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR1() uint64 {
	if m != nil && m.R1 != nil {
		return *m.R1
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR2() uint64 {
	if m != nil && m.R2 != nil {
		return *m.R2
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR3() uint64 {
	if m != nil && m.R3 != nil {
		return *m.R3
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR4() uint64 {
	if m != nil && m.R4 != nil {
		return *m.R4
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR5() uint64 {
	if m != nil && m.R5 != nil {
		return *m.R5
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR6() uint64 {
	if m != nil && m.R6 != nil {
		return *m.R6
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR7() uint64 {
	if m != nil && m.R7 != nil {
		return *m.R7
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR8() uint64 {
	if m != nil && m.R8 != nil {
		return *m.R8
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR9() uint64 {
	if m != nil && m.R9 != nil {
		return *m.R9
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR10() uint64 {
	if m != nil && m.R10 != nil {
		return *m.R10
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR11() uint64 {
	if m != nil && m.R11 != nil {
		return *m.R11
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR12() uint64 {
	if m != nil && m.R12 != nil {
		return *m.R12
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR13() uint64 {
	if m != nil && m.R13 != nil {
		return *m.R13
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR14() uint64 {
	if m != nil && m.R14 != nil {
		return *m.R14
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR15() uint64 {
	if m != nil && m.R15 != nil {
		return *m.R15
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR16() uint64 {
	if m != nil && m.R16 != nil {
		return *m.R16
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR17() uint64 {
	if m != nil && m.R17 != nil {
		return *m.R17
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR18() uint64 {
	if m != nil && m.R18 != nil {
		return *m.R18
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR19() uint64 {
	if m != nil && m.R19 != nil {
		return *m.R19
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR20() uint64 {
	if m != nil && m.R20 != nil {
		return *m.R20
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR21() uint64 {
	if m != nil && m.R21 != nil {
		return *m.R21
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR22() uint64 {
	if m != nil && m.R22 != nil {
		return *m.R22
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR23() uint64 {
	if m != nil && m.R23 != nil {
		return *m.R23
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR24() uint64 {
	if m != nil && m.R24 != nil {
		return *m.R24
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR25() uint64 {
	if m != nil && m.R25 != nil {
		return *m.R25
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR26() uint64 {
	if m != nil && m.R26 != nil {
		return *m.R26
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR27() uint64 {
	if m != nil && m.R27 != nil {
		return *m.R27
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR28() uint64 {
	if m != nil && m.R28 != nil {
		return *m.R28
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR29() uint64 {
	if m != nil && m.R29 != nil {
		return *m.R29
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR30() uint64 {
	if m != nil && m.R30 != nil {
		return *m.R30
	}
	return 0
}

func (m *UserMipsRegsEntry) GetR31() uint64 {
	if m != nil && m.R31 != nil {
		return *m.R31
	}
	return 0
}

func (m *UserMipsRegsEntry) GetLo() uint64 {
	if m != nil && m.Lo != nil {
		return *m.Lo
	}
	return 0
}

func (m *UserMipsRegsEntry) GetHi() uint64 {
	if m != nil && m.Hi != nil {
		return *m.Hi
	}
	return 0
}

func (m *UserMipsRegsEntry) GetCp0Epc() uint64 {
	if m != nil && m.Cp0Epc != nil {
		return *m.Cp0Epc
	}
	return 0
}

func (m *UserMipsRegsEntry) GetCp0Badvaddr() uint64 {
	if m != nil && m.Cp0Badvaddr != nil {
		return *m.Cp0Badvaddr
	}
	return 0
}

func (m *UserMipsRegsEntry) GetCp0Status() uint64 {
	if m != nil && m.Cp0Status != nil {
		return *m.Cp0Status
	}
	return 0
}

func (m *UserMipsRegsEntry) GetCp0Cause() uint64 {
	if m != nil && m.Cp0Cause != nil {
		return *m.Cp0Cause
	}
	return 0
}

type UserMipsFpregsEntry struct {
	R0                   *uint64  `protobuf:"varint,1,req,name=r0" json:"r0,omitempty"`
	R1                   *uint64  `protobuf:"varint,2,req,name=r1" json:"r1,omitempty"`
	R2                   *uint64  `protobuf:"varint,3,req,name=r2" json:"r2,omitempty"`
	R3                   *uint64  `protobuf:"varint,4,req,name=r3" json:"r3,omitempty"`
	R4                   *uint64  `protobuf:"varint,5,req,name=r4" json:"r4,omitempty"`
	R5                   *uint64  `protobuf:"varint,6,req,name=r5" json:"r5,omitempty"`
	R6                   *uint64  `protobuf:"varint,7,req,name=r6" json:"r6,omitempty"`
	R7                   *uint64  `protobuf:"varint,8,req,name=r7" json:"r7,omitempty"`
	R8                   *uint64  `protobuf:"varint,9,req,name=r8" json:"r8,omitempty"`
	R9                   *uint64  `protobuf:"varint,10,req,name=r9" json:"r9,omitempty"`
	R10                  *uint64  `protobuf:"varint,11,req,name=r10" json:"r10,omitempty"`
	R11                  *uint64  `protobuf:"varint,12,req,name=r11" json:"r11,omitempty"`
	R12                  *uint64  `protobuf:"varint,13,req,name=r12" json:"r12,omitempty"`
	R13                  *uint64  `protobuf:"varint,14,req,name=r13" json:"r13,omitempty"`
	R14                  *uint64  `protobuf:"varint,15,req,name=r14" json:"r14,omitempty"`
	R15                  *uint64  `protobuf:"varint,16,req,name=r15" json:"r15,omitempty"`
	R16                  *uint64  `protobuf:"varint,17,req,name=r16" json:"r16,omitempty"`
	R17                  *uint64  `protobuf:"varint,18,req,name=r17" json:"r17,omitempty"`
	R18                  *uint64  `protobuf:"varint,19,req,name=r18" json:"r18,omitempty"`
	R19                  *uint64  `protobuf:"varint,20,req,name=r19" json:"r19,omitempty"`
	R20                  *uint64  `protobuf:"varint,21,req,name=r20" json:"r20,omitempty"`
	R21                  *uint64  `protobuf:"varint,22,req,name=r21" json:"r21,omitempty"`
	R22                  *uint64  `protobuf:"varint,23,req,name=r22" json:"r22,omitempty"`
	R23                  *uint64  `protobuf:"varint,24,req,name=r23" json:"r23,omitempty"`
	R24                  *uint64  `protobuf:"varint,25,req,name=r24" json:"r24,omitempty"`
	R25                  *uint64  `protobuf:"varint,26,req,name=r25" json:"r25,omitempty"`
	R26                  *uint64  `protobuf:"varint,27,req,name=r26" json:"r26,omitempty"`
	R27                  *uint64  `protobuf:"varint,28,req,name=r27" json:"r27,omitempty"`
	R28                  *uint64  `protobuf:"varint,29,req,name=r28" json:"r28,omitempty"`
	R29                  *uint64  `protobuf:"varint,30,req,name=r29" json:"r29,omitempty"`
	R30                  *uint64  `protobuf:"varint,31,req,name=r30" json:"r30,omitempty"`
	R31                  *uint64  `protobuf:"varint,32,req,name=r31" json:"r31,omitempty"`
	Lo                   *uint64  `protobuf:"varint,33,req,name=lo" json:"lo,omitempty"`
	Hi                   *uint64  `protobuf:"varint,34,req,name=hi" json:"hi,omitempty"`
	FpuFcr31             *uint32  `protobuf:"varint,35,req,name=fpu_fcr31,json=fpuFcr31" json:"fpu_fcr31,omitempty"`
	FpuId                *uint32  `protobuf:"varint,36,req,name=fpu_id,json=fpuId" json:"fpu_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMipsFpregsEntry) Reset()         { *m = UserMipsFpregsEntry{} }
func (m *UserMipsFpregsEntry) String() string { return proto.CompactTextString(m) }
func (*UserMipsFpregsEntry) ProtoMessage()    {}
func (*UserMipsFpregsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c3e8602be8a6a2b, []int{1}
}

func (m *UserMipsFpregsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMipsFpregsEntry.Unmarshal(m, b)
}
func (m *UserMipsFpregsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMipsFpregsEntry.Marshal(b, m, deterministic)
}
func (m *UserMipsFpregsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMipsFpregsEntry.Merge(m, src)
}
func (m *UserMipsFpregsEntry) XXX_Size() int {
	return xxx_messageInfo_UserMipsFpregsEntry.Size(m)
}
func (m *UserMipsFpregsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMipsFpregsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserMipsFpregsEntry proto.InternalMessageInfo

func (m *UserMipsFpregsEntry) GetR0() uint64 {
	if m != nil && m.R0 != nil {
		return *m.R0
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR1() uint64 {
	if m != nil && m.R1 != nil {
		return *m.R1
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR2() uint64 {
	if m != nil && m.R2 != nil {
		return *m.R2
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR3() uint64 {
	if m != nil && m.R3 != nil {
		return *m.R3
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR4() uint64 {
	if m != nil && m.R4 != nil {
		return *m.R4
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR5() uint64 {
	if m != nil && m.R5 != nil {
		return *m.R5
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR6() uint64 {
	if m != nil && m.R6 != nil {
		return *m.R6
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR7() uint64 {
	if m != nil && m.R7 != nil {
		return *m.R7
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR8() uint64 {
	if m != nil && m.R8 != nil {
		return *m.R8
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR9() uint64 {
	if m != nil && m.R9 != nil {
		return *m.R9
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR10() uint64 {
	if m != nil && m.R10 != nil {
		return *m.R10
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR11() uint64 {
	if m != nil && m.R11 != nil {
		return *m.R11
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR12() uint64 {
	if m != nil && m.R12 != nil {
		return *m.R12
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR13() uint64 {
	if m != nil && m.R13 != nil {
		return *m.R13
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR14() uint64 {
	if m != nil && m.R14 != nil {
		return *m.R14
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR15() uint64 {
	if m != nil && m.R15 != nil {
		return *m.R15
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR16() uint64 {
	if m != nil && m.R16 != nil {
		return *m.R16
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR17() uint64 {
	if m != nil && m.R17 != nil {
		return *m.R17
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR18() uint64 {
	if m != nil && m.R18 != nil {
		return *m.R18
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR19() uint64 {
	if m != nil && m.R19 != nil {
		return *m.R19
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR20() uint64 {
	if m != nil && m.R20 != nil {
		return *m.R20
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR21() uint64 {
	if m != nil && m.R21 != nil {
		return *m.R21
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR22() uint64 {
	if m != nil && m.R22 != nil {
		return *m.R22
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR23() uint64 {
	if m != nil && m.R23 != nil {
		return *m.R23
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR24() uint64 {
	if m != nil && m.R24 != nil {
		return *m.R24
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR25() uint64 {
	if m != nil && m.R25 != nil {
		return *m.R25
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR26() uint64 {
	if m != nil && m.R26 != nil {
		return *m.R26
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR27() uint64 {
	if m != nil && m.R27 != nil {
		return *m.R27
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR28() uint64 {
	if m != nil && m.R28 != nil {
		return *m.R28
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR29() uint64 {
	if m != nil && m.R29 != nil {
		return *m.R29
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR30() uint64 {
	if m != nil && m.R30 != nil {
		return *m.R30
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetR31() uint64 {
	if m != nil && m.R31 != nil {
		return *m.R31
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetLo() uint64 {
	if m != nil && m.Lo != nil {
		return *m.Lo
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetHi() uint64 {
	if m != nil && m.Hi != nil {
		return *m.Hi
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetFpuFcr31() uint32 {
	if m != nil && m.FpuFcr31 != nil {
		return *m.FpuFcr31
	}
	return 0
}

func (m *UserMipsFpregsEntry) GetFpuId() uint32 {
	if m != nil && m.FpuId != nil {
		return *m.FpuId
	}
	return 0
}

type ThreadInfoMips struct {
	ClearTidAddr         *uint64              `protobuf:"varint,1,req,name=clear_tid_addr,json=clearTidAddr" json:"clear_tid_addr,omitempty"`
	Tls                  *uint64              `protobuf:"varint,2,req,name=tls" json:"tls,omitempty"`
	Gpregs               *UserMipsRegsEntry   `protobuf:"bytes,3,req,name=gpregs" json:"gpregs,omitempty"`
	Fpregs               *UserMipsFpregsEntry `protobuf:"bytes,4,req,name=fpregs" json:"fpregs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ThreadInfoMips) Reset()         { *m = ThreadInfoMips{} }
func (m *ThreadInfoMips) String() string { return proto.CompactTextString(m) }
func (*ThreadInfoMips) ProtoMessage()    {}
func (*ThreadInfoMips) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c3e8602be8a6a2b, []int{2}
}

func (m *ThreadInfoMips) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadInfoMips.Unmarshal(m, b)
}
func (m *ThreadInfoMips) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadInfoMips.Marshal(b, m, deterministic)
}
func (m *ThreadInfoMips) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadInfoMips.Merge(m, src)
}
func (m *ThreadInfoMips) XXX_Size() int {
	return xxx_messageInfo_ThreadInfoMips.Size(m)
}
func (m *ThreadInfoMips) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadInfoMips.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadInfoMips proto.InternalMessageInfo

func (m *ThreadInfoMips) GetClearTidAddr() uint64 {
	if m != nil && m.ClearTidAddr != nil {
		return *m.ClearTidAddr
	}
	return 0
}

func (m *ThreadInfoMips) GetTls() uint64 {
	if m != nil && m.Tls != nil {
		return *m.Tls
	}
	return 0
}

func (m *ThreadInfoMips) GetGpregs() *UserMipsRegsEntry {
	if m != nil {
		return m.Gpregs
	}
	return nil
}

func (m *ThreadInfoMips) GetFpregs() *UserMipsFpregsEntry {
	if m != nil {
		return m.Fpregs
	}
	return nil
}

func init() {
	proto.RegisterType((*UserMipsRegsEntry)(nil), "user_mips_regs_entry")
	proto.RegisterType((*UserMipsFpregsEntry)(nil), "user_mips_fpregs_entry")
	proto.RegisterType((*ThreadInfoMips)(nil), "thread_info_mips")
}

func init() {
	proto.RegisterFile("core-mips.proto", fileDescriptor_3c3e8602be8a6a2b)
}

var fileDescriptor_3c3e8602be8a6a2b = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xd5, 0x34, 0x49, 0xd3, 0xcd, 0x47, 0xc3, 0xd2, 0x34, 0x43, 0x4b, 0xa1, 0x0d, 0x1f,
	0x42, 0x42, 0x84, 0x78, 0xed, 0x38, 0xce, 0x09, 0x11, 0x04, 0x12, 0xd7, 0xc0, 0x89, 0x8b, 0xe5,
	0xda, 0x4e, 0x62, 0x29, 0xc4, 0x96, 0xed, 0x20, 0xf1, 0x26, 0xbc, 0x0d, 0xef, 0xc0, 0x91, 0xa7,
	0xc1, 0x3b, 0xb3, 0x03, 0x97, 0x5e, 0xb9, 0xf5, 0xb6, 0xbf, 0xdf, 0xec, 0x44, 0x7f, 0x29, 0xff,
	0x4d, 0xc4, 0x49, 0x98, 0xe6, 0xf1, 0xab, 0xaf, 0x49, 0x56, 0x8c, 0xb3, 0x3c, 0x2d, 0xd3, 0x73,
	0x91, 0x66, 0xa5, 0x39, 0x8f, 0x7e, 0x37, 0xc4, 0xe9, 0xbe, 0x88, 0x73, 0x5f, 0xcf, 0xfd, 0x3c,
	0x5e, 0x17, 0x7e, 0xbc, 0x2b, 0xf3, 0xef, 0xb2, 0x27, 0x6a, 0xf9, 0x04, 0x0e, 0xae, 0x6a, 0x2f,
	0xea, 0xcb, 0xea, 0x84, 0x6c, 0x41, 0xcd, 0xb0, 0x85, 0xac, 0xe0, 0xd0, 0xb0, 0x42, 0xb6, 0xa1,
	0x6e, 0xd8, 0x46, 0x76, 0xa0, 0x61, 0xd8, 0x41, 0x9e, 0x42, 0xd3, 0xf0, 0x14, 0xd9, 0x85, 0x23,
	0xc3, 0x2e, 0xf2, 0x0c, 0x5a, 0x86, 0x67, 0xc8, 0x1e, 0x1c, 0x1b, 0xf6, 0x90, 0xe7, 0x20, 0x0c,
	0xcf, 0x65, 0x5f, 0x1c, 0xe6, 0xd6, 0x04, 0xda, 0x28, 0xf4, 0x91, 0x8c, 0x05, 0x1d, 0x36, 0x16,
	0x19, 0x05, 0x5d, 0x36, 0x8a, 0x8c, 0x0d, 0x3d, 0x36, 0x36, 0x19, 0x07, 0x4e, 0xd8, 0x38, 0x64,
	0xa6, 0xd0, 0x67, 0x33, 0x25, 0xe3, 0xc2, 0x3d, 0x36, 0x2e, 0x99, 0x19, 0x48, 0x36, 0x33, 0x32,
	0x1e, 0xdc, 0x67, 0xe3, 0x91, 0x99, 0xc3, 0x29, 0x1b, 0xca, 0xac, 0x26, 0x30, 0x30, 0x46, 0x51,
	0x66, 0x65, 0xc1, 0x19, 0x1b, 0xca, 0xac, 0x14, 0x0c, 0xd9, 0x50, 0x66, 0x65, 0x03, 0xb0, 0xa1,
	0xcc, 0xca, 0x81, 0x07, 0x6c, 0x28, 0xb3, 0x9a, 0xc2, 0x39, 0x1b, 0xca, 0xac, 0x5c, 0xb8, 0x60,
	0x43, 0x99, 0xd5, 0x0c, 0x1e, 0xb2, 0xa1, 0xcc, 0xca, 0x83, 0x4b, 0x36, 0x94, 0x59, 0xcd, 0xe1,
	0x11, 0x1b, 0xca, 0x6c, 0x4f, 0xe0, 0xb1, 0x31, 0x36, 0x65, 0xb6, 0x2d, 0xb8, 0x62, 0x83, 0x5d,
	0xd8, 0xa6, 0x70, 0x4d, 0xdf, 0xcd, 0x36, 0xd5, 0xbc, 0x49, 0x60, 0x44, 0xbc, 0x49, 0xe4, 0x50,
	0x1c, 0x85, 0xd9, 0xc4, 0x8f, 0xb3, 0x10, 0x9e, 0xa0, 0x6c, 0x56, 0xf8, 0x3e, 0x0b, 0xe5, 0xb5,
	0xe8, 0xe8, 0xc1, 0x4d, 0x10, 0x7d, 0x0b, 0xa2, 0x28, 0x87, 0xa7, 0x38, 0x6d, 0x57, 0x6e, 0x61,
	0x94, 0xbc, 0x14, 0x42, 0x5f, 0x29, 0xca, 0xa0, 0xdc, 0x17, 0xf0, 0x0c, 0x2f, 0x1c, 0x57, 0xe6,
	0x13, 0x0a, 0x79, 0x21, 0x34, 0xf8, 0x61, 0x50, 0x95, 0x18, 0x9e, 0xe3, 0xb4, 0x55, 0x89, 0x77,
	0x9a, 0x47, 0x3f, 0x1a, 0xe2, 0xec, 0x5f, 0xb9, 0x57, 0xd9, 0x5d, 0xbd, 0xef, 0xea, 0xfd, 0x1f,
	0xea, 0x5d, 0x75, 0x70, 0x95, 0xed, 0xfd, 0x55, 0xa8, 0xf7, 0x74, 0xc1, 0xbb, 0xcb, 0x56, 0x25,
	0x3e, 0x68, 0x96, 0x03, 0xd1, 0xd4, 0xc3, 0x24, 0xc2, 0x72, 0x77, 0x97, 0x8d, 0x8a, 0x3e, 0x46,
	0xa3, 0x9f, 0x07, 0xa2, 0x5f, 0x6e, 0xf2, 0x38, 0x88, 0xfc, 0x64, 0xb7, 0x4a, 0xb1, 0xa1, 0xf2,
	0xa5, 0xe8, 0x85, 0xdb, 0x38, 0xc8, 0xfd, 0x32, 0x89, 0x7c, 0x7c, 0x10, 0x58, 0xd0, 0x45, 0xe3,
	0xd7, 0x9b, 0x5a, 0xeb, 0x60, 0xd9, 0xc1, 0xe1, 0xe7, 0x24, 0x7a, 0xab, 0x1f, 0x46, 0x95, 0xb3,
	0xdc, 0x16, 0xa6, 0xb2, 0xfa, 0x28, 0x1d, 0xd1, 0x5c, 0x63, 0xc7, 0xb1, 0xb7, 0x6d, 0x35, 0x18,
	0xdf, 0xf6, 0xcb, 0xce, 0x9f, 0x66, 0xee, 0x4a, 0x57, 0x07, 0xc4, 0xad, 0x3a, 0x6e, 0x0d, 0xc7,
	0xb7, 0x3f, 0x99, 0xbf, 0x7b, 0x24, 0x17, 0xfd, 0x2f, 0xbd, 0xd7, 0xf8, 0x1f, 0x72, 0x93, 0xec,
	0xa2, 0x64, 0xb7, 0x2e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x14, 0x92, 0x13, 0xdc, 0x69, 0x06,
	0x00, 0x00,
}
