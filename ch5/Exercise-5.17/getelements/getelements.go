package getelements

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, elements ...string) []*html.Node {
	output := []*html.Node{}
	var scan func(n *html.Node)
	scan = func(n *html.Node) {
		if contains(elements, n.Data) {
			output = append(output, n)
		}
		if c := n.FirstChild; c != nil {
			scan(c)
		}
		if s := n.NextSibling; s != nil {
			scan(s)
		}
	}
	scan(doc)
	return output
}

func contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
