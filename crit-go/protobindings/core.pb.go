// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

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

//
// These match the SECCOMP_MODE_* flags from <linux/seccomp.h>.
type SeccompMode int32

const (
	SeccompMode_disabled SeccompMode = 0
	SeccompMode_strict   SeccompMode = 1
	SeccompMode_filter   SeccompMode = 2
)

var SeccompMode_name = map[int32]string{
	0: "disabled",
	1: "strict",
	2: "filter",
}

var SeccompMode_value = map[string]int32{
	"disabled": 0,
	"strict":   1,
	"filter":   2,
}

func (x SeccompMode) Enum() *SeccompMode {
	p := new(SeccompMode)
	*p = x
	return p
}

func (x SeccompMode) String() string {
	return proto.EnumName(SeccompMode_name, int32(x))
}

func (x *SeccompMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SeccompMode_value, data, "SeccompMode")
	if err != nil {
		return err
	}
	*x = SeccompMode(value)
	return nil
}

func (SeccompMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

type CoreEntryMarch int32

const (
	CoreEntry_UNKNOWN CoreEntryMarch = 0
	CoreEntry_X86_64  CoreEntryMarch = 1
	CoreEntry_ARM     CoreEntryMarch = 2
	CoreEntry_AARCH64 CoreEntryMarch = 3
	CoreEntry_PPC64   CoreEntryMarch = 4
	CoreEntry_S390    CoreEntryMarch = 5
	CoreEntry_MIPS    CoreEntryMarch = 6
)

var CoreEntryMarch_name = map[int32]string{
	0: "UNKNOWN",
	1: "X86_64",
	2: "ARM",
	3: "AARCH64",
	4: "PPC64",
	5: "S390",
	6: "MIPS",
}

var CoreEntryMarch_value = map[string]int32{
	"UNKNOWN": 0,
	"X86_64":  1,
	"ARM":     2,
	"AARCH64": 3,
	"PPC64":   4,
	"S390":    5,
	"MIPS":    6,
}

func (x CoreEntryMarch) Enum() *CoreEntryMarch {
	p := new(CoreEntryMarch)
	*p = x
	return p
}

func (x CoreEntryMarch) String() string {
	return proto.EnumName(CoreEntryMarch_name, int32(x))
}

func (x *CoreEntryMarch) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CoreEntryMarch_value, data, "CoreEntryMarch")
	if err != nil {
		return err
	}
	*x = CoreEntryMarch(value)
	return nil
}

func (CoreEntryMarch) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{5, 0}
}

