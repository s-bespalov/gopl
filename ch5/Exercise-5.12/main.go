package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// forEachNode calls the function pre(x) and post(x) for each node
// x in tree rooted at n. Both functions are optional.
// pre is called before the children are visited(preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func outline(doc *html.Node) {
	var depth int
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	forEachNode(doc, startElement, endElement)
}

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot get an HTML from a web site %s, error: %v\n", arg, err)
			continue
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in parsing an HTML from %s, error: %v\n", arg, err)
			continue
		}
		outline(doc)
	}
}
