package outline2

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline2(t *testing.T) {
	resp, ok := getTestWebPage(t)
	if !ok {
		return
	}
	n, err := html.Parse(resp)
	if err != nil {
		t.Errorf("Error when parsing web page %v", err)
		return
	}

	buf := bytes.Buffer{}
	ForEachNode(n, StartElement, EndElement, &buf, 0)
	//fmt.Println(string(buf.Bytes()))

	n, err = html.Parse(&buf)
	if err != nil {
		t.Errorf("Outline2, parsing ForEachNode result failed: %v", err)
	}
	ForEachNode(n, StartElement, EndElement, os.Stdout, 0)

}

func getTestWebPage(t *testing.T) (io.ReadCloser, bool) {
	url := "https://habr.com"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Outline2: cant reach %s", url)
		return nil, false
	}
	if resp.StatusCode != 200 {
		t.Errorf("Outline2: http error, status code: %d, url: %s", resp.StatusCode, url)
		return nil, false
	}
	return resp.Body, true
}
