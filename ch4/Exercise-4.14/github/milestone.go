package github

type Milestone struct {
	URL         string `json:"html_url"`
	Id          int
	Title       string
	Description string
	State       string
}
