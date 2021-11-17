package github

type User struct {
	Login string
	URL   string `json:"html_url"`
	Id    int
}
