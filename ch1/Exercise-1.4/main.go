package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type counter struct {
	files map[string]bool
	count int
}

func main() {
	counts := make(map[string]*counter)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			var files []string
			for key := range n.files {
				files = append(files, key)
			}
			fmt.Printf("%d\t%s\t%s\n", n.count, line, strings.Join(files, ","))
		}
	}
}

func countLines(f *os.File, counts map[string]*counter) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		if counts[s] == nil {
			counts[s] = &counter{make(map[string]bool), 0}
		}
		counts[s].files[f.Name()] = true
		counts[s].count++
	}
}
