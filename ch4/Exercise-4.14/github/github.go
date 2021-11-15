// provides go api for github
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API = "https://api.github.com"
const URL = API + "/repos/%s/%s/issues"

func GetIssuesMilestonesUsers(owner, repo string) (issues *[]Issue, milestones *[]Milestone, users *[]User, err error) {
	url := fmt.Sprintf(URL, owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status not ok, status:%d", resp.StatusCode)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return
	}
	m := make([]Milestone, 0)
	u := make([]User, 0)
	for _, v := range *issues {
		if v.Milestone.Id != 0 {
			m = append(m, v.Milestone)
		}
		if v.User.Id != 0 {
			u = append(u, v.User)
		}

	}
	milestones, users = &m, &u
	return
}
