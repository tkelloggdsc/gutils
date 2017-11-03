package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for _, dir := range os.Args[1:] {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("%s is not a directory", dir)
		} else {
			listFiles(dir)
		}
	}
}

func listFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print(err)
	} else {
		for _, file := range files {
			fmt.Printf(file.Name() + " ")
		}
		fmt.Print("\n")
	}
}
