package main

import (
	"log"
	"os"
	"time"
)

func main() {
	res := make(chan int)

	ping := func(in chan struct{}, out chan struct{}, done chan struct{}) func() {
		var cnt int
		return func() {
			for {
				select {
				case <-in:
					cnt++
					go func() {
						out <- struct{}{}
					}()
				case <-done:
					res <- cnt
					return
				}
			}
		}
	}

	upstream := make(chan struct{})
	downstream := make(chan struct{})
	done := make(chan struct{})
	p1 := ping(upstream, downstream, done)
	p2 := ping(downstream, upstream, done)

	go p1()
	go p2()
	upstream <- struct{}{}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		log.Println("cancelled")
		close(done)
	}()

	go func() {
		t := time.NewTicker(1 * time.Second)
		<-t.C
		close(done)
	}()

	log.Printf("count %d \n", <-res)
}
