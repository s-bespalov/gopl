package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Exercise 1.8
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("fetch: %v\n", err)
		}

		// Exercise 1.9
		fmt.Printf("status code: %d\n", resp.StatusCode)

		// Exercise 1.7
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("fetch: reading %s: %v\n", url, err)
		}
	}
}
