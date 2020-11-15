package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var searchID string
var found *html.Node

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if pre(n) {
			return true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) {
			return true
		}
	}
	if post != nil {
		if post(n) {
			return true
		}
	}
	return false
}

func compareID(n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == searchID {
				found = n
				return true
			}
		}
	}
	return false
}

//ElementByID returns first html.Node object in HTML document with given ID
func ElementByID(doc *html.Node, id string) *html.Node {
	searchID = id
	if forEachNode(doc, compareID, nil) {
		return found
	}
	return nil
}

func main() {
	url := os.Args[1]
	elementID := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get an HTML from a web site %s, error: %v\n", url, err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in parsing an HTML from %s, error: %v\n", url, err)
		os.Exit(1)
	}
	n := ElementByID(doc, elementID)
	if n == nil {
		fmt.Printf("Element with id=%q not found at %s\n", elementID, url)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", n)
}
