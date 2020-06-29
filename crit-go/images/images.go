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
	"reflect"
	"strconv"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	//"github.com/golang/protobuf/jsonpb"
	 //"google.golang.org/protobuf/encoding/protojson"
	//"strings"
	//"bytes"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

var jsonmap map[string]interface{}
var bynamemap map[string]interface{}
var byvalmap map[string]interface{}

func Load(f *os.File, pretty bool, nopl bool) {
	/*
	   Convert criu image from binary format to dict(json).
	   Takes a file-like object to read criu image from.
	   Returns criu image in dict(json) format.
	*/
	//image := make(map[string]int)
	//m, handler := rhandler(f)
	rhandler(f)
	//image["magic"] = m
	//image["entries"] = 5
	//j := binary.LittleEndian.Uint32(f)
	//fmt.Println(j)
}

func rhandler(f *os.File) {
	buffer, err := ioutil.ReadAll(f)
	check(err)
	img_magic := strconv.FormatUint(uint64(binary.LittleEndian.Uint32(buffer[0:4])), 10)
	fmt.Println(reflect.TypeOf(img_magic))
	fmt.Println(reflect.TypeOf(bynamemap["IMG_COMMON"]))
	magicjson()
	//if img_magic == bynamemap[IMG_COMMON]
	//conv := bynamemap["IMG_COMMON"]
	// type converting because json being json unmarhsalls as float64
	if img_magic == bynamemap["IMG_COMMON"] || img_magic == bynamemap["IMG_SERVICE"] {
		img_magic = strconv.FormatUint(uint64(binary.LittleEndian.Uint32(buffer[4:8])), 10)
		fmt.Println("True")
	}
	/*go run ma
	if img_magic == uint32(bynamemap["IMG_COMMON"].(float64)) || img_magic == uint32(bynamemap["IMG_SERVICE"].(float64)) {
		img_magic = binary.LittleEndian.Uint32(buffer[4:8])
		fmt.Println("True")
	}
	*/
	m := byvalmap[img_magic]
	if m == nil {
		fmt.Println("Unknown magic error,check input files magic")
	}
	fmt.Println(m, buffer[8])
	size := buffer[8]
	fmt.Println(size)
	//buf := readNextBytes(f, 43)
	fmt.Println(buffer, "buffer")
	//buff := bytes.NewBuffer(buf)
	//buff := strconv.FormatUint(uint64(binary.LittleEndian.Uint32(buffer[43:])), 10)
	//fmt.Println(buff)
	fmt.Printf("Data as hex: %x\n", buffer)
   	fmt.Printf("Data as string: %s\n", buffer)
   	fmt.Println("Number of bytes read:", len(buffer))
   	//myString := string(buffer[55:98])
   	//fmt.Println(myString)
   	//fmt.Println(reflect.TypeOf(strconv.Itoa(int(buffer[43:93]))
	//address := 43
	protobin := &protobindings.FdinfoEntry{}
	//r := strings.NewReader(string(buffer[55:98]))
	//fmt.Println(reflect.TypeOf(r))
	//bbyte := []byte("D\x00\x00\x00\x08\x01\x10\x02\x1a>\x08\x02\x10\x00\x18\x00*\n\x08\x00\x10\x00\x18\x00 \x00(\x002'/lib/x86_64-l")
	/*
	if err := jsonpb.Unmarshal(r, protobin); err != nil {
    	panic(err)
	}
	*/

    fmt.Printf("%x ", buffer)
    /*
	for i := 0; i < len(bbyte); i++ {
        fmt.Printf("%x ", bbyte[i])
    }
    */
	err = proto.UnmarshalOptions{AllowPartial: true}.Unmarshal(buffer, protobin)
	if err != nil {
		fmt.Println(err)
	}
	if err := prototext.Unmarshal(buffer, protobin); err != nil {
    	fmt.Println(err, "unable to work wiht prototext")
	}
	if err := protojson.Unmarshal(buffer, protobin); err != nil {
    	fmt.Println(err)
	}
	//if err := proto.Unmarshal(buffer[43:98], protobin)
	if err := proto.Unmarshal(buffer[12:20], &protobindings.FdinfoEntry{}); err != nil {
        fmt.Println("Failed to parse data into proto fdinfoentry:", err)
	}
	/*
	if err := proto.Unmarshal(buffer, &protobindings.FileEntry{}); err != nil {
        fmt.Println("Failed to parse data into proto filentry:", err)
	}
	*/

	fmt.Println(protobin.GetId())
	fmt.Println(protobin.GetFlags())
	fmt.Println(protobin.GetType())
	fmt.Println(protobin.GetFd())
	err = protojson.UnmarshalOptions{AllowPartial: false}.Unmarshal(buffer, &protobindings.FileEntry{}); 
	if err != nil {
		fmt.Println("Failed to conver prot to json:", err)
	}
	hex, err := protojson.Marshal(protobin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex)
	//fmt.Println(buf)
	/*
		handler := handlers[m]
		if handler == nil {
			fmt.Println("Unknwon handler,Handler not found")
		}

		return m, handler
		/*
			if err != nil {
				fmt.Println("Image magic is missing in file")
			}
	*/
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
	check(err)
	datatest := jsonmap["byname"].(map[string]interface{})
	dataval := jsonmap["byval"].(map[string]interface{})
	bynamemap = datatest
	byvalmap = dataval
	//data2 := datatest["AUTOFS"]
	fmt.Println(reflect.TypeOf(bynamemap["IMG_COMMON"]))
	fmt.Println(bynamemap["IMG_COMMON"])
	//err = json.Unmarshal(jsondata, &objmap)
	//check(err)
}
