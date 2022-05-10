package findlinks

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type myReader struct {
	input string
	idx   int
}

func (r *myReader) Read(p []byte) (n int, err error) {
	if r.idx >= len(r.input) {
		return 0, io.EOF
	}
	n = 0
	for i := 0; i < len(p); i++ {
		p[i] = r.input[r.idx]
		n++
		r.idx++
		if r.idx >= len(r.input) {
			return
		}
	}
	return
}

func newReader(str string) *myReader {
	r := myReader{str, 0}
	return &r
}

func FindLinks(h string) []string {
	var links []string
	doc, err := html.Parse(newReader(h))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	links = append(links, visit(nil, doc)...)
	return links
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
