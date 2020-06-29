package gocrit

import (
	"crit-go/images"
	"fmt"
	//"io/ioutil"
	//"io"
	"os"
	"reflect"
	//"os"
)

func Check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func Decode(inloc string, outloc string, pretty bool, nopl bool) {
	//fmt.Println(inloc)
	//img, err :=
	images.Load(inf(inloc), pretty, nopl)
	fmt.Println("test placeholder - Decode called ")
	//fmt.Println(reflect.TypeOf(data))
}

func Encode(inloc string, outlock string) {
	fmt.Println("test placeholder - Encode called ")
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
	if inloc == "" {
		imgfile := os.Stdin
		/*
			if err != nil {
				fmt.Println("failed reading data from stdin: %s", err)
			}
		*/
		fmt.Println(reflect.TypeOf(imgfile))
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
