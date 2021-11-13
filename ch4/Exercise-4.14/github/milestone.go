package github

type Milestone struct {
	Url         string `json:html_url`
	Id          int
	Title       string
	Description string
	State       string
}
