package main

import (
	"flag"
	"fmt"
)

var tempC = CelsiusFlag("temp", 20.0, "the celsius temperature")
var tempK = KelvinFlag("kelvin", 295.0, "absolute temperature")

func main() {
	flag.Parse()
	fmt.Println(*tempK)
}
