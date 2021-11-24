package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}
	histogram := make(map[string]int)
	visit(&histogram, doc)
	keys := make([]string, 0, len(histogram))
	for k := range histogram {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s:\t%d\n", k, histogram[k])
	}
}

func visit(h *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		k := *h
		k[n.Data]++
	}
	if n.FirstChild != nil {
		visit(h, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(h, n.NextSibling)
	}
}
