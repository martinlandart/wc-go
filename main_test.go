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
		{
			Name: "line count",
			Args: []string{"-l", testFile},
		},
	}
	for _, testCase := range testCases {
		wc := exec.Command("wc", testCase.Args...)
		gowc := exec.Command("./gowc", testCase.Args...)
		AssertEqualCommandOutput(t, testCase.Name, gowc, wc)
	}
}

func AssertEqualCommandOutput(t *testing.T, testName string, command *exec.Cmd, target *exec.Cmd) {
	t.Helper()
	execute := func(out io.Writer, cmd *exec.Cmd) error {
		cmd.Stdout = out
		err := cmd.Run()
		if err != nil {
			t.Fatalf("%s: gowc unexpected error %q", testName, err)
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
		t.Fatalf("%s: got '%s' want '%s'", testName, got.Bytes(), want.Bytes())
	}
}

func RequireNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}
}
