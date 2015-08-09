package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

type testCase struct { //Test case for Param
	cmdline string
	code int
}

var testCases []testCase //Test cases for Param


func TestMain(m *testing.M) {
	//Test cases for CLI
	testCases = []testCase{
		{cmdline: "./demo4 https://github.com/technoweenie", code: ExitCodeOK },
		{cmdline: "./demo4 -c t https://github.com/technoweenie", code: ExitCodeOK },
		{cmdline: "./demo4 -c t noturl", code: ExitCodeError },
		{cmdline: "./demo4  noturl", code: ExitCodeError },
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./demo4 -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("demo4 version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_Arg(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	for _, test := range testCases {
		args := strings.Split(test.cmdline, " ")
		status := cli.Run(args)
		if status != test.code {
			t.Errorf("expected %d to eq %d", status, test.code)
		}
	}
}
