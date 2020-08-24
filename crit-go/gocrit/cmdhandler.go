package gocrit

import (
	"crit-go/images"
	"fmt"
	//"io/ioutil"
	//"io"
	"os"
	//"reflect"
	//"os"

	"encoding/json"
	//"io/ioutil"
)

func Check(e error) {
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

func Decode(inloc string, outloc string, pretty bool, nopl bool) {
	//fmt.Println(inloc)
	//img, err :=
	img := images.Load(inf(inloc), pretty, nopl)
	fmt.Println("test placeholder - Decode called ")
	//fmt.Println(reflect.TypeOf(data))
	ouf := outf(outloc)
	if pretty == true {
		//file, _ := json.MarshalIndent(img, "", "   ")
		encoder := json.NewEncoder(ouf)
		encoder.SetIndent("", "    ")
		err := encoder.Encode(&img)
		checkfile(err, ouf)
		ouf.Close()
	} else {
		encoder := json.NewEncoder(ouf)
		err := encoder.Encode(&img)
		checkfile(err, ouf)
		ouf.Close()
	}
}

func Encode(inloc string, outloc string) {
	// Encodes the json file into binary img file
	// map holding the json file data to be converted to binary
	/*
		var jtb map[string]interface{}
		fmt.Println(inloc)
		file, err := ioutil.ReadFile(inloc)
		fmt.Println(reflect.TypeOf(file))
		if err != nil {
			fmt.Println("Unable to read Input json file")
			fmt.Println(err)
			os.Exit(1)
		}
		if err := json.Unmarshal(file, &jtb); err != nil {
			Check(err)
		}
	*/
	images.Dump(inf(inloc), outf(outloc))
	fmt.Println("test placeholder - encode called ")
}

func Info(inloc string) {
	fmt.Println("test placeholder - Info called ")
}

func Explore(dir string, what string) {
	switch what {
	case "ps":
		explore_ps(dir)
	case "fds":
		explore_fds(dir)
	case "mss":
		explore_mems(dir)
	case "rss":
		explore_rss(dir)
	}
}

func Show(inloc string) {
	fmt.Println("test placeholder - show called")

}

func inf(inloc string) *os.File {
	/*
		return a pointer to the file or stdin
	*/

	if inloc == "" {
		imgfile := os.Stdin
		/*
			if err != nil {
				fmt.Println("failed reading data from stdin: %s", err)
			}
		*/
		//fmt.Println(reflect.TypeOf(imgfile))
		return imgfile
	} else {
		imgfile, err := os.Open(inloc)
		/*
			defer func() {
				if err := imgfile.Close(); err != nil {
					Check(err)
				}
			}()
		*/
		if err != nil {
			fmt.Println("Failed to open input file: %s", err)
		}

		//imgfile, err := ioutil.ReadAll(file)
		//fmt.Println(reflect.TypeOf(imgfile))
		return imgfile
	}
}

func outf(outloc string) *os.File {
	/*
		return a pointer to the file or stdout
	*/
	if outloc == "" {
		outfile := os.Stdout
		/*
			if err != nil {
				fmt.Println("failed reading data from stdin: %s", err)
			}
		*/
		return outfile
	} else {
		outfile, err := os.OpenFile(outloc, os.O_CREATE|os.O_WRONLY, 0644)
		/*
			defer func() {
				if err := imgfile.Close(); err != nil {
					Check(err)
				}
			}()
		*/
		if err != nil {
			fmt.Println("Failed to open output file: %s", err)
		}

		//imgfile, err := ioutil.ReadAll(file)
		//fmt.Println(reflect.TypeOf(imgfile))
		return outfile
	}
}

/*
func outf(outloc string) (file *os.file) {
	fmt.Println("placeholder")
	return nil
}

func dinf(dir string, name string) (file *os.file) {
	fmt.Println("placeholder")
	return nil
}

*/

func explore_ps(dir string) {
	fmt.Println("placeholder")
}

func explore_fds(dir string) {
	fmt.Println("placeholder")
}

func explore_mems(dir string) {
	fmt.Println("placeholder")
}

func explore_rss(dir string) {
	fmt.Println("placeholder")
}
