package main

import (
	"bytes"
	"os/exec"
	"reflect"
	"testing"
)

func TestByteCount(t *testing.T) {
	wc := exec.Command("wc", "-c", "test.txt")

	var wcout bytes.Buffer
	wc.Stdout = &wcout

	err := wc.Run()
	if err != nil {
		t.Errorf("wc unexpected error %q", err)
	}

	var goWcOut bytes.Buffer

	err = GoWc(&goWcOut, "-c", "test.txt")
	if err != nil {
		t.Errorf("gowc unexpected error %q", err)
	}

	got := goWcOut.Bytes()
	want := wcout.Bytes()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
