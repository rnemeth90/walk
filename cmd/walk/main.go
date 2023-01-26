package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type config struct {
	extension string
	size      int64
	list      bool
	del       bool
}

var (
	root      string
	list      bool
	extension string
	size      int64
	del       bool
)

func init() {
	flag.StringVar(&root, "r", "", "Root directory to search")
	flag.BoolVar(&list, "l", false, "List files only")
	flag.StringVar(&extension, "e", "", "File extension to search for")
	flag.Int64Var(&size, "s", 0, "Minimum file size to search for")
	flag.BoolVar(&del, "d", false, "delete found files")
}

func main() {
	flag.Parse()

	c := config{
		extension: extension,
		list:      list,
		size:      size,
		del:       del,
	}

	if err := run(root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(rootdir string, w io.Writer, c config) error {
	return filepath.Walk(rootdir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, c.extension, c.size, info) {
			return nil
		}

		if c.list {
			return listFile(path, w)
		}

		if c.del {
			return delete(path)
		}

		return listFile(path, w)
	})
}
