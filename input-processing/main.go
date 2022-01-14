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

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// TODO: Look for lines in the STDIN reader that contain "error" and output them.

	const ERR = "error"

	for {
		readLine, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
		}

		readLineAsStr := string(readLine)
		if strings.Contains(readLineAsStr, ERR) {
			fmt.Println(readLineAsStr)
		}

		if isPrefix == false {
			return
		}

	}
}
