package main

import (
	"flag"
	"fmt"
	"natsuPacker/packager"
	"os"
)

var fileInput string

func main() {
	flag.StringVar(&fileInput, "input", "", "Input File (path/to/file)")
	flag.Parse()
	if fileInput == "" {
		flag.PrintDefaults()
		return
	}
	_, err := os.Open(fileInput)
	if err != nil {
		fmt.Println("Invalid Input File!,Make sure you put the file location correctly")
	} else {
		packager.Packager(fileInput)
	}
}
