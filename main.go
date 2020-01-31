package main

import (
	"errors"
	"fmt"
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

	fmt.Printf("Listing for '%s'\n", absPath)

	var counter struct {
		filesCount int
		dirsCount  int
	}

	filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			counter.dirsCount++
		} else {
			counter.filesCount++
		}
		//fmt.Println("-", path)

		return nil
	})

	fmt.Printf("Total: %d files in %d directories\n", counter.filesCount, counter.dirsCount)
}

func printWorkingDirOrFail() {
	if dir, err := os.Getwd(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("working dir: %s\n", dir)
	}
}

func changeWorkingDirectoryOrFail(dir string) {
	if err := os.Chdir(dir); err != nil {
		fmt.Printf("Can't change working directory %e\n", err)
		os.Exit(1)
	}
}

func factorial(x int) (int, error) {

	if x < 0 {
		return -1, errors.New(fmt.Sprintf("Can't calculate factorial for negative value: %d", x))
	}

	if x == 0 {
		return 1, nil
	}
	res := 1

	for i := 2; i <= x; i++ {
		res *= i
	}

	return res, nil
}
