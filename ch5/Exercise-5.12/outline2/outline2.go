package outline2

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

var voidTags map[string]bool = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"commang": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}

func Outline(n *html.Node, out io.Writer) {
	depth := 0
	w := out
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Fprintf(w, "%*s<%s", depth*2, " ", n.Data)
			for _, a := range n.Attr {
				fmt.Fprintf(w, " %s='%s'", a.Key, html.EscapeString(a.Val))
			}
			if n.FirstChild != nil && hasChild(n.FirstChild) {
				depth++
				fmt.Fprint(w, ">\n")
			} else if voidTags[n.Data] {
				fmt.Fprint(w, "/>\n")
			} else {
				fmt.Fprint(w, ">")
			}
		}
	}
	post := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.FirstChild != nil && hasChild(n.FirstChild) {
				depth--
				fmt.Fprintf(w, "%*s</%s>\n", depth*2, " ", n.Data)
				return
			} else if voidTags[n.Data] {
				return
			}
			fmt.Fprintf(w, "</%s>\n", n.Data)
		}
	}
	forEachNode(n, pre, post)
}

// forEachNode calls the function pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func hasChild(n *html.Node) bool {
	if n.Type == html.ElementNode {
		return true
	}
	if n.FirstChild != nil && hasChild(n.FirstChild) {
		return true
	}
	if n.NextSibling != nil && hasChild(n.NextSibling) {
		return true
	}
	return false
}
