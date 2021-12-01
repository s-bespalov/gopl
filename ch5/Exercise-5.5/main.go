package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var ignoredTags map[string]bool = map[string]bool{
	"script": true,
	"style":  true,
	"air":    true,
}

func main() {
	URL := os.Args[1]
	w, i, e := CountWordsAdnImages(URL)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("words: %d\nimages: %d\n", w, i)
}

func CountWordsAdnImages(url string) (words, images int, err error) {
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
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		input := bufio.NewScanner(bytes.NewBufferString(n.Data))
		input.Split(bufio.ScanWords)
		if input.Scan() {
			words++
		}

	}
	if n.Type != html.ElementNode || !ignoredTags[n.Data] {
		if n.FirstChild != nil {
			tw, ti := countWordsAndImages(n.FirstChild)
			words, images = words+tw, images+ti
		}
	}
	if n.NextSibling != nil {
		tw, ti := countWordsAndImages(n.NextSibling)
		words, images = words+tw, images+ti
	}
	return
}
