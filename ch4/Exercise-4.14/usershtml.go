package main

import "html/template"

var usersList = template.Must(template.New("userlist").Parse(`
<!DOCTYPE html>
<html>
    <head>
        <title>users</title>
    </head>
    <body>
        <table>
            <tr style="text-align: center;">
                <th><a href="/?owner={{.Owner}}&repo={{.Repo}}">Issues</a></th>
                <th><a href="/milestones?owner={{.Owner}}&repo={{.Repo}}">Milestones</a></th>
                <th>Users</th>
            </tr>
        </table>
        <h1>{{.Count}} Users</h1>
        <table>
            <tr style="text-align: left">
                <th>Id</th>
                <th>User</th>
            </tr>
			{{range .Items}}
            <tr>
                <td><a href="{{.URL}}">{{.Id}}</a></td>
                <td><a href="{{.URL}}">{{.Login}}</a></td>
            </tr>
			{{end}}
        </table>
    </body>
</html>
`))
