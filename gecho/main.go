package main

import "os"
import "fmt"

func main() {
	const sep = " "
	for _, val := range os.Args[1:] {
		fmt.Print(val + sep)
	}
}
