package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exercise 5.3: %v\n", err)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode && n.Data != "" {
		fmt.Printf("%#v\n", n.Data)
	}

	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" || n.Data == "air" {
			if c := n.NextSibling; c != nil {
				visit(c)
			}

			return
		}
	}

	if c := n.FirstChild; c != nil {
		visit(c)
	}

	if c := n.NextSibling; c != nil {
		visit(c)
	}
}
