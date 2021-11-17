package main

import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<!DOCTYPE html>
<html>
    <head>
        <title>Issues</title>
    </head>
    <body>
        <table>
            <tr style="text-align: center;">
                <th>Issues</th>
                <th><a href="/milestones?owner={{.Owner}}&repo={{.Repo}}">Milestones</a></th>
                <th><a href="/users?owner={{.Owner}}&repo={{.Repo}}">Users</a></th>
            </tr>
        </table>
        <h1>{{.Count}} Issues</h1>
        <table>
            <tr style="text-align: left">
                <th>#</th>
                <th>State</th>
                <th>User</th>
                <th>Title</th>
            </tr>
			{{range .Items}}
            <tr>
                <td><a href="{{.URL}}">{{.Number}}</a></td>
                <td>{{.State}}</td>
                <td><a href="{{.User.URL}}">{{.User.Login}}</a></td>
                <td><a href="{{.URL}}">{{.Title}}</a></td>
            </tr>
			{{end}}
        </table>
    </body>
</html>
`))
