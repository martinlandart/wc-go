package main

import (
	"bytes"
	"io"
	"os/exec"
	"reflect"
	"testing"
)

var testFile string = "test.txt"

func TestBehaviourMatchesWc(t *testing.T) {
	testCases := []struct {
		Name string
		Args []string
	}{
		{
			Name: "byte count",
			Args: []string{"-c", testFile},
		},
	}
	for _, testCase := range testCases {
		wc := exec.Command("wc", testCase.Args...)
		gowc := exec.Command("./gowc", testCase.Args...)
		AssertEqualCommandOutput(t, gowc, wc)
	}
}

func AssertEqualCommandOutput(t *testing.T, command *exec.Cmd, target *exec.Cmd) {
	t.Helper()
	execute := func(out io.Writer, cmd *exec.Cmd) error {
		cmd.Stdout = out
		err := cmd.Run()
		if err != nil {
			t.Fatalf("gowc unexpected error %q", err)
		}
		return nil
	}
	var got bytes.Buffer
	var want bytes.Buffer

	err := execute(&got, command)
	RequireNoErr(t, err)
	err = execute(&want, target)
	RequireNoErr(t, err)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q want %q", got, want)
	}
}

func RequireNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}
}

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
