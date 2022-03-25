package getelements

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	testElementsByTagName("https://yandex.ru", t)
	testElementsByTagName("https://habr.com", t)
	testElementsByTagName("https://github.com", t)
}

func testElementsByTagName(url string, t *testing.T) {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("error when loading site %s: %s", url, err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("loading %s, status code: %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	pbytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error when reading %s: %s", url, err)
	}
	page := string(pbytes)
	buf := bytes.NewBuffer(pbytes)
	doc, err := html.Parse(buf)
	if err != nil {
		t.Fatalf("error when parsing %s: %s", url, err)
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	if len(images) == 0 && strings.Contains(page, "<img") {
		t.Errorf("no img elements in result, but page %s has them", url)
	}
	if len(headings) == 0 {
		flag := false
		if strings.Contains(page, "<h1") {
			flag = true
		}
		if strings.Contains(page, "<h2") {
			flag = true
		}
		if strings.Contains(page, "<h3") {
			flag = true
		}
		if strings.Contains(page, "<h4") {
			flag = true
		}
		if flag {
			t.Errorf("no headings elements in result, but page %s has them", url)
		}
	}
	for _, img := range images {
		if img.Data != "img" {
			t.Errorf("%s in images", img.Data)
		}
	}
	for _, h := range headings {
		if !contains([]string{"h1", "h2", "h3", "h4"}, h.Data) {
			t.Errorf("%s in headings", h.Data)
		}
	}
}
