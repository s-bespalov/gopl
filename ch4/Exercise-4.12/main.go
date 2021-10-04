package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/s-bespalov/gopl/ch4/Exercise-4.12/xkcd"
)

func main() {
	var comics *[]xkcd.Comic
	var err error

	flag.Parse()

	switch flag.Arg(0) {
	case "download":
		_, comics, err = xkcd.DownloadAll(5)
		check(err)
		err = xkcd.SaveAll(comics)
		check(err)
		fmt.Printf("%d comics downloaded and saved\n", len(*comics))
	case "show":
		comics, err = xkcd.ReadAll()
		check(err)
		for _, c := range *comics {
			fmt.Printf("%d\t%s\n", c.Num, c.Title)
		}
	case "search":
		num, title, content := getSearchTerms()
		comics, err := xkcd.ReadAll()
		check(err)
		comics = xkcd.Search(comics, num, title, content)
		printResults(comics)
	}
}

func getSearchTerms() (num int, title, content string) {
	args := flag.Args()[1:]
	var err error
	for _, a := range args {
		t := strings.Split(a, ":")
		if t[0] == "num" {
			num, err = strconv.Atoi(t[1])
			check(err)
		}
		if t[0] == "title" {
			title = strings.ToLower(t[1])
		}
		if t[1] == "content" {
			content = strings.ToLower(t[1])
		}
	}
	return
}

func printResults(comics *[]xkcd.Comic) {
	for _, c := range *comics {
		fmt.Printf("#%-5d  %15.15s  %.100q\n", c.Num, c.Title, c.Transcript)
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
