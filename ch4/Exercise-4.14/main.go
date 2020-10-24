package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.14/github2"
)

var IssueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Milestone</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	{{if .Milestone}}
	<td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
	{{else}}
	<td></td>
	{{end}}
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		args := []string{"repo:golang/go"}
		issues, _ := github2.SearchIssues(args)
		if err := IssueList.Execute(w, issues); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
