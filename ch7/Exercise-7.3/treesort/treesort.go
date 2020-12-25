package treesort

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root)
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
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

func (t *tree) String() string {
	r := fmt.Sprintf("%d\n", t.value)
	next := []*tree{t.left, t.right}
	for len(next) > 0 {
		tnext := []*tree{}
		for _, node := range next {
			r = fmt.Sprintf("%s %d", r, node.value)
			if node.left != nil {
				tnext = append(tnext, node.left)
			}
			if node.right != nil {
				tnext = append(tnext, node.right)
			}
		}
		next = tnext
		if len(next) > 0 {
			r += "\n"
		}
	}
	return r
}
