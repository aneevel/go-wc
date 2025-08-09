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

	// If this is a regular file, we can simply check the Stat() info
	if fileData.Mode().IsRegular() {
		return fileData.Size()
	}

	// For pipes/stdin, read all data to count bytes
	data, err := io.ReadAll(file)
	check(err)

	return int64(len(data))
}

func lineCount(file *os.File) int {
	count := 0

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		count++
	}

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

		if *bytesPtr {

			fileHandle, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			bytesCount := byteCount(fileHandle)

			fmt.Printf("%d %s\n", bytesCount, strings.Join(fileArg, ""))
		} else if *linesPtr {

			fileHandle, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			count := lineCount(fileHandle)
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))

		} else if *wordsPtr {

			fileHandle, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			count := wordCount(fileHandle)
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))

		} else if *charsPtr {

			fileHandle, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			count := charCount(fileHandle)
			fmt.Printf("%d %s\n", count, strings.Join(fileArg, ""))

		} else {

			fileHandle, err := os.Open(strings.Join(fileArg, ""))
			check(err)

			linesCount := lineCount(fileHandle)

			fileHandle, err = os.Open(strings.Join(fileArg, ""))
			check(err)

			bytesCount := byteCount(fileHandle)

			fileHandle, err = os.Open(strings.Join(fileArg, ""))
			check(err)

			wordCount := wordCount(fileHandle)

			fmt.Printf("%d\t%d\t%d %s\n", linesCount, wordCount, bytesCount, strings.Join(fileArg, ""))
		}
	} else {
		if *bytesPtr {

			bytesCount := byteCount(os.Stdin)
			fmt.Printf("%d\n", bytesCount)

		} else if *linesPtr {

			count := lineCount(os.Stdin)
			fmt.Printf("%d\n", count)

		} else if *wordsPtr {

			count := wordCount(os.Stdin)
			fmt.Printf("%d\n", count)

		} else if *charsPtr {

			count := charCount(os.Stdin)
			fmt.Printf("%d\n", count)

		} else {

			// Read all data from stdin once
			data, err := io.ReadAll(os.Stdin)
			check(err)

			lineFileHandle, err := os.Create("lineFile.txt")
			check(err)
			_, err = lineFileHandle.Write(data)
			check(err)
			_ = lineFileHandle.Close()

			byteFileHandle, err := os.Create("byteFile.txt")
			check(err)
			_, err = byteFileHandle.Write(data)
			check(err)
			_ = byteFileHandle.Close()

			wordFileHandle, err := os.Create("wordFile.txt")
			check(err)
			_, err = wordFileHandle.Write(data)
			check(err)
			_ = wordFileHandle.Close()

			// Reopen files for reading
			lineFileHandle, err = os.Open("lineFile.txt")
			check(err)
			linesCount := lineCount(lineFileHandle)
			_ = lineFileHandle.Close()

			byteFileHandle, err = os.Open("byteFile.txt")
			check(err)
			bytesCount := byteCount(byteFileHandle)
			_ = byteFileHandle.Close()

			wordFileHandle, err = os.Open("wordFile.txt")
			check(err)
			wordCount := wordCount(wordFileHandle)
			_ = wordFileHandle.Close()

			// Lets avoid leaving a bunch of artifacts
			_ = os.Remove("lineFile.txt")
			_ = os.Remove("byteFile.txt")
			_ = os.Remove("wordFile.txt")

			fmt.Printf("%d\t%d\t%d\n", linesCount, wordCount, bytesCount)
		}
	}
}
