package main

import (
	"bytes"
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

func lineCount(fileName string) int {
	count := 0
	separator := []byte{'\n'}

	data, err := os.ReadFile(fileName)
	check(err)

	count += bytes.Count(data, separator)

	return count

}

func charCount(fileName string) int {
	count := 0

	data, err := os.ReadFile(fileName)
	check(err)

	count += len(bytes.Runes(data))

	return count
}

func main() {

	bytesPtr := flag.Bool("c", false, "Prints the byte count")
	linesPtr := flag.Bool("l", false, "Prints the number of lines in a file")
	charsPtr := flag.Bool("w", false, "Prints the number of characters in a file")
	flag.Parse()

	fileArg := os.Args[2:]

	if len(fileArg) > 0 {
		if *bytesPtr {

			file, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			fileData, err := file.Stat()
			check(err)

			fmt.Printf("%d %s\n", fileData.Size(), fileData.Name())
		} else if *linesPtr {

			count := lineCount(strings.Join(fileArg, ""))
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))
		} else if *charsPtr {

			count := charCount(strings.Join(fileArg, ""))
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))
		}
	}

}
