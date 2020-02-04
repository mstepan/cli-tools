package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Specify file name")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Can't open file %s, error: %s\n", filePath, err)
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var linesCount = 0

	for err == nil {
		for prefixOnly := true; err == nil && prefixOnly; {
			_, prefixOnly, err = reader.ReadLine()
		}

		if err == nil {
			// prefixOnly = false, line fully read
			linesCount++
		}
	}

	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Printf("Lines: %d\n", linesCount)

}
