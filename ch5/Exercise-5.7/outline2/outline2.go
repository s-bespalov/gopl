package outline2

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// forEachNode calls the function pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func ForEachNode(n *html.Node, pre, post func(n *html.Node), out io.Writer, startDepth int) {
	w = out
	if startDepth != -1 {
		depth = startDepth
	}
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post, out, -1)
	}

	if post != nil {
		post(n)
	}
}

var depth int
var w io.Writer

func StartElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, " %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild != nil && n.FirstChild.Type == html.ElementNode {
			//if n.FirstChild != nil  {
			fmt.Fprint(w, ">\n")
			depth++
		} else {
			fmt.Fprint(w, ">\n")
		}
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		//if n.Type == html.ElementNode && (n.FirstChild != nil && n.FirstChild.Type == html.ElementNode) {
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
