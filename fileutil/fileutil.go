package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	srcFile = "src/query.log"
	dstFile = "dst/query.log"
)

func main() {
	// if err := moveFile(srcFile, dstFile); err != nil {
	// panic(err)
	// }

	files, err := listGzipFile("src/*.gz")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}

func moveFile(src, dst string) error {
	return os.Rename(src, dst)
}

func listGzipFile(path string) ([]string, error) {
	return filepath.Glob(path)
}
