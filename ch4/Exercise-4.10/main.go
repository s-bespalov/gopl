// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.10/github"
)

func main() {
	lastMonth := bytes.NewBufferString("Last month:\n")
	lastYear := bytes.NewBufferString("Last year:\n")
	older := bytes.NewBufferString("Older thar year:\n")

	today := time.Now()
	monthAgo := today.AddDate(0, -1, 0)
	yearAgo := today.AddDate(-1, 0, 0)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		var out *bytes.Buffer
		if item.CreatedAt.After(monthAgo) {
			out = lastMonth
		} else if item.CreatedAt.After(yearAgo) {
			out = lastYear
		} else {
			out = older
		}
		fmt.Fprintf(out, "#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	lastMonth.WriteTo(os.Stdout)
	lastYear.WriteTo(os.Stdout)
	older.WriteTo(os.Stdout)
}
