package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.14/github"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		issues, err := github.GetIssues("golang", "go")
		if err != nil {
			fmt.Fprintln(rw, err)
			fmt.Fprintln(os.Stderr, err)
			return
		}
		var issuesData struct {
			Count int
			Items *[]github.Issue
		}
		issuesData.Count = len(*issues)
		issuesData.Items = issues
		if err = issueList.Execute(rw, issuesData); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
