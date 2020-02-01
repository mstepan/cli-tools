package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Specify file as a parameter")
		return
	}

	rawPath := os.Args[1]
	absPath, err := filepath.Abs(rawPath)

	if err != nil {
		fmt.Printf("Can't make abs path from file %s\n", rawPath)
		return
	}

	fmt.Printf("Reading content for '%s'\n", absPath)

	var (
		bufSize = 256
		buf     = make([]byte, bufSize)
	)

	file, err := os.Open(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for true {
		n, readErr := file.Read(buf)

		if readErr == nil {
			fmt.Print(string(buf[:n]))
		} else {
			if readErr == io.EOF {
				break
			}
			fmt.Println("Error during read process", readErr)
			return
		}
	}
}

