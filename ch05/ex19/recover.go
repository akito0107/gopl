package main

import "fmt"

func main() {
	fmt.Printf("value is %d\n", run())
}

func run() (val int) {

	defer func() {
		i := recover()
		val = i.(int)
	}()

	panic(1)
}
