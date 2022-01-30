package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"com.github/s-bespalov/gopl/ch5/Exercise-5.8/search"
	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	id := os.Args[2]
	resp, err := http.Get(url)
	check(err)
	if resp.StatusCode != 200 {
		log.Fatalf("url: %s, status code not OK, status code:%d", url, resp.StatusCode)
	}
	h, err := html.Parse(resp.Body)
	check(err)
	n, err := search.ElsementByID(h, id)
	check(err)
	fmt.Printf("<%s", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s='%s'", a.Key, a.Val)
	}
	fmt.Print(">\n")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
