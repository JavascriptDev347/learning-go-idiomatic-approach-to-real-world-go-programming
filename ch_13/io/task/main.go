package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

const FILE_NAME = "data.txt.gz"

func main() {
	r, closer, err := BuildGZipReader(FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer closer()
	fmt.Println(r)
}

func BuildGZipReader(fileName string) (*gzip.Reader, func(), error) {

	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gz, func() {
		gz.Close()
		r.Close()
	}, nil
}
