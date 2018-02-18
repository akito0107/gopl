package main

import (
	"io/ioutil"
	"runtime"
	"testing"
)

func BenchmarkRun1(b *testing.B) {
	runtime.GOMAXPROCS(1)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}

func BenchmarkRun2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}

func BenchmarkRun4(b *testing.B) {
	runtime.GOMAXPROCS(4)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}

func BenchmarkRun8(b *testing.B) {
	runtime.GOMAXPROCS(8)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}

func BenchmarkRun16(b *testing.B) {
	runtime.GOMAXPROCS(16)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}

func BenchmarkRun32(b *testing.B) {
	runtime.GOMAXPROCS(32)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}

func BenchmarkRun64(b *testing.B) {
	runtime.GOMAXPROCS(64)
	defer func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}()
	for i := 0; i < b.N; i++ {
		Run(ioutil.Discard)
	}
}
