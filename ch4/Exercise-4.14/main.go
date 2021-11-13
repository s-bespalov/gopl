package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.14/github"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		issues, err := github.GetIssues("golang", "go")
		if err != nil {
			fmt.Fprintln(rw, err)
			return
		}
		fmt.Fprintln(rw, issues)
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
