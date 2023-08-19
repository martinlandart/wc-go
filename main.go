package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

}

func GoWc(writer io.Writer, args ...string) error {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		return err
	}

	fmt.Fprintf(writer, "%v test.txt\n", len(file))
	return nil
}
