package main

import "html/template"

var milestoneList = template.Must(template.New("milestonelist").Parse(`
<!DOCTYPE html>
<html>
    <head>
        <title>Milestones</title>
    </head>
    <body>
        <table>
            <tr style="text-align: center;">
                <th><a href="/">Issues</a></th>
                <th>Milestones</th>
                <th><a href="/users">Users</a></th>
            </tr>
        </table>
        <h1>{{.Count}} Milestones</h1>
        <table>
            <tr style="text-align: left">
                <th>Id</th>
                <th>State</th>
                <th>Title</th>
                <th>Description</th>
            </tr>
			{{range .Items}}
            <tr>
                <td><a href="{{.URL}}">{{.Id}}</a></td>
                <td>{{.State}}</td>
                <td><a href="{{.URL}}">{{.Title}}</a></td>
                <td>{{.Description}}</td>
            </tr>
			{{end}}
        </table>
    </body>
</html>
`))
