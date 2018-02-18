package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("value is %d\n", run())
	run2()
}

func run() (val int) {

	defer func() {
		i := recover()
		val = i.(int)
	}()
	panic(1)
}

func run2() {
	panic(1)
	defer func() {
		log.Println("----")
	}()
}
