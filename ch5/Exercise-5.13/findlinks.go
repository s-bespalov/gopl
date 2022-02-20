package main

import (
	"com/github/s-bespalov/gopl/ch5/Exercise-5.13/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const rootDir = "downloads"

// breasthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawlSites(worklist []string) {
	saved := make(map[string]bool)
	crawl := func(URL string) []string {
		fmt.Println(URL)
		list, err := links.Extract(URL)
		if err != nil {
			log.Print(err)
		}
		u, err := url.Parse(URL)
		if err != nil {
			return list
		}
		h := u.Hostname()
		for _, URL = range list {
			u2, err := url.Parse(URL)
			if err != nil {
				log.Print(err)
				return list
			}
			if saved[URL] {
				return list
			}
			if h != u2.Hostname() {
				continue
			}
			resp, err := http.Get(URL)
			if err != nil || resp.StatusCode != http.StatusOK {
				return list
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return list
			}
			h = strings.ReplaceAll(h, ".", "_")
			p := fmt.Sprintf("%s/%s/%s", rootDir, h, u2.Path)
			f := fmt.Sprintf("%s/%d_raw", p, time.Now().Unix())
			err = os.MkdirAll(p, 0777)
			if err != nil {
				return list
			}
			os.WriteFile(f, data, 0666)
			saved[URL] = true
		}
		return list
	}
	breadthFirst(crawl, worklist)
}
