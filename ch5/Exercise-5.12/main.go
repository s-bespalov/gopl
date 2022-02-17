package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"com/github/s-bespalov/gopl/ch5/Exercise-5.12/outline2"

	"golang.org/x/net/html"
)

func main() {
	body, ok := getWebPage(os.Args[1])
	if !ok {
		return
	}
	n, err := html.Parse(body)
	if err != nil {
		log.Fatal(err)
	}
	outline2.Outline(n, os.Stdout)
}

func getWebPage(url string) (io.ReadCloser, bool) {
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
