package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var trackedElements = map[string]bool{
	"a":      true,
	"img":    true,
	"script": true,
	"link":   true,
}
var trackedAttributes = map[string]bool{
	"href": true,
	"src":  true,
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && trackedElements[n.Data] {
		for _, a := range n.Attr {
			if trackedAttributes[a.Key] && strings.HasPrefix(a.Val, "http") {
				links = append(links, a.Val)
			}
		}
	}
	if c := n.FirstChild; c != nil {
		links = visit(links, c)
	}

	if c := n.NextSibling; c != nil {
		links = visit(links, c)
	}

	return links
}
