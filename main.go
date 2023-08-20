package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
}

func HandleGetByteCountCommand(w io.Writer, filename string) error {
	bc, _ := ByteCount(filename)
	fmt.Fprintf(w, "%v %s\n", bc, filename)
	return nil
}

func ByteCount(filename string) (uint, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return uint(len(file)), nil
}
