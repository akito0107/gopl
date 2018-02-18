package main

import (
	"flag"
	"log"
	"math/rand"
)

var numOfWorker = flag.Int("n", 100, "number of workers")

const INITIAL = 10000

func main() {
	flag.Parse()

	task := func(in interface{}) interface{} {
		limit := in.(int64)
		r := rand.New(rand.NewSource(limit))

		return r.Int63()
	}
	root := NewWorker(task)
	in := make(chan interface{})

	root.In = in
	root.Run()

	l := root
	for i := 0; i < *numOfWorker-1; i++ {
		w := NewWorker(task)
		l.AppendChild(w)
		w.Run()
		l = w
	}

	in <- int64(INITIAL)
	log.Println(<-l.Out)
	close(in)
}

type Worker struct {
	Id  int64
	In  chan interface{}
	Out chan interface{}
	f   func(arg interface{}) interface{}
}

type task func(arg interface{}) interface{}

func NewWorker(f task) *Worker {
	return &Worker{
		Out: make(chan interface{}),
		f:   f,
	}
}

func (w *Worker) Run() {
	go func() {
		for {
			a, ok := <-w.In
			log.Printf("Id %d got message %+v \n", w.Id, a)
			if !ok {
				log.Printf("Id %d got close \n", w.Id)
				close(w.Out)
				return
			}
			w.Out <- w.f(a)
		}
	}()
}

func (w *Worker) AppendChild(child *Worker) {
	child.In = w.Out
	child.Id = w.Id + 1
}
