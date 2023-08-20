package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	filename := args[len(args)-1]

	var byteCountFlag bool
	flag.BoolVar(&byteCountFlag, "c", false, "print the byte counts")

	var lineCountFlag bool
	flag.BoolVar(&lineCountFlag, "l", false, "print the newline counts")

	flag.Parse()

	if byteCountFlag {
		HandleCommand(os.Stdout, filename, ByteCount)
	}
	if lineCountFlag {
		HandleCommand(os.Stdout, filename, LineCount)
	}

	fmt.Println(filename)
}

func LineCount(file []byte) (int, error) {
	return strings.Count(string(file), "\n"), nil
}

func HandleCommand(w io.Writer, filename string, command func([]byte) (int, error)) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	bc, err := command(file)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%v ", bc)
	if err != nil {
		return err
	}

	return nil
}

func ByteCount(file []byte) (int, error) {
	return len(file), nil
}
