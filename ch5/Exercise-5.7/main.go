package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"com.github/s-bespalov/gopl/ch5/Exercise-5.7/outline2"
	"golang.org/x/net/html"
)

func main() {
	body, ok := getWebPage()
	if !ok {
		return
	}
	n, err := html.Parse(body)
	if err != nil {
		log.Fatal(err)
	}
	outline2.ForEachNode(n, outline2.StartElement, outline2.EndElement, os.Stdout, 0)
}

func getWebPage() (io.ReadCloser, bool) {
	url := "https://habr.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Outline2: cant reach %s", url)
		return nil, false
	}
	if resp.StatusCode != 200 {
		fmt.Errorf("Outline2: http error, status code: %d, url: %s", resp.StatusCode, url)
		return nil, false
	}
	return resp.Body, true
}
