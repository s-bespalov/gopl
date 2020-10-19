package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.14/github2"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		github2.GetRepository("gopl", "go")
		fmt.Fprint(w, "Hello World")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
