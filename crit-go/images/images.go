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
	//"crit-go/protobindings"
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
	//"bytes"
	//"bytes"
	"errors"
	"reflect"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func checkfile(e error, f *os.File) {
	if e != nil {
		f.Close()
		fmt.Println(e)
		os.Exit(1)
	}
}

const sizeof_u16 = 2
const sizeof_u32 = 4
const sizeof_u64 = 8

var jsonmap map[string]interface{}
var bynamemap map[string]interface{}
var byvalmap map[string]interface{}

func Load(f *os.File, pretty bool, nopl bool) map[string]interface{} {
	image := make(map[string]interface{})
	/*
	   Convert criu image from binary format to dict(json).
	   Takes a file-like object to read criu image from.
	   Returns criu image in dict(json) format.
	*/
	magicjson()
	m := magichandler(f)
	fmt.Println(m)
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
	case m == "PIPES_DATA":
		handler := &pipes_data_extra_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "FIFO_DATA":
		handler := &pipes_data_extra_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "SK_QUEUES":
		handler := &sk_queues_extra_handler{}
		handler.m = m
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "IPCNS_SHM":
		handler := &ipc_shm_set_handler{}
		handler.m = m
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "IPCNS_SEM":
		handler := &ipc_sem_set_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "IPCNS_MSG":
		handler := &ipc_msg_queue_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		image["entries"] = handler.Load(f, pretty, nopl)
	case m == "TCP_STREAM":
		handler := &tcp_stream_extra_handler{}
		handler.m = m
		image["entries"] = handler.Load(f, pretty, nopl)
	default:
		handler := &entry_handler{}
		handler.m = m
		//fmt.Println(reflect.TypeOf(handler.m), "this is handler")
		image["entries"] = handler.Load(f, pretty, nopl)
	}
	f.Close()
	return image
}

