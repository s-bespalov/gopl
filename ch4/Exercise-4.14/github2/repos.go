// Package github2 provides a Go API for Github repository tracker.
package github2

import "time"

const reposUrl = "https://api.github.com/repos"

type Repository struct {
	Name            string
	Description     string
	HTMLURL         string `json:"html_url`
	issuesURL       string `json:issues_url`
	contributorsURL string `json:contributors_url`
	milestonesURL   string `json:milestones_url`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
