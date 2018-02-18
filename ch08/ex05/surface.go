package main

import (
	"fmt"
	"math"
	"sync"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmls='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	buf := make([]string, cells*cells)
	resChan := make(chan result, cells*cells)
	done := make(chan struct{})
	ws := newWorkers(10)
	wg := &sync.WaitGroup{}

	go func(buf []string) {
		defer func() {
			done <- struct{}{}
		}()
		for {
			res, ok := <-resChan
			if !ok {
				return
			}
			buf[res.num] = res.result
		}
	}(buf)

	for i := 0; i < cells; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			w := ws.Get()
			w.Work(cells, i, resChan)
			ws.Put(w)
		}(i)
	}

	wg.Wait()
	close(resChan)

	<-done

	for _, r := range buf {
		fmt.Print(r)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

type worker struct{}

type result struct {
	num    int
	result string
}

func (w *worker) Work(numOfCells int, i int, resultChan chan<- result) {
	for j := 0; j < numOfCells; j++ {
		ax, ay := corner(i+1, j)
		bx, by := corner(i, j)
		cx, cy := corner(i, j+1)
		dx, dy := corner(i+1, j+1)
		res := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy)

		resultChan <- result{
			num:    (i * numOfCells) + j,
			result: res,
		}
	}
}

type workers struct {
	limit chan struct{}
	pool  sync.Pool
}

func newWorkers(n int) *workers {
	ws := workers{}
	ws.limit = make(chan struct{}, n)
	ws.pool = sync.Pool{New: func() interface{} {
		return &worker{}
	}}
	return &ws
}

func (ws *workers) Get() *worker {
	ws.limit <- struct{}{}
	return ws.pool.Get().(*worker)
}

func (ws *workers) Put(w *worker) {
	ws.pool.Put(w)
	<-ws.limit
}
