package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	x := 2.0
	if xf, e := strconv.ParseFloat(r.URL.Query().Get("x"), 64); e == nil && xf > 0.0 {
		x = xf
	}
	y := 2.0
	if yf, e := strconv.ParseFloat(r.URL.Query().Get("y"), 64); e == nil && yf > 0.0 {
		y = yf
	}
	res := 1
	if resi, e := strconv.Atoi(r.URL.Query().Get("res")); e == nil && resi > 0 {
		res = resi
	}
	w.Header().Set("Content-Type", "image/png")
	draw(x, y, res, w)
}
