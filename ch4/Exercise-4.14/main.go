package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Test")
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
