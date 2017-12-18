package main

import (
	"os"
	"strconv"

	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"fmt"

	"github.com/akito0107/gopl/ch04/ex12/crawler"
)

var index map[string]string

func init() {
	index = map[string]string{}
	files, err := ioutil.ReadDir("../data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		var c crawler.Comic
		f, err := ioutil.ReadFile("../data/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(f, &c)
		if err != nil {
			log.Fatal(err)
		}
		index[strconv.Itoa(c.Num)] = c.Transcript
	}

}

func main() {
	subcommand := os.Args[1]
	if subcommand == "crawl" {
		crawler.Crawl()
		return
	}
	if subcommand == "search" {
		search(os.Args[2], os.Stdout)
		return
	}
	log.Fatal("unsupported subcommand: [crawl|search]")
}

func search(query string, out io.Writer) {
	if t, ok := index[query]; ok {
		fmt.Fprintf(out, "%s\n", t)
	} else {
		fmt.Fprint(out, "not found \n")
	}
}
