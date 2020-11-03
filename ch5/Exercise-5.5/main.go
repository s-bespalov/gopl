package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, arg := range os.Args[1:] {
		words, images, err := CountWordsAndImages(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Count words and images error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\nWords: %d, Images: %d\n", arg, words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode && n.Data != "" {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
		if err := input.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading html text node:", err)
			os.Exit(1)
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if nextNode := n.FirstChild; nextNode != nil {
		nextWords, nextImages := countWordsAndImages(nextNode)
		words += nextWords
		images += nextImages
	}
	if nextNode := n.NextSibling; nextNode != nil {
		nextWords, nextImages := countWordsAndImages(nextNode)
		words += nextWords
		images += nextImages
	}
	return
}
