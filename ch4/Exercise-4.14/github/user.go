package github

type User struct {
	Login string
	Url   string `json:"html_url"`
	Id    int
}
