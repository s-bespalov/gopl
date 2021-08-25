package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.11/github"
)

const creditsFile = "credits"
const tmpFolder = "tmp/"
const creditsPath = tmpFolder + creditsFile
const jsonIndent = "  "
const jsonPrefix = ""

var token string
var user string
var mode string
var editor string

func init() {
	flag.StringVar(&user, "u", "", "user name")
	flag.StringVar(&token, "t", "", "github access token")
	flag.StringVar(&mode, "m", "", "mode, search/read/update/create/close")
	flag.StringVar(&editor, "e", "gedit", "editor, default gedit")
}

func readCredits() {
	data, err := ioutil.ReadFile(creditsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "issues editor:", err)
		return
	}
	d := strings.Split(string(data), "\n")
	token = d[0]
	user = d[1]
}

func saveCredits() {
	os.Mkdir(tmpFolder, 0777)
	data := fmt.Sprintf("%s\n%s", token, user)
	err := ioutil.WriteFile(creditsPath, []byte(data), 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "issue editor:", err)
	}

}

func read() {
	checkParams()
	issue, err := github.ReadIssue(flag.Args())
	check(err)
	fmt.Printf("Number: #%d\nAuthor: %s\nTitle:%s\nBody:\n%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func search() {
	result, err := github.SearchIssues(flag.Args())
	check(err)
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("less than a month old:")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours() < 720 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("less than a year old:")
	for _, item := range result.Items {
		sc := time.Since(item.CreatedAt).Hours()
		if sc >= 720 && sc < 8760 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("more than a year old:")
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours() >= 8760 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

func update() {
	checkParams()
	issue, err := github.ReadIssue(flag.Args())
	check(err)
	data, err := json.MarshalIndent(issue, jsonPrefix, jsonIndent)
	check(err)

	f := fmt.Sprintf("%s%d%d.json", tmpFolder, time.Now().Unix(), issue.Number)
	err = ioutil.WriteFile(f, data, 0644)
	check(err)
	openEditor(f)

	data, err = ioutil.ReadFile(f)
	check(err)
	json.Unmarshal(data, &issue)

	_, err = github.PatchIssue(flag.Args(), issue)
	check(err)
	fmt.Println("Issue updated")
}

func checkParams() {
	if flag.Arg(0) == "" || flag.Arg(1) == "" || flag.Arg(2) == "" {
		log.Fatalln("Arguments should have owner, repo, issue number. current:", flag.Args())
	}
}

func openEditor(f string) {
	e, err := exec.LookPath(editor)
	check(err)
	p, err := os.StartProcess(e, []string{e, f}, &os.ProcAttr{Env: os.Environ()})
	check(err)
	_, err = p.Wait()
	check(err)
}

func clearTmpFiles() {
	ds, err := os.ReadDir(tmpFolder)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, e := range ds {
		if e.Name() == creditsFile {
			continue
		}
		err = os.RemoveAll(e.Name())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	flag.Parse()
	if user == "" || token == "" {
		tmpU, tmpT := user, token
		readCredits()
		if tmpT != "" {
			token = tmpT
			saveCredits()
		}
		if tmpU != "" {
			user = tmpU
			saveCredits()
		}
	} else {
		saveCredits()
	}
	github.OAuth(user, token)

	switch mode {
	case "read":
		read()
	case "search":
		search()
	case "update":
		update()
	}

	clearTmpFiles()
}
