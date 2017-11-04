package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// command line options
type options struct {
	detailed bool
}

func main() {
	detailed := flag.Bool("List detailed", true, "List file details")

	flag.Parse()

	opts := options{*detailed}

	for _, dir := range flag.Args() {
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

	for _, file := range files {
		// w := tabwriter.NewWriter(os.Stdout, 0, 0, 300, '.', tabwriter.Debug|tabwriter.AlignRight)
		listFile(file, opts)
		// w.Flush()
	}

	fmt.Print("\n")
}

func listFile(file os.FileInfo, opts options) {
	if opts.detailed {
		size := fileSize(file)
		fmt.Printf("%s %s %s\n", file.Mode(), size, file.Name())
	} else {
		fmt.Printf(file.Name() + " ")
	}
}

// TODO: work in progress
func fileSize(file os.FileInfo) string {
	// s := file.Size()
	// b := make([]byte, 8)
	// println
	// binary.LittleEndian.PutUint64(b, uint64(file.Size()))
	return "size"
}

// size := humanize.Bytes(file.Size())
