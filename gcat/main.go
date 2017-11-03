package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, path := range os.Args[1:] {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Print("File does not exist")
		} else {
			echoFile(path)
		}
	}
}

func echoFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
