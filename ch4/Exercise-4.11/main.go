package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.11/github"
)

const creditsFile = "credits"
const tmpFolder = "tmp"
const creditsPath = tmpFolder + "/" + creditsFile

var t string
var u string
var m string

func init() {
	flag.StringVar(&u, "u", "", "user name")
	flag.StringVar(&t, "t", "", "github access token")
	flag.StringVar(&m, "m", "search", "mode, search/read/update/create/close")
}

func readCredits() {
	data, err := ioutil.ReadFile(creditsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "issues editor:", err)
		return
	}
	d := strings.Split(string(data), "\n")
	t = d[0]
	u = d[1]
}

func saveCredits() {
	os.Mkdir(tmpFolder, 0777)
	data := fmt.Sprintf("%s\n%s", t, u)
	err := ioutil.WriteFile(creditsPath, []byte(data), 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "issue editor:", err)
	}

}

func read() {
	if flag.Arg(0) == "" || flag.Arg(1) == "" || flag.Arg(2) == "" {
		log.Fatalln("Arguments should have owner, repo, issue number. current:", flag.Args())
	}
	issue, err := github.ReadIssue(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number: #%d\nAuthor: %s\nTitle:%s\nBody:\n%s\n", issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func search() {
	result, err := github.SearchIssues(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
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

func main() {
	flag.Parse()
	if u == "" || t == "" {
		tmpU, tmpT := u, t
		readCredits()
		if tmpT != "" {
			t = tmpT
			saveCredits()
		}
		if tmpU != "" {
			u = tmpU
			saveCredits()
		}
	} else {
		saveCredits()
	}
	github.OAuth(u, t)

	switch m {
	case "read":
		read()
	case "search":
		search()
	}
}
