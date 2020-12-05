package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// ElementsByTagName returns all html nodes from document that match one of names
func ElementsByTagName(doc *html.Node, name ...string) (r []*html.Node) {
	if len(name) == 0 {
		return
	}

	seek := make(map[string]bool)
	for _, val := range name {
		seek[val] = true
	}

	var handle func(n *html.Node)
	handle = func(n *html.Node) {
		if n.Type == html.ElementNode && seek[n.Data] {
			r = append(r, n)
		}
		if next := n.FirstChild; next != nil {
			handle(next)
		}
		if next := n.NextSibling; next != nil {
			handle(next)
		}
	}
	handle(doc)
	return
}

func main() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http") {
		log.Fatalf("Incorrect url %s\n", url)
	}
	tags := os.Args[2:]
	if len(tags) == 0 {
		log.Fatalln("tags not set")
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP GET %s returned bad status code %d\n", url, resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	found := ElementsByTagName(doc, tags...)
	for _, element := range found {
		fmt.Println(element)
	}
}
