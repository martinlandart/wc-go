package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	filename := args[len(args)-1]
	HandleGetByteCountCommand(os.Stdout, filename)

}

func HandleGetByteCountCommand(w io.Writer, filename string) error {
	bc, err := ByteCount(filename)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%v %s\n", bc, filename)
	if err != nil {
		return err
	}

	return nil
}

func ByteCount(filename string) (uint, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return uint(len(file)), nil
}
