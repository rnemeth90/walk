package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {

	testCases := []struct {
		testName string
		file     string
		ext      string
		size     int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
		{"FilterExtensionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", ".sh", 0, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", ".log", 10, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			info, err := os.Stat(test.file)
			if err != nil {
				t.Fatal(err)
			}

			f := filterOut(test.file, test.ext, test.size, info)

			if f != test.expected {
				t.Errorf("Expected %t, got %t instead\n", test.expected, f)
			}

		})
	}
}
