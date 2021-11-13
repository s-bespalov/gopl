// provides go api for github
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API = "https://api.github.com"
const URL = API + "/repos/%s/%s/issues"

func GetIssues(owner, repo string) (*[]Issue, error) {
	url := fmt.Sprintf(URL, owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status not ok, status:%d", resp.StatusCode)
	}
	issues := make([]Issue, 0)
	if err = json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return &issues, nil
}
