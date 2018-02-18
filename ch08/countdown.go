package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case time.After(10 * time.Second):
		// do nothing
	case <-abort:
		fmt.Println("Launch abort")
		return
	}
	// launch()
}
