package main

import (
	"io/ioutil"
)

func main() {}

func scanPath(path string) ([]string, error) {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, f := range fs {
		if f.IsDir() == false {
			result = append(result, f.Name())
		}
	}
	return nil, nil
}
