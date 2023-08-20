package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:]
	filename := args[len(args)-1]

	var byteCountFlag bool
	flag.BoolVar(&byteCountFlag, "c", false, "print the byte counts")

	var lineCountFlag bool
	flag.BoolVar(&lineCountFlag, "l", false, "print the newline counts")

	var wordCountFlag bool
	flag.BoolVar(&wordCountFlag, "w", false, "print the word counts")

	var charCountFlag bool
	flag.BoolVar(&charCountFlag, "m", false, "print the character counts")

	flag.Parse()

	if byteCountFlag {
		_ = HandleCommand(os.Stdout, filename, ByteCount)
	}
	if lineCountFlag {
		_ = HandleCommand(os.Stdout, filename, LineCount)
	}
	if wordCountFlag {
		_ = HandleCommand(os.Stdout, filename, WordCount)
	}
	if charCountFlag {
		_ = HandleCommand(os.Stdout, filename, CharCount)
	}

	fmt.Println(filename)
}

func CharCount(file []byte) (int, error) {
	str := string(file)
	counter := 0
	strings.FieldsFunc(str, func(r rune) bool {
		counter++
		return false
	})
	return counter, nil
}

func WordCount(file []byte) (int, error) {
	words := strings.FieldsFunc(string(file), func(r rune) bool {
		return unicode.IsSpace(r)
	})
	return len(words), nil
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
