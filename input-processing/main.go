package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	_, err := readLargeInput(os.Stdin, "error")
	if err != nil {
		return
	}
}

func readLargeInput(rd io.Reader, filterStr string) (result []string, err error) {

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(rd)

	// TODO: Look for lines in the STDIN reader that contain "error" and output them.

	if reader == nil {
		return
	}

	lineNum := 0
	for {
		readLine, isPrefix, err1 := reader.ReadLine()
		if err1 != nil {
			err = err1

			return
		}

		readLineAsStr := string(readLine)
		if strings.Contains(readLineAsStr, filterStr) {
			fmt.Println(readLineAsStr)
			result = append(result, readLineAsStr)
		}

		lineNum++

		if isPrefix == false {
			fmt.Printf("Totally %v number of line was read\n", lineNum)

			return
		}
	}
}
