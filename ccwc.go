package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func byteCount(file *os.File) int64 {

	fileData, err := file.Stat()
	check(err)

	return fileData.Size()

}

func lineCount(file *os.File) int {
	count := 0
	separator := []byte{'\n'}

	data, err := io.ReadAll(file)
	check(err)

	count += bytes.Count(data, separator)

	return count

}

func wordCount(file *os.File) int {
	count := 0

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}

	return count
}

func charCount(file *os.File) int {
	count := 0

	data, err := io.ReadAll(file)
	check(err)

	count += len(bytes.Runes(data))

	return count
}

func main() {

	bytesPtr := flag.Bool("c", false, "Prints the byte count")
	linesPtr := flag.Bool("l", false, "Prints the number of lines in a file")
	wordsPtr := flag.Bool("w", false, "Prints the number of words in a file")
	charsPtr := flag.Bool("m", false, "Prints the number of characters in a file")
	flag.Parse()

	fileArg := flag.Args()

	// If given args, file handle is the arg
	if len(fileArg) > 0 {

		fileHandle, err := os.Open(strings.Join(fileArg, ""))
		check(err)

		if *bytesPtr {
			bytesCount := byteCount(fileHandle)

			fmt.Printf("%d %s\n", bytesCount, strings.Join(fileArg, ""))
		} else if *linesPtr {

			count := lineCount(fileHandle)
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))

		} else if *wordsPtr {

			count := wordCount(fileHandle)
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))

		} else if *charsPtr {

			count := charCount(fileHandle)
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))

		} else {

			bytesCount := byteCount(fileHandle)
			linesCount := lineCount(fileHandle)
			wordCount := wordCount(fileHandle)

			fmt.Printf("%d\t%d\t%d %s\n", linesCount, wordCount, bytesCount, strings.Join(fileArg, ""))
		}
	}
}
