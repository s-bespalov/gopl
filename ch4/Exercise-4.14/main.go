package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.14/github"
)

var owner, repo string
var issues *[]github.Issue
var milestones *[]github.Milestone
var users *[]github.User

func getIssues(o, r string) {
	if owner == o && repo == r && issues != nil && milestones != nil && users != nil {
		return
	}
	owner, repo = o, r
	var err error
	fmt.Println("request new data from github")
	issues, milestones, users, err = github.GetIssuesMilestonesUsers(o, r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func internalError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write([]byte("500 - server error"))
}

func parseQuery(req *http.Request) (err error) {
	if err = req.ParseForm(); err != nil {
		return 
	}
	var o, r string
	for k, v := range req.Form {
		if k == "owner" {
			o = v[0]
		}
		if k == "repo" {
			r = v[0]
		}
	}
	getIssues(o, r)
	return nil
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if err := parseQuery(r); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if issues == nil {
			internalError(rw)
			return
		}
		var issuesData struct {
			Count int
			Items *[]github.Issue
			Owner string
			Repo  string
		}
		issuesData.Count = len(*issues)
		issuesData.Items = issues
		issuesData.Owner = owner
		issuesData.Repo = repo
		if err := issueList.Execute(rw, issuesData); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})
	http.HandleFunc("/milestones", func(rw http.ResponseWriter, r *http.Request) {
		if err := parseQuery(r); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if milestones == nil {
			internalError(rw)
			return
		}
		var milestonesData struct {
			Count int
			Items *[]github.Milestone
			Owner string
			Repo  string
		}
		milestonesData.Count = len(*milestones)
		milestonesData.Items = milestones
		milestonesData.Owner = owner
		milestonesData.Repo = repo
		if err := milestoneList.Execute(rw, milestonesData); err != nil {
			fmt.Fprintln(os.Stderr, err)
			internalError(rw)
		}
	})
	http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		if err := parseQuery(r); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if users == nil {
			internalError(rw)
			return
		}
		var usersData struct {
			Count int
			Items *[]github.User
			Owner string
			Repo  string
		}
		usersData.Count = len(*users)
		usersData.Items = users
		usersData.Owner = owner
		usersData.Repo = repo
		if err := usersList.Execute(rw, usersData); err != nil {
			fmt.Fprintln(os.Stderr, err)
			internalError(rw)
		}
	})
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
