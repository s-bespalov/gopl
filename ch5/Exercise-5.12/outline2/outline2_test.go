package outline2

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestGolangOrg(t *testing.T) {
	testOutline2("https://golang.org", t)
}

func TestHabrCom(t *testing.T) {
	testOutline2("https://habr.com", t)
}

func TestGithubCom(t *testing.T) {
	testOutline2("https://github.com", t)
}

func TestFacebook(t *testing.T) {
	testOutline2("https://facebook.com", t)
}

func TestUnsplashCom(t *testing.T) {
	testOutline2("https://unsplash.com", t)
}

func TestYaRu(t *testing.T) {
	testOutline2("https://ya.ru", t)
}

func testOutline2(page string, t *testing.T) {
	resp, err := http.Get(page)
	if err != nil {
		t.Errorf("Outline2: cant reach %s", page)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Outline2: http error, status code: %d, url: %s", resp.StatusCode, page)
	}
	// poarse site
	n, err := html.Parse(resp.Body)
	if err != nil {
		t.Errorf("Error when parsing web page %v", err)
		return
	}
	buf := bytes.Buffer{}
	Outline(n, &buf)
	output := buf.String()

	// parse the output of ForEachNode
	n2, err := html.Parse(&buf)
	if err != nil {
		t.Errorf("Outline2, parsing ForEachNode result failed: %v", err)
	}
	buf2 := bytes.Buffer{}
	Outline(n2, &buf2)
	if buf2.String() != output {
		t.Errorf(`the output of ForEachNode on the result of ForEachNode should be the same, file test_fail_buff.html should be equal test_fail_buff2.html, url:%s`, page)
		ioutil.WriteFile("test_fail_buff.html", []byte(output), 0644)
		ioutil.WriteFile("test_fail_buff2.html", buf2.Bytes(), 0644)
	}
}
