package search

import "golang.org/x/net/html"

func ElsementByID(h *html.Node, id string) (*html.Node, error) {
	var r *html.Node
	forEachNode(h, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					r = n
					return false
				}
			}
		}
		return true
	}, nil)
	return r, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if !pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return
		}
	}
}
