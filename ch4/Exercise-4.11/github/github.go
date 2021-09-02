// Provides Go API for working with Github Issues
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const api = "https://api.github.com"
const readIssue = api + "/repos/%s/%s/issues/%s"
const pathIssue = readIssue
const postIssue = api + "/repos/%s/%s/issues"
const searchIssueUrl = api + "/search/issues"

var auth [2]string

type Issue struct {
	Number    int       `json:"number"`
	HTMLURL   string    `json:"html_url"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Body      string    `json:"body"`
	User      *User     `json:"user"`
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
	url := searchIssueUrl + "?per_page=100&q=" + q
	r, err := httpHelper(http.MethodGet, url, nil, &IssuesSearchResult{})
	if err != nil {
		return nil, err
	}
	result := r.(*IssuesSearchResult)
	return result, nil
}

func ReadIssue(params []string) (*Issue, error) {
	escapeParams(&params)
	if len(params) < 3 {
		return nil, fmt.Errorf("read issuue, missed params")
	}
	url := fmt.Sprintf(readIssue, params[0], params[1], params[2])
	r, err := httpHelper(http.MethodGet, url, nil, &Issue{})
	if err != nil {
		return nil, err
	}
	result := r.(*Issue)
	return result, nil
}

func PatchIssue(params []string, issue *Issue) (*Issue, error) {
	escapeParams(&params)
	if len(params) < 3 {
		return nil, fmt.Errorf("path issuue, missed params")
	}
	url := fmt.Sprintf(pathIssue, params[0], params[1], params[2])
	body, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	r, err := httpHelper(http.MethodPatch, url, bytes.NewBuffer(body), &Issue{})
	if err != nil {
		return nil, err
	}
	result := r.(*Issue)
	return result, nil
}

func CloseIssue(params []string) (*Issue, error) {
	escapeParams(&params)
	if len(params) < 3 {
		return nil, fmt.Errorf("path issuue, missed params")
	}
	url := fmt.Sprintf(pathIssue, params[0], params[1], params[2])
	r, err := httpHelper(http.MethodPatch, url, bytes.NewBuffer([]byte("{\"state\":\"close\"}")), &Issue{})
	if err != nil {
		return nil, err
	}
	result := r.(*Issue)
	return result, nil
}

func PostIssue(params []string, issue *Issue) (*Issue, error) {
	escapeParams(&params)
	if len(params) < 2 {
		return nil, fmt.Errorf("post issuue, missed params")
	}
	url := fmt.Sprintf(postIssue, params[0], params[1])
	body, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	r, err := httpHelper(http.MethodPost, url, bytes.NewBuffer(body), &Issue{})
	if err != nil {
		return nil, err
	}
	result := r.(*Issue)
	return result, nil
}

func escapeParams(params *[]string) {
	p := *params
	for i, v := range p {
		p[i] = url.QueryEscape(v)
	}
}

func httpHelper(method, url string, body io.Reader, response interface{}) (interface{}, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	if auth[0] != "" && auth[1] != "" {
		req.SetBasicAuth(auth[0], auth[1])
	}
	c := http.Client{}
	r, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		if method != http.MethodPost && r.StatusCode != http.StatusCreated {
			return nil, fmt.Errorf("search query failed: %s", r.Status)
		}
	}
	if err := json.NewDecoder(r.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil
}