func Dump(ilf *os.File, ouf *os.File) {
	/*
	   Convert criu image from dict(json) format to binary.
	   Takes an image in dict(json) format and file-like
	   object to write to.
	*/
	magicjson()
	var img map[string]interface{}
	//fmt.Println(inloc)
	file, err := ioutil.ReadAll(ilf)
	if err != nil {
		fmt.Println("Error reading Input Json file: %s", err)
		ilf.Close()
		ouf.Close()
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &img); err != nil {
		check(err)
	}
	m := img["magic"].(string)
	entries := img["entries"]
	fmt.Println(reflect.TypeOf(entries), "extracted magic")
	fmt.Println(entries, "image entries")
	//strconv.Atoi(bynamemap[m].(string))
	//magic_val := bynamemap[m].(string)
	magic_val, err := strconv.ParseUint(bynamemap[m].(string), 10, 64)
	fmt.Println(magic_val, "this is magic val")
	if m != "INVENTORY" {
		if m == "STATS" || m == "IRMAP_CACHE" {
			//bnvmap := bynamemap["IMG_SERVICE"] //.(float64)
			bnvmap, err := strconv.ParseUint(bynamemap["IMG_SERVICE"].(string), 10, 64)
			checkfile(err, ouf)
			//fmt.Println(jsondata)
			//_ = new(bytes.Buffer)
			//err := binary.Write(ouf, binary.LittleEndian, jsondata)
			//checkfile(err, ouf)
			bs := make([]byte, 4)
			binary.LittleEndian.PutUint32(bs, uint32(bnvmap))
			n2, err := ouf.Write(bs)
			fmt.Println(n2, bs, "n2")
			checkfile(err, ouf)
		} else {
			jsondata, err := strconv.ParseUint(bynamemap["IMG_COMMON"].(string), 10, 64) //.(string) //.(float64)
			checkfile(err, ouf)
			bnvmap, err := strconv.ParseUint(bynamemap["IMG_COMMON"].(string), 10, 64)
			checkfile(err, ouf)
			//bdata := []byte(jsondata)
			fmt.Println(reflect.TypeOf(jsondata))
			fmt.Println(uint(jsondata), "uint")
			fmt.Println(uint64(jsondata), "uint64")
			fmt.Println(int64(jsondata), "int64")
			fmt.Println(jsondata, "int")
			//_ = new(bytes.Buffer)
			bs := make([]byte, 4)
			binary.LittleEndian.PutUint32(bs, uint32(bnvmap))
			//cv := make([]byte, 8)
			//binary.PutVarint(cv, int64(jsondata))
			//err = binary.Write(ouf, binary.LittleEndian, jsondata)
			//n2, err := ouf.Write(bs)
			//fmt.Println(n2, bs, "n2")
			n2, err := ouf.Write(bs)
			fmt.Println(n2, bs, "n2")
			checkfile(err, ouf)
		}
	}
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(magic_val))
	n2, err := ouf.Write(bs)
	fmt.Println(n2, bs, "n2")
	if err != nil {
		fmt.Println(err)
	}
	checkfile(err, ouf)
	//err := binary.Write(f, binary.LittleEndian, magic_val)
	//checkfile(err, f)
	/*
			_ = new(bytes.Buffer)
			err := binary.Write(f, binary.LittleEndian, byvalmap)
		checkfile(err, f)
	*/
	switch {
	case m == "GHOST_FILE":
		handler := &ghost_file_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "PAGEMAP":
		handler := &pagemap_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "PIPES_DATA":
		handler := &pipes_data_extra_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "FIFO_DATA":
		handler := &pipes_data_extra_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		//handler.dump(jsonmap["entries"], f)
	case m == "SK_QUEUES":
		handler := &sk_queues_extra_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "IPCNS_SHM":
		handler := &ipc_shm_set_handler{}
		handler.m = m
		///image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "IPCNS_SEM":
		handler := &ipc_sem_set_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "IPCNS_MSG":
		handler := &ipc_msg_queue_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	case m == "TCP_STREAM":
		handler := &tcp_stream_extra_handler{}
		handler.m = m
		//image["entries"] = handler.Load(f, pretty, nopl)
		handler.Dump(img, ouf)
	default:
		handler := &entry_handler{}
		handler.m = m
		//fmt.Println(reflect.TypeOf(handler.m), "this is handler")
		//image["entries"] = handler.Dump(jsonmap, f)
		handler.Dump(img, ouf)
	}
}

func magicjson() {
	/*
		Populates the byname and byval magic maps
	*/
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
	//datatest := jsonmap["byname"].(map[string]interface{})
	//dataval := jsonmap["byval"].(map[string]interface{})
	//fmt.Println(datatest)
	//fmt.Println(dataval)
	bynamemap = jsonmap["byname"].(map[string]interface{})
	byvalmap = jsonmap["byval"].(map[string]interface{})
	//data2 := datatest["AUTOFS"]
	//fmt.Println(reflect.TypeOf(bynamemap["IMG_COMMON"]))
	//fmt.Println(bynamemap["IMG_COMMON"])
	//err = json.Unmarshal(jsondata, &objmap)
	//check(err)
}

func magichandler(f *os.File) string {
	/*
		Fetches the magic of a .img file for the
		load function.
	*/
	fmt.Println(f.Stat())
	bufl := make([]byte, 4)
	buffer, err := f.Read(bufl)
	checkfile(err, f)
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
	checkfile(err, f)
	bufl = make([]byte, 4)
	buffer, err = f.Read(bufl)
	checkfile(err, f)
	fmt.Println(buffer, "otherbuffer")
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
		f.Close()
		os.Exit(1)
	}
	return m.(string)
}

func roundup(size uint32, sizeof uint32) uint32 {
	round := ((size - 1) | (sizeof - 1) + 1)
	return round
}

func parseInt(i interface{}) (int, error) {
	s, ok := i.(string)
	if !ok {
		return 0, errors.New("not string")
	}
	return strconv.Atoi(s)
}
