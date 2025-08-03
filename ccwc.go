package main

import (
	"flag"
	"fmt"
)

func main() {

	bytesPtr := flag.Bool("c", true, "Prints the byte count")

	flag.Parse()

	fmt.Println("bytesPtr:", *bytesPtr)
}
