package main

import (
	"bytes"
	"reflect"
	"testing"
)

// func TestByteCount(t *testing.T) {
// 	wc := exec.Command("wc", "-c", "test.txt")
// 	var wcout bytes.Buffer
// 	wc.Stdout = &wcout
// 	err := wc.Run()
// 	if err != nil {
// 		t.Fatalf("wc unexpected error %q", err)
// 	}
//
// 	var goWcOut bytes.Buffer
// 	os.Args = []string{"-c", "test.txt"}
// 	err = GoWc(&goWcOut)
// 	if err != nil {
// 		t.Fatalf("gowc unexpected error %q", err)
// 	}
// 	got := goWcOut.Bytes()
// 	want := wcout.Bytes()
// 	if !reflect.DeepEqual(got, want) {
// 		t.Fatalf("got %q want %q", got, want)
// 	}
// }

var testFile string = "test.txt"

func TestHandleGetByteCountCommand(t *testing.T) {
	var got bytes.Buffer
	err := HandleGetByteCountCommand(&got, testFile)
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}

	want := bytes.NewBufferString("335191 test.txt\n")
	if !reflect.DeepEqual(got, *want) {
		t.Fatalf("got %s want %s", got.Bytes(), want.Bytes())
	}
}

func TestByteCount(t *testing.T) {
	want := uint(335191)
	got, err := ByteCount(testFile)
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
