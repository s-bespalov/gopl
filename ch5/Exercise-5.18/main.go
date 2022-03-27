package main

import (
	"com/github/s-bespalov/gopl/ch5/Exercise-5.18/fetch"
	"fmt"
	"os"
)

func main() {
	url := os.Args[1]
	if _, _, err := fetch.Fetch(url); err != nil {
		fmt.Println(err)
	}
}
