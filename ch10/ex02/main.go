package main

import (
	"log"

	"github.com/akito0107/gopl/ch10/ex02/archive"
	_ "github.com/akito0107/gopl/ch10/ex02/tar"
	_ "github.com/akito0107/gopl/ch10/ex02/zip"
)

func main() {
	filename := "sample.tar"

	reader, err := archive.Decode(filename)
	if err != nil {
		log.Fatal(err)
	}
	files := reader.Files()

	for _, f := range files {
		log.Printf("%v\n", f.FileInfo().Name())
	}
}
