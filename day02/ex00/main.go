package main

import (
	"flag"
	"fmt"
	"os"
)

var p = fmt.Println
var pf = fmt.Printf

var (
	mode int
	ext  string
	path string
)

const (
	file = 1 << iota
	dir
	link
)

func init() {
	var (
		f bool
		d bool
		l bool
	)

	flag.BoolVar(&f, "f", false, "Find regular files")
	flag.BoolVar(&d, "d", false, "Find directories")
	flag.BoolVar(&l, "sl", false, "Find symlinks")
	flag.StringVar(&ext, "ext", "", "File extension")
	flag.Parse()
	if len(flag.Args()) == 1 {
		path = flag.Arg(0)
	} else if len(flag.Args()) == 0 {
		exeFileName, _ := os.Executable()
		pf("Missing required parameter 'path'\nUsage of %s:\n", exeFileName)
		flag.PrintDefaults()
		os.Exit(2)
	} else {
		exeFileName, _ := os.Executable()
		pf("Junk in args: %v\nUsage of %s:\n", flag.Args(), exeFileName)
		flag.PrintDefaults()
		os.Exit(2)
	}

	fmt.Printf("f = %v\n", f)
	fmt.Printf("d = %v\n", d)
	fmt.Printf("l = %v\n", l)
	fmt.Printf("ext  = %s\n", ext)
	fmt.Printf("path = %s\n", path)
}

func main() {

	fmt.Println("+++++++const+++++++")
	fmt.Println(file)
	fmt.Println(dir)
	fmt.Println(link)

}

func gracefulExit(msg string, val int) {
	p(msg)
	os.Exit(val)
}