type TaskCoreEntry struct {
	TaskState   *uint32           `protobuf:"varint,1,req,name=task_state,json=taskState" json:"task_state,omitempty"`
	ExitCode    *uint32           `protobuf:"varint,2,req,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	Personality *uint32           `protobuf:"varint,3,req,name=personality" json:"personality,omitempty"`
	Flags       *uint32           `protobuf:"varint,4,req,name=flags" json:"flags,omitempty"`
	BlkSigset   *uint64           `protobuf:"varint,5,req,name=blk_sigset,json=blkSigset" json:"blk_sigset,omitempty"`
	Comm        *string           `protobuf:"bytes,6,req,name=comm" json:"comm,omitempty"`
	Timers      *TaskTimersEntry  `protobuf:"bytes,7,opt,name=timers" json:"timers,omitempty"`
	Rlimits     *TaskRlimitsEntry `protobuf:"bytes,8,opt,name=rlimits" json:"rlimits,omitempty"`
	CgSet       *uint32           `protobuf:"varint,9,opt,name=cg_set,json=cgSet" json:"cg_set,omitempty"`
	SignalsS    *SignalQueueEntry `protobuf:"bytes,10,opt,name=signals_s,json=signalsS" json:"signals_s,omitempty"`
	// These two are deprecated, should be per-thread
	OldSeccompMode   *SeccompMode `protobuf:"varint,11,opt,name=old_seccomp_mode,json=oldSeccompMode,enum=SeccompMode" json:"old_seccomp_mode,omitempty"`
	OldSeccompFilter *uint32      `protobuf:"varint,12,opt,name=old_seccomp_filter,json=oldSeccompFilter" json:"old_seccomp_filter,omitempty"`
	Loginuid         *uint32      `protobuf:"varint,13,opt,name=loginuid" json:"loginuid,omitempty"`
	OomScoreAdj      *int32       `protobuf:"varint,14,opt,name=oom_score_adj,json=oomScoreAdj" json:"oom_score_adj,omitempty"`
	Sigactions       []*SaEntry   `protobuf:"bytes,15,rep,name=sigactions" json:"sigactions,omitempty"`
	ChildSubreaper   *bool        `protobuf:"varint,18,opt,name=child_subreaper,json=childSubreaper" json:"child_subreaper,omitempty"`
	// Reserved for container relative start time
	//optional uint64		start_time	= 19;
	BlkSigsetExtended    *uint64  `protobuf:"varint,20,opt,name=blk_sigset_extended,json=blkSigsetExtended" json:"blk_sigset_extended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCoreEntry) Reset()         { *m = TaskCoreEntry{} }
func (m *TaskCoreEntry) String() string { return proto.CompactTextString(m) }
func (*TaskCoreEntry) ProtoMessage()    {}
func (*TaskCoreEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

func (m *TaskCoreEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCoreEntry.Unmarshal(m, b)
}
func (m *TaskCoreEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCoreEntry.Marshal(b, m, deterministic)
}
func (m *TaskCoreEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCoreEntry.Merge(m, src)
}
func (m *TaskCoreEntry) XXX_Size() int {
	return xxx_messageInfo_TaskCoreEntry.Size(m)
}
func (m *TaskCoreEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCoreEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCoreEntry proto.InternalMessageInfo

func (m *TaskCoreEntry) GetTaskState() uint32 {
	if m != nil && m.TaskState != nil {
		return *m.TaskState
	}
	return 0
}

func (m *TaskCoreEntry) GetExitCode() uint32 {
	if m != nil && m.ExitCode != nil {
		return *m.ExitCode
	}
	return 0
}

func (m *TaskCoreEntry) GetPersonality() uint32 {
	if m != nil && m.Personality != nil {
		return *m.Personality
	}
	return 0
}

func (m *TaskCoreEntry) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *TaskCoreEntry) GetBlkSigset() uint64 {
	if m != nil && m.BlkSigset != nil {
		return *m.BlkSigset
	}
	return 0
}

func (m *TaskCoreEntry) GetComm() string {
	if m != nil && m.Comm != nil {
		return *m.Comm
	}
	return ""
}

func (m *TaskCoreEntry) GetTimers() *TaskTimersEntry {
	if m != nil {
		return m.Timers
	}
	return nil
}

func (m *TaskCoreEntry) GetRlimits() *TaskRlimitsEntry {
	if m != nil {
		return m.Rlimits
	}
	return nil
}

func (m *TaskCoreEntry) GetCgSet() uint32 {
	if m != nil && m.CgSet != nil {
		return *m.CgSet
	}
	return 0
}

func (m *TaskCoreEntry) GetSignalsS() *SignalQueueEntry {
	if m != nil {
		return m.SignalsS
	}
	return nil
}

func (m *TaskCoreEntry) GetOldSeccompMode() SeccompMode {
	if m != nil && m.OldSeccompMode != nil {
		return *m.OldSeccompMode
	}
	return SeccompMode_disabled
}

func (m *TaskCoreEntry) GetOldSeccompFilter() uint32 {
	if m != nil && m.OldSeccompFilter != nil {
		return *m.OldSeccompFilter
	}
	return 0
}

func (m *TaskCoreEntry) GetLoginuid() uint32 {
	if m != nil && m.Loginuid != nil {
		return *m.Loginuid
	}
	return 0
}

func (m *TaskCoreEntry) GetOomScoreAdj() int32 {
	if m != nil && m.OomScoreAdj != nil {
		return *m.OomScoreAdj
	}
	return 0
}

func (m *TaskCoreEntry) GetSigactions() []*SaEntry {
	if m != nil {
		return m.Sigactions
	}
	return nil
}

func (m *TaskCoreEntry) GetChildSubreaper() bool {
	if m != nil && m.ChildSubreaper != nil {
		return *m.ChildSubreaper
	}
	return false
}

func (m *TaskCoreEntry) GetBlkSigsetExtended() uint64 {
	if m != nil && m.BlkSigsetExtended != nil {
		return *m.BlkSigsetExtended
	}
	return 0
}

type TaskKobjIdsEntry struct {
	VmId                 *uint32  `protobuf:"varint,1,req,name=vm_id,json=vmId" json:"vm_id,omitempty"`
	FilesId              *uint32  `protobuf:"varint,2,req,name=files_id,json=filesId" json:"files_id,omitempty"`
	FsId                 *uint32  `protobuf:"varint,3,req,name=fs_id,json=fsId" json:"fs_id,omitempty"`
	SighandId            *uint32  `protobuf:"varint,4,req,name=sighand_id,json=sighandId" json:"sighand_id,omitempty"`
	PidNsId              *uint32  `protobuf:"varint,5,opt,name=pid_ns_id,json=pidNsId" json:"pid_ns_id,omitempty"`
	NetNsId              *uint32  `protobuf:"varint,6,opt,name=net_ns_id,json=netNsId" json:"net_ns_id,omitempty"`
	IpcNsId              *uint32  `protobuf:"varint,7,opt,name=ipc_ns_id,json=ipcNsId" json:"ipc_ns_id,omitempty"`
	UtsNsId              *uint32  `protobuf:"varint,8,opt,name=uts_ns_id,json=utsNsId" json:"uts_ns_id,omitempty"`
	MntNsId              *uint32  `protobuf:"varint,9,opt,name=mnt_ns_id,json=mntNsId" json:"mnt_ns_id,omitempty"`
	UserNsId             *uint32  `protobuf:"varint,10,opt,name=user_ns_id,json=userNsId" json:"user_ns_id,omitempty"`
	CgroupNsId           *uint32  `protobuf:"varint,11,opt,name=cgroup_ns_id,json=cgroupNsId" json:"cgroup_ns_id,omitempty"`
	TimeNsId             *uint32  `protobuf:"varint,12,opt,name=time_ns_id,json=timeNsId" json:"time_ns_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskKobjIdsEntry) Reset()         { *m = TaskKobjIdsEntry{} }
func (m *TaskKobjIdsEntry) String() string { return proto.CompactTextString(m) }
func (*TaskKobjIdsEntry) ProtoMessage()    {}
func (*TaskKobjIdsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1}
}

func (m *TaskKobjIdsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskKobjIdsEntry.Unmarshal(m, b)
}
func (m *TaskKobjIdsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskKobjIdsEntry.Marshal(b, m, deterministic)
}
func (m *TaskKobjIdsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskKobjIdsEntry.Merge(m, src)
}
func (m *TaskKobjIdsEntry) XXX_Size() int {
	return xxx_messageInfo_TaskKobjIdsEntry.Size(m)
}
func (m *TaskKobjIdsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskKobjIdsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TaskKobjIdsEntry proto.InternalMessageInfo

func (m *TaskKobjIdsEntry) GetVmId() uint32 {
	if m != nil && m.VmId != nil {
		return *m.VmId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetFilesId() uint32 {
	if m != nil && m.FilesId != nil {
		return *m.FilesId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetFsId() uint32 {
	if m != nil && m.FsId != nil {
		return *m.FsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetSighandId() uint32 {
	if m != nil && m.SighandId != nil {
		return *m.SighandId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetPidNsId() uint32 {
	if m != nil && m.PidNsId != nil {
		return *m.PidNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetNetNsId() uint32 {
	if m != nil && m.NetNsId != nil {
		return *m.NetNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetIpcNsId() uint32 {
	if m != nil && m.IpcNsId != nil {
		return *m.IpcNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetUtsNsId() uint32 {
	if m != nil && m.UtsNsId != nil {
		return *m.UtsNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetMntNsId() uint32 {
	if m != nil && m.MntNsId != nil {
		return *m.MntNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetUserNsId() uint32 {
	if m != nil && m.UserNsId != nil {
		return *m.UserNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetCgroupNsId() uint32 {
	if m != nil && m.CgroupNsId != nil {
		return *m.CgroupNsId
	}
	return 0
}

func (m *TaskKobjIdsEntry) GetTimeNsId() uint32 {
	if m != nil && m.TimeNsId != nil {
		return *m.TimeNsId
	}
	return 0
}

type ThreadSasEntry struct {
	SsSp                 *uint64  `protobuf:"varint,1,req,name=ss_sp,json=ssSp" json:"ss_sp,omitempty"`
	SsSize               *uint64  `protobuf:"varint,2,req,name=ss_size,json=ssSize" json:"ss_size,omitempty"`
	SsFlags              *uint32  `protobuf:"varint,3,req,name=ss_flags,json=ssFlags" json:"ss_flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadSasEntry) Reset()         { *m = ThreadSasEntry{} }
func (m *ThreadSasEntry) String() string { return proto.CompactTextString(m) }
func (*ThreadSasEntry) ProtoMessage()    {}
func (*ThreadSasEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{2}
}

func (m *ThreadSasEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadSasEntry.Unmarshal(m, b)
}
func (m *ThreadSasEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadSasEntry.Marshal(b, m, deterministic)
}
func (m *ThreadSasEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadSasEntry.Merge(m, src)
}
func (m *ThreadSasEntry) XXX_Size() int {
	return xxx_messageInfo_ThreadSasEntry.Size(m)
}
func (m *ThreadSasEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadSasEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadSasEntry proto.InternalMessageInfo

func (m *ThreadSasEntry) GetSsSp() uint64 {
	if m != nil && m.SsSp != nil {
		return *m.SsSp
	}
	return 0
}

func (m *ThreadSasEntry) GetSsSize() uint64 {
	if m != nil && m.SsSize != nil {
		return *m.SsSize
	}
	return 0
}

func (m *ThreadSasEntry) GetSsFlags() uint32 {
	if m != nil && m.SsFlags != nil {
		return *m.SsFlags
	}
	return 0
}

type ThreadCoreEntry struct {
	FutexRla             *uint64           `protobuf:"varint,1,req,name=futex_rla,json=futexRla" json:"futex_rla,omitempty"`
	FutexRlaLen          *uint32           `protobuf:"varint,2,req,name=futex_rla_len,json=futexRlaLen" json:"futex_rla_len,omitempty"`
	SchedNice            *int32            `protobuf:"zigzag32,3,opt,name=sched_nice,json=schedNice" json:"sched_nice,omitempty"`
	SchedPolicy          *uint32           `protobuf:"varint,4,opt,name=sched_policy,json=schedPolicy" json:"sched_policy,omitempty"`
	SchedPrio            *uint32           `protobuf:"varint,5,opt,name=sched_prio,json=schedPrio" json:"sched_prio,omitempty"`
	BlkSigset            *uint64           `protobuf:"varint,6,opt,name=blk_sigset,json=blkSigset" json:"blk_sigset,omitempty"`
	Sas                  *ThreadSasEntry   `protobuf:"bytes,7,opt,name=sas" json:"sas,omitempty"`
	PdeathSig            *uint32           `protobuf:"varint,8,opt,name=pdeath_sig,json=pdeathSig" json:"pdeath_sig,omitempty"`
	SignalsP             *SignalQueueEntry `protobuf:"bytes,9,opt,name=signals_p,json=signalsP" json:"signals_p,omitempty"`
	Creds                *CredsEntry       `protobuf:"bytes,10,opt,name=creds" json:"creds,omitempty"`
	SeccompMode          *SeccompMode      `protobuf:"varint,11,opt,name=seccomp_mode,json=seccompMode,enum=SeccompMode" json:"seccomp_mode,omitempty"`
	SeccompFilter        *uint32           `protobuf:"varint,12,opt,name=seccomp_filter,json=seccompFilter" json:"seccomp_filter,omitempty"`
	Comm                 *string           `protobuf:"bytes,13,opt,name=comm" json:"comm,omitempty"`
	BlkSigsetExtended    *uint64           `protobuf:"varint,14,opt,name=blk_sigset_extended,json=blkSigsetExtended" json:"blk_sigset_extended,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ThreadCoreEntry) Reset()         { *m = ThreadCoreEntry{} }
func (m *ThreadCoreEntry) String() string { return proto.CompactTextString(m) }
func (*ThreadCoreEntry) ProtoMessage()    {}
func (*ThreadCoreEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{3}
}

func (m *ThreadCoreEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadCoreEntry.Unmarshal(m, b)
}
func (m *ThreadCoreEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadCoreEntry.Marshal(b, m, deterministic)
}
func (m *ThreadCoreEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadCoreEntry.Merge(m, src)
}
func (m *ThreadCoreEntry) XXX_Size() int {
	return xxx_messageInfo_ThreadCoreEntry.Size(m)
}
func (m *ThreadCoreEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadCoreEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadCoreEntry proto.InternalMessageInfo

func (m *ThreadCoreEntry) GetFutexRla() uint64 {
	if m != nil && m.FutexRla != nil {
		return *m.FutexRla
	}
	return 0
}

func (m *ThreadCoreEntry) GetFutexRlaLen() uint32 {
	if m != nil && m.FutexRlaLen != nil {
		return *m.FutexRlaLen
	}
	return 0
}

func (m *ThreadCoreEntry) GetSchedNice() int32 {
	if m != nil && m.SchedNice != nil {
		return *m.SchedNice
	}
	return 0
}

func (m *ThreadCoreEntry) GetSchedPolicy() uint32 {
	if m != nil && m.SchedPolicy != nil {
		return *m.SchedPolicy
	}
	return 0
}

func (m *ThreadCoreEntry) GetSchedPrio() uint32 {
	if m != nil && m.SchedPrio != nil {
		return *m.SchedPrio
	}
	return 0
}

func (m *ThreadCoreEntry) GetBlkSigset() uint64 {
	if m != nil && m.BlkSigset != nil {
		return *m.BlkSigset
	}
	return 0
}

func (m *ThreadCoreEntry) GetSas() *ThreadSasEntry {
	if m != nil {
		return m.Sas
	}
	return nil
}

func (m *ThreadCoreEntry) GetPdeathSig() uint32 {
	if m != nil && m.PdeathSig != nil {
		return *m.PdeathSig
	}
	return 0
}

func (m *ThreadCoreEntry) GetSignalsP() *SignalQueueEntry {
	if m != nil {
		return m.SignalsP
	}
	return nil
}

func (m *ThreadCoreEntry) GetCreds() *CredsEntry {
	if m != nil {
		return m.Creds
	}
	return nil
}

func (m *ThreadCoreEntry) GetSeccompMode() SeccompMode {
	if m != nil && m.SeccompMode != nil {
		return *m.SeccompMode
	}
	return SeccompMode_disabled
}

func (m *ThreadCoreEntry) GetSeccompFilter() uint32 {
	if m != nil && m.SeccompFilter != nil {
		return *m.SeccompFilter
	}
	return 0
}

func (m *ThreadCoreEntry) GetComm() string {
	if m != nil && m.Comm != nil {
		return *m.Comm
	}
	return ""
}

func (m *ThreadCoreEntry) GetBlkSigsetExtended() uint64 {
	if m != nil && m.BlkSigsetExtended != nil {
		return *m.BlkSigsetExtended
	}
	return 0
}

type TaskRlimitsEntry struct {
	Rlimits              []*RlimitEntry `protobuf:"bytes,1,rep,name=rlimits" json:"rlimits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskRlimitsEntry) Reset()         { *m = TaskRlimitsEntry{} }
func (m *TaskRlimitsEntry) String() string { return proto.CompactTextString(m) }
func (*TaskRlimitsEntry) ProtoMessage()    {}
func (*TaskRlimitsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{4}
}

func (m *TaskRlimitsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRlimitsEntry.Unmarshal(m, b)
}
func (m *TaskRlimitsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRlimitsEntry.Marshal(b, m, deterministic)
}
func (m *TaskRlimitsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRlimitsEntry.Merge(m, src)
}
func (m *TaskRlimitsEntry) XXX_Size() int {
	return xxx_messageInfo_TaskRlimitsEntry.Size(m)
}
func (m *TaskRlimitsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRlimitsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRlimitsEntry proto.InternalMessageInfo

func (m *TaskRlimitsEntry) GetRlimits() []*RlimitEntry {
	if m != nil {
		return m.Rlimits
	}
	return nil
}

type CoreEntry struct {
	Mtype                *CoreEntryMarch    `protobuf:"varint,1,req,name=mtype,enum=CoreEntryMarch" json:"mtype,omitempty"`
	ThreadInfo           *ThreadInfoX86     `protobuf:"bytes,2,opt,name=thread_info,json=threadInfo" json:"thread_info,omitempty"`
	TiArm                *ThreadInfoArm     `protobuf:"bytes,6,opt,name=ti_arm,json=tiArm" json:"ti_arm,omitempty"`
	TiAarch64            *ThreadInfoAarch64 `protobuf:"bytes,8,opt,name=ti_aarch64,json=tiAarch64" json:"ti_aarch64,omitempty"`
	TiPpc64              *ThreadInfoPpc64   `protobuf:"bytes,9,opt,name=ti_ppc64,json=tiPpc64" json:"ti_ppc64,omitempty"`
	TiS390               *ThreadInfoS390    `protobuf:"bytes,10,opt,name=ti_s390,json=tiS390" json:"ti_s390,omitempty"`
	TiMips               *ThreadInfoMips    `protobuf:"bytes,11,opt,name=ti_mips,json=tiMips" json:"ti_mips,omitempty"`
	Tc                   *TaskCoreEntry     `protobuf:"bytes,3,opt,name=tc" json:"tc,omitempty"`
	Ids                  *TaskKobjIdsEntry  `protobuf:"bytes,4,opt,name=ids" json:"ids,omitempty"`
	ThreadCore           *ThreadCoreEntry   `protobuf:"bytes,5,opt,name=thread_core,json=threadCore" json:"thread_core,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CoreEntry) Reset()         { *m = CoreEntry{} }
func (m *CoreEntry) String() string { return proto.CompactTextString(m) }
func (*CoreEntry) ProtoMessage()    {}
func (*CoreEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{5}
}

func (m *CoreEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreEntry.Unmarshal(m, b)
}
func (m *CoreEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreEntry.Marshal(b, m, deterministic)
}
func (m *CoreEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreEntry.Merge(m, src)
}
func (m *CoreEntry) XXX_Size() int {
	return xxx_messageInfo_CoreEntry.Size(m)
}
func (m *CoreEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CoreEntry proto.InternalMessageInfo

func (m *CoreEntry) GetMtype() CoreEntryMarch {
	if m != nil && m.Mtype != nil {
		return *m.Mtype
	}
	return CoreEntry_UNKNOWN
}

func (m *CoreEntry) GetThreadInfo() *ThreadInfoX86 {
	if m != nil {
		return m.ThreadInfo
	}
	return nil
}

func (m *CoreEntry) GetTiArm() *ThreadInfoArm {
	if m != nil {
		return m.TiArm
	}
	return nil
}

func (m *CoreEntry) GetTiAarch64() *ThreadInfoAarch64 {
	if m != nil {
		return m.TiAarch64
	}
	return nil
}

func (m *CoreEntry) GetTiPpc64() *ThreadInfoPpc64 {
	if m != nil {
		return m.TiPpc64
	}
	return nil
}

func (m *CoreEntry) GetTiS390() *ThreadInfoS390 {
	if m != nil {
		return m.TiS390
	}
	return nil
}

func (m *CoreEntry) GetTiMips() *ThreadInfoMips {
	if m != nil {
		return m.TiMips
	}
	return nil
}

func (m *CoreEntry) GetTc() *TaskCoreEntry {
	if m != nil {
		return m.Tc
	}
	return nil
}

func (m *CoreEntry) GetIds() *TaskKobjIdsEntry {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *CoreEntry) GetThreadCore() *ThreadCoreEntry {
	if m != nil {
		return m.ThreadCore
	}
	return nil
}

func init() {
	proto.RegisterEnum("SeccompMode", SeccompMode_name, SeccompMode_value)
	proto.RegisterEnum("CoreEntryMarch", CoreEntryMarch_name, CoreEntryMarch_value)
	proto.RegisterType((*TaskCoreEntry)(nil), "task_core_entry")
	proto.RegisterType((*TaskKobjIdsEntry)(nil), "task_kobj_ids_entry")
	proto.RegisterType((*ThreadSasEntry)(nil), "thread_sas_entry")
	proto.RegisterType((*ThreadCoreEntry)(nil), "thread_core_entry")
	proto.RegisterType((*TaskRlimitsEntry)(nil), "task_rlimits_entry")
	proto.RegisterType((*CoreEntry)(nil), "core_entry")
}

func init() {
	proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe)
}

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 1208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0x89, 0x9d, 0x38, 0xc7, 0x49, 0xea, 0x4e, 0x8b, 0x30, 0x05, 0xa4, 0x12, 0xfe, 0xca,
	0x8a, 0x0d, 0x4b, 0xb7, 0x94, 0xe5, 0x02, 0xad, 0xd2, 0x8a, 0x15, 0x15, 0xb4, 0x44, 0x8e, 0x56,
	0x20, 0xb8, 0xb0, 0x5c, 0x7b, 0x92, 0x4c, 0x37, 0xfe, 0xc1, 0x33, 0x59, 0xb5, 0xbc, 0x02, 0x0f,
	0xc1, 0x05, 0x6f, 0xb4, 0xaf, 0xc0, 0x8b, 0xec, 0x99, 0x1f, 0xa7, 0x4e, 0xb7, 0xd2, 0x5e, 0x65,
	0xce, 0xf7, 0x7d, 0x73, 0x32, 0x73, 0xbe, 0x99, 0x33, 0x06, 0x88, 0xf3, 0x92, 0x8e, 0x8a, 0x32,
	0x17, 0xf9, 0xde, 0x40, 0x8e, 0x1f, 0x5e, 0x3f, 0x39, 0xde, 0x88, 0xa3, 0x32, 0x35, 0x31, 0xd1,
	0x71, 0x54, 0xc6, 0x8b, 0xe3, 0x23, 0x83, 0x79, 0x0a, 0x2b, 0x8a, 0x78, 0x8d, 0x6c, 0x29, 0x84,
	0x3f, 0xfe, 0xfe, 0xd1, 0x06, 0x90, 0xb2, 0x82, 0x1b, 0xa0, 0x57, 0x2e, 0x59, 0xca, 0x84, 0x89,
	0x5c, 0xc1, 0x52, 0x5a, 0x56, 0x41, 0x5c, 0xd2, 0xa4, 0xd2, 0x39, 0x3c, 0x32, 0xa3, 0x3e, 0x67,
	0x73, 0x96, 0xcd, 0x72, 0x13, 0x42, 0x5e, 0x08, 0x23, 0x1a, 0xfe, 0x6b, 0xc3, 0x96, 0x88, 0xf8,
	0x8b, 0x50, 0xfe, 0x4b, 0x48, 0x33, 0x51, 0xde, 0x90, 0x2f, 0x00, 0x14, 0xc4, 0x45, 0x24, 0xa8,
	0xdf, 0xd8, 0x6f, 0x1e, 0xf4, 0x4f, 0x9c, 0x57, 0x4f, 0xed, 0xc3, 0xd6, 0x9c, 0x66, 0x41, 0x57,
	0x72, 0x53, 0x49, 0x91, 0x0f, 0xa0, 0x4b, 0xaf, 0x99, 0xc0, 0xb9, 0x09, 0xf5, 0x9b, 0x52, 0x17,
	0x38, 0x12, 0x38, 0xc5, 0x98, 0xec, 0x83, 0x5b, 0xd0, 0x92, 0xe7, 0x59, 0xb4, 0x64, 0xe2, 0xc6,
	0x6f, 0x29, 0xba, 0x0e, 0x91, 0x5d, 0xb0, 0x67, 0xcb, 0x68, 0xce, 0x7d, 0x4b, 0x71, 0x3a, 0x20,
	0x9f, 0x02, 0x5c, 0x2e, 0xf1, 0xcf, 0xd9, 0x9c, 0x53, 0xe1, 0xdb, 0x48, 0x59, 0x27, 0xf6, 0xab,
	0xa7, 0x4d, 0xa7, 0x11, 0x74, 0x91, 0x98, 0x2a, 0x9c, 0x10, 0xb0, 0xe2, 0x3c, 0x4d, 0xfd, 0x36,
	0xf2, 0xdd, 0x40, 0x8d, 0xc9, 0x03, 0x68, 0xab, 0x62, 0x70, 0xbf, 0xb3, 0xdf, 0x38, 0x70, 0x0f,
	0xc9, 0x48, 0x6d, 0x43, 0x63, 0x7a, 0x6f, 0x81, 0x51, 0x90, 0x87, 0xd0, 0xd1, 0x65, 0xe4, 0xbe,
	0xa3, 0xc4, 0x3b, 0x5a, 0x6c, 0x40, 0xa3, 0xae, 0x34, 0xe4, 0x5d, 0x68, 0xc7, 0xf3, 0x50, 0x2e,
	0xa8, 0x8b, 0x6a, 0x5c, 0x6b, 0x3c, 0x9f, 0xe2, 0x2a, 0x1e, 0x41, 0x17, 0xd7, 0x89, 0xdb, 0xe1,
	0x21, 0xf7, 0xc1, 0xe4, 0xd1, 0x48, 0xf8, 0xd7, 0x8a, 0xae, 0x4c, 0x45, 0x03, 0xc7, 0xa8, 0xa6,
	0xe4, 0x3b, 0xf0, 0xf2, 0x65, 0x82, 0x99, 0x62, 0x5c, 0x72, 0x11, 0xa6, 0xb2, 0x72, 0x2e, 0x4e,
	0x1c, 0x1c, 0xf6, 0x47, 0x75, 0x30, 0x18, 0xa0, 0x6c, 0xaa, 0x81, 0x73, 0x59, 0xce, 0xaf, 0x80,
	0xd4, 0x27, 0xce, 0xd8, 0x52, 0xd0, 0xd2, 0xef, 0xa9, 0xd5, 0x78, 0xb7, 0xda, 0x67, 0x0a, 0x27,
	0x7b, 0xe0, 0x2c, 0x73, 0xf4, 0x7c, 0xc5, 0x12, 0xbf, 0xaf, 0x34, 0xeb, 0x98, 0x0c, 0xa1, 0x9f,
	0xe7, 0x69, 0xc8, 0x95, 0xe3, 0x51, 0x72, 0xe5, 0x0f, 0x50, 0x60, 0x07, 0x2e, 0x82, 0x53, 0x89,
	0x8d, 0x93, 0x2b, 0xf2, 0x25, 0x00, 0x2e, 0x39, 0x8a, 0x05, 0xcb, 0x33, 0xee, 0x6f, 0xed, 0xb7,
	0x70, 0x67, 0xdd, 0x11, 0x8f, 0xcc, 0x7e, 0x6a, 0x24, 0x9e, 0x96, 0xad, 0x78, 0xc1, 0xe4, 0xd2,
	0x56, 0x97, 0x25, 0x8d, 0xd0, 0x60, 0x9f, 0x60, 0x42, 0x27, 0x18, 0x28, 0x78, 0x5a, 0xa1, 0xe4,
	0x5b, 0xd8, 0xb9, 0x35, 0x36, 0xa4, 0xd7, 0x82, 0x66, 0x09, 0x4d, 0xfc, 0x5d, 0x14, 0xaf, 0x1d,
	0xde, 0x5e, 0x3b, 0xfc, 0xa3, 0xe1, 0x87, 0xff, 0x37, 0x61, 0x47, 0x59, 0xf3, 0x22, 0xbf, 0xbc,
	0x0a, 0x59, 0x62, 0xbc, 0x21, 0x3b, 0x60, 0xbf, 0x4c, 0x31, 0xd6, 0x07, 0x34, 0xb0, 0x5e, 0xa6,
	0x67, 0x09, 0x79, 0x1f, 0x1c, 0xac, 0x0c, 0xe5, 0x12, 0xd7, 0x07, 0xb2, 0xa3, 0x62, 0xa4, 0x50,
	0x3f, 0x53, 0xb8, 0x3e, 0x89, 0xd6, 0x4c, 0x82, 0x1f, 0xa9, 0x7d, 0x2e, 0xa2, 0x2c, 0x91, 0x8c,
	0x3e, 0x87, 0x5d, 0x83, 0x20, 0xbd, 0x07, 0xdd, 0x82, 0x25, 0x61, 0xa6, 0xe6, 0xd9, 0xaa, 0x8e,
	0x1d, 0x04, 0x2e, 0xb8, 0xe6, 0x32, 0xdc, 0x87, 0xe6, 0xda, 0x9a, 0x43, 0xa0, 0xe2, 0x58, 0x11,
	0x1b, 0xae, 0xa3, 0x39, 0x04, 0x2a, 0x6e, 0x85, 0x07, 0x4c, 0x73, 0x8e, 0xe6, 0x10, 0xa8, 0xb8,
	0x34, 0xab, 0x72, 0xea, 0x93, 0xd6, 0x41, 0x40, 0x71, 0x1f, 0x02, 0xac, 0x38, 0x2d, 0x0d, 0x09,
	0xda, 0x54, 0x89, 0x28, 0x76, 0x1f, 0x7a, 0xf1, 0xbc, 0xcc, 0x57, 0x85, 0xe1, 0x5d, 0xc5, 0x83,
	0xc6, 0xaa, 0xf9, 0xf2, 0xec, 0x1b, 0x5e, 0x1f, 0x1c, 0x47, 0x22, 0x92, 0x1d, 0xfe, 0x09, 0x9e,
	0x58, 0xa0, 0x51, 0x68, 0x63, 0x54, 0xab, 0x30, 0xc7, 0x83, 0x5d, 0xa8, 0x0a, 0x5b, 0x81, 0xc5,
	0xf9, 0xb4, 0x20, 0xef, 0x41, 0x47, 0x82, 0xec, 0x6f, 0x7d, 0xe3, 0xad, 0xa0, 0x8d, 0x30, 0x46,
	0xb2, 0xf4, 0x48, 0xe8, 0x0b, 0xad, 0x4b, 0x8c, 0xc2, 0x67, 0x32, 0x1c, 0xfe, 0x63, 0xc1, 0xb6,
	0xc9, 0x5e, 0x6b, 0x33, 0xd8, 0x3d, 0x66, 0x2b, 0x41, 0xaf, 0xf1, 0xce, 0x45, 0xe6, 0x2f, 0x1c,
	0x05, 0x04, 0xcb, 0x48, 0x1e, 0xd2, 0x35, 0x19, 0x2e, 0x69, 0x66, 0xdc, 0x74, 0x2b, 0xc1, 0x2f,
	0x34, 0x53, 0xe6, 0xc5, 0x0b, 0x8a, 0xfe, 0xb0, 0x98, 0xe2, 0x7f, 0x36, 0x0e, 0xb6, 0xd1, 0x3c,
	0x89, 0x5c, 0x20, 0x40, 0x3e, 0x86, 0x9e, 0xa6, 0x8b, 0x7c, 0xc9, 0xe2, 0x1b, 0x74, 0x57, 0x6e,
	0xd9, 0x55, 0xd8, 0x44, 0x41, 0xb7, 0x19, 0x8a, 0x92, 0xe5, 0xc6, 0x60, 0x9d, 0x61, 0x82, 0x80,
	0xa4, 0x6b, 0xad, 0x48, 0x7a, 0x6c, 0xd5, 0x7b, 0xd0, 0x27, 0xd0, 0xc2, 0x62, 0x99, 0x66, 0xb3,
	0x3d, 0xba, 0x5b, 0xbf, 0x40, 0xb2, 0x32, 0x47, 0x91, 0xd0, 0x48, 0x2c, 0x64, 0x1a, 0xe3, 0x77,
	0x57, 0x23, 0x98, 0xa6, 0xde, 0x41, 0x0a, 0xe5, 0xf8, 0x5b, 0x3a, 0xc8, 0x04, 0x2b, 0x63, 0xab,
	0x2e, 0x6f, 0xfa, 0x4d, 0x6f, 0xa4, 0x22, 0x23, 0xd3, 0x14, 0x66, 0xed, 0xbd, 0xbd, 0xc3, 0xb8,
	0xbc, 0xd6, 0x5e, 0x3e, 0x83, 0xc1, 0xbd, 0xad, 0xa5, 0xcf, 0x37, 0xfa, 0x4a, 0xd5, 0x76, 0x65,
	0x4f, 0xa9, 0xda, 0xee, 0xe8, 0xfe, 0x7b, 0x3d, 0x50, 0xe5, 0xba, 0xe7, 0x42, 0xff, 0x00, 0xe4,
	0xcd, 0x56, 0x8b, 0x6d, 0x64, 0xdd, 0x90, 0x1b, 0xaa, 0xdd, 0xf4, 0x47, 0x3a, 0xbe, 0xd3, 0x8a,
	0x87, 0xff, 0x59, 0xfa, 0xd5, 0x5d, 0xcf, 0xb3, 0x53, 0x71, 0x53, 0xe8, 0x77, 0x6a, 0x80, 0x36,
	0xdc, 0x72, 0xa3, 0x54, 0xbe, 0xb5, 0x81, 0xe6, 0xc9, 0x37, 0xe0, 0x1a, 0x87, 0xe4, 0x53, 0x88,
	0xe7, 0x49, 0x56, 0xcf, 0x1b, 0xd5, 0xb0, 0x10, 0xdf, 0xee, 0x00, 0x34, 0x70, 0x86, 0x31, 0xe6,
	0xc6, 0xe7, 0x22, 0xc4, 0x17, 0x5c, 0x79, 0x7f, 0x57, 0x8d, 0x78, 0x60, 0x0b, 0x36, 0x2e, 0x53,
	0xf2, 0x58, 0xde, 0xad, 0xd0, 0x3c, 0xed, 0xe6, 0x41, 0xd9, 0xdd, 0x14, 0x6b, 0x0e, 0x5f, 0x4f,
	0x36, 0xd6, 0x43, 0x7c, 0x82, 0xf0, 0xfa, 0x85, 0xea, 0xed, 0x37, 0xce, 0x93, 0x8d, 0x29, 0x8a,
	0x09, 0x3a, 0x82, 0x4d, 0xe4, 0x00, 0x5f, 0x37, 0x1c, 0x86, 0xf2, 0xc3, 0xc0, 0x38, 0xbf, 0xbd,
	0xa1, 0x96, 0x84, 0x7c, 0xdd, 0xa6, 0xf8, 0x6b, 0xb4, 0xf2, 0x9b, 0x41, 0x59, 0x7f, 0x57, 0x2b,
	0x09, 0xa9, 0x3d, 0xc7, 0x5f, 0xec, 0x1c, 0x4d, 0x11, 0xab, 0xdb, 0xa3, 0x36, 0xb8, 0xf9, 0x2d,
	0x10, 0x20, 0x47, 0x3e, 0x87, 0x16, 0xb6, 0x5d, 0x75, 0x7f, 0xd4, 0xb6, 0xde, 0x6c, 0xc6, 0x81,
	0x14, 0x60, 0x15, 0xdc, 0xda, 0x2d, 0x57, 0xd7, 0xa9, 0xb6, 0xa7, 0x5a, 0x52, 0x53, 0xe3, 0x53,
	0x44, 0x86, 0xcf, 0xd1, 0x3f, 0x59, 0x10, 0xe2, 0x42, 0xe7, 0xf9, 0xc5, 0xcf, 0x17, 0xbf, 0xfe,
	0x76, 0xe1, 0xbd, 0x43, 0x00, 0xda, 0xbf, 0x3f, 0x39, 0x0e, 0x8f, 0x8f, 0xbc, 0x06, 0xe9, 0x40,
	0x6b, 0x1c, 0x9c, 0x7b, 0x4d, 0xa9, 0x18, 0x8f, 0x83, 0xd3, 0x9f, 0x10, 0x6d, 0x91, 0x2e, 0xd8,
	0x93, 0xc9, 0x29, 0x0e, 0x2d, 0xe2, 0x80, 0x25, 0x77, 0xed, 0xd9, 0x72, 0x74, 0x7e, 0x36, 0x99,
	0x7a, 0xed, 0x07, 0x47, 0x9b, 0x37, 0x80, 0xf4, 0xc0, 0x49, 0x18, 0x8f, 0x2e, 0x97, 0x34, 0xd1,
	0xe9, 0xb9, 0x28, 0x59, 0x2c, 0x30, 0x3d, 0x8e, 0xf5, 0x89, 0xf7, 0x9a, 0x27, 0xde, 0x1f, 0x83,
	0xaf, 0xd5, 0x77, 0xd1, 0x25, 0xcb, 0x12, 0x96, 0xcd, 0xf9, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x26, 0xa7, 0xea, 0x9b, 0xe1, 0x09, 0x00, 0x00,
}