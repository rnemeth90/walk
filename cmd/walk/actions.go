package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// filter out and ignore the path
func filterOut(path string, ext string, size int64, fInfo fs.FileInfo) bool {
	if fInfo.IsDir() || fInfo.Size() < size {
		return true
	}

	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false
}

func listFile(file string, writer io.Writer) error {
	_, err := fmt.Fprintln(writer, file)
	return err
}

func delete(file string) error {
	return os.Remove(file)
}
