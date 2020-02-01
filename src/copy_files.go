package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Why we choose 128KB as a buffer size, check this link: https://eklitzke.org/efficient-file-copying-on-linux
// This one is also interesting https://lwn.net/Articles/372384/
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
		return
	}

	defer srcFile.Close()

	var buf = make([]byte, bufSize)

	destFile, destErr := os.Create(destPath)

	if destErr != nil {
		fmt.Printf("Can't open dest file %s, %s\n", destFileName, destErr)
		return
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
				return
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

// Returns the next power of 2 bigger than 'value'
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

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
