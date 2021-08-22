// Provides Go API for working with Github Issues
package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const api = "https://api.github.com"
const readIssues = api + "/repos/%s/%s/issues/%s"
const IssueUrl = api + "/search/issues"

var auth string

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Body      string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// set OAuth username and access token
func OAuth(u, t string) {
	auth = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", u, t)))
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueUrl + "?per_page=100&q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
