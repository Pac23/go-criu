package images

import (
    //"bufio"
    "encoding/binary"
    //"encoding/json"
    "fmt"
    //"io"
    //"io/ioutil"
    "os"
    //"strings"
    //"internal"
    //"reflect"
    "bufio"
    //"strconv"
    //"buffer"
    "reflect"
    //"io/ioutil"
    //"io"
    "crit-go/protobindings"
    "google.golang.org/protobuf/proto"
)

/*
type Entry_handler interface {
    Load() []interface{}
}
*/

type entry_handler struct {
    m string
}

type ipc_shm_handler struct{}

type ipc_msg_queue_handler struct{}

type ipc_sem_set_handler struct{}

type tcp_stream_extra_handler struct{}

type sk_queues_extra_handler struct{}

type pipes_data_extra_handler struct{}

type ghost_file_handler struct {
    m string
}

type pagemap_handler struct {
    m string
}

func (m *entry_handler) Load(f *os.File, pretty bool, nopl bool) string {
    /*
       Convert criu image entries from binary format to dict(json).
       Takes a file-like object and returnes a list with entries in
       dict(json) format.
    */
    //var entries []interface{}

    fmt.Println("load method called")
    fmt.Println(f.Stat())
    mval := m.m
    fmt.Println(reflect.TypeOf(mval), "this is handler")
    protobinding := pbuffhandler[mval] //pbuffhandler[m]
    readbuffer := make([]byte, 4)
    r := bufio.NewReader(f)
    for {
        n, err := r.Read(readbuffer)
        if err != nil {
            fmt.Println("Error reading file:", err)
            break
        }
        fmt.Println("bytes read: ", n)
        fmt.Println("bytestream : ", readbuffer)
        fmt.Println("bytestream to string: ", string(readbuffer))
        //fmt.Println(string(b[0:n]))
        //size := strconv.FormatUint(uint64(binary.LittleEndian.Uint32(readbuffer)), 10)
        size := uint64(binary.LittleEndian.Uint32(readbuffer))
        fmt.Println(reflect.TypeOf(size))
        //fmt.Println("size", size)
        internalrbuffer := make([]byte, size)
        n, err = r.Read(internalrbuffer)

        if err := proto.Unmarshal(internalrbuffer, protobinding); err != nil {
            fmt.Println("Failed to parse data into proto: ", err)
        }
        //fmt.Println("bytes read later: ", n)
        //fmt.Println("bytestream later : ", internalrbuffer)
        //fmt.Println("bytestream to string later: ", string(internalrbuffer))
        //size := strconv.FormatUint(uint64(binary.LittleEndian.Uint32(readbuffer)), 10)
        //fmt.Println("size", size)
    }
    var entries = "test"
    return entries
}

func (m *ghost_file_handler) Load(f *os.File, pretty bool, nopl bool) string {
    /*
       Convert criu image entries from binary format to dict(json).
       Takes a file-like object and returnes a list with entries in
       dict(json) format.
    */
    //var entries []interface{}

    fmt.Println("load method of ghost file handler called")
    return "hi"
}

func (m *pagemap_handler) Load(f *os.File, pretty bool, nopl bool) string {
    /*
       Convert criu image entries from binary format to dict(json).
       Takes a file-like object and returnes a list with entries in
       dict(json) format.
    */
    //var entries []interface{}

    fmt.Println("load method of pagemap called")
    return "hi"
}
