package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	bytesPtr := flag.Bool("c", true, "Prints the byte count")
	flag.Parse()

	fileArg := os.Args[1:]
	if len(fileArg) > 0 {
		if *bytesPtr {
			file, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			fileData, err := file.Stat()
			check(err)

			fmt.Printf("%d %s\n", fileData.Size(), fileData.Name())
		}
	}

}
