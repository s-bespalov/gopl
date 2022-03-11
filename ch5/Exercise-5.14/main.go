package main

import (
	"fmt"
	"os"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// prints path, if item is directory returns a list of entries
func process(item string) []string {
	fmt.Println(item)
	fi, err := os.Lstat(item)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	if fi.IsDir() {
		items, err := os.ReadDir(item)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}

		r := []string{}
		for _, v := range items {
			p := fmt.Sprintf("%s/%s", item, v.Name())
			r = append(r, p)
		}
		return r
	}
	return nil
}

func main() {
	worklist := process(os.Args[1])
	breadthFirst(process, worklist)
}
