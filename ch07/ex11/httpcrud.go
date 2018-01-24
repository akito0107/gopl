package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set item name\n")
		return
	}
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if priceStr == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set price\n")
		return
	}
	db[name] = dollars(price)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s: %s\n", name, dollars(price))
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set item name\n")
		return
	}
	if _, ok := db[name]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set exists name\n")
		return
	}
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if priceStr == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set price\n")
		return
	}
	db[name] = dollars(price)
	w.WriteHeader(http.StatusNoContent)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set item name\n")
		return
	}
	if _, ok := db[name]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "please set exists name\n")
		return
	}
	delete(db, name)
	w.WriteHeader(http.StatusNoContent)
}
