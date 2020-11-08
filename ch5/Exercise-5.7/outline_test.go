package main

import (
	"bytes"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestCanParseOutput(t *testing.T) {
	testSites := []string{
		"https://vc.ru",
		"https://facebook.com",
		"https://golang.org",
	}
	for _, site := range testSites {
		resp, err := http.Get(site)
		if err != nil {
			t.Errorf("Error in connection to a test site %s, error: %v\n", site, err)
			continue
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Fatalf("Error in parsing html from a test site %s, error: %v\n", site, err)
			continue
		}
		var buf bytes.Buffer
		output = &buf
		forEachNode(doc, startElement, endElement)
		_, err = html.Parse(&buf)
		if err != nil {
			t.Errorf("Can not parse the generated HTML file, error: %v", err)
		}
	}
}
