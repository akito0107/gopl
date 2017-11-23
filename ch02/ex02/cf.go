package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/akito0107/gopl/ch02/tempconv"
)

func main() {
	var in string
	if len(os.Args) == 2 {
		in = os.Args[1]
	} else if len(os.Args) == 1 {
		fmt.Scan(&in)
	}
	s, err := convert(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", s)
}

func convert(input string) (string, error) {
	if f, ok := parseFloat(input, "kg"); ok {
		return KToP(Kilo(f)).String(), nil
	}
	if f, ok := parseFloat(input, "lb"); ok {
		return PToK(Pound(f)).String(), nil
	}
	if f, ok := parseFloat(input, "m"); ok {
		return MToF(Metre(f)).String(), nil
	}
	if f, ok := parseFloat(input, "ft"); ok {
		return FToM(Feet(f)).String(), nil
	}
	if f, ok := parseFloat(input, "c"); ok {
		return tempconv.CToF(tempconv.Celsius(f)).String(), nil
	}
	if f, ok := parseFloat(input, "f"); ok {
		return tempconv.Fahrenheit(tempconv.Fahrenheit(f)).String(), nil
	}
	return "", fmt.Errorf("unknown format %s \n", input)
}

func parseFloat(input, suffix string) (float64, bool) {
	if strings.HasSuffix(input, suffix) {
		input = strings.TrimRight(input, suffix)
		f, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return 0, false
		}
		return f, true
	}

	return 0, false
}
