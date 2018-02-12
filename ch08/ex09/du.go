package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	log.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walk(root, &n, os.Stdout)
	}
	n.Wait()
}

func printDiskUsage(w io.Writer, root string, nfiles, nbytes int64) {
	fmt.Fprintf(w, "root: %s, %d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func walk(dir string, n *sync.WaitGroup, w io.Writer) {
	defer n.Done()
	fileSizes := make(chan int64)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go walkDir(dir, wg, fileSizes)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(w, dir, nfiles, nbytes)
		}
	}
	printDiskUsage(w, dir, nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
