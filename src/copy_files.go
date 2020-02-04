package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
As of May 2014, 128KiB is determined to be the minimum block size to best minimize system call overhead.
See: http://git.savannah.gnu.org/cgit/coreutils.git/tree/src/ioblksize.h and
https://eklitzke.org/efficient-file-copying-on-linux
*/
const BufSizeInBytes int64 = 128 * 1024
const DebugMode = false

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Specify 2 arguments src and dest as a parameter")
		return
	}

	srcPath, _ := filepath.Abs(os.Args[1])
	srcFileName := extractFileName(srcPath)

	destPath, _ := filepath.Abs(os.Args[2])
	destFileName := extractFileName(destPath)

	bufSize := calculateBufSize(srcPath)

	if DebugMode {
		fmt.Printf("copy '%s' to '%s' using %d bytes as buffer\n", srcFileName, destFileName, bufSize)
	}

	srcFile, srcErr := os.Open(srcPath)

	if srcErr != nil {
		fmt.Printf("Can't open src file %s, %s\n", srcFileName, srcErr)
		panic(srcErr)
	}

	defer srcFile.Close()

	var buf = make([]byte, bufSize)

	destFile, writeErr := os.Create(destPath)

	if writeErr != nil {
		fmt.Printf("Can't open dest file %s, %s\n", destFileName, writeErr)
		panic(writeErr)
	}

	defer destFile.Close()

	for true {
		copiedBytes, err := srcFile.Read(buf)

		if err == nil {
			_, writeErr := destFile.Write(buf[:copiedBytes])

			if writeErr != nil {
				panic(writeErr)
			}
		} else {
			if err != io.EOF {
				fmt.Printf("Can't read from file %s, %s\n", srcFileName, err)
				panic(err)
			}
			break
		}
	}
}

func calculateBufSize(path string) int64 {
	if fileInfo, err := os.Stat(path); err == nil {
		return minInt64(ceilPowerOf2(fileInfo.Size()), BufSizeInBytes)
	}
	return BufSizeInBytes
}

func minInt64(first int64, second int64) int64 {
	if first <= second {
		return first
	}
	return second
}

/*
Returns the next power of 2 bigger than 'value'.
Example: 20 in decimal = 010100 in binary
1. Extend right most one: 010[100] => 010[111]
2. Fill right side with all 1s till the left most 1 value: 0[10100] => 0[11111]
3. Add one to result: 011111 => 100000
*/
func ceilPowerOf2(value int64) int64 {
	value--
	value |= value >> 1
	value |= value >> 2
	value |= value >> 4
	value |= value >> 8
	value |= value >> 16
	value |= value >> 32
	value++
	return value
}

func extractFileName(path string) string {
	_, file := filepath.Split(path)
	return file
}
