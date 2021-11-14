package github

type Issue struct {
	Id        int
	URL       string `json:"html_url"`
	Number    int
	CreatedAt string `json:"created_at"`
	State     string
	Title     string
	Body      string
	Milestone Milestone
	User      User
}
