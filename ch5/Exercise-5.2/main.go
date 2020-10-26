package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exercise 5.2: %v\n", err)
		os.Exit(1)
	}
	histogram := make(map[string]int)
	visit(histogram, doc)
	for el, count := range histogram {
		fmt.Printf("%s\t%d\n", el, count)
	}
}

func visit(elements map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	if c := n.FirstChild; c != nil {
		visit(elements, c)
	}

	if c := n.NextSibling; c != nil {
		visit(elements, c)
	}
}
