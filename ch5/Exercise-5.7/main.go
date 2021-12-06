package main

import (
	"log"
	"os"

	"com.github/s-bespalov/gopl/ch5/Exercise-5.7/outline2"
	"golang.org/x/net/html"
)

func main() {
	n, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	outline2.ForEachNode(n, outline2.StartElement, outline2.EndElement)
}
