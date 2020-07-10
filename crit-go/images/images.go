package images

import (
	//"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	//"strings"
	"crit-go/protobindings"
	//"reflect"
	"strconv"
	//"github.com/thedevsaddam/gojsonq/v2"
	//"google.golang.org/protobuf/proto"
	//"google.golang.org/protobuf/encoding/protojson"
	//"google.golang.org/protobuf/encoding/prototext"
	//"github.com/golang/protobuf/jsonpb"
	//"google.golang.org/protobuf/encoding/protojson"
	//"strings"
	//"bytes"
	"reflect"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

var pbuffhandler = map[string]interface{}{
	"INVENTORY":      &protobindings.InventoryEntry{},
	"CORE":           &protobindings.CoreEntry{},
	"IDS":            &protobindings.TaskKobjIdsEntry{},
	"CREDS":          &protobindings.CredsEntry{},
	"UTSNS":          &protobindings.UtsnsEntry{},
	"TIMENS":         &protobindings.TimensEntry{},
	"IPC_VAR":        &protobindings.IpcVarEntry{},
	"FS":             &protobindings.FsEntry{},
	"MM":             &protobindings.MmEntry{},
	"CGROUP":         &protobindings.CgroupEntry{},
	"TCP_STREAM":     &protobindings.TcpStreamEntry{},
	"STATS":          &protobindings.StatsEntry{},
	"PSTREE":         &protobindings.PstreeEntry{},
	"REG_FILES":      &protobindings.RegFileEntry{},
	"NS_FILES":       &protobindings.NsFileEntry{},
	"EVENTFD_FILE":   &protobindings.EventfdFileEntry{},
	"EVENTPOLL_FILE": &protobindings.EventpollFileEntry{},
	"EVENTPOLL_TFD":  &protobindings.EventpollTfdEntry{},
	"SIGNALFD":       &protobindings.SignalfdEntry{},
	"TIMERFD":        &protobindings.TimerfdEntry{},
	"INOTIFY_FILE":   &protobindings.InotifyFileEntry{},
	"INOTIFY_WD":     &protobindings.InotifyWdEntry{},
	"FANOTIFY_FILE":  &protobindings.FanotifyFileEntry{},
	"FANOTIFY_MARK":  &protobindings.FanotifyMarkEntry{},
	"VMAS":           &protobindings.VmaEntry{},
	"PIPES":          &protobindings.PipeEntry{},
	"FIFO":           &protobindings.FifoEntry{},
	"SIGACT":         &protobindings.SaEntry{},
	"NETLINK_SK":     &protobindings.NetlinkSkEntry{},
	"REMAP_FPATH":    &protobindings.RemapFilePathEntry{},
	"MNTS":           &protobindings.MntEntry{},
	"TTY_FILES":      &protobindings.TtyFileEntry{},
	"TTY_INFO":       &protobindings.TtyInfoEntry{},
	"TTY_DATA":       &protobindings.TtyDataEntry{},
	"RLIMIT":         &protobindings.RlimitEntry{},
	"TUNFILE":        &protobindings.TunfileEntry{},
	"EXT_FILES":      &protobindings.ExtFileEntry{},
	"IRMAP_CACHE":    &protobindings.IrmapCacheEntry{},
	"FILE_LOCKS":     &protobindings.FileLockEntry{},
	"FDINFO":         &protobindings.FdinfoEntry{},
	"UNIXSK":         &protobindings.UnixSkEntry{},
	"INETSK":         &protobindings.InetSkEntry{},
	"PACKETSK":       &protobindings.PacketSockEntry{},
	"ITIMERS":        &protobindings.ItimerEntry{},
	"POSIX_TIMERS":   &protobindings.PosixTimerEntry{},
	"NETDEV":         &protobindings.NetDeviceEntry{},
	"PIPES_DATA":     &protobindings.PipeDataEntry{},
	"FIFO_DATA":      &protobindings.PipeDataEntry{},
	"SK_QUEUES":      &protobindings.SkPacketEntry{},
	"IPCNS_SHM":      &protobindings.IpcShmEntry{},
	"IPCNS_SEM":      &protobindings.IpcSemEntry{},
	"IPCNS_MSG":      &protobindings.IpcMsgEntry{},
	"NETNS":          &protobindings.NetnsEntry{},
	"USERNS":         &protobindings.UsernsEntry{},
	"SECCOMP":        &protobindings.SeccompEntry{},
	"AUTOFS":         &protobindings.AutofsEntry{},
	"FILES":          &protobindings.FileEntry{},
	"CPUINFO":        &protobindings.CpuinfoEntry{},
	"MEMFD_FILE":     &protobindings.MemfdFileEntry{},
	"MEMFD_INODE":    &protobindings.MemfdInodeEntry{},
}

var jsonmap map[string]interface{}
var bynamemap map[string]interface{}
var byvalmap map[string]interface{}

func Load(f *os.File, pretty bool, nopl bool) {
	image := make(map[string]interface{})
	/*
	   Convert criu image from binary format to dict(json).
	   Takes a file-like object to read criu image from.
	   Returns criu image in dict(json) format.
	*/
	magicjson()
	m := magichandler(f)
	image["magic"] = m
	switch {
	case m == "GHOST_FILE":
		handler := &ghost_file_handler{}
		handler.m = m
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "PAGEMAP":
		handler := &pagemap_handler{}
		handler.m = m
		image["entries"] = handler.Load(f, pretty, nopl)
	default:
		handler := &entry_handler{}
		handler.m = m
		fmt.Println(reflect.TypeOf(handler.m), "this is handler")
		image["entries"] = handler.Load(f, pretty, nopl)
	}
}
func magicjson() {
	//var jsondata io.Reader
	//jsondata, err := os.Open("./magic-gen/magic.json")
	jsondata, err := ioutil.ReadFile("./magic-gen/magic.json")
	check(err)
	/*
		d := json.NewDecoder(jsondata)
		//d.UseNumber()
		err = d.Decode(&jsonmap)
	*/
	err = json.Unmarshal(jsondata, &jsonmap)
	//fmt.Println(reflect.TypeOf(jsonmap))
	check(err)
	datatest := jsonmap["byname"].(map[string]interface{})
	dataval := jsonmap["byval"].(map[string]interface{})
	//fmt.Println(datatest)
	//fmt.Println(dataval)
	bynamemap = datatest
	byvalmap = dataval
	//data2 := datatest["AUTOFS"]
	//fmt.Println(reflect.TypeOf(bynamemap["IMG_COMMON"]))
	//fmt.Println(bynamemap["IMG_COMMON"])
	//err = json.Unmarshal(jsondata, &objmap)
	//check(err)
}

func magichandler(f *os.File) string {
	fmt.Println(f.Stat())
	bufl := make([]byte, 4)
	buffer, err := f.Read(bufl)
	check(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	//buffer, err := ioutil.ReadAll(f)
	//check(err)
	fmt.Println(buffer, "read buffer in other function")
	img_magic := strconv.FormatUint(uint64(binary.LittleEndian.Uint32(bufl)), 10)
	fmt.Println(img_magic)
	//fmt.Println(bynamemap["IMG_COMMON"])
	//fmt.Println(reflect.TypeOf(img_magic))
	//fmt.Println(reflect.TypeOf(bynamemap["IMG_COMMON"]))
	//if img_magic == bynamemap[IMG_COMMON]
	//conv := bynamemap["IMG_COMMON"]
	// type converting because json being json unmarhsalls as float64
	_, err = f.Seek(4, 0)
	check(err)
	bufl = make([]byte, 4)
	buffer, err = f.Read(bufl)
	fmt.Println(buffer)
	if img_magic == bynamemap["IMG_COMMON"] || img_magic == bynamemap["IMG_SERVICE"] {
		img_magic = strconv.FormatUint(uint64(binary.LittleEndian.Uint32(bufl)), 10)
		//fmt.Println("True")
		fmt.Println(img_magic, "this is image magic")
	}
	m := byvalmap[img_magic]
	//fmt.Println(reflect.TypeOf(m))
	//fmt.Println(m)
	if m == nil {
		fmt.Println("Unknown magic error,please input appropriate files")
		os.Exit(1)
	}
	return m.(string)
}

/*
var handlers = map[string]interface{}{
	"INVENTORY":      &entry_handler{},
	"CORE":           &Entry_handler{},
	"IDS":            &entry_handler{},
	"CREDS":          &entry_handler{},
	"UTSNS":          &entry_handler{},
	"TIMENS":         &entry_handler{},
	"IPC_VAR":        &entry_handler{},
	"FS":             &entry_handler{},
	"GHOST_FILE":     &ghost_file_handler{}, // ghost file
	"MM":             &entry_handler{},
	"CGROUP":         &entry_handler{},
	"TCP_STREAM":     &entry_handler{},
	"STATS":          &entry_handler{},
	"PAGEMAP":        &pagemap_handler{}, // Special one
	"PSTREE":         &entry_handler{},
	"REG_FILES":      &entry_handler{},
	"NS_FILES":       &entry_handler{},
	"EVENTFD_FILE":   &entry_handler{},
	"EVENTPOLL_FILE": &entry_handler{},
	"EVENTPOLL_TFD":  &entry_handler{},
	"SIGNALFD":       &entry_handler{},
	"TIMERFD":        &entry_handler{},
	"INOTIFY_FILE":   &entry_handler{},
	"INOTIFY_WD":     &entry_handler{},
	"FANOTIFY_FILE":  &entry_handler{},
	"FANOTIFY_MARK":  &entry_handler{},
	"VMAS":           &entry_handler{},
	"PIPES":          &entry_handler{},
	"FIFO":           &entry_handler{},
	"SIGACT":         &entry_handler{},
	"NETLINK_SK":     &entry_handler{},
	"REMAP_FPATH":    &entry_handler{},
	"MNTS":           &entry_handler{},
	"TTY_FILES":      &entry_handler{},
	"TTY_INFO":       &entry_handler{},
	"TTY_DATA":       &entry_handler{},
	"RLIMIT":         &entry_handler{},
	"TUNFILE":        &entry_handler{},
	"EXT_FILES":      &entry_handler{},
	"IRMAP_CACHE":    &entry_handler{},
	"FILE_LOCKS":     &entry_handler{},
	"FDINFO":         &entry_handler{},
	"UNIXSK":         &entry_handler{},
	"INETSK":         &entry_handler{},
	"PACKETSK":       &entry_handler{},
	"ITIMERS":        &entry_handler{},
	"POSIX_TIMERS":   &entry_handler{},
	"NETDEV":         &entry_handler{},
	"PIPES_DATA":     &entry_handler{},
	"FIFO_DATA":      &entry_handler{},
	"SK_QUEUES":      &entry_handler{},
	"IPCNS_SHM":      &entry_handler{},
	"IPCNS_SEM":      &entry_handler{},
	"IPCNS_MSG":      &entry_handler{},
	"NETNS":          &entry_handler{},
	"USERNS":         &entry_handler{},
	"SECCOMP":        &entry_handler{},
	"AUTOFS":         &entry_handler{},
	"FILES":          &entry_handler{},
	"CPUINFO":        &entry_handler{},
	"MEMFD_FILE":     &entry_handler{},
	"MEMFD_INODE":    &entry_handler{},
}

var extra_handlers = map[string]interface{}{
	"TCP_STREAM": &tcp_stream_extra_handler{},
	"PIPES_DATA": &pipes_data_extra_handler{},
	"SK_QUEUES":  &sk_queues_extra_handler{},
	"IPCNS_SHM":  &ipc_shm_handler{},
	"IPCNS_SEM":  &ipc_sem_set_handler{},
	"IPCNS_MSG":  &ipc_msg_queue_handler{},
}
*/
