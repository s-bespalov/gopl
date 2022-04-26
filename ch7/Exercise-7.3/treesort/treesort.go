package treesort

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Tree struct {
	value       int
	left, right *Tree
}

// Sort values in place
func Sort(values []int) *Tree {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *Tree) print(prefix1, prefix2 string, needline bool, w io.Writer) {
	var val string
	if t == nil {
		if prefix2 == "" {
			return
		}
		val = "-"
	} else {
		val = fmt.Sprintf("%d", t.value)
	}
	if needline {
		l, _ := fmt.Fprintf(w, "%s├%s%s\n", prefix1, prefix2, val)
		l -= (len(prefix1) + 5)
		prefix1 += "│" + strings.Repeat(" ", l)
	} else {
		l, _ := fmt.Fprintf(w, "%s└%s%s\n", prefix1, prefix2, val)
		l -= (len(prefix1) + 5)
		prefix1 += strings.Repeat(" ", l)
	}
	if t != nil && (t.left != nil || t.right != nil) {
		t.right.print(prefix1, "R:", true, w)
		t.left.print(prefix1, "L:", false, w)
	}
}

func (t *Tree) String() string {
	var buf bytes.Buffer
	t.print("", "", false, &buf)
	return buf.String()
}
