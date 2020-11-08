package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int
var output io.Writer

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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		tag := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			tag = fmt.Sprintf("%s %s='%s'", tag, attr.Key, attr.Val)
		}
		fmt.Fprintf(output, "%s>\n", tag)
		depth++
	}
	if n.Type == html.CommentNode {
		fmt.Fprintf(output, "<!-- %s -->\n", n.Data)
	}
	if n.Type == html.TextNode && n.Data != "" {
		fmt.Fprintf(output, "%*s%#v\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(output, "%*s</%s>\n", depth*2, "", n.Data)
	}
}

func init() {
	output = os.Stdout
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
		forEachNode(doc, startElement, endElement)
	}
}
