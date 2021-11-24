package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

var ignoredTags map[string]bool = map[string]bool{
	"script": true,
	"style":  true,
	"air":    true,
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("printtext: %v\n", err)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode && !ignoredTags[n.Parent.Data] {
		fmt.Printf("%#v\n", n.Data)
	}
	if n.FirstChild != nil {
		visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
}
