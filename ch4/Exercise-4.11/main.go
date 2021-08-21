package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const creditsFile = "credits"
const creditsFolder = "tmp"
const creditsPath = creditsFolder + "/" + creditsFile

var t string
var u string
var m string

func init() {
	flag.StringVar(&u, "u", "", "user name")
	flag.StringVar(&t, "t", "", "github access token")
	flag.StringVar(&m, "m", "read", "mode, read/update/create/close")
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
	os.Mkdir(creditsFolder, 0777)
	data := fmt.Sprintf("%s\n%s", t, u)
	err := ioutil.WriteFile(creditsPath, []byte(data), 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "issue editor:", err)
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

	switch m {
	case "read":
		fmt.Println("Reading issue", flag.Arg(0))
		issue := flag.Arg(0)
		owner := flag.Arg(1)
		repo := flag.Arg(3)
		if issue == "" || owner == "" || repo == "" {
			log.Fatalln("Input should have owner, repo, issue number. current:", owner, repo, issue)
		}
	}
}
