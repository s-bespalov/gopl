package github

type User struct {
	login string
	url   string `json:html_url`
}
