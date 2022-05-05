package findlinks

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestFindLinks(t *testing.T) {
	url := "https://go.dev"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("url: %s, status code:%s\n", url, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	html := string(body)
	links := FindLinks(html)
	allLinks := strings.Join(links, ", ")
	if !strings.Contains(allLinks, "https://github.com/golang") {
		t.Fail()
	}
}
