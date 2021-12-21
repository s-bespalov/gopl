package search

import (
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	url := "https://github.com"
	id := "home-community"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("error when trying to fetch %s: %v", url, err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("http get %s, status code NOT ok, status: %d", url, resp.StatusCode)
		return
	}
	n, err := html.Parse(resp.Body)
	if err != nil {
		t.Errorf("can't parse response code for %s: %v", url, err)
		return
	}

	element, err := ElsementByID(n, id)
	if err != nil {
		t.Errorf("error when searching element by id: %v", err)
		return
	}
	if element == nil {
		t.Errorf("not found element with id: %s on page %s", id, url)
		return
	}
	for _, attr := range element.Attr {
		if attr.Key == "id" && attr.Val == id {
			return
		}
	}
	t.Errorf("found elemend don't have id: %s", id)
}
