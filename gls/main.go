package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

// command line options
type options struct {
	detailed bool
}

func main() {
	detailed := flag.Bool("l", false, "List file details")

	flag.Parse()

	opts := options{*detailed}

	for _, dir := range flag.Args() {
		println(dir)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		listFiles(dir, opts)
	}
}

func listFiles(dir string, opts options) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// find the length of the biggest one and make them all that
	// prepend spaces to the beginning until they equal the largest one's size
	for _, file := range files {
		listFile(file, opts)
	}

	fmt.Print("\n")
}

func listFile(file os.FileInfo, opts options) {
	if opts.detailed {
		size := fileSize(file)
		modified := file.ModTime().Format(time.RFC822)
		fmt.Printf("%s %s %s ", file.Mode(), modified, size)

		if file.IsDir() {
			color.Set(color.FgHiBlue, color.Bold)
			fmt.Printf("%s\n", file.Name())
			color.Unset()
		} else {
			fmt.Printf("%s\n", file.Name())
		}

	} else {
		fmt.Printf(file.Name() + " ")
	}
}

func fileSize(file os.FileInfo) string {
	s := file.Size()
	return humanize.Bytes(uint64(s))
}
