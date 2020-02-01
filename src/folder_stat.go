package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Specify folder")
		return
	}

	absPath, err := filepath.Abs(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Calculating statistics for '%s'...\n", absPath)

	var counter struct {
		filesCount int
		dirsCount  int
		size       int64
	}

	_ = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			counter.dirsCount++
		} else {
			counter.filesCount++
		}

		counter.size += info.Size()

		return nil
	})

	printer := message.NewPrinter(language.English)

	_, _ = printer.Printf("Total: %d files in %d directories, size %d MB\n", counter.filesCount, counter.dirsCount,
		sizeInMb(counter.size))
}

func sizeInMb(sizeInBytes int64) int64 {
	return sizeInBytes / 1024 / 1024
}
