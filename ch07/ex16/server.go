package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	calcHandler := func(w http.ResponseWriter, req *http.Request) {
		formula := req.URL.Query().Get("formula")
		log.Println(formula)
		exec, err := Parse(formula)
		if err != nil {
			fmt.Fprintf(w, "error %+v", err)
			return
		}
		res := exec.Eval(Env{})
		fmt.Fprintf(w, "%g\n", res)
	}
	http.HandleFunc("/", calcHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
