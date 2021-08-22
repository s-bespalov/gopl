// Provides Go API for working with Github Issues
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const api = "https://api.github.com"
const readIssues = api + "/repos/%s/%s/issues/%s"
const assignedIssues = api + "/repos/%s/%s/issues"
const IssueUrl = api + "/search/issues"

var auth [2]string

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
	auth = [2]string{u, t}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	url := IssueUrl + "?per_page=100&q=" + q
	t := struct{}(IssuesSearchResult{})
	r, err := get(url, t)
	if err != nil {
		return nil, err
	}
	result := r.(IssuesSearchResult)

	return &result, nil
}

func ReadIssue(params []string) (*Issue, error) {
	escapeParams(&params)
	url := fmt.Sprintf(readIssues, params[0], params[1], params[2])
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("read issue failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func escapeParams(params *[]string) {
	p := *params
	for i, v := range p {
		p[i] = url.QueryEscape(v)
	}
}

func get(url string, r *struct{}) (*interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth[0] != "" && auth[1] != "" {
		req.SetBasicAuth(auth[0], auth[1])
	}
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}
	return r, nil
}
