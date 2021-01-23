package main

import (
	"flag"
	"fmt"
	"os"
)

var p = fmt.Println
var pf = fmt.Printf
var sf = fmt.Sprintf

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
		f  bool
		d  bool
		sl bool
	)

	exeFileName, _ := os.Executable()

	flag.BoolVar(&f, "f", false, "Find regular files")
	flag.BoolVar(&d, "d", false, "Find directories")
	flag.BoolVar(&sl, "sl", false, "Find symlinks")
	flag.StringVar(&ext, "ext", "", "File extension")
	flag.Parse()
	if len(flag.Args()) == 1 {
		path = flag.Arg(0)
	} else if len(flag.Args()) == 0 {
		inputError(sf("Missing required parameter 'path'\nUsage of %s:\n", exeFileName))
	} else {
		inputError(sf("Junk in args: %v\nUsage of %s:\n", flag.Args(), exeFileName))
	}

	if !f && !d && !sl {
		mode = file + dir + link
	}
	if f {
		mode += file
	}
	if d {
		mode += dir
	}
	if sl {
		mode += link
	}

	if ext != "" && mode != file {
		inputError(sf("-ext working only with -f\nUsage of %s:\n", exeFileName))
	}

	fmt.Printf("f = %v\n", f)
	fmt.Printf("d = %v\n", d)
	fmt.Printf("sl = %v\n", sl)
	fmt.Printf("ext  = %s\n", ext)
	fmt.Printf("path = %s\n", path)
}

func main() {
	fmt.Println("+++++++const+++++++")
	fmt.Println(file)
	fmt.Println(dir)
	fmt.Println(link)

	// if ext != "" {
	// 	findWithExt()
	// }

	findFiles(path)
}

// func findWithExt() {
// 	// fileInfo, err := os.Stat(path)
// 	// if err != nil {
// 	// 	gracefulExit(err.Error())
// 	// }
// 	// if fileInfo.Mode().IsDir() == false {
// 	// 	gracefulExit("Incorrect path")
// 	// }

// 	files, err := os.Readdir()
// }

func findFiles(pth string) { // os.Stat
	fd, err := os.Open(pth)
	if err != nil {
		p(err)
	} else {
		// p(fd)
		// x, err := fd.Readdir(-1)
		// // p(x)
		// for _, i := range x {
		// 	p(i)
		// }
		// p(err)
		files, err := fd.Readdir(-1)
		if err != nil {
			// p(err)
		} else {
			for _, file := range files {
				find
			}
		}
	}
}

func inputError(msg string) {
	p(msg)
	flag.PrintDefaults()
	os.Exit(2)
}

func gracefulExit(msg string) {
	p(msg)
	os.Exit(1)
}
