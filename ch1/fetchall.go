// Fetchall fetches URLs from file in parallel and reports their times and sizes.package ch1
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := make(map[string]bool)

	// read urls from file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		url := input.Text()
		if !strings.HasPrefix(url, "http://") {
			if !strings.HasPrefix(url, "https://") {
				url = "http://" + url
			}
		}
		urls[url] = true
	}

	// fetch urls
	for url := range urls {
		go fetch(url, ch) // start a gorutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	//create file
	path := strings.ReplaceAll(url, "http://", "")
	path = strings.ReplaceAll(path, "https://", "")
	path = strings.ReplaceAll(path, ".", "_") + ".txt"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	output := bufio.NewWriter(f)
	nbytes, err := io.Copy(output, resp.Body)
	resp.Body.Close() // don't leak resources
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v, url, err")
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
