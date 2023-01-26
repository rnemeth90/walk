package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var (
	binName string = "walk"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")

	resultCode := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(resultCode)
}

func TestRun(t *testing.T) {
	testCases := []struct {
		testName string
		root     string
		cfg      config
		expected string
	}{
		{testName: "NoFilter", root: "testdata", cfg: config{extension: "", size: 0, list: true}, expected: "testdata/dir.log\ntestdata/dir2/script.sh\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var b bytes.Buffer

			if err := run(tc.root, &b, tc.cfg); err != nil {
				t.Fatal(err)
			}

			result := b.String()

			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
